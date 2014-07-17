// Code generated by protoc-gen-go.
// source: protocol/packetid.proto
// DO NOT EDIT!

package protocol

import proto "code.google.com/p/goprotobuf/proto"
import json "encoding/json"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type SrvlibPacketID int32

const (
	SrvlibPacketID_PACKET_SS_REGISTE         SrvlibPacketID = -2000
	SrvlibPacketID_PACKET_SS_MULTICAST       SrvlibPacketID = -2001
	SrvlibPacketID_PACKET_SS_BROADCAST       SrvlibPacketID = -2002
	SrvlibPacketID_PACKET_SS_TRANSIT         SrvlibPacketID = -2003
	SrvlibPacketID_PACKET_SS_REDIRECT        SrvlibPacketID = -2004
	SrvlibPacketID_PACKET_SS_SERVICE_REGISTE SrvlibPacketID = -2005
	SrvlibPacketID_PACKET_SS_SERVICE_INFO    SrvlibPacketID = -2006
	SrvlibPacketID_PACKET_SS_SERVICE_SHUT    SrvlibPacketID = -2007
)

var SrvlibPacketID_name = map[int32]string{
	-2000: "PACKET_SS_REGISTE",
	-2001: "PACKET_SS_MULTICAST",
	-2002: "PACKET_SS_BROADCAST",
	-2003: "PACKET_SS_TRANSIT",
	-2004: "PACKET_SS_REDIRECT",
	-2005: "PACKET_SS_SERVICE_REGISTE",
	-2006: "PACKET_SS_SERVICE_INFO",
	-2007: "PACKET_SS_SERVICE_SHUT",
}
var SrvlibPacketID_value = map[string]int32{
	"PACKET_SS_REGISTE":         -2000,
	"PACKET_SS_MULTICAST":       -2001,
	"PACKET_SS_BROADCAST":       -2002,
	"PACKET_SS_TRANSIT":         -2003,
	"PACKET_SS_REDIRECT":        -2004,
	"PACKET_SS_SERVICE_REGISTE": -2005,
	"PACKET_SS_SERVICE_INFO":    -2006,
	"PACKET_SS_SERVICE_SHUT":    -2007,
}

func (x SrvlibPacketID) Enum() *SrvlibPacketID {
	p := new(SrvlibPacketID)
	*p = x
	return p
}
func (x SrvlibPacketID) String() string {
	return proto.EnumName(SrvlibPacketID_name, int32(x))
}
func (x SrvlibPacketID) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *SrvlibPacketID) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(SrvlibPacketID_value, data, "SrvlibPacketID")
	if err != nil {
		return err
	}
	*x = SrvlibPacketID(value)
	return nil
}

func init() {
	proto.RegisterEnum("protocol.SrvlibPacketID", SrvlibPacketID_name, SrvlibPacketID_value)
}
