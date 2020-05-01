package cmd

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

var AppName = "job_service"

type Config struct {
	AMQPUrl string `json:"amqp_url"`
}

func DefaultConfiguration() *Config {
	return &Config{
		AMQPUrl: "amqp://guest:guest@localhost:5672",
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
