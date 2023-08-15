package main

import (
	"fmt"
	"kafka_events/internal/handlers"
	"kafka_events/internal/services"
	"kafka_events/pkg/database"
	k "kafka_events/pkg/kafka"
)

func main() {
	topic := "events"

	database.Init()
	defer database.Close()

	go k.InitKafkaConsumer(topic)
	defer k.CloseConsumer()
	fmt.Println("foo")
	eventHandler := &handlers.EventHandler{}
	services.StartKafkaService(eventHandler) // need to rename
}
