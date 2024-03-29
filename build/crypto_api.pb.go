// Code generated by protoc-gen-go. DO NOT EDIT.
// source: crypto_api.proto

package build

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

type Ticker struct {
	Base                 string   `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	Target               string   `protobuf:"bytes,2,opt,name=target,proto3" json:"target,omitempty"`
	Price                float32  `protobuf:"fixed32,3,opt,name=price,proto3" json:"price,omitempty"`
	Volume               float32  `protobuf:"fixed32,4,opt,name=volume,proto3" json:"volume,omitempty"`
	Change               float32  `protobuf:"fixed32,5,opt,name=change,proto3" json:"change,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Ticker) Reset()         { *m = Ticker{} }
func (m *Ticker) String() string { return proto.CompactTextString(m) }
func (*Ticker) ProtoMessage()    {}
func (*Ticker) Descriptor() ([]byte, []int) {
	return fileDescriptor_crypto_api_ccf304f17f8188e3, []int{0}
}
func (m *Ticker) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Ticker.Unmarshal(m, b)
}
func (m *Ticker) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Ticker.Marshal(b, m, deterministic)
}
func (dst *Ticker) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Ticker.Merge(dst, src)
}
func (m *Ticker) XXX_Size() int {
	return xxx_messageInfo_Ticker.Size(m)
}
func (m *Ticker) XXX_DiscardUnknown() {
	xxx_messageInfo_Ticker.DiscardUnknown(m)
}

var xxx_messageInfo_Ticker proto.InternalMessageInfo

func (m *Ticker) GetBase() string {
	if m != nil {
		return m.Base
	}
	return ""
}

func (m *Ticker) GetTarget() string {
	if m != nil {
		return m.Target
	}
	return ""
}

func (m *Ticker) GetPrice() float32 {
	if m != nil {
		return m.Price
	}
	return 0
}

func (m *Ticker) GetVolume() float32 {
	if m != nil {
		return m.Volume
	}
	return 0
}

func (m *Ticker) GetChange() float32 {
	if m != nil {
		return m.Change
	}
	return 0
}

type CryptonatorResponse struct {
	Ticker               *Ticker  `protobuf:"bytes,1,opt,name=ticker,proto3" json:"ticker,omitempty"`
	Timestamp            int64    `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Success              bool     `protobuf:"varint,3,opt,name=success,proto3" json:"success,omitempty"`
	Error                string   `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CryptonatorResponse) Reset()         { *m = CryptonatorResponse{} }
func (m *CryptonatorResponse) String() string { return proto.CompactTextString(m) }
func (*CryptonatorResponse) ProtoMessage()    {}
func (*CryptonatorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_crypto_api_ccf304f17f8188e3, []int{1}
}
func (m *CryptonatorResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CryptonatorResponse.Unmarshal(m, b)
}
func (m *CryptonatorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CryptonatorResponse.Marshal(b, m, deterministic)
}
func (dst *CryptonatorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CryptonatorResponse.Merge(dst, src)
}
func (m *CryptonatorResponse) XXX_Size() int {
	return xxx_messageInfo_CryptonatorResponse.Size(m)
}
func (m *CryptonatorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CryptonatorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CryptonatorResponse proto.InternalMessageInfo

func (m *CryptonatorResponse) GetTicker() *Ticker {
	if m != nil {
		return m.Ticker
	}
	return nil
}

func (m *CryptonatorResponse) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *CryptonatorResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *CryptonatorResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func init() {
	proto.RegisterType((*Ticker)(nil), "build.Ticker")
	proto.RegisterType((*CryptonatorResponse)(nil), "build.CryptonatorResponse")
}

func init() { proto.RegisterFile("crypto_api.proto", fileDescriptor_crypto_api_ccf304f17f8188e3) }

var fileDescriptor_crypto_api_ccf304f17f8188e3 = []byte{
	// 219 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x44, 0x90, 0xcd, 0x4a, 0xc4, 0x30,
	0x14, 0x46, 0xe9, 0xaf, 0xf6, 0x8a, 0x20, 0x51, 0x24, 0x0b, 0x17, 0x52, 0x10, 0x5c, 0x75, 0xa1,
	0x8f, 0xe0, 0x1b, 0x04, 0xf7, 0x92, 0xc6, 0x4b, 0x0d, 0xb6, 0x4d, 0xb8, 0x49, 0x05, 0xe7, 0x11,
	0xe6, 0xa9, 0xa7, 0xbd, 0xe9, 0x30, 0xbb, 0x9c, 0x93, 0xc0, 0x77, 0x08, 0xdc, 0x19, 0xfa, 0xf7,
	0xd1, 0x7d, 0x69, 0x6f, 0x3b, 0x4f, 0x2e, 0x3a, 0x51, 0xf5, 0x8b, 0x1d, 0xbf, 0xdb, 0x03, 0xd4,
	0x9f, 0xd6, 0xfc, 0x22, 0x09, 0x01, 0x65, 0xaf, 0x03, 0xca, 0xec, 0x39, 0x7b, 0x6d, 0x14, 0x9f,
	0xc5, 0x23, 0xd4, 0x51, 0xd3, 0x80, 0x51, 0xe6, 0x6c, 0x77, 0x12, 0x0f, 0x50, 0x79, 0xb2, 0x06,
	0x65, 0xb1, 0xea, 0x5c, 0x25, 0xd8, 0x5e, 0xff, 0xb9, 0x71, 0x99, 0x50, 0x96, 0xac, 0x77, 0xda,
	0xbc, 0xf9, 0xd1, 0xf3, 0x80, 0xb2, 0x4a, 0x3e, 0x51, 0x7b, 0xcc, 0xe0, 0xfe, 0x83, 0xbb, 0x66,
	0x1d, 0x1d, 0x29, 0x0c, 0xde, 0xcd, 0xeb, 0xea, 0xcb, 0xba, 0xca, 0x4d, 0xdc, 0x72, 0xf3, 0x76,
	0xdb, 0x71, 0x6b, 0x97, 0x42, 0xd5, 0x7e, 0x29, 0x9e, 0xa0, 0x89, 0x76, 0xc2, 0x10, 0xf5, 0xe4,
	0xb9, 0xaf, 0x50, 0x17, 0x21, 0x24, 0x5c, 0x85, 0xc5, 0x18, 0x0c, 0x81, 0x23, 0xaf, 0xd5, 0x19,
	0xb7, 0x78, 0x24, 0x72, 0xc4, 0x95, 0x8d, 0x4a, 0xd0, 0xd7, 0xfc, 0x2d, 0xef, 0xa7, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x6b, 0x8f, 0xaa, 0x2f, 0x2a, 0x01, 0x00, 0x00,
}
