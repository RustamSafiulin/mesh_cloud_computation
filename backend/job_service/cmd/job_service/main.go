package main

import (
	"fmt"
	"github.com/RustamSafiulin/mesh_cloud_computation/backend/common/messaging"
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

	err := messagingClient.SubscribeToQueue("task_queue", appName, onMessage)
	failOnError(err, "Could not start subscribe to task_queue")

	logrus.Info("Successfully initialize messaging for " + appName)
}

func onMessage(delivery amqp.Delivery) {
	logrus.Infof("Got a message: %v\n", string(delivery.Body))
	time.Sleep(time.Millisecond * 100)
}

func main()  {
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