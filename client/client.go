package main

import (
	"../protocol"
	"encoding/binary"
	"fmt"
	"github.com/golang/protobuf/proto"
	"net"
)

func main() {
	conn, ec := net.Dial("tcp", "localhost:8080")

	if ec != nil {
		fmt.Println("Failed to connect to localhost:8080")
		return
	}

	packet := &protocol.Packet{TestString: proto.String("Hello World!")}

	binary.Write(conn, binary.BigEndian, uint32(proto.Size(packet)))
	packetData, ec := proto.Marshal(packet)

	if ec != nil {
		fmt.Println("Failed to marshel packet")
		return
	}

	conn.Write(packetData)
}
