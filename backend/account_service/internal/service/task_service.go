package service

import (
	"github.com/RustamSafiulin/3d_reconstruction_service/account_service/internal/model"
	"github.com/RustamSafiulin/3d_reconstruction_service/common/messaging"
	"github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"time"

	"github.com/RustamSafiulin/3d_reconstruction_service/account_service/internal/storage"
	"github.com/sarulabs/di"
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

func (s *TaskService) CreateNewTask(accountId string, t *model.TaskDto) error {
	logrus.Info("Create new task")

	task := &model.Task{
		AccountID:   bson.ObjectIdHex(accountId),
		Description: t.Description,
		DataUrl:     t.DataUrl,
		CreatedAt:   time.Now().Unix(),
		State:       model.StateCreated,
	}

	err := s.taskStorage.Insert(task)
	return err
}

func (s *TaskService) UploadTaskData() error {
	logrus.Info("UploadTaskData")
}

func (s *TaskService) StartTask() error {

	logrus.Info("StartTask")

	err := s.mqClient.PublishOnQueue([]byte{}, "task_queue")
	return err
}

func (s *TaskService) GetAllAccountTasks(accountId string) (*[]model.TaskDto, error) {
	logrus.Info("GetAllAccountTasks")

	tasks, err := s.taskStorage.FindAll(accountId)
	return tasks, err
}

func (s *TaskService) GetTaskInfo(id string) (*model.TaskDto, error) {
	logrus.Info("GetTaskInfo")

	task, err := s.taskStorage.FindById(id)
	taskDto := &model.TaskDto{
		ID:          task.ID.String(),
		AccountID:   task.AccountID.String(),
		Description: task.Description,
		DataUrl:     task.DataUrl,
		StartedAt:   task.StartedAt,
		CompletedAt: task.CompletedAt,
		State:       task.State,
	}
	
	return taskDto, err
}

func (s *TaskService) DeleteTask(id string) error {
	logrus.Info("DeleteTask")

	err := s.taskStorage.Delete(id)
	return err
}




