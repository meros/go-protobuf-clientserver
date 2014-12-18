package main

import (
	"fmt"
	"github.com/meros/go-protobuf-clientserver/connchan"
	"github.com/meros/go-protobuf-clientserver/protocoder"
	"github.com/meros/go-protobuf-clientserver/protocol"
	"net"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	for {
		packet := &protocol.Packet{}
		err := protocoder.ReadPacket(conn, packet)

		if err != nil {
			fmt.Println("No more data on socket!")
			return
		}

		fmt.Println("Got packet with TestString: ", packet.GetTestString())
	}
}

func main() {
	fmt.Println("Starting server, listening on port 8080 for 1 connection")

	connchan, err := connchan.Create()

	if err != nil {
		fmt.Println("Failed to open connection channel")
		return
	}

	for conn := range connchan {
		go handleConnection(conn)
	}
}
