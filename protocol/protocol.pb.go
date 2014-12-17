// Code generated by protoc-gen-go.
// source: protocol
// DO NOT EDIT!

/*
Package protocol is a generated protocol buffer package.

It is generated from these files:
	protocol

It has these top-level messages:
	Packet
*/
package protocol

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type Packet struct {
	TestString       *string `protobuf:"bytes,1,opt,name=testString" json:"testString,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Packet) Reset()         { *m = Packet{} }
func (m *Packet) String() string { return proto.CompactTextString(m) }
func (*Packet) ProtoMessage()    {}

func (m *Packet) GetTestString() string {
	if m != nil && m.TestString != nil {
		return *m.TestString
	}
	return ""
}

func init() {
}
