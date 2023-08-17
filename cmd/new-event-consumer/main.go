package main

import (
	"kafka_events/internal/handlers"
	"kafka_events/internal/services"
	"kafka_events/pkg/database/cassandra"
	k "kafka_events/pkg/kafka"
)

func main() {
	topic := "new-events"

	cassandra.Init()
	defer cassandra.Close()

	go k.InitKafkaConsumer(topic)
	defer k.CloseConsumer()

	eventHandler := &handlers.EventHandler{}
	services.StartKafkaService(eventHandler)
}
