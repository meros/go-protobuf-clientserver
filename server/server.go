package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/meros/go-protobuf-clientserver/connchan"
	"github.com/meros/go-protobuf-clientserver/protocoder"
	"github.com/meros/go-protobuf-clientserver/protocol"
	"net"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	packetReader := protocoder.NewPacketReader(conn)
	packetWriter := protocoder.NewPacketWriter(conn)

	req := new(protocol.Req)
	resp := new(protocol.Resp)

	for {
		err := packetReader.Read(req)

		if err != nil {
			fmt.Println("Failed to read packet!")
			return
		}

		resp.CalcSum = &protocol.Resp_CalcSum{
			Sum: proto.Int32(req.GetCalcSum().GetA() + req.GetCalcSum().GetB())}

		packetWriter.Write(resp)
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
