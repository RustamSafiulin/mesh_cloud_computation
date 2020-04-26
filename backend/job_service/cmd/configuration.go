package cmd

var AppName = "job_service"

type Config struct {
	AMQPUrl string `json:"amqp_url"`
}

func DefaultConfiguration() *Config {
	return &Config{
		AMQPUrl: "amqp://rust:123@rabbitmq/",
	}
}
