package service

import (
	"fmt"
	"github.com/RustamSafiulin/mesh_cloud_computation/backend/account_service/internal/dto"
	"github.com/RustamSafiulin/mesh_cloud_computation/backend/account_service/internal/model"
	"github.com/RustamSafiulin/mesh_cloud_computation/backend/account_service/internal/storage"
	"github.com/RustamSafiulin/mesh_cloud_computation/backend/common/errors_helper"
	"github.com/RustamSafiulin/mesh_cloud_computation/backend/common/messaging"
	"github.com/globalsign/mgo"
	"github.com/pkg/errors"
	"github.com/sarulabs/di"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type TaskService struct {
	client      *http.Client
	taskStorage storage.BaseTaskStorage
	mqClient	*messaging.AmqpClient
}

func PrepareTaskServiceDef(store storage.BaseTaskStorage, mqClient *messaging.AmqpClient) di.Def {
	return di.Def{
		Name:  "TaskService",
		Build: func(ctn di.Container) (i interface{}, e error) {
			return &TaskService{client: &http.Client{}, taskStorage: store, mqClient: mqClient}, nil
		},
	}
}

func (s *TaskService) CreateNewTask(accountId string, taskCreationDto *dto.TaskCreationDto) (*model.Task, error) {
	logrus.Info("Create new task")

	task := dto.TaskCreationToTask(taskCreationDto)
	err := s.taskStorage.Insert(task)

	if err != nil {
		return nil, errors.WithMessage(errors_helper.ErrStorageError, fmt.Sprintf("Reason: %s", err.Error()))
	}

	return task, err
}

func (s *TaskService) UploadTaskData(taskId string, r *http.Request) error {
	logrus.Info("UploadTaskData")

	r.ParseMultipartForm(32 << 20)
	file, header, err := r.FormFile("task_data")
	if err != nil {
		return errors.WithMessage(errors_helper.ErrParseFormFileHeader, fmt.Sprintf("Reason: %s", err.Error()))
	}

	defer file.Close()

	taskDataRelativePath := strings.Join([]string{"./uploads/", header.Filename, "_", taskId}, "")
	taskDataAbsolutePath, _ := filepath.Abs(taskDataRelativePath)
	f, err := os.Create(taskDataAbsolutePath)

	if err != nil {
		return errors.WithMessage(errors_helper.ErrFileCreation, fmt.Sprintf("File path: %s, reason: %s", taskDataAbsolutePath, err.Error()))
	}

	defer f.Close()
	if _, err := io.Copy(f, file); err != nil {
		return errors.WithMessage(errors_helper.ErrWriteFile, fmt.Sprintf("File path: %s, reason: %s", taskDataAbsolutePath, err.Error()))
	}

	return nil
}

func (s *TaskService) StartTask() error {

	logrus.Info("StartTask")

	err := s.mqClient.PublishOnQueue([]byte{}, "task_queue")
	return err
}

func (s *TaskService) GetAllAccountTasks(accountId string) ([]model.Task, error) {
	logrus.Info("GetAllAccountTasks")

	tasks, err := s.taskStorage.FindAll(accountId)
	if err != nil {
		return nil, errors.WithMessage(errors_helper.ErrStorageError, fmt.Sprintf("Reason: %s", err.Error()))
	}

	return tasks, err
}

func (s *TaskService) GetTaskInfo(id string) (*model.Task, error) {
	logrus.Info("GetTaskInfo")

	task, err := s.taskStorage.FindById(id)

	if err != nil {
		if err == mgo.ErrNotFound {
			return nil, errors.WithMessage(errors_helper.ErrTaskNotExists, fmt.Sprintf("Task ID: %s, Reason: %s", id, err.Error()))
		}

		return nil, errors.WithMessage(errors_helper.ErrStorageError, fmt.Sprintf("Reason: %s", err.Error()))
	}

	return task, nil
}

func (s *TaskService) DeleteTask(id string) error {
	logrus.Info("DeleteTask")

	err := s.taskStorage.Delete(id)
	if err != nil {

		if err == mgo.ErrNotFound {
			return errors.WithMessage(errors_helper.ErrTaskNotExists, fmt.Sprintf("Task ID: %s, Reason: %s", id, err.Error()))
		}

		return errors.WithMessage(errors_helper.ErrStorageError, fmt.Sprintf("Reason: %s", err.Error()))
	}

	return nil
}




