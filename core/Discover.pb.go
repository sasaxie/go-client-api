// Code generated by protoc-gen-go. DO NOT EDIT.
// source: core/Discover.proto

package core // import "github.com/stanche/go-client-api/core"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Endpoint struct {
	Address              []byte   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Port                 int32    `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	NodeId               []byte   `protobuf:"bytes,3,opt,name=nodeId,proto3" json:"nodeId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Endpoint) Reset()         { *m = Endpoint{} }
func (m *Endpoint) String() string { return proto.CompactTextString(m) }
func (*Endpoint) ProtoMessage()    {}
func (*Endpoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_Discover_5033a4cb958eef49, []int{0}
}
func (m *Endpoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Endpoint.Unmarshal(m, b)
}
func (m *Endpoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Endpoint.Marshal(b, m, deterministic)
}
func (dst *Endpoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Endpoint.Merge(dst, src)
}
func (m *Endpoint) XXX_Size() int {
	return xxx_messageInfo_Endpoint.Size(m)
}
func (m *Endpoint) XXX_DiscardUnknown() {
	xxx_messageInfo_Endpoint.DiscardUnknown(m)
}

var xxx_messageInfo_Endpoint proto.InternalMessageInfo

func (m *Endpoint) GetAddress() []byte {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *Endpoint) GetPort() int32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *Endpoint) GetNodeId() []byte {
	if m != nil {
		return m.NodeId
	}
	return nil
}

