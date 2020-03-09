package handler

import (
	"github.com/RustamSafiulin/3d_reconstruction_service/account_service/internal/service"
	"github.com/sarulabs/di"
	"net/http"
)

// TaskHandler handle task routes
type TaskHandler struct {
	ctn di.Container
}

func NewTaskHandler(ctn di.Container) *TaskHandler {
	return &TaskHandler{ctn: ctn}
}

func (h *TaskHandler) CreateTaskHandler(w http.ResponseWriter, r *http.Request) {

	service := h.ctn.Get("TaskService").(*service.TaskService)
	service.CreateNewTask()
}

func (h *TaskHandler) UploadTaskDataHandler(w http.ResponseWriter, r *http.Request) {

	service := h.ctn.Get("TaskService").(*service.TaskService)
	service.UploadTaskData()
}

func (h *TaskHandler) StartTaskHandler(w http.ResponseWriter, r *http.Request) {

	service := h.ctn.Get("TaskService").(*service.TaskService)
	service.StartTask()
}

func (h *TaskHandler) GetAllAccountTasksHandler(w http.ResponseWriter, r *http.Request) {

	service := h.ctn.Get("TaskService").(*service.TaskService)
	service.GetAllAccountTasks()
}

func (h *TaskHandler) GetTaskHandler(w http.ResponseWriter, r *http.Request) {

	service := h.ctn.Get("TaskService").(*service.TaskService)
	service.GetTaskInfo()
}

func (h *TaskHandler) DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {

	service := h.ctn.Get("TaskService").(*service.TaskService)
	service.DeleteTask()
}

