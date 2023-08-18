package main

import (
	"fmt"
	"kafka_events/internal/routes"
	"kafka_events/pkg/kafka"
	"log"
)

func main() {
	kafka.InitKafkaProducer()
	defer kafka.CloseProducer()

	router := routes.SetupNewEventRoutes()
	fmt.Println("Server started on port 4000")
	log.Fatal(router.Run(":4000"))
}
