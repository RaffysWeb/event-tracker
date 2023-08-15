package main

import (
	"fmt"
	"kafka_events/internal/routes"
	"kafka_events/pkg/kafka"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	kafka.InitKafkaProducer()
	defer kafka.CloseProducer()

	app := fiber.New(fiber.Config{})
	routes.SetupEventRoutes(app)

	fmt.Println("Server started on port 3000")
	log.Fatal(app.Listen(":3000"))
}
