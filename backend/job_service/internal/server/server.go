package server

import (
	"encoding/json"
	"github.com/RustamSafiulin/mesh_cloud_computation/backend/common/helpers"
	"github.com/RustamSafiulin/mesh_cloud_computation/backend/common/messaging"
	rpc_model "github.com/RustamSafiulin/mesh_cloud_computation/backend/common/messaging/model"
	"github.com/RustamSafiulin/mesh_cloud_computation/backend/common/middleware"
	"github.com/RustamSafiulin/mesh_cloud_computation/backend/common/utils"
	"github.com/RustamSafiulin/mesh_cloud_computation/backend/job_service/cmd"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"net/http"
	"time"
)

type Server struct {
	r *mux.Router
	client *http.Client
	messageClient *messaging.AmqpClient
	workerPool *utils.WorkerPool
}

func NewServer(messageClient *messaging.AmqpClient, workerPool *utils.WorkerPool) *Server {

	s := &Server{
		r:             mux.NewRouter(),
		messageClient: messageClient,
		workerPool: workerPool,
		client: &http.Client{},
	}

	s.messageClient.SubscribeToQueue(rpc_model.TasksStartQueueName, cmd.AppName, s.onStartTask)

	return s
}

func (s *Server) Start()  {

	err := http.ListenAndServe(":8082", helpers.EnableCors(s.r))
	if err != nil {
		logrus.WithError(err).Fatal("Error during start Http server on port 8082")
	}
}

func (s *Server) SetupRoutes() {

	s.r.Use(middleware.PanicRecoveryMiddleware)

	api := s.r.PathPrefix("/api/v1/").Subrouter()
	api.HandleFunc("/results/download", s.DownloadTaskResult).Methods("GET")
}

func (s *Server) DownloadTaskResult(w http.ResponseWriter, r *http.Request) {
}

func (s *Server) onStartTask(delivery amqp.Delivery) {

	s.workerPool.QueueWorkItem(func() {
		taskStartInfo := &rpc_model.TaskStartInfo{}
		err := json.Unmarshal(delivery.Body, taskStartInfo)

		if err != nil {
			logrus.Debugf("Error was caused during parse json body of TaskStartInfo, Reason: %s", err.Error())
			return
		}

		logrus.Debugf("Task start info: %s", taskStartInfo.ToString())
		s.longJobFunc(*taskStartInfo)
	})
}

func (s *Server) longJobFunc(taskStartInfo rpc_model.TaskStartInfo) {

	logrus.Debugln("Stub wait with timeout 30s")
	time.Sleep(time.Second * 30)

	taskResultInfo := &rpc_model.TaskResultInfo{
		TaskID:       taskStartInfo.TaskID,
		Result:       rpc_model.TaskResultCompleted,
		WorkerHostIP: "localhost",
		WorkerPort:   8082,
	}

	bytes, _ := json.Marshal(taskResultInfo)
	s.messageClient.PublishOnQueue(bytes, rpc_model.TasksResultQueueName)
}

