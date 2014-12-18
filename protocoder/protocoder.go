package protocoder

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/golang/protobuf/proto"
	"io"
)

func WritePacket(
	writer io.Writer,
	packet proto.Message) error {

	err := binary.Write(writer, binary.BigEndian, uint32(proto.Size(packet)))
	if err != nil {
		return err
	}

	packetData, err := proto.Marshal(packet)
	if err != nil {
		return err
	}

	// TODO: early write termination?
	_, err = writer.Write(packetData)
	if err != nil {
		return err
	}

	return nil
}

func ReadPacket(
	reader io.Reader,
	packet proto.Message) error {
	// First read delimeter
	lengthDelimBuf := make([]byte, 4)
	var delimBytesRead uint32 = 0
	for delimBytesRead < 4 {
		n, err := reader.Read(lengthDelimBuf[delimBytesRead:])

		if err != nil {
			return errors.New("Error while reading tcp connection!")
		}

		delimBytesRead += uint32(n)
	}

	var lengthDelim uint32
	err := binary.Read(
		bytes.NewReader(lengthDelimBuf),
		binary.BigEndian,
		&lengthDelim)

	if err != nil {
		return errors.New("Failed to read uint32 for length delimiter")
	}

	fmt.Println("Got valid delimiter: ", lengthDelim)

	// Now read packet data
	packetData := make([]byte, lengthDelim)
	var packetBytesRead uint32 = 0

	for packetBytesRead < lengthDelim {
		n, err := reader.Read(packetData[packetBytesRead:])

		if err != nil {
			return errors.New("Error while reading tcp connection!")
		}

		packetBytesRead += uint32(n)
	}

	ec := proto.Unmarshal(packetData, packet)

	if ec != nil {
		return errors.New("Error while parsing protocol buffer")
	}

	return nil
}
