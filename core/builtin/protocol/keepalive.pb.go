// Code generated by protoc-gen-go.
// source: protocol/keepalive.proto
// DO NOT EDIT!

package protocol

import proto "code.google.com/p/goprotobuf/proto"
import json "encoding/json"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type SSPacketKeepAlive struct {
	PacketId         *CoreBuiltinPacketID `protobuf:"varint,1,req,enum=protocol.CoreBuiltinPacketID,def=-1005" json:"PacketId,omitempty"`
	Flag             *int32               `protobuf:"varint,2,req" json:"Flag,omitempty"`
	XXX_unrecognized []byte               `json:"-"`
}

func (m *SSPacketKeepAlive) Reset()         { *m = SSPacketKeepAlive{} }
func (m *SSPacketKeepAlive) String() string { return proto.CompactTextString(m) }
func (*SSPacketKeepAlive) ProtoMessage()    {}

const Default_SSPacketKeepAlive_PacketId CoreBuiltinPacketID = CoreBuiltinPacketID_PACKET_SS_KEEPALIVE

func (m *SSPacketKeepAlive) GetPacketId() CoreBuiltinPacketID {
	if m != nil && m.PacketId != nil {
		return *m.PacketId
	}
	return Default_SSPacketKeepAlive_PacketId
}

func (m *SSPacketKeepAlive) GetFlag() int32 {
	if m != nil && m.Flag != nil {
		return *m.Flag
	}
	return 0
}

func init() {
}