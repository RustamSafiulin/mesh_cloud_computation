package dto

import (
	"github.com/RustamSafiulin/mesh_cloud_computation/backend/account_service/internal/model"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type TaskCreationDto struct {
	AccountID   string `json:"account_id,omitempty"`
	Description string `json:"description,omitempty"`
}

func TaskCreationToTask(t *TaskCreationDto) *model.Task {
	return &model.Task{
		AccountID:   bson.ObjectIdHex(t.AccountID),
		Description: t.Description,
		CreatedAt:   time.Now().Unix(),
		State:       model.StateCreated,
	}
}