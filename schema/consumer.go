package schema

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func Consumer() {

	// 1) Create the consumer as you would
	// normally do using Confluent's Go client
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	defer c.Close()

	c.SubscribeTopics([]string{"myTopic"}, nil)

	// 2) Create a instance of the client to retrieve the schemas for each message
	// schemaRegistryClient := srclient.CreateSchemaRegistryClient("http://localhost:8081")

	for {
		msg, err := c.ReadMessage(-1)
		if err == nil {
			// 3) Recover the schema id from the message and use the
			// client to retrieve the schema from Schema Registry.
			// Then use it to deserialize the record accordingly.
			// schemaID := binary.BigEndian.Uint32(msg.Value[1:5])
			// schema, err := schemaRegistryClient.GetSchema(int(schemaID))
			// if err != nil {
			// 	panic(fmt.Sprintf("Error getting the schema with id '%d' %s", schemaID, err))
			// }
			// native, _, _ := schema.Codec().NativeFromBinary(msg.Value[5:])
			// value, _ := schema.Codec().TextualFromNative(nil, native)
			fmt.Printf("Here is the message %s\n", string(msg.Value))
		} else {
			fmt.Printf("Error consuming the message: %v (%v)\n", err, msg)
		}
	}

}
