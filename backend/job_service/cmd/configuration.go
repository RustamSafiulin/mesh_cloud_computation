package cmd

type Config struct {
	AMQPUrl    string `json:"amqp_url"`
}

func DefaultConfiguration() *Config {
	return &Config{
		AMQPUrl:"amqp://guest:guest@localhost:5672",
	}
}