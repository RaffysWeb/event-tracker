package kafka

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var consumer *kafka.Consumer
var messageChannel = make(chan *kafka.Message, 5)

func InitKafkaConsumer(topic string) {
	var err error

	configMap := &kafka.ConfigMap{
		"bootstrap.servers": kafkaBrokers,
		"group.id":          "consumer-group",
		"auto.offset.reset": "earliest",
	}

	consumer, err = kafka.NewConsumer(configMap)
	if err != nil {
		log.Fatalf("Failed to create Kafka consumer: %s", err)
	}

	err = consumer.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		log.Fatalf("Failed to subscribe to topic %s: %s", topic, err)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigChan
		log.Printf("Received signal: %v. Shutting down consumer...\n", sig)
		CloseConsumer()
		close(messageChannel)
		os.Exit(0)
	}()

	for {
		msg, err := consumer.ReadMessage(-1)
		if err != nil {
			log.Printf("Error while reading message: %v\n", err)
			continue
		}

		log.Printf("Received message from Kafka topic %s: %s\n", *msg.TopicPartition.Topic, string(msg.Value))
		messageChannel <- msg
	}
}

func GetMessageChannel() <-chan *kafka.Message {
	return messageChannel
}

func CloseConsumer() {
	if consumer == nil {
		log.Println("Kafka consumer is not initialized.")
		return
	}

	consumer.Close()
	log.Println("Kafka consumer closed successfully.")
}
