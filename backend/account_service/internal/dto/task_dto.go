package dto

import (
	"github.com/RustamSafiulin/mesh_cloud_computation/backend/account_service/internal/model"
	"github.com/sirupsen/logrus"
)

type TaskDto struct {
	ID 		    string `json:"id,omitempty"`
	AccountID   string `json:"account_id,omitempty"`
	Description string `json:"description,omitempty"`
	StartedAt   int64  `json:"started_at,omitempty"`
	CompletedAt int64  `json:"completed_at,omitempty"`
	State       int    `json:"state,omitempty"`
}

func TaskDtoFromTask(task *model.Task) *TaskDto {
	logrus.Debug(task.ID)

	taskDto := &TaskDto{
		ID:          task.ID.Hex(),
		AccountID:   task.AccountID.Hex(),
		Description: task.Description,
		StartedAt:   task.StartedAt,
		CompletedAt: task.CompletedAt,
		State:       task.State,
	}

	return taskDto
}

func TaskDtoListFromTaskList(tasks []model.Task) []TaskDto {

	var result []TaskDto
	for _, task := range tasks {
		taskDto := TaskDtoFromTask(&task)
		result = append(result, *taskDto)
	}

	return result
}