package handlers

import (
	"encoding/json"
	"fmt"
	k "kafka_events/pkg/kafka"
	"kafka_events/pkg/models"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gocql/gocql"
	"github.com/gofiber/fiber/v2"
)

// TODO: move this to config or elsewhere
var topic = "events"

type EventHandler struct {
	// Add any additional fields or dependencies your handler needs
}

func CreateEventHandler(c *fiber.Ctx) error {
	var event models.Event
	if err := c.BodyParser(&event); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse JSON",
		})
	}

	event.ID = gocql.TimeUUID().String()
	event.CreatedAt = time.Now()

	k.ProduceMessage(topic, event)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "Event event sent to Kafka",
	})
}

func (h *EventHandler) HandleMessage(msg *kafka.Message) error {
	var event models.Event

	err := json.Unmarshal(msg.Value, &event)
	if err != nil {
		return err
	}

	if err := event.SaveEvent(); err != nil {
		fmt.Println("SaveEvent", err)

		return err
	}

	return nil
}
