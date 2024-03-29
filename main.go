package main

import (
	"github.com/muling3/confluent-go/schema"
)

func main() {
	// // produce the messages
	// p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost"})
	// if err != nil {
	// 	panic(err)
	// }

	// defer p.Close()

	// // Delivery report handler for produced messages
	// go func() {
	// 	for e := range p.Events() {
	// 		switch ev := e.(type) {
	// 		case *kafka.Message:
	// 			if ev.TopicPartition.Error != nil {
	// 				fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
	// 			} else {
	// 				fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
	// 			}
	// 		}
	// 	}
	// }()

	// // Produce messages to topic (asynchronously)
	// topic := "myTopic"
	// for _, word := range []string{"Welcome", "to", "the", "Confluent", "Kafka", "Golang", "client"} {
	// 	p.Produce(&kafka.Message{
	// 		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
	// 		Value:          []byte(word),
	// 	}, nil)
	// }

	// // Wait for message deliveries before shutting down
	// p.Flush(15 * 1000)

	// // consuming the messages
	// go func() {
	// 	c, err := kafka.NewConsumer(&kafka.ConfigMap{
	// 		"bootstrap.servers": "localhost",
	// 		"group.id":          "myGroup",
	// 		"auto.offset.reset": "earliest",
	// 	})

	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	c.SubscribeTopics([]string{"myTopic"}, nil)

	// 	// A signal handler or similar could be used to set this to false to break the loop.
	// 	run := true

	// 	for run {
	// 		msg, err := c.ReadMessage(time.Second)
	// 		if err == nil {
	// 			fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
	// 		} else if !err.(kafka.Error).IsTimeout() {
	// 			// The client will automatically try to recover from all errors.
	// 			// Timeout is not considered an error because it is raised by
	// 			// ReadMessage in absence of messages.
	// 			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
	// 		}
	// 	}

	// 	c.Close()
	// }()

	// select {}

	// produce messages using a schema
	go func() { schema.Producer() }()
	go func() { schema.Consumer() }()
	select {}
}
