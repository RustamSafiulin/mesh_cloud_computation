package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/RustamSafiulin/mesh_cloud_computation/backend/account_service/internal/dto"
	"github.com/RustamSafiulin/mesh_cloud_computation/backend/account_service/internal/model"
	"github.com/RustamSafiulin/mesh_cloud_computation/backend/account_service/internal/storage"
	"github.com/RustamSafiulin/mesh_cloud_computation/backend/common/errors_helper"
	"github.com/RustamSafiulin/mesh_cloud_computation/backend/common/messaging"
	rpc_model "github.com/RustamSafiulin/mesh_cloud_computation/backend/common/messaging/model"
	"github.com/RustamSafiulin/mesh_cloud_computation/backend/common/utils"
	"github.com/pkg/errors"
	"github.com/sarulabs/di"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type TaskService struct {
	client      *http.Client
	taskStorage storage.BaseTaskStorage
	mqClient    *messaging.AmqpClient
}

func PrepareTaskServiceDef(store storage.BaseTaskStorage, mqClient *messaging.AmqpClient) di.Def {
	return di.Def{
		Name: "TaskService",
		Build: func(ctn di.Container) (i interface{}, e error) {
			return newTaskService(store, mqClient), nil
		},
	}
}

func newTaskService(store storage.BaseTaskStorage, mqClient *messaging.AmqpClient) *TaskService {

	ts := &TaskService{client: &http.Client{}, taskStorage: store, mqClient: mqClient}
	ts.createMessagingSubscriptions()
	return ts
}

func (s *TaskService) createMessagingSubscriptions() {
	s.mqClient.SubscribeToQueue(rpc_model.TasksResultQueueName, "account_service", s.onTaskCompletedHandler)
}

func (s *TaskService) onTaskCompletedHandler(delivery amqp.Delivery) {

	taskResult := &rpc_model.TaskResultInfo{}
	err := json.Unmarshal(delivery.Body, taskResult)
	if err != nil {
		logrus.Debugf("Error was caused during parse task result info from worker. Reason: %s", err.Error())
		return
	}

	logrus.Debugf("Task result info: %s", taskResult.ToString())
}

func (s *TaskService) CreateNewTask(accountId string, taskCreationDto *dto.TaskCreationDto) (*model.Task, error) {

	task, err := s.taskStorage.Insert(dto.TaskCreationToTask(taskCreationDto))

	if err != nil {
		return task, errors.WithMessage(errors_helper.ErrStorageError, fmt.Sprintf("Reason: %s", err.Error()))
	}

	return task, err
}

func (s *TaskService) DownloadTaskData(taskId string) (*model.TaskFile, error) {

	tf, err := s.taskStorage.FindTaskFile(taskId)
	if err != nil {
		if err == mgo.ErrNotFound {
			return tf, errors.WithMessage(errors_helper.ErrTaskDataFileNotExists, fmt.Sprintf("Task ID: %s, Reason: %s", taskId, err.Error()))
		}

		return tf, errors.WithMessage(errors_helper.ErrStorageError, fmt.Sprintf("Reason: %s", err.Error()))
	}

	return tf, err
}

