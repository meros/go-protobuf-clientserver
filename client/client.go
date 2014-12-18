package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/meros/go-protobuf-clientserver/protocoder"
	"github.com/meros/go-protobuf-clientserver/protocol"
	"net"
)

func main() {
	conn, ec := net.Dial("tcp", "localhost:8080")

	if ec != nil {
		fmt.Println("Failed to connect to localhost:8080")
		return
	}

	packet := &protocol.Packet{TestString: proto.String("Hello World!")}
	err := protocoder.WritePacket(conn, packet)
	if err != nil {
		fmt.Println("Failed to send packet! Closing socket...")
		return
	}

	packet = &protocol.Packet{TestString: proto.String("And yet another string...")}
	err = protocoder.WritePacket(conn, packet)
	if err != nil {
		fmt.Println("Failed to send packet! Closing socket...")
		return
	}
}
