package main

import (
	"encoding/json"
	"fmt"
	"github.com/RustamSafiulin/mesh_cloud_computation/backend/common/messaging"
	rpc_model "github.com/RustamSafiulin/mesh_cloud_computation/backend/common/messaging/model"
	"github.com/RustamSafiulin/mesh_cloud_computation/backend/job_service/cmd"
	"github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var appName = "job_service"

var messagingClient *messaging.AmqpClient

func handleSigterm(handleExit func()) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	go func() {
		<-c
		handleExit()
		os.Exit(1)
	}()
}


func failOnError(err error, msg string) {
	if err != nil {
		logrus.Errorf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func ConfigureMessaging(cfg *cmd.Config)  {
	messagingClient = &messaging.AmqpClient{}
	messagingClient.ConnectToBroker(cfg.AMQPUrl)

	err := messagingClient.SubscribeToQueue(rpc_model.TasksStartQueueName, appName, onMessage)
	failOnError(err, "Could not start subscribe to task_queue")

	logrus.Info("Successfully initialize messaging for " + appName)
}

func onMessage(delivery amqp.Delivery) {

	taskStartInfo := &rpc_model.TaskStartInfo{}
	err := json.Unmarshal(delivery.Body, taskStartInfo)

	if err != nil {
		logrus.Debugf("Error was caused during parse json body of TaskStartInfo, Reason: %s", err.Error())
		return
	}

	logrus.Debugf("Task start info: %s", taskStartInfo.ToString())

	taskResultInfo := &rpc_model.TaskResultInfo{
		TaskID:      taskStartInfo.TaskID,
		Result:      rpc_model.TaskResultCompleted,
		WorkerHostIP: "localhost",
		WorkerPort:   8082,
	}

	bytes, _ := json.Marshal(taskResultInfo)
	messagingClient.PublishOnQueue(bytes, rpc_model.TasksResultQueueName)

	time.Sleep(time.Millisecond * 100)
}

func main()  {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.Info("Starting " + appName + "...")

	cfg := cmd.DefaultConfiguration()

	ConfigureMessaging(cfg)

	handleSigterm(func() {
		if messagingClient != nil {
			messagingClient.Close()
		}
	})

	err := http.ListenAndServe(":8082", nil)
	if err != nil {
		logrus.WithError(err).Fatal("Error during start Http server on :8082")
	}
}