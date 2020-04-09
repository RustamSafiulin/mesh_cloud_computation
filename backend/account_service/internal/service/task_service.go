package service

import (
	"github.com/RustamSafiulin/mesh_cloud_computation/backend/account_service/internal/dto"
	"github.com/RustamSafiulin/mesh_cloud_computation/backend/account_service/internal/model"
	"github.com/RustamSafiulin/mesh_cloud_computation/backend/account_service/internal/storage"
	"github.com/RustamSafiulin/mesh_cloud_computation/backend/common/errors_helper"
	"github.com/RustamSafiulin/mesh_cloud_computation/backend/common/messaging"
	"github.com/globalsign/mgo"
	"github.com/sarulabs/di"
	"github.com/sirupsen/logrus"
	"net/http"
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
		return nil, errors_helper.NewApplicationError(errors_helper.ErrStorageError)
	}

	return task, err
}

func (s *TaskService) UploadTaskData() error {
	logrus.Info("UploadTaskData")
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
		return nil, errors_helper.NewApplicationError(errors_helper.ErrStorageError)
	}

	return tasks, err
}

func (s *TaskService) GetTaskInfo(id string) (*model.Task, error) {
	logrus.Info("GetTaskInfo")

	task, err := s.taskStorage.FindById(id)

	if err != nil {
		if err == mgo.ErrNotFound {
			return nil, errors_helper.NewApplicationError(errors_helper.ErrTaskNotExists, id)
		}

		return nil, errors_helper.NewApplicationError(errors_helper.ErrStorageError)
	}

	return task, nil
}

func (s *TaskService) DeleteTask(id string) error {
	logrus.Info("DeleteTask")

	err := s.taskStorage.Delete(id)
	if err != nil {

		if err == mgo.ErrNotFound {
			return errors_helper.NewApplicationError(errors_helper.ErrStorageError)
		}

		return errors_helper.NewApplicationError(errors_helper.ErrStorageError)
	}

	return nil
}




