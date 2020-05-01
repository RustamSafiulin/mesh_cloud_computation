package cmd

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

var DbName = "service_3d_db"

type Config struct {
	AMQPUrl    string `json:"rabbitmq_url"`
	MongoDBUrl string `json:"mongo_db_url"`
}

func DefaultConfiguration() *Config {
	return &Config{
		AMQPUrl:    "amqp://guest:guest@localhost:5672",
		MongoDBUrl: "mongodb://127.0.0.1",
	}
}

func ReadConfiguration(configFilePath string) *Config {
	configFile, err := os.Open(configFilePath)
	if err != nil {
		return DefaultConfiguration()
	}
	defer configFile.Close()

	byteValue, _ := ioutil.ReadAll(configFile)
	var config Config

	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		return DefaultConfiguration()
	}

	return &config
}
