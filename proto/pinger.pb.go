// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pinger.proto

package pinger

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

type Account struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Email                string   `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Token                string   `protobuf:"bytes,3,opt,name=token,proto3" json:"token,omitempty"`
	Password             string   `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	Name                 string   `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Account) Reset()         { *m = Account{} }
func (m *Account) String() string { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()    {}
func (*Account) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0fa9ff67baabf60, []int{0}
}

func (m *Account) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Account.Unmarshal(m, b)
}
func (m *Account) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Account.Marshal(b, m, deterministic)
}
func (m *Account) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Account.Merge(m, src)
}
func (m *Account) XXX_Size() int {
	return xxx_messageInfo_Account.Size(m)
}
func (m *Account) XXX_DiscardUnknown() {
	xxx_messageInfo_Account.DiscardUnknown(m)
}

var xxx_messageInfo_Account proto.InternalMessageInfo

func (m *Account) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Account) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Account) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *Account) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *Account) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Site struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Url                  string   `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Site) Reset()         { *m = Site{} }
func (m *Site) String() string { return proto.CompactTextString(m) }
func (*Site) ProtoMessage()    {}
func (*Site) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0fa9ff67baabf60, []int{1}
}

func (m *Site) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Site.Unmarshal(m, b)
}
func (m *Site) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Site.Marshal(b, m, deterministic)
}
func (m *Site) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Site.Merge(m, src)
}
func (m *Site) XXX_Size() int {
	return xxx_messageInfo_Site.Size(m)
}
func (m *Site) XXX_DiscardUnknown() {
	xxx_messageInfo_Site.DiscardUnknown(m)
}

var xxx_messageInfo_Site proto.InternalMessageInfo

func (m *Site) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Site) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type GetSitesRequest struct {
	Limit                int32    `protobuf:"varint,1,opt,name=limit,proto3" json:"limit,omitempty"`
	Offset               int32    `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetSitesRequest) Reset()         { *m = GetSitesRequest{} }
func (m *GetSitesRequest) String() string { return proto.CompactTextString(m) }
func (*GetSitesRequest) ProtoMessage()    {}
func (*GetSitesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0fa9ff67baabf60, []int{2}
}

func (m *GetSitesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSitesRequest.Unmarshal(m, b)
}
func (m *GetSitesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSitesRequest.Marshal(b, m, deterministic)
}
func (m *GetSitesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSitesRequest.Merge(m, src)
}
func (m *GetSitesRequest) XXX_Size() int {
	return xxx_messageInfo_GetSitesRequest.Size(m)
}
func (m *GetSitesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSitesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetSitesRequest proto.InternalMessageInfo

func (m *GetSitesRequest) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *GetSitesRequest) GetOffset() int32 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func init() {
	proto.RegisterType((*Account)(nil), "Account")
	proto.RegisterType((*Site)(nil), "Site")
	proto.RegisterType((*GetSitesRequest)(nil), "GetSitesRequest")
}

func init() { proto.RegisterFile("pinger.proto", fileDescriptor_e0fa9ff67baabf60) }

var fileDescriptor_e0fa9ff67baabf60 = []byte{
	// 193 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8f, 0xbd, 0x6e, 0xc2, 0x30,
	0x14, 0x85, 0x95, 0x1f, 0xa7, 0xed, 0x55, 0xd5, 0x56, 0x56, 0x55, 0x59, 0x9d, 0x50, 0xa6, 0x4c,
	0x2c, 0x3c, 0x00, 0x62, 0x62, 0x37, 0x4f, 0x10, 0x92, 0x1b, 0x64, 0x91, 0xd8, 0xc1, 0xbe, 0x16,
	0xaf, 0x8f, 0x7c, 0x03, 0x88, 0xed, 0x7c, 0x9f, 0x7d, 0x74, 0x6c, 0xf8, 0x9c, 0x8d, 0x3d, 0xa1,
	0x5f, 0xcf, 0xde, 0x91, 0xab, 0x23, 0xbc, 0xed, 0xba, 0xce, 0x45, 0x4b, 0xf2, 0x0b, 0x72, 0xd3,
	0xab, 0x6c, 0x95, 0x35, 0x1f, 0x3a, 0x37, 0xbd, 0xfc, 0x05, 0x81, 0x53, 0x6b, 0x46, 0x95, 0xb3,
	0x5a, 0x20, 0x59, 0x72, 0x67, 0xb4, 0xaa, 0x58, 0x2c, 0x83, 0xfc, 0x87, 0xf7, 0xb9, 0x0d, 0xe1,
	0xea, 0x7c, 0xaf, 0x4a, 0x3e, 0x78, 0xb2, 0x94, 0x50, 0xda, 0x76, 0x42, 0x25, 0xd8, 0x73, 0xae,
	0x1b, 0x28, 0x0f, 0x86, 0xf0, 0x65, 0xb3, 0xe0, 0xcd, 0x1f, 0x28, 0xa2, 0x7f, 0x2c, 0xa6, 0x58,
	0x6f, 0xe1, 0x7b, 0x8f, 0x94, 0x2e, 0x07, 0x8d, 0x97, 0x88, 0x81, 0xd2, 0x13, 0x46, 0x33, 0x19,
	0xe2, 0x9e, 0xd0, 0x0b, 0xc8, 0x3f, 0xa8, 0xdc, 0x30, 0x04, 0x24, 0x6e, 0x0b, 0x7d, 0xa7, 0x63,
	0xc5, 0x1f, 0xdd, 0xdc, 0x02, 0x00, 0x00, 0xff, 0xff, 0x56, 0x17, 0x6c, 0xfb, 0xf8, 0x00, 0x00,
	0x00,
}
