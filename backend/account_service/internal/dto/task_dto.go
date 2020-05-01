package dto

import (
	"github.com/RustamSafiulin/mesh_cloud_computation/backend/account_service/internal/model"
)

type TaskDto struct {
	ID 		    string `json:"id,omitempty"`
	AccountID   string `json:"account_id,omitempty"`
	Description string `json:"description,omitempty"`
	StartedAt   int64  `json:"started_at"`
	CompletedAt int64  `json:"completed_at"`
	State       int    `json:"state"`
	StateText   string `json:"state_text"`
}

func TaskDtoFromTask(task *model.Task) *TaskDto {

	taskDto := &TaskDto{
		ID:          task.ID.Hex(),
		AccountID:   task.AccountID.Hex(),
		Description: task.Description,
		StartedAt:   task.StartedAt,
		CompletedAt: task.CompletedAt,
		State:       task.State,
		StateText:   model.GetStateStringFromState(task.State),
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