# Kafka Avro Example Project

This project demonstrates a simple Kafka setup using Docker Compose, along with a Kafka producer and consumer written in Go. The producer sends messages to a Kafka topic, and the consumer reads messages from the topic. The messages are serialized and deserialized using Avro schemas.

## Prerequisites

- Docker and Docker Compose installed on your machine
- Go installed on your machine
- `avrogen` installed on your machine (`go get github.com/actgardner/gogen-avro/cmd/avrogen`)



## Getting Started

### Step 1: Create generated User struct from Avro schema
Run the following command to generate the Go struct for the `User` schema defined in `avroschema/user.avsc`:

```bash
avrogen -pkg avro -o avroschema/user.go -tags json:snake,yaml:upper-camel avroschema/user.avsc
```

This command does the following:

- `pkg avro`: Specifies the package name for the generated Go code.
- `-o avroschema/user.go`: Specifies the output file where the generated Go code will be saved.
- `-tags json:snake,yaml:upper-camel`: Adds struct tags for JSON and YAML serialization with specified naming conventions.
After running this command, the User struct will be created in avroschema/user.go. This struct will be used for Avro serialization and deserialization in your Kafka producer and consumer.
- `avroschema/user.avsc`: Specifies the Avro schema file.

### Step 2: Start Kafka and Zookeeper using Docker Compose

1. Navigate to the project directory.
2. Run the following command to start the services:

```bash
docker-compose up -d
```

This command will start the following services:

Zookeeper
Kafka
Kafka UI (a web interface to monitor your Kafka cluster)
You can access the Kafka UI at http://localhost:8080.

### Step 3: Run the Kafka Producer
1. Navigate to the producer directory.
2. Run the following command to start the producer:

```bash
go run main.go
```

### Step 4: Run the Kafka Consumer
1. Navigate to the consumer directory.
2. go run main.go

```bash
go run main.go
```

The consumer will read messages from the Kafka topic and print them to the console.