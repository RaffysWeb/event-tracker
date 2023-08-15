package services

import (
	k "kafka_events/pkg/kafka"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type MessageHandler interface {
	HandleMessage(*kafka.Message) error
}

func StartKafkaService(handler MessageHandler) {
	ch := k.GetMessageChannel()

	for msg := range ch {
		if err := handler.HandleMessage(msg); err != nil {
			log.Printf("Error while processing message: %v\n", err)
		}
	}
}
