package handler

import (
	"bytes"
	"encoding/json"
	"github.com/RustamSafiulin/mesh_cloud_computation/backend/account_service/internal/dto"
	"github.com/RustamSafiulin/mesh_cloud_computation/backend/account_service/internal/model"
	"github.com/RustamSafiulin/mesh_cloud_computation/backend/account_service/internal/service"
	"github.com/RustamSafiulin/mesh_cloud_computation/backend/account_service/internal/storage"
	"github.com/RustamSafiulin/mesh_cloud_computation/backend/common/middleware"
	"github.com/sarulabs/di"
	"github.com/stretchr/testify/assert"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestCreateTask(t *testing.T) {

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

	mockTaskStorage := &storage.MockedTaskStorage{}
	ctn, err := createTestDiContainer([]di.Def{
		service.PrepareTaskServiceDef(mockTaskStorage, nil),
	})

	assert.Nil(t, err)

	taskHandler := NewTaskHandler(ctn)

	data, _ := json.Marshal(taskDto)
	req, err := http.NewRequest("POST", "/api/v1/tasks", bytes.NewBuffer(data))

	ctx := middleware.NewAccountIDContext(req.Context(), bson.NewObjectId().Hex())
	req = req.WithContext(ctx)

	if assert.Nil(t, err) {
		rr := httptest.NewRecorder()
		h := http.HandlerFunc(taskHandler.CreateTaskHandler)
		mockTaskStorage.On("Insert", task).Return(nil)
		h.ServeHTTP(rr, req)
		mockTaskStorage.AssertExpectations(t)
		assert.Equal(t, rr.Code, 200)
	}
}

func createTestDiContainer(defs []di.Def) (di.Container, error) {

	builder, err := di.NewBuilder()
	if err != nil {
		return nil, err
	}

	if err := builder.Add(defs...); err != nil {
		return nil, err
	}

	return builder.Build(), nil
}