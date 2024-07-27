package main

import (
	"context"
	"log"

	"github.com/hamba/avro/v2"
	"github.com/segmentio/kafka-go"

	avroschema "kafka-example/avroschema"
)

func main() {
	// make a new reader that consumes from my-topic, partition 0, at offset 42
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"localhost:29092"},
		Topic:     "my-topic",
		GroupID:   "consumer1",
		Partition: 0,
		MinBytes:  10e3, // 10KB
		MaxBytes:  10e6, // 10MB
	})

	for {
		data, err := r.ReadMessage(context.Background())
		if err != nil {
			break
		}

		schema, err := avroschema.GetSchema()
		if err != nil {
			return
		}

		out := avroschema.User{}
		err = avro.Unmarshal(*schema, data.Value, &out)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("message at offset %d: %s = %s\n", data.Offset, string(data.Key), out)
	}

	if err := r.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}
}
