package protocoder

import (
	"encoding/binary"
	"github.com/golang/protobuf/proto"
	"io"
)

type PacketWriter struct {
	writer io.Writer
}

func NewPacketWriter(writer io.Writer) PacketWriter {
	return PacketWriter{writer}
}

func (self *PacketWriter) Write(
	packet proto.Message) error {

	err := binary.Write(self.writer, binary.BigEndian, uint32(proto.Size(packet)))
	if err != nil {
		return err
	}

	packetData, err := proto.Marshal(packet)
	if err != nil {
		return err
	}

	_, err = self.writer.Write(packetData)
	if err != nil {
		return err
	}

	return nil
}
