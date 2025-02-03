package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"github.com/IBM/sarama"
)

func producer(broker, topic string) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer([]string{broker}, config)
	if err != nil {
		log.Fatalf("Failed to create producer: %v", err)
	}
	defer producer.Close()

	for i := 0; i < 1000; i++ {
		msg := &sarama.ProducerMessage{
			Topic: topic,
			Value: sarama.StringEncoder(fmt.Sprintf("Message %d", i)),
		}
		_, _, err := producer.SendMessage(msg)
		if err != nil {
			log.Printf("Failed to send message: %v", err)
		}
	}
}

func consumer(broker, topic string, numPartitions int) {
	consumer, err := sarama.NewConsumer([]string{broker}, nil)
	if err != nil {
		log.Fatalf("Cannot create consumer: %v", err)
	}
	defer consumer.Close()

	partitionConsumers := make([]sarama.PartitionConsumer, numPartitions)

	for partition := 0; partition < numPartitions; partition++ {
		partitionConsumer, err := consumer.ConsumePartition(topic, int32(partition), sarama.OffsetNewest)
		if err != nil {
			log.Fatalf("Cannot create partition consumer: %v", err)
		}
		defer partitionConsumer.Close()
		partitionConsumers[partition] = partitionConsumer

		go func(partition int) {
			for count := 1; ; count++ {
				msg := <-partitionConsumers[partition].Messages()
				if count%100 == 0 {
					log.Printf("%s: received %d messages from partition %d, last (%s)",
						topic, count, partition, string(msg.Value))
				}
			}
		}(partition)
	}
	select {}
}

func main() {
	topic := os.Getenv("TOPIC")
	if topic == "" {
		log.Fatalf("Unknown topic")
	}

	role := os.Getenv("ROLE")
	broker := os.Getenv("KAFKA_BROKER")
	partitions, err := strconv.Atoi(os.Getenv("PARTITIONS"))
	if err != nil {
		log.Fatalf("Invalid partition count")
	}

	if role == "producer" {
		producer(broker, topic)
	} else if role == "consumer" {
		consumer(broker, topic, partitions)
	} else {
		log.Fatalf("Unknown role %s", role)
	}
}
