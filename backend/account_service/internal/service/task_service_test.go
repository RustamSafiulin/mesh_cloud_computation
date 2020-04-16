package service

import (
	"github.com/RustamSafiulin/mesh_cloud_computation/backend/account_service/internal/dto"
	"github.com/RustamSafiulin/mesh_cloud_computation/backend/account_service/internal/model"
	"github.com/RustamSafiulin/mesh_cloud_computation/backend/account_service/internal/storage"
	"github.com/RustamSafiulin/mesh_cloud_computation/backend/common/errors_helper"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"gopkg.in/mgo.v2/bson"
	"testing"
	"time"
)

func TestCreateNewTask(t *testing.T) {

	mockedStorage := &storage.MockedTaskStorage{}
	taskService := &TaskService{taskStorage: mockedStorage, mqClient: nil, client: nil}

	taskDto := &dto.TaskCreationDto{
		AccountID:   bson.NewObjectId().Hex(),
		Description: "test description",
	}

	task := &model.Task{
		AccountID:   bson.ObjectIdHex(taskDto.AccountID),
		Description: taskDto.Description,
		CreatedAt:   time.Now().Unix(),
		CompletedAt: time.Now().Unix(),
		StartedAt:   time.Now().Unix(),
		State:       model.StateCreated,
	}

	mockedStorage.On("Insert", task).Return(nil)

	createdTask, err := taskService.CreateNewTask(bson.NewObjectId().Hex(), taskDto)

	mockedStorage.AssertExpectations(t)
	assert.ObjectsAreEqualValues(task, createdTask)
	assert.Nil(t, err)
}

func TestCreateNewTaskFailed(t *testing.T) {

	mockedStorage := &storage.MockedTaskStorage{}
	taskService := &TaskService{taskStorage: mockedStorage, mqClient: nil, client: nil}

	taskDto := &dto.TaskCreationDto{
		AccountID:   bson.NewObjectId().Hex(),
		Description: "test description",
	}

	task := &model.Task{
		AccountID:   bson.ObjectIdHex(taskDto.AccountID),
		Description: taskDto.Description,
		CreatedAt:   time.Now().Unix(),
		CompletedAt: time.Now().Unix(),
		StartedAt:   time.Now().Unix(),
		State:       model.StateCreated,
	}

	mockedStorage.On("Insert", task).Return(errors.New("Some task creation error"))

	createdTask, err := taskService.CreateNewTask(bson.NewObjectId().Hex(), taskDto)

	mockedStorage.AssertExpectations(t)
	assert.Nil(t, createdTask)
	assert.NotNil(t, err)
	assert.Equal(t, errors_helper.ErrStorageError, errors.Cause(err))
}