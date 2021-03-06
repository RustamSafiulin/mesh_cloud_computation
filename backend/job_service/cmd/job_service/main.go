package main

import (
	"flag"
	"github.com/RustamSafiulin/mesh_cloud_computation/backend/common/messaging"
	"github.com/RustamSafiulin/mesh_cloud_computation/backend/common/utils"
	"github.com/RustamSafiulin/mesh_cloud_computation/backend/job_service/cmd"
	"github.com/RustamSafiulin/mesh_cloud_computation/backend/job_service/internal/server"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

var (
	appName = "job_service"
	configFilePathArg = flag.String("config_path", "config.json", "Configuration file path")
)

func ConfigureMessaging(cfg *cmd.Config)  *messaging.AmqpClient {
	messagingClient := &messaging.AmqpClient{}
	messagingClient.ConnectToBroker(cfg.AMQPUrl)

	logrus.Info("Successfully initialize messaging for: " + appName)
	return messagingClient
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

func main()  {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.Info("Starting " + cmd.AppName + "...")

	flag.Parse()

	var cfg *cmd.Config
	if *configFilePathArg == "" {
		cfg = cmd.DefaultConfiguration()
	} else {
		cfg = cmd.ReadConfiguration(*configFilePathArg)
	}

	mc := ConfigureMessaging(cfg)
	workerPool, _ := utils.NewWorkerPool(20)

	s := server.NewServer(mc, workerPool)
	s.SetupRoutes()

	handleSigterm(func() {
		mc.Close()
		workerPool.WaitForCompletion()
	})

	s.Start()
}