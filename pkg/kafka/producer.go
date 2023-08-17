package kafka

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var producer *kafka.Producer

func InitKafkaProducer() {
	configMap := &kafka.ConfigMap{
		"bootstrap.servers": kafkaBrokers,
	}
	var err error
	producer, err = kafka.NewProducer(configMap)
	if err != nil {
		log.Fatalf("Failed to create Kafka producer: %s", err)
	}

	go func() {
		for e := range producer.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					log.Printf("Failed to produce message: %v\n", ev.TopicPartition.Error)
				} else {
					log.Printf("Produced message to Kafka: %s\n", string(ev.Value))
				}
			}
		}
	}()
}

func ProduceMessage(topic string, message []byte) {
	if producer == nil {
		log.Println("Kafka producer is not initialized.")
		return
	}

	deliveryChan := make(chan kafka.Event)

	err := producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          message,
	}, deliveryChan)

	if err != nil {
		log.Fatalf("Failed to produce message: %s", err)
	}

	e := <-deliveryChan
	m := e.(*kafka.Message)

	if m.TopicPartition.Error != nil {
		log.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
	} else {
		log.Printf("Delivered message to topic %s [%d] at offset %v\n",
			*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
	}
}

func CloseProducer() {
	if producer == nil {
		log.Println("Kafka producer is not initialized.")
		return
	}

	producer.Close()
	log.Println("Kafka producer closed successfully.")
}
