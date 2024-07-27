package avro

import (
	"io"
	"os"

	"github.com/hamba/avro/v2"
)

func (u *User) GetAvroBytes() ([]byte, error) {
	file, err := os.Open("avroschema/user.avsc")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	schema, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	schemaData, err := avro.Parse(string(schema))
	if err != nil {
		return nil, err
	}

	data, err := avro.Marshal(schemaData, u)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func GetSchema() (*avro.Schema, error) {
	file, err := os.Open("avroschema/user.avsc")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	schema, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	schemaData, err := avro.Parse(string(schema))
	if err != nil {
		return nil, err
	}

	return &schemaData, nil
}
