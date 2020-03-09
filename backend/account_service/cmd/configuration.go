package cmd

var DbName = "service_3d_db"

type Config struct {
	AMQPUrl    string `json:"amqp_url"`
	MongoDBUrl string `json:"mongo_db_url"`
}

func DefaultConfiguration() *Config {
	return &Config{
		MongoDBUrl:"mongodb://127.0.0.1",
		AMQPUrl:"amqp://guest:guest@localhost:5672",
	}
}

func ReadConfiguration(configFilePath string) *Config {
	return &Config{}
}

