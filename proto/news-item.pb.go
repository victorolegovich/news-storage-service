// Code generated by protoc-gen-go. DO NOT EDIT.
// source: news-item.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type NewsItem struct {
	ID                   string   `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Header               string   `protobuf:"bytes,2,opt,name=Header,proto3" json:"Header,omitempty"`
	CreationDate         string   `protobuf:"bytes,3,opt,name=CreationDate,proto3" json:"CreationDate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewsItem) Reset()         { *m = NewsItem{} }
func (m *NewsItem) String() string { return proto.CompactTextString(m) }
func (*NewsItem) ProtoMessage()    {}
func (*NewsItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_bac58885563324fe, []int{0}
}

func (m *NewsItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewsItem.Unmarshal(m, b)
}
func (m *NewsItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewsItem.Marshal(b, m, deterministic)
}
func (m *NewsItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewsItem.Merge(m, src)
}
func (m *NewsItem) XXX_Size() int {
	return xxx_messageInfo_NewsItem.Size(m)
}
func (m *NewsItem) XXX_DiscardUnknown() {
	xxx_messageInfo_NewsItem.DiscardUnknown(m)
}

var xxx_messageInfo_NewsItem proto.InternalMessageInfo

func (m *NewsItem) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *NewsItem) GetHeader() string {
	if m != nil {
		return m.Header
	}
	return ""
}

func (m *NewsItem) GetCreationDate() string {
	if m != nil {
		return m.CreationDate
	}
	return ""
}

func init() {
	proto.RegisterType((*NewsItem)(nil), "proto.NewsItem")
}

func init() { proto.RegisterFile("news-item.proto", fileDescriptor_bac58885563324fe) }

var fileDescriptor_bac58885563324fe = []byte{
	// 115 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcf, 0x4b, 0x2d, 0x2f,
	0xd6, 0xcd, 0x2c, 0x49, 0xcd, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a,
	0x61, 0x5c, 0x1c, 0x7e, 0xa9, 0xe5, 0xc5, 0x9e, 0x25, 0xa9, 0xb9, 0x42, 0x7c, 0x5c, 0x4c, 0x9e,
	0x2e, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0x4c, 0x9e, 0x2e, 0x42, 0x62, 0x5c, 0x6c, 0x1e,
	0xa9, 0x89, 0x29, 0xa9, 0x45, 0x12, 0x4c, 0x60, 0x31, 0x28, 0x4f, 0x48, 0x89, 0x8b, 0xc7, 0xb9,
	0x28, 0x35, 0xb1, 0x24, 0x33, 0x3f, 0xcf, 0x25, 0xb1, 0x24, 0x55, 0x82, 0x19, 0x2c, 0x8b, 0x22,
	0x96, 0xc4, 0x06, 0x36, 0xde, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x52, 0x6f, 0x06, 0x17, 0x78,
	0x00, 0x00, 0x00,
}
