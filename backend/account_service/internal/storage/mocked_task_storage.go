package storage

import (
	"github.com/RustamSafiulin/mesh_cloud_computation/backend/account_service/internal/model"
	"github.com/stretchr/testify/mock"
)

type MockedTaskStorage struct {
	mock.Mock
}

func (s *MockedTaskStorage) FindById(id string) (*model.Task, error)  {
	args := s.Mock.Called(id)
	return args.Get(0).(*model.Task), args.Error(1)
}

func (s *MockedTaskStorage) Insert(t *model.Task) (*model.Task, error)  {
	args := s.Mock.Called(t)
	return args.Get(0).(*model.Task), args.Error(1)
}

func (s *MockedTaskStorage) FindAllByAccount(accountId string) ([]model.Task, error)  {
	args := s.Mock.Called(accountId)
	return args.Get(0).([]model.Task), args.Error(1)
}

func (s *MockedTaskStorage) Delete(id string) error  {
	args := s.Mock.Called(id)
	return args.Error(0)
}

func (s *MockedTaskStorage) Update(t *model.Task) error {
	args := s.Mock.Called(t)
	return args.Error(0)
}

func (s *MockedTaskStorage) InsertTaskFile(tf *model.TaskFile) (*model.TaskFile, error) {
	args := s.Mock.Called(tf)
	return args.Get(0).(*model.TaskFile), args.Error(1)
}

func (s *MockedTaskStorage) DeleteTaskFile(taskFileId string) error {
	args := s.Mock.Called(taskFileId)
	return args.Error(0)
}

func (s *MockedTaskStorage) FindTaskFile(taskId string) (*model.TaskFile, error) {
	args := s.Mock.Called(taskId)
	return args.Get(0).(*model.TaskFile), args.Error(1)
}