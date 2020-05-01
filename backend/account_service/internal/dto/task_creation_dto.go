package dto

import (
	"github.com/RustamSafiulin/mesh_cloud_computation/backend/account_service/internal/model"
	"time"
)

type TaskCreationDto struct {
	Description string `json:"description,omitempty"`
}

func TaskCreationToTask(t *TaskCreationDto) *model.Task {
	return &model.Task{
		Description: t.Description,
		CreatedAt:   time.Now().Unix(),
		StartedAt:   0,
		CompletedAt: 0,
		State:       model.StateCreated,
	}
}