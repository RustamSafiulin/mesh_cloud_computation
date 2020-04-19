package main

import (
	"github.com/RustamSafiulin/mesh_cloud_computation/backend/account_service/cmd"
	"github.com/RustamSafiulin/mesh_cloud_computation/backend/account_service/internal/handler"
	"github.com/RustamSafiulin/mesh_cloud_computation/backend/account_service/internal/server"
	"github.com/RustamSafiulin/mesh_cloud_computation/backend/account_service/internal/service"
	"github.com/RustamSafiulin/mesh_cloud_computation/backend/account_service/internal/storage"
	"github.com/RustamSafiulin/mesh_cloud_computation/backend/common/messaging"
	"github.com/sarulabs/di"
	"github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2"
	"os"
	"os/signal"
	"syscall"
)

func ConfigureMessaging(cfg *cmd.Config) *messaging.AmqpClient {

	messageClient := &messaging.AmqpClient{}
	messageClient.ConnectToBroker(cfg.AMQPUrl)
	return messageClient
}

func ConfigureMongoSession(cfg *cmd.Config) *mgo.Session {
	s, err := mgo.Dial(cfg.MongoDBUrl)
	if err != nil {
		logrus.Fatal(err.Error())
	}

	return s
}

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

func ConfigureServices(defs []di.Def) (di.Container, error) {

	builder, err := di.NewBuilder()
	if err != nil {
		return nil, err
	}

	if err := builder.Add(defs...); err != nil {
		return nil, err
	}

	return builder.Build(), nil
}

func main() {

	logrus.Info("Starting account_service")

	config := cmd.DefaultConfiguration()
	mc := ConfigureMessaging(config)
	mgoSession := ConfigureMongoSession(config)

	accountStorage := storage.NewAccountStorage(mgoSession)
	taskStorage := storage.NewTaskStorage(mgoSession)
	ctn, err := ConfigureServices([]di.Def{
		service.PrepareAccountServiceDef(accountStorage),
		service.PrepareTaskServiceDef(taskStorage, mc),
	})

	if err != nil {
		panic("Cannot configure services")
	}

	s := server.NewServer(handler.NewAccountHandler(ctn), handler.NewTaskHandler(ctn))
	s.SetupRoutes()
	handleSigterm(func() {
		mc.Close()
		mgoSession.Close()
	})

	s.Start()
}

