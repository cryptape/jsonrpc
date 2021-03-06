// Code generated by protoc-gen-go. DO NOT EDIT.
// source: communication.proto

package libproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type MsgType int32

const (
	MsgType_REQUEST            MsgType = 0
	MsgType_HEADER             MsgType = 1
	MsgType_BLOCK              MsgType = 2
	MsgType_STATUS             MsgType = 3
	MsgType_MSG                MsgType = 4
	MsgType_RESPONSE           MsgType = 5
	MsgType_VERIFY_TX_REQ      MsgType = 6
	MsgType_VERIFY_TX_RESP     MsgType = 7
	MsgType_VERIFY_BLK_REQ     MsgType = 8
	MsgType_VERIFY_BLK_RESP    MsgType = 9
	MsgType_BLOCK_TXHASHES     MsgType = 10
	MsgType_BLOCK_TXHASHES_REQ MsgType = 11
	MsgType_BLOCK_WITH_PROOF   MsgType = 12
	MsgType_BLOCK_TXS          MsgType = 13
	MsgType_RICH_STATUS        MsgType = 14
)

var MsgType_name = map[int32]string{
	0:  "REQUEST",
	1:  "HEADER",
	2:  "BLOCK",
	3:  "STATUS",
	4:  "MSG",
	5:  "RESPONSE",
	6:  "VERIFY_TX_REQ",
	7:  "VERIFY_TX_RESP",
	8:  "VERIFY_BLK_REQ",
	9:  "VERIFY_BLK_RESP",
	10: "BLOCK_TXHASHES",
	11: "BLOCK_TXHASHES_REQ",
	12: "BLOCK_WITH_PROOF",
	13: "BLOCK_TXS",
	14: "RICH_STATUS",
}
var MsgType_value = map[string]int32{
	"REQUEST":            0,
	"HEADER":             1,
	"BLOCK":              2,
	"STATUS":             3,
	"MSG":                4,
	"RESPONSE":           5,
	"VERIFY_TX_REQ":      6,
	"VERIFY_TX_RESP":     7,
	"VERIFY_BLK_REQ":     8,
	"VERIFY_BLK_RESP":    9,
	"BLOCK_TXHASHES":     10,
	"BLOCK_TXHASHES_REQ": 11,
	"BLOCK_WITH_PROOF":   12,
	"BLOCK_TXS":          13,
	"RICH_STATUS":        14,
}

func (x MsgType) String() string {
	return proto.EnumName(MsgType_name, int32(x))
}
func (MsgType) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

type OperateType int32

const (
	OperateType_BROADCAST OperateType = 0
	OperateType_SINGLE    OperateType = 1
	OperateType_SUBTRACT  OperateType = 2
)

var OperateType_name = map[int32]string{
	0: "BROADCAST",
	1: "SINGLE",
	2: "SUBTRACT",
}
var OperateType_value = map[string]int32{
	"BROADCAST": 0,
	"SINGLE":    1,
	"SUBTRACT":  2,
}

func (x OperateType) String() string {
	return proto.EnumName(OperateType_name, int32(x))
}
func (OperateType) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

type Message struct {
	CmdId   uint32      `protobuf:"varint,1,opt,name=cmd_id,json=cmdId" json:"cmd_id,omitempty"`
	Type    MsgType     `protobuf:"varint,2,opt,name=type,enum=MsgType" json:"type,omitempty"`
	Origin  uint32      `protobuf:"varint,3,opt,name=origin" json:"origin,omitempty"`
	Operate OperateType `protobuf:"varint,4,opt,name=operate,enum=OperateType" json:"operate,omitempty"`
	Content []byte      `protobuf:"bytes,5,opt,name=content,proto3" json:"content,omitempty"`
}

func (m *Message) Reset()                    { *m = Message{} }
func (m *Message) String() string            { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()               {}
func (*Message) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *Message) GetCmdId() uint32 {
	if m != nil {
		return m.CmdId
	}
	return 0
}

func (m *Message) GetType() MsgType {
	if m != nil {
		return m.Type
	}
	return MsgType_REQUEST
}

func (m *Message) GetOrigin() uint32 {
	if m != nil {
		return m.Origin
	}
	return 0
}

func (m *Message) GetOperate() OperateType {
	if m != nil {
		return m.Operate
	}
	return OperateType_BROADCAST
}

func (m *Message) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

func init() {
	proto.RegisterType((*Message)(nil), "Message")
	proto.RegisterEnum("MsgType", MsgType_name, MsgType_value)
	proto.RegisterEnum("OperateType", OperateType_name, OperateType_value)
}

