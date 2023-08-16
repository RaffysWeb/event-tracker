package kafka

import (
	"fmt"
	"log"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/confluentinc/confluent-kafka-go/schemaregistry"
	"github.com/confluentinc/confluent-kafka-go/schemaregistry/serde"
	"github.com/confluentinc/confluent-kafka-go/schemaregistry/serde/avro"
)

var producer *kafka.Producer

func InitKafkaProducer(topic string) {
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

func ProduceMessage(topic string, message interface{}) {
	if producer == nil {
		log.Println("Kafka producer is not initialized.")
		return
	}

	serializer := buildSerializer(topic)
	deliveryChan := make(chan kafka.Event)

	serializedMessage, err := serializer.Serialize(topic, &message)

	if err != nil {
		fmt.Printf("Failed to serialize payload: %s\n", err)
	}

	err = producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          serializedMessage,
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

func buildSerializer(topic string) *avro.GenericSerializer {
	client, err := schemaregistry.NewClient(schemaregistry.NewConfig(registryURL))

	if err != nil {
		log.Fatalf("Failed to create schema registry client: %s\n", err)
	}

	ser, err := avro.NewGenericSerializer(client, serde.ValueSerde, avro.NewSerializerConfig())

	if err != nil {
		log.Fatalf("Failed to create serializer: %s\n", err)
	}

	return ser
}
