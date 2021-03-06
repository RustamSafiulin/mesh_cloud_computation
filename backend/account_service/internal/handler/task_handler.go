package handler

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/RustamSafiulin/mesh_cloud_computation/backend/account_service/internal/dto"
	"github.com/RustamSafiulin/mesh_cloud_computation/backend/account_service/internal/service"
	"github.com/RustamSafiulin/mesh_cloud_computation/backend/common/errors_helper"
	"github.com/RustamSafiulin/mesh_cloud_computation/backend/common/helpers"
	"github.com/RustamSafiulin/mesh_cloud_computation/backend/common/middleware"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"github.com/sarulabs/di"
	"github.com/sirupsen/logrus"
)

// TaskHandler handle task routes
type TaskHandler struct {
	ctn di.Container
}

func NewTaskHandler(ctn di.Container) *TaskHandler {
	return &TaskHandler{ctn: ctn}
}

func (h *TaskHandler) CreateTaskHandler(w http.ResponseWriter, r *http.Request) {

	var taskCreationDto dto.TaskCreationDto

	err := helpers.ReadJSONBody(r, &taskCreationDto)
	if err != nil {
		helpers.WriteJSONResponse(w, http.StatusBadRequest, dto.ErrorMsgResponse{err.Error()})
		return
	}

	accountId, ok := middleware.AccountIDFromContext(r.Context())
	if !ok {
		helpers.WriteJSONResponse(w, http.StatusNotFound, dto.ErrorMsgResponse{errors_helper.ErrAccountIdNotFoundInContext.Error()})
		return
	}

	service := h.ctn.Get("TaskService").(*service.TaskService)

	task, err := service.CreateNewTask(accountId, &taskCreationDto)
	if err != nil {
		helpers.WriteJSONResponse(w, http.StatusInternalServerError, dto.ErrorMsgResponse{err.Error()})
	} else {
		taskDto := dto.TaskDtoFromTask(task)
		helpers.WriteJSONResponse(w, http.StatusOK, taskDto)
	}
}

func (h *TaskHandler) UploadTaskDataHandler(w http.ResponseWriter, r *http.Request) {

	service := h.ctn.Get("TaskService").(*service.TaskService)

	var taskId = mux.Vars(r)["task_id"]
	taskFileInfo, err := service.UploadTaskData(taskId, r)

	if err != nil {
		helpers.WriteJSONResponse(w, http.StatusInternalServerError, dto.ErrorMsgResponse{err.Error()})
	} else {
		helpers.WriteJSONResponse(w, http.StatusOK, dto.TaskFileDtoFromTaskFile(taskFileInfo))
	}
}

func (h *TaskHandler) DownloadTaskDataHandler(w http.ResponseWriter, r *http.Request) {

	service := h.ctn.Get("TaskService").(*service.TaskService)

	var taskId = mux.Vars(r)["task_id"]

	tf, err := service.DownloadTaskData(taskId)
	if err != nil {
		helpers.WriteJSONResponse(w, http.StatusInternalServerError, err.Error())
		return
	} else {

		fileHandle, err := os.Open(tf.Path)
		if err != nil {
			fullErr := errors.WithMessage(errors_helper.ErrFileCreation, fmt.Sprintf("Task ID: %s, Reason: %s", taskId, err.Error()))
			helpers.WriteJSONResponse(w, http.StatusInternalServerError, dto.ErrorMsgResponse{fullErr.Error()})
			return
		}
		defer fileHandle.Close()

		w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", tf.Name))
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Length", fmt.Sprintf("%d", tf.Size))
		io.Copy(w, fileHandle)
	}
}

func (h *TaskHandler) StartTaskHandler(w http.ResponseWriter, r *http.Request) {

	service := h.ctn.Get("TaskService").(*service.TaskService)

	var taskId = mux.Vars(r)["task_id"]
	updatedTask, err := service.StartTask(taskId)

	if err != nil {
		logrus.Debugf("Error was caused when start task. Reason: %s", err.Error())
	} else {
		updatedTaskDto := dto.TaskDtoFromTask(updatedTask)
		helpers.WriteJSONResponse(w, http.StatusOK, updatedTaskDto)
	}
}

func (h *TaskHandler) StopTaskHandler(w http.ResponseWriter, r *http.Request) {

	service := h.ctn.Get("TaskService").(*service.TaskService)

	var taskId = mux.Vars(r)["task_id"]
	updatedTask, err := service.StopTask(taskId)

	if err != nil {
		logrus.Debugf("Error was caused during stop task. Reason: %s", err.Error())
	} else {
		updatedTaskDto := dto.TaskDtoFromTask(updatedTask)
		helpers.WriteJSONResponse(w, http.StatusOK, updatedTaskDto)
	}
}

func (h *TaskHandler) GetAllAccountTasksHandler(w http.ResponseWriter, r *http.Request) {

	service := h.ctn.Get("TaskService").(*service.TaskService)

	accountId, ok := middleware.AccountIDFromContext(r.Context())
	if !ok {
		helpers.WriteJSONResponse(w, http.StatusNotFound, dto.ErrorMsgResponse{errors_helper.ErrAccountIdNotFoundInContext.Error()})
	}

	tasks, err := service.GetAllAccountTasks(accountId)

	if err != nil {
		helpers.WriteJSONResponse(w, http.StatusInternalServerError, dto.ErrorMsgResponse{err.Error()})
	} else {

		taskDtos := dto.TaskDtoListFromTaskList(tasks)

		if len(taskDtos) == 0 {
			helpers.WriteJSONResponse(w, http.StatusNoContent, taskDtos)
			return
		}

		helpers.WriteJSONResponse(w, http.StatusOK, taskDtos)
	}
}

func (h *TaskHandler) GetTaskHandler(w http.ResponseWriter, r *http.Request) {

	service := h.ctn.Get("TaskService").(*service.TaskService)

	var taskId = mux.Vars(r)["task_id"]
	task, err := service.GetTaskInfo(taskId)

	if err != nil {

		if errors.Cause(err) == errors_helper.ErrTaskNotExists {
			helpers.WriteJSONResponse(w, http.StatusNotFound, dto.ErrorMsgResponse{err.Error()})
			return
		}

		helpers.WriteJSONResponse(w, http.StatusInternalServerError, dto.ErrorMsgResponse{err.Error()})
	} else {

		taskDto := dto.TaskDtoFromTask(task)
		helpers.WriteJSONResponse(w, http.StatusOK, taskDto)
	}
}

func (h *TaskHandler) DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {

	service := h.ctn.Get("TaskService").(*service.TaskService)

	var taskId = mux.Vars(r)["task_id"]
	err := service.DeleteTask(taskId)

	if err != nil {

		if errors.Cause(err) == errors_helper.ErrTaskNotExists {
			helpers.WriteJSONResponse(w, http.StatusNotFound, dto.ErrorMsgResponse{err.Error()})
			return
		}

		helpers.WriteJSONResponse(w, http.StatusInternalServerError, dto.ErrorMsgResponse{err.Error()})
	} else {
		w.WriteHeader(http.StatusOK)
	}
}
