package config

import (
	"os"
	"strconv"
)

type Config struct {
	PostgresDsn      string
	RedisAddress     string
	RedisPassword    string
	ElasticsearchUrl string
	GrpcPort         int
	KafkaBrokers     string
}

func LoadConfig() (*Config, error) {
	grpcPort, err := strconv.Atoi(os.Getenv("GRPC_PORT"))
	if err != nil {
		panic("Could not parse GRPC_PORT")
	}

	return &Config{
		PostgresDsn:      os.Getenv("POSTGRES_DSN"),
		RedisAddress:     os.Getenv("REDIS_ADDRESS"),
		RedisPassword:    os.Getenv("REDIS_PASSWORD"),
		ElasticsearchUrl: os.Getenv("ELASTICSEARCH_URL"),
		GrpcPort:         grpcPort,
		KafkaBrokers:     os.Getenv("KAFKA_BROKERS"),
	}, nil
}