func (s *TaskService) UploadTaskData(taskId string, r *http.Request) (*model.TaskFile, error) {

	r.ParseMultipartForm(32 << 20)
	file, header, err := r.FormFile("task_data")
	if err != nil {
		return nil, errors.WithMessage(errors_helper.ErrParseFormFileHeader, fmt.Sprintf("Reason: %s", err.Error()))
	}

	defer file.Close()

	uploadsDir := strings.Join([]string{"./uploads/", taskId}, "")
	if _, err := os.Stat(uploadsDir); os.IsNotExist(err) {

		if err = os.Mkdir(uploadsDir, os.ModePerm); err != nil {
			return nil, errors.WithMessage(errors_helper.ErrCreateDirectory, fmt.Sprintf("Directory path: %s, Reason: %s", uploadsDir, err.Error()))
		}
	}

	taskDataRelativePath := strings.Join([]string{uploadsDir, header.Filename}, "//")
	taskDataAbsolutePath, _ := filepath.Abs(taskDataRelativePath)
	f, err := os.Create(taskDataAbsolutePath)

	if err != nil {
		return nil, errors.WithMessage(errors_helper.ErrFileCreation, fmt.Sprintf("File path: %s, reason: %s", taskDataAbsolutePath, err.Error()))
	}

	defer f.Close()
	if _, err := io.Copy(f, file); err != nil {
		return nil, errors.WithMessage(errors_helper.ErrWriteFile, fmt.Sprintf("File path: %s, reason: %s", taskDataAbsolutePath, err.Error()))
	}

	md5hash, err := utils.CalculateMD5HashFormFile(file)
	if err != nil {
		logrus.Debugf("Something wrong when calculate md5 hash from file. Reason: %s", err.Error())
	}

	taskFileInfo := &model.TaskFile{
		TaskID:    bson.ObjectIdHex(taskId),
		Path:      taskDataAbsolutePath,
		Name:      header.Filename,
		Size:      header.Size,
		MD5:       md5hash,
		CreatedAt: time.Now().Unix(),
	}

	taskFile, err := s.taskStorage.InsertTaskFile(taskFileInfo)
	if err != nil {
		return nil, errors.WithMessage(errors_helper.ErrStorageError, fmt.Sprintf("Reason: %s", err.Error()))
	}

	return taskFile, nil
}

func (s *TaskService) StartTask(taskId string) (*model.Task, error) {

	task, err := s.GetTaskInfo(taskId)

	if err != nil {
		return task, err
	}

	tf, err := s.taskStorage.FindTaskFile(taskId)
	if err != nil {
		if err == mgo.ErrNotFound {
			return task, errors.WithMessage(errors_helper.ErrTaskDataFileNotExists, fmt.Sprintf("Task ID: %s, Reason: %s", taskId, err.Error()))
		}

		return task, errors.WithMessage(errors_helper.ErrStorageError, fmt.Sprintf("Reason: %s", err.Error()))
	}

	taskDataInfo := &rpc_model.TaskStartInfo{
		FileID:        tf.ID.Hex(),
		TaskID:        tf.TaskID.Hex(),
		FileName:      tf.Name,
		ServiceHostIP: "localhost",
		ServicePort:   8081,
	}

	data, err := json.Marshal(taskDataInfo)

	err = s.mqClient.PublishOnQueue(data, rpc_model.TasksStartQueueName)
	if err != nil {
		return nil, errors.WithMessage(errors_helper.ErrStartComputationTask, fmt.Sprintf("Reason: %s", err.Error()))
	}

	task.State = model.StateRunning
	task.StartedAt = time.Now().Unix()
	err = s.taskStorage.Update(task)
	if err != nil {
		return nil, errors.WithMessage(errors_helper.ErrStorageError, fmt.Sprintf("Reason: %s", err.Error()))
	}

	return task, err
}

func (s *TaskService) GetAllAccountTasks(accountId string) ([]model.Task, error) {

	tasks, err := s.taskStorage.FindAllByAccount(accountId)
	if err != nil {
		return nil, errors.WithMessage(errors_helper.ErrStorageError, fmt.Sprintf("Reason: %s", err.Error()))
	}

	return tasks, err
}

func (s *TaskService) GetTaskInfo(id string) (*model.Task, error) {

	task, err := s.taskStorage.FindById(id)

	if err != nil {
		if err == mgo.ErrNotFound {
			return task, errors.WithMessage(errors_helper.ErrTaskNotExists, fmt.Sprintf("Task ID: %s, Reason: %s", id, err.Error()))
		}

		return task, errors.WithMessage(errors_helper.ErrStorageError, fmt.Sprintf("Reason: %s", err.Error()))
	}

	return task, nil
}

func (s *TaskService) DeleteTask(id string) error {

	err := s.taskStorage.Delete(id)
	if err != nil {

		if err == mgo.ErrNotFound {
			return errors.WithMessage(errors_helper.ErrTaskNotExists, fmt.Sprintf("Task ID: %s, Reason: %s", id, err.Error()))
		}

		return errors.WithMessage(errors_helper.ErrStorageError, fmt.Sprintf("Reason: %s", err.Error()))
	}

	return nil
}
