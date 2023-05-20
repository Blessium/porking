package kafka

import (
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func GetConfiguration() *kafka.ConfigMap{
    return &kafka.ConfigMap{
        "bootstrap.servers": "porking-kafka:9092",
        "client.id": "sussy",
    } 
}
