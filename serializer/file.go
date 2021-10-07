package serializer

import (
	"fmt"
	"io/ioutil"

	"github.com/golang/protobuf/proto"
)

func WriteProtobufToBinaryFile(message proto.Message, filename string) error {
	data, err := proto.Marshal(message)

	if err != nil {
		return fmt.Errorf("Cannot marshall proto messasge to binary: %w", err)
	}

	err = ioutil.WriteFile(filename, data, 0644)

	if err != nil {
		return fmt.Errorf("Cannot write binary to file: %w", err)
	}

	return nil
}

func WriteProtobufToJSON(message proto.Message, filename string) error {
	data, err := ProtobufToJSON(message)
	if err != nil {
		return fmt.Errorf("Cannot marshal proto message to JSON: %w", err)
	}

	err = ioutil.WriteFile(filename, []byte(data), 0644)

	if err != nil {
		return fmt.Errorf("Cannot write JSON data to file: %w", err)
	}

	return nil
}


func ReadProtobuffFromBinaryFile(filename string, message proto.Message) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("Cannot read binary data from file: %w", err)
	}

	err = proto.Unmarshal(data, message)

	if err != nil {
		return fmt.Errorf("Cannot unmarshal binary to protobuf: %w", err)
	}

	return nil
}