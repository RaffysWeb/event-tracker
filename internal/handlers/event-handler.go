package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"kafka_events/pkg/database/influxdb"
	k "kafka_events/pkg/kafka"
	"kafka_events/pkg/models"
	"log"
	"net/http"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/gin-gonic/gin"
	"github.com/gocql/gocql"
	"github.com/gorilla/websocket"
)

// TODO: move this to config or elsewhere
var topic = "new_events"

type EventHandler struct{}

func NewEventHandler() *EventHandler {
	return &EventHandler{}
}

// HandleLiveEventsWebSocket handles the live events request
func HandleLiveEventsWebSocket(c *gin.Context, db *influxdb.DB) {
	conn, err := websocket.Upgrade(c.Writer, c.Request, nil, 1024, 1024)
	if err != nil {
		log.Println("WebSocket upgrade error:", err)
		c.Status(http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	fmt.Println("HandleLiveEventsWebSocket")
	for {
		now := time.Now()
		fiveMinutesAgo := now.Add(-5 * time.Minute)

		query := fmt.Sprintf(`
		from(bucket: "events")
			|> range(start: %v, stop: %v)
			|> filter(fn: (r) => r["_measurement"] == "kafka_consumer")
	`, fiveMinutesAgo.UTC().Format(time.RFC3339), now.UTC().Format(time.RFC3339))

		queryAPI := db.QueryAPI()

		result, err := queryAPI.Query(context.Background(), query)
		if err != nil {
			fmt.Println(err)
			break
		}
		defer result.Close()

		var events []models.Event

		for result.Next() {
			row := result.Record().Values()
			bytes, err := json.Marshal(row)
			if err != nil {
				fmt.Println("Error marshaling:", err)
				continue
			}

			var event models.Event
			err = json.Unmarshal(bytes, &event)

			if err != nil {
				fmt.Println("Error unmarshaling:", err)
				continue
			}

			// Append the event to the events slice
			events = append(events, event)
		}

		// Send the data to the WebSocket client.
		if err := conn.WriteJSON(events); err != nil {
			log.Println("WebSocket error:", err)
			return
		}

		time.Sleep(time.Second * 5)
	}
}

// HandleKafkaMessage handles the event message from kafka and saves it to cassandra
func (h *EventHandler) HandleKafkaMessage(msg *kafka.Message) error {
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

// CreateEventHandler handles the new events and sends them to kafka
func (h *EventHandler) CreateEventHandler(c *gin.Context) {
	var event models.Event
	if err := c.ShouldBindJSON(&event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to parse JSON",
		})
		return
	}

	event.ID = gocql.TimeUUID()
	event.CreatedAt = time.Now()

	jsonBytes, err := json.Marshal(event)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to marshal event object to JSON",
		})
		return
	}

	k.ProduceMessage(topic, jsonBytes)

	c.JSON(http.StatusCreated, gin.H{
		"message": "Event event sent to Kafka",
	})
}
