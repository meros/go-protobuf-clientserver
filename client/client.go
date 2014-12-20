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

	packetWriter := protocoder.NewPacketWriter(conn)
	packetReader := protocoder.NewPacketReader(conn)

	req := &protocol.Req{
		CalcSum: &protocol.Req_CalcSum{
			A: proto.Int32(42),
			B: proto.Int32(42)}}

	err := packetWriter.Write(req)
	if err != nil {
		fmt.Println("Failed to send packet! Closing socket...")
		return
	}

	fmt.Println("A: ", req.GetCalcSum().GetA())
	fmt.Println("B: ", req.GetCalcSum().GetB())

	resp := &protocol.Resp{}
	err = packetReader.Read(resp)
	if err != nil {
		fmt.Println("Failed to read response")
		return
	}

	fmt.Println("Sum: ", resp.GetCalcSum().GetSum())
}
