package protocoder

import (
	"bytes"
	"encoding/binary"
	"errors"
	"github.com/golang/protobuf/proto"
	"io"
)

type PacketReader struct {
	buffer bytes.Buffer
	reader io.Reader
}

func NewPacketReader(reader io.Reader) PacketReader {
	return PacketReader{reader: reader}
}

func (self *PacketReader) Read(
	packet proto.Message) error {

	defer self.buffer.Truncate(0)

	// First read delimeter
	self.buffer.Truncate(0)
	_, err := io.CopyN(&self.buffer, self.reader, 4)

	if err != nil {
		return errors.New("Failed to read length delimiter from reader")
	}

	var lengthDelim uint32
	err = binary.Read(
		bytes.NewReader(self.buffer.Bytes()[0:4]),
		binary.BigEndian,
		&lengthDelim)

	if err != nil {
		return errors.New("Failed to read uint32 for length delimiter")
	}

	// Now read packet data
	self.buffer.Truncate(0)
	_, err = io.CopyN(&self.buffer, self.reader, int64(lengthDelim))

	if err != nil {
		return errors.New("Failed to read packet data from reader")
	}

	err = proto.Unmarshal(self.buffer.Bytes(), packet)

	if err != nil {
		return errors.New("Error while parsing protocol buffer")
	}

	return nil
}
