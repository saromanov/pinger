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
	UserId               string   `protobuf:"bytes,3,opt,name=userId,proto3" json:"userId,omitempty"`
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

func (m *Site) GetUserId() string {
	if m != nil {
		return m.UserId
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

type GetSiteRequest struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetSiteRequest) Reset()         { *m = GetSiteRequest{} }
func (m *GetSiteRequest) String() string { return proto.CompactTextString(m) }
func (*GetSiteRequest) ProtoMessage()    {}
func (*GetSiteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0fa9ff67baabf60, []int{3}
}

func (m *GetSiteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetSiteRequest.Unmarshal(m, b)
}
func (m *GetSiteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetSiteRequest.Marshal(b, m, deterministic)
}
func (m *GetSiteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetSiteRequest.Merge(m, src)
}
func (m *GetSiteRequest) XXX_Size() int {
	return xxx_messageInfo_GetSiteRequest.Size(m)
}
func (m *GetSiteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetSiteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetSiteRequest proto.InternalMessageInfo

func (m *GetSiteRequest) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type GetStatsRequest struct {
	SiteID               int64    `protobuf:"varint,1,opt,name=siteID,proto3" json:"siteID,omitempty"`
	UserID               string   `protobuf:"bytes,2,opt,name=userID,proto3" json:"userID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetStatsRequest) Reset()         { *m = GetStatsRequest{} }
func (m *GetStatsRequest) String() string { return proto.CompactTextString(m) }
func (*GetStatsRequest) ProtoMessage()    {}
func (*GetStatsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0fa9ff67baabf60, []int{4}
}

func (m *GetStatsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetStatsRequest.Unmarshal(m, b)
}
func (m *GetStatsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetStatsRequest.Marshal(b, m, deterministic)
}
func (m *GetStatsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetStatsRequest.Merge(m, src)
}
func (m *GetStatsRequest) XXX_Size() int {
	return xxx_messageInfo_GetStatsRequest.Size(m)
}
func (m *GetStatsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetStatsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetStatsRequest proto.InternalMessageInfo

func (m *GetStatsRequest) GetSiteID() int64 {
	if m != nil {
		return m.SiteID
	}
	return 0
}

func (m *GetStatsRequest) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

type CountStatRequest struct {
	SiteID               int64    `protobuf:"varint,1,opt,name=siteID,proto3" json:"siteID,omitempty"`
	UserID               int64    `protobuf:"varint,2,opt,name=userID,proto3" json:"userID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CountStatRequest) Reset()         { *m = CountStatRequest{} }
func (m *CountStatRequest) String() string { return proto.CompactTextString(m) }
func (*CountStatRequest) ProtoMessage()    {}
func (*CountStatRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0fa9ff67baabf60, []int{5}
}

func (m *CountStatRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CountStatRequest.Unmarshal(m, b)
}
func (m *CountStatRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CountStatRequest.Marshal(b, m, deterministic)
}
func (m *CountStatRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CountStatRequest.Merge(m, src)
}
func (m *CountStatRequest) XXX_Size() int {
	return xxx_messageInfo_CountStatRequest.Size(m)
}
func (m *CountStatRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CountStatRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CountStatRequest proto.InternalMessageInfo

func (m *CountStatRequest) GetSiteID() int64 {
	if m != nil {
		return m.SiteID
	}
	return 0
}

func (m *CountStatRequest) GetUserID() int64 {
	if m != nil {
		return m.UserID
	}
	return 0
}

func init() {
	proto.RegisterType((*Account)(nil), "Account")
	proto.RegisterType((*Site)(nil), "Site")
	proto.RegisterType((*GetSitesRequest)(nil), "GetSitesRequest")
	proto.RegisterType((*GetSiteRequest)(nil), "GetSiteRequest")
	proto.RegisterType((*GetStatsRequest)(nil), "GetStatsRequest")
	proto.RegisterType((*CountStatRequest)(nil), "CountStatRequest")
}

func init() { proto.RegisterFile("pinger.proto", fileDescriptor_e0fa9ff67baabf60) }

var fileDescriptor_e0fa9ff67baabf60 = []byte{
	// 250 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xcd, 0x4a, 0xc3, 0x40,
	0x14, 0x85, 0x69, 0xfe, 0xd4, 0x8b, 0xd4, 0x32, 0x48, 0x19, 0x5c, 0x95, 0x59, 0xb9, 0x72, 0xe3,
	0x03, 0x68, 0xb5, 0x20, 0xdd, 0xc6, 0x27, 0x88, 0xcd, 0xad, 0x0c, 0x36, 0x99, 0x38, 0x73, 0x07,
	0x5f, 0x5f, 0xe6, 0xe6, 0x26, 0x88, 0xbb, 0xee, 0xe6, 0x3b, 0x49, 0xbe, 0x73, 0xc8, 0xc0, 0xf5,
	0x60, 0xfb, 0x4f, 0xf4, 0x0f, 0x83, 0x77, 0xe4, 0x4c, 0x84, 0x8b, 0xed, 0xe1, 0xe0, 0x62, 0x4f,
	0x6a, 0x09, 0x99, 0x6d, 0xf5, 0x62, 0xb3, 0xb8, 0xbf, 0xaa, 0x33, 0xdb, 0xaa, 0x5b, 0x28, 0xb1,
	0x6b, 0xec, 0x49, 0x67, 0x1c, 0x8d, 0x90, 0x52, 0x72, 0x5f, 0xd8, 0xeb, 0x7c, 0x4c, 0x19, 0xd4,
	0x1d, 0x5c, 0x0e, 0x4d, 0x08, 0x3f, 0xce, 0xb7, 0xba, 0xe0, 0x07, 0x33, 0x2b, 0x05, 0x45, 0xdf,
	0x74, 0xa8, 0x4b, 0xce, 0xf9, 0x6c, 0x9e, 0xa1, 0x78, 0xb7, 0x84, 0x7f, 0x3a, 0x73, 0xee, 0x5c,
	0x41, 0x1e, 0xfd, 0xd4, 0x98, 0x8e, 0x6a, 0x0d, 0x55, 0x0c, 0xe8, 0xf7, 0xad, 0x14, 0x0a, 0x99,
	0x27, 0xb8, 0x79, 0x43, 0x4a, 0x92, 0x50, 0xe3, 0x77, 0xc4, 0x40, 0x69, 0xda, 0xc9, 0x76, 0x96,
	0xd8, 0x57, 0xd6, 0x23, 0x24, 0x81, 0x3b, 0x1e, 0x03, 0x12, 0x5b, 0xcb, 0x5a, 0xc8, 0x6c, 0x60,
	0x29, 0x82, 0xe9, 0xfb, 0x7f, 0x63, 0xcc, 0x76, 0xac, 0xa0, 0x86, 0xe6, 0x8a, 0x35, 0x54, 0xc1,
	0x12, 0xee, 0x77, 0xf2, 0x9a, 0xd0, 0xbc, 0x72, 0x27, 0xd3, 0x85, 0xcc, 0x0b, 0xac, 0x5e, 0xd3,
	0xcf, 0x4d, 0x92, 0xf3, 0x1c, 0xf9, 0xe4, 0xf8, 0xa8, 0xf8, 0xa6, 0x1e, 0x7f, 0x03, 0x00, 0x00,
	0xff, 0xff, 0xa2, 0x6a, 0xe1, 0x6f, 0xb9, 0x01, 0x00, 0x00,
}
