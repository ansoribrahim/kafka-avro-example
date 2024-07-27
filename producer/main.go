package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/segmentio/kafka-go"

	avroschema "kafka-example/avroschema"
)

func main() {
	// Define Kafka topic and partition
	topic := "my-topic"
	partition := 0

	// Function to create a new Kafka connection
	createKafkaConnection := func() *kafka.Conn {
		conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:29092", topic, partition)
		if err != nil {
			log.Fatalf("failed to dial leader: %v", err)
		}
		return conn
	}

	// Create initial Kafka connection
	conn := createKafkaConnection()
	defer func() {
		if err := conn.Close(); err != nil {
			log.Fatalf("failed to close writer: %v", err)
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter messages to send to Kafka (type 'exit' to quit):")

	for {
		if !scanner.Scan() {
			break
		}
		text := scanner.Text()

		if text == "exit" {
			break
		}

		conn.SetWriteDeadline(time.Now().Add(10 * time.Second))

		req := avroschema.User{
			Name: text,
			Age:  30,
		}

		bytes, err := req.GetAvroBytes()
		if err != nil {
			return
		}

		_, err = conn.WriteMessages(
			kafka.Message{Key: []byte("user-message"), Value: bytes},
		)
		if err != nil {
			log.Printf("failed to write message: %v", err)
			log.Println("reconnecting to Kafka...")
			conn = createKafkaConnection()
		} else {
			fmt.Println("Message sent:", text)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading input: %v", err)
	}
}
