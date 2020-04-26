package cmd

var DbName = "service_3d_db"

type Config struct {
	AMQPUrl    string `json:"amqp_url"`
	MongoDBUrl string `json:"mongo_db_url"`
}

func DefaultConfiguration() *Config {
	return &Config{
		MongoDBUrl: "mongodb://rust:123@mongodb:27017/service_3d_db",
		AMQPUrl:    "amqp://rust:123@rabbitmq/",
	}
}

func ReadConfiguration(configFilePath string) *Config {
	return &Config{}
}
