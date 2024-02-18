package schema

import (
	"fmt"
	"os"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/google/uuid"
	"github.com/linkedin/goavro/v2"
)

type Person struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstname"`
	Email     string `json:"email"`
}

func Producer() {
	topic := "myTopic"

	// Create Avro codec
	schema, err := os.ReadFile("person.avsc")
	if err != nil {
		fmt.Printf("Failed to read svro schema file: %s\n", err)
	}

	codec, err := goavro.NewCodec(string(schema))
	if err != nil {
		fmt.Printf("Failed to create Avro codec: %s\n", err)
	}

	// 1) Create the producer as you would normally do using Confluent's Go client
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost",
		// "schema.registry.url": "http://localhost:8081",
	})

	if err != nil {
		panic(err)
	}
	defer p.Close()

	go func() {
		for event := range p.Events() {
			switch ev := event.(type) {
			case *kafka.Message:
				message := ev
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Error delivering the message '%s'\n", message.Key)
				} else {
					fmt.Printf("Message '%s' delivered successfully!\n", message.Key)
				}
			}
		}
	}()

	// 2) Fetch the latest version of the schema, or create a new one if it is the first
	// schemaRegistryClient := srclient.CreateSchemaRegistryClient("http://localhost:8081")
	// schema, err := schemaRegistryClient.GetLatestSchema(topic)

	// if err != nil {
	// 	log.Printf("ERROR GETTING LAST SCHEMA %v", err)
	// }

	// if schema == nil {
	// 	schemaBytes, _ := os.ReadFile("person.avsc")
	// 	schema, err = schemaRegistryClient.CreateSchema(topic, string(schemaBytes), srclient.Avro)
	// 	if err != nil {
	// 		panic(fmt.Sprintf("Error creating the schema %s", err))
	// 	}
	// }
	// schemaIDBytes := make([]byte, 4)
	// binary.BigEndian.PutUint32(schemaIDBytes, uint32(schema.ID()))

	// 3) Serialize the record using the schema provided by the client,
	// making sure to include the schema id as part of the record.
	// person := Person{ID: 1, FirstName: "Alexander", Email: "alexander@yahuu.com"}
	// value, _ := json.Marshal(person)
	// native, _, _ := schema.Codec().NativeFromTextual(value)
	// log.Println("GOT HERE 69")
	// valueBytes, _ := schema.Codec().BinaryFromNative(nil, native)
	// Encode message data to Avro binary
	value := map[string]interface{}{
		"id":        "1",
		"firstName": "alexander",
		"email":     "alexander@yahuu.com",
	}
	avroBinary, err := codec.BinaryFromNative(nil, value)
	if err != nil {
		fmt.Printf("Failed to encode Avro binary: %s\n", err)
	}

	// var recordValue []byte
	// recordValue = append(recordValue, byte(0))
	// recordValue = append(recordValue, schemaIDBytes...)
	// recordValue = append(recordValue, valueBytes...)

	key, _ := uuid.NewUUID()
	p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic: &topic, Partition: kafka.PartitionAny},
		Key: []byte(key.String()), Value: avroBinary}, nil)

	p.Flush(15 * 1000)

}
