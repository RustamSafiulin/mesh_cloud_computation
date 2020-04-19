package dto

import "github.com/RustamSafiulin/mesh_cloud_computation/backend/account_service/internal/model"

type TaskFileDto struct {
	ID        string `json:"id,omitempty"`
	TaskID    string `json:"task_id,omitempty"`
	Name      string `json:"name,omitempty"`
	Size      int64  `json:"size,omitempty"`
}

func TaskFileDtoFromTaskFile(tf *model.TaskFile) *TaskFileDto {

	return &TaskFileDto{
		ID:     tf.ID.Hex(),
		TaskID: tf.TaskID.Hex(),
		Name:   tf.Name,
		Size:   tf.Size,
	}
}

func TaskFileDtoListFromTaskFileList(taskFiles []model.TaskFile) []TaskFileDto {

	var result []TaskFileDto
	for _, taskFile := range taskFiles {
		taskFileDto := TaskFileDtoFromTaskFile(&taskFile)
		result = append(result, *taskFileDto)
	}

	return result
}
