package main

import (
	"fmt"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "Usage: %s <broker> <topic>\n",
			os.Args[0])
		os.Exit(1)
	}

	fmt.Println(kafka.LibraryVersion())

	broker := os.Args[1]
	topic := os.Args[2]

	p, err := kafka.NewProducer(
		&kafka.ConfigMap{"bootstrap.servers": broker,
			"plugin.library.paths": "monitoring-interceptor", "debug": "all"})
	if err != nil {
		panic(err)
	}

	err = p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte("hi"),
		Headers:        []kafka.Header{{Key: "myTestHeader", Value: []byte("header values are binary")}},
	}, nil)

	p.Close()
}