func init() { proto.RegisterFile("communication.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 365 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0x4f, 0xaf, 0x9a, 0x40,
	0x14, 0xc5, 0x0b, 0x0a, 0xe8, 0x05, 0x7c, 0xd3, 0xfb, 0xda, 0x17, 0x16, 0x5d, 0x98, 0x2e, 0x1a,
	0xe3, 0xc2, 0x45, 0x9b, 0x74, 0x8f, 0x38, 0x0a, 0xf1, 0x0f, 0x38, 0x33, 0xb6, 0x76, 0x45, 0x2c,
	0x10, 0xc3, 0x82, 0x3f, 0x51, 0xba, 0x70, 0xdb, 0x4f, 0xd1, 0x8f, 0xdb, 0x38, 0x4a, 0xa2, 0xcb,
	0xf3, 0x9b, 0x73, 0xe6, 0xe6, 0x9e, 0x0b, 0xaf, 0x49, 0x55, 0x14, 0x7f, 0xca, 0x3c, 0x39, 0x34,
	0x79, 0x55, 0x4e, 0xea, 0x53, 0xd5, 0x54, 0x9f, 0xff, 0x29, 0x60, 0xac, 0xb3, 0xf3, 0xf9, 0x70,
	0xcc, 0xf0, 0x23, 0xe8, 0x49, 0x91, 0xc6, 0x79, 0xea, 0x28, 0x43, 0x65, 0x64, 0x33, 0x2d, 0x29,
	0xd2, 0x20, 0xc5, 0x4f, 0xd0, 0x6d, 0x2e, 0x75, 0xe6, 0xa8, 0x43, 0x65, 0x34, 0xf8, 0xda, 0x9b,
	0xac, 0xcf, 0x47, 0x71, 0xa9, 0x33, 0x26, 0x29, 0xbe, 0x81, 0x5e, 0x9d, 0xf2, 0x63, 0x5e, 0x3a,
	0x1d, 0x19, 0xba, 0x2b, 0xfc, 0x02, 0x46, 0x55, 0x67, 0xa7, 0x43, 0x93, 0x39, 0x5d, 0x19, 0xb4,
	0x26, 0xe1, 0x4d, 0xcb, 0x70, 0xfb, 0x88, 0x0e, 0x18, 0x49, 0x55, 0x36, 0x59, 0xd9, 0x38, 0xda,
	0x50, 0x19, 0x59, 0xac, 0x95, 0xe3, 0xbf, 0x2a, 0x18, 0xf7, 0x59, 0x68, 0x82, 0xc1, 0xe8, 0x76,
	0x47, 0xb9, 0x20, 0xef, 0x10, 0x40, 0xf7, 0xa9, 0x3b, 0xa3, 0x8c, 0x28, 0xd8, 0x07, 0x6d, 0xba,
	0x0a, 0xbd, 0x25, 0x51, 0xaf, 0x98, 0x0b, 0x57, 0xec, 0x38, 0xe9, 0xa0, 0x01, 0x9d, 0x35, 0x5f,
	0x90, 0x2e, 0x5a, 0xd0, 0x63, 0x94, 0x47, 0xe1, 0x86, 0x53, 0xa2, 0xe1, 0x7b, 0xb0, 0x7f, 0x50,
	0x16, 0xcc, 0x7f, 0xc5, 0x62, 0x1f, 0x33, 0xba, 0x25, 0x3a, 0x22, 0x0c, 0x1e, 0x11, 0x8f, 0x88,
	0xf1, 0xc0, 0xa6, 0xab, 0xa5, 0xf4, 0xf5, 0xf0, 0x15, 0x5e, 0x9e, 0x18, 0x8f, 0x48, 0xff, 0x6a,
	0x94, 0xd3, 0x63, 0xb1, 0xf7, 0x5d, 0xee, 0x53, 0x4e, 0x00, 0xdf, 0x00, 0x9f, 0x99, 0xfc, 0xc0,
	0xc4, 0x0f, 0x40, 0x6e, 0xfc, 0x67, 0x20, 0xfc, 0x38, 0x62, 0x61, 0x38, 0x27, 0x16, 0xda, 0xd0,
	0x6f, 0xdd, 0x9c, 0xd8, 0xf8, 0x02, 0x26, 0x0b, 0x3c, 0x3f, 0xbe, 0x2f, 0x32, 0x18, 0x7f, 0x07,
	0xf3, 0xa1, 0x36, 0x69, 0x67, 0xa1, 0x3b, 0xf3, 0xdc, 0xb6, 0x09, 0x1e, 0x6c, 0x16, 0x2b, 0x4a,
	0x94, 0xeb, 0xa6, 0x7c, 0x37, 0x15, 0xcc, 0xf5, 0x04, 0x51, 0x7f, 0xeb, 0xf2, 0xbc, 0xdf, 0xfe,
	0x07, 0x00, 0x00, 0xff, 0xff, 0xf7, 0xb4, 0x6f, 0x7d, 0xf5, 0x01, 0x00, 0x00,
}