type PingMessage struct {
	From                 *Endpoint `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To                   *Endpoint `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	Version              int32     `protobuf:"varint,3,opt,name=version,proto3" json:"version,omitempty"`
	Timestamp            int64     `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *PingMessage) Reset()         { *m = PingMessage{} }
func (m *PingMessage) String() string { return proto.CompactTextString(m) }
func (*PingMessage) ProtoMessage()    {}
func (*PingMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_Discover_5033a4cb958eef49, []int{1}
}
func (m *PingMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PingMessage.Unmarshal(m, b)
}
func (m *PingMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PingMessage.Marshal(b, m, deterministic)
}
func (dst *PingMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PingMessage.Merge(dst, src)
}
func (m *PingMessage) XXX_Size() int {
	return xxx_messageInfo_PingMessage.Size(m)
}
func (m *PingMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_PingMessage.DiscardUnknown(m)
}

var xxx_messageInfo_PingMessage proto.InternalMessageInfo

func (m *PingMessage) GetFrom() *Endpoint {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *PingMessage) GetTo() *Endpoint {
	if m != nil {
		return m.To
	}
	return nil
}

func (m *PingMessage) GetVersion() int32 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *PingMessage) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type PongMessage struct {
	From                 *Endpoint `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	Echo                 int32     `protobuf:"varint,2,opt,name=echo,proto3" json:"echo,omitempty"`
	Timestamp            int64     `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *PongMessage) Reset()         { *m = PongMessage{} }
func (m *PongMessage) String() string { return proto.CompactTextString(m) }
func (*PongMessage) ProtoMessage()    {}
func (*PongMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_Discover_5033a4cb958eef49, []int{2}
}
func (m *PongMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PongMessage.Unmarshal(m, b)
}
func (m *PongMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PongMessage.Marshal(b, m, deterministic)
}
func (dst *PongMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PongMessage.Merge(dst, src)
}
func (m *PongMessage) XXX_Size() int {
	return xxx_messageInfo_PongMessage.Size(m)
}
func (m *PongMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_PongMessage.DiscardUnknown(m)
}

var xxx_messageInfo_PongMessage proto.InternalMessageInfo

func (m *PongMessage) GetFrom() *Endpoint {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *PongMessage) GetEcho() int32 {
	if m != nil {
		return m.Echo
	}
	return 0
}

func (m *PongMessage) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type FindNeighbours struct {
	From                 *Endpoint `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	TargetId             []byte    `protobuf:"bytes,2,opt,name=targetId,proto3" json:"targetId,omitempty"`
	Timestamp            int64     `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *FindNeighbours) Reset()         { *m = FindNeighbours{} }
func (m *FindNeighbours) String() string { return proto.CompactTextString(m) }
func (*FindNeighbours) ProtoMessage()    {}
func (*FindNeighbours) Descriptor() ([]byte, []int) {
	return fileDescriptor_Discover_5033a4cb958eef49, []int{3}
}
func (m *FindNeighbours) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindNeighbours.Unmarshal(m, b)
}
func (m *FindNeighbours) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindNeighbours.Marshal(b, m, deterministic)
}
func (dst *FindNeighbours) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindNeighbours.Merge(dst, src)
}
func (m *FindNeighbours) XXX_Size() int {
	return xxx_messageInfo_FindNeighbours.Size(m)
}
func (m *FindNeighbours) XXX_DiscardUnknown() {
	xxx_messageInfo_FindNeighbours.DiscardUnknown(m)
}

var xxx_messageInfo_FindNeighbours proto.InternalMessageInfo

func (m *FindNeighbours) GetFrom() *Endpoint {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *FindNeighbours) GetTargetId() []byte {
	if m != nil {
		return m.TargetId
	}
	return nil
}

func (m *FindNeighbours) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type Neighbours struct {
	From                 *Endpoint   `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	Neighbours           []*Endpoint `protobuf:"bytes,2,rep,name=neighbours,proto3" json:"neighbours,omitempty"`
	Timestamp            int64       `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Neighbours) Reset()         { *m = Neighbours{} }
func (m *Neighbours) String() string { return proto.CompactTextString(m) }
func (*Neighbours) ProtoMessage()    {}
func (*Neighbours) Descriptor() ([]byte, []int) {
	return fileDescriptor_Discover_5033a4cb958eef49, []int{4}
}
func (m *Neighbours) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Neighbours.Unmarshal(m, b)
}
func (m *Neighbours) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Neighbours.Marshal(b, m, deterministic)
}
func (dst *Neighbours) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Neighbours.Merge(dst, src)
}
func (m *Neighbours) XXX_Size() int {
	return xxx_messageInfo_Neighbours.Size(m)
}
func (m *Neighbours) XXX_DiscardUnknown() {
	xxx_messageInfo_Neighbours.DiscardUnknown(m)
}

var xxx_messageInfo_Neighbours proto.InternalMessageInfo

func (m *Neighbours) GetFrom() *Endpoint {
	if m != nil {
		return m.From
	}
	return nil
}

func (m *Neighbours) GetNeighbours() []*Endpoint {
	if m != nil {
		return m.Neighbours
	}
	return nil
}

func (m *Neighbours) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type BackupMessage struct {
	Flag                 bool     `protobuf:"varint,1,opt,name=flag,proto3" json:"flag,omitempty"`
	Priority             int32    `protobuf:"varint,2,opt,name=priority,proto3" json:"priority,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BackupMessage) Reset()         { *m = BackupMessage{} }
func (m *BackupMessage) String() string { return proto.CompactTextString(m) }
func (*BackupMessage) ProtoMessage()    {}
func (*BackupMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_Discover_5033a4cb958eef49, []int{5}
}
func (m *BackupMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BackupMessage.Unmarshal(m, b)
}
func (m *BackupMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BackupMessage.Marshal(b, m, deterministic)
}
func (dst *BackupMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BackupMessage.Merge(dst, src)
}
func (m *BackupMessage) XXX_Size() int {
	return xxx_messageInfo_BackupMessage.Size(m)
}
func (m *BackupMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_BackupMessage.DiscardUnknown(m)
}

var xxx_messageInfo_BackupMessage proto.InternalMessageInfo

func (m *BackupMessage) GetFlag() bool {
	if m != nil {
		return m.Flag
	}
	return false
}

func (m *BackupMessage) GetPriority() int32 {
	if m != nil {
		return m.Priority
	}
	return 0
}

func init() {
	proto.RegisterType((*Endpoint)(nil), "protocol.Endpoint")
	proto.RegisterType((*PingMessage)(nil), "protocol.PingMessage")
	proto.RegisterType((*PongMessage)(nil), "protocol.PongMessage")
	proto.RegisterType((*FindNeighbours)(nil), "protocol.FindNeighbours")
	proto.RegisterType((*Neighbours)(nil), "protocol.Neighbours")
	proto.RegisterType((*BackupMessage)(nil), "protocol.BackupMessage")
}

func init() { proto.RegisterFile("core/Discover.proto", fileDescriptor_Discover_5033a4cb958eef49) }

var fileDescriptor_Discover_5033a4cb958eef49 = []byte{
	// 363 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xcf, 0x4a, 0xeb, 0x40,
	0x14, 0xc6, 0xc9, 0x9f, 0xf6, 0xe6, 0x9e, 0xf6, 0xde, 0x0b, 0x73, 0x41, 0x82, 0xb8, 0x08, 0x01,
	0xa5, 0x9b, 0x26, 0x50, 0x1f, 0x40, 0x08, 0x2a, 0x74, 0xa1, 0x94, 0x2c, 0xdd, 0x4d, 0x93, 0xe9,
	0x74, 0xb0, 0x99, 0x13, 0x66, 0xa6, 0x45, 0x5f, 0xc0, 0xbd, 0x6f, 0x2c, 0x19, 0x3b, 0xf5, 0x0f,
	0x4a, 0xd5, 0x55, 0xce, 0x97, 0x7c, 0x27, 0xbf, 0xef, 0x1c, 0x0e, 0xfc, 0xaf, 0x50, 0xb1, 0xfc,
	0x5c, 0xe8, 0x0a, 0x37, 0x4c, 0x65, 0xad, 0x42, 0x83, 0x24, 0xb2, 0x8f, 0x0a, 0x57, 0xe9, 0x0c,
	0xa2, 0x0b, 0x59, 0xb7, 0x28, 0xa4, 0x21, 0x31, 0xfc, 0xa2, 0x75, 0xad, 0x98, 0xd6, 0xb1, 0x97,
	0x78, 0xa3, 0x61, 0xe9, 0x24, 0x21, 0x10, 0xb6, 0xa8, 0x4c, 0xec, 0x27, 0xde, 0xa8, 0x57, 0xda,
	0x9a, 0x1c, 0x40, 0x5f, 0x62, 0xcd, 0xa6, 0x75, 0x1c, 0x58, 0xf3, 0x56, 0xa5, 0x8f, 0x1e, 0x0c,
	0x66, 0x42, 0xf2, 0x2b, 0xa6, 0x35, 0xe5, 0x8c, 0x9c, 0x40, 0xb8, 0x50, 0xd8, 0xd8, 0x5f, 0x0e,
	0x26, 0x24, 0x73, 0xe8, 0xcc, 0x71, 0x4b, 0xfb, 0x9d, 0xa4, 0xe0, 0x1b, 0xb4, 0x84, 0x8f, 0x5d,
	0xbe, 0xc1, 0x2e, 0xe1, 0x86, 0x29, 0x2d, 0x50, 0x5a, 0x68, 0xaf, 0x74, 0x92, 0x1c, 0xc1, 0x6f,
	0x23, 0x1a, 0xa6, 0x0d, 0x6d, 0xda, 0x38, 0x4c, 0xbc, 0x51, 0x50, 0xbe, 0xbc, 0x48, 0x39, 0x0c,
	0x66, 0xf8, 0xfd, 0x48, 0x04, 0x42, 0x56, 0x2d, 0xd1, 0x8d, 0xdd, 0xd5, 0x6f, 0x41, 0xc1, 0x7b,
	0x90, 0x82, 0xbf, 0x97, 0x42, 0xd6, 0xd7, 0x4c, 0xf0, 0xe5, 0x1c, 0xd7, 0x4a, 0x7f, 0x99, 0x75,
	0x08, 0x91, 0xa1, 0x8a, 0x33, 0x33, 0xad, 0x2d, 0x6f, 0x58, 0xee, 0xf4, 0x1e, 0xe6, 0x83, 0x07,
	0xf0, 0x03, 0xe0, 0x04, 0x40, 0xee, 0xba, 0x62, 0x3f, 0x09, 0x3e, 0x71, 0xbf, 0x72, 0xed, 0x09,
	0x72, 0x06, 0x7f, 0x0a, 0x5a, 0xdd, 0xae, 0x5b, 0xb7, 0x67, 0x02, 0xe1, 0x62, 0x45, 0xb9, 0x8d,
	0x12, 0x95, 0xb6, 0xee, 0xe6, 0x6c, 0x95, 0x40, 0x25, 0xcc, 0xfd, 0x76, 0xaf, 0x3b, 0x5d, 0x14,
	0xf0, 0x0f, 0x15, 0xcf, 0x8c, 0x42, 0xf9, 0x1c, 0x44, 0x17, 0x91, 0xbb, 0xdc, 0x9b, 0x63, 0x2e,
	0xcc, 0x72, 0x3d, 0xcf, 0x2a, 0x6c, 0x72, 0x4d, 0x35, 0xbd, 0x13, 0x2c, 0xe7, 0x38, 0xae, 0x56,
	0x82, 0x49, 0x33, 0xa6, 0xad, 0xc8, 0xbb, 0x4b, 0x9f, 0xf7, 0x6d, 0xe3, 0xe9, 0x53, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x98, 0xdf, 0x99, 0x2a, 0xf8, 0x02, 0x00, 0x00,
}
