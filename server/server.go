package main

import (
	"../protocol"
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/golang/protobuf/proto"
	"net"
)

func readPacket(conn net.Conn) error {
	// First read delimeter
	lengthDelimBuf := make([]byte, 4)
	var delimBytesRead uint32 = 0
	for delimBytesRead < 4 {
		n, err := conn.Read(lengthDelimBuf[delimBytesRead:])

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
		n, err := conn.Read(packetData[packetBytesRead:])

		if err != nil {
			return errors.New("Error while reading tcp connection!")
		}

		packetBytesRead += uint32(n)
	}

	packet := &protocol.Packet{}
	ec := proto.Unmarshal(packetData, packet)

	if ec != nil {
		return errors.New("Error while parsing protocol buffer")
	}

	fmt.Println(packet.GetTestString())

	return nil
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	for {
		ec := readPacket(conn)
		if ec != nil {
			fmt.Println("Closing socket due to error")
			return
		}
	}
}

func main() {
	fmt.Println("Starting server, listening on port 8080 for 1 connection")

	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Failed to create tcp listener!")

		return
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Failed to accept new tcp client!")

			return
		}

		go handleConnection(conn)
	}
}
