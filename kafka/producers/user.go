package producers

import (
    "github.com/confluentinc/confluent-kafka-go/v2/kafka"
    "encoding/json"
)

type UserEvent struct {
    EventType string `json:"event_type"`
    UserID uint `json:"user_id"`
    Email string `json:"email"`
}

type UserProducer struct {
    config *kafka.ConfigMap `di.inject:"kafkaConfiguration"`
}

type IUserProducer interface {
    Produce(u UserEvent) error; 
}

func (u *UserProducer) Produce(e UserEvent) error {
    producer, err := kafka.NewProducer(u.config)
    if err != nil {
        return err
    }
    defer producer.Close()

    topic := "UserManagement"

    event, err := json.Marshal(e)
    if err != nil {
        return err
    }

    message := &kafka.Message {
      TopicPartition: kafka.TopicPartition {    
          Topic: &topic, 
          Partition: kafka.PartitionAny,
      },  
      Value: event,
    }

    if err := producer.Produce(message, producer.Events()); err != nil {
        return err
    }

    return nil
}
