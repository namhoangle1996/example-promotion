// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service_user/balance.proto

package service_user

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

type BalanceDTO struct {
	Id                   string   `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	AmountAvailable      int64    `protobuf:"varint,2,opt,name=AmountAvailable,proto3" json:"AmountAvailable,omitempty"`
	AmountFreeze         int64    `protobuf:"varint,3,opt,name=AmountFreeze,proto3" json:"AmountFreeze,omitempty"`
	Type                 string   `protobuf:"bytes,4,opt,name=Type,proto3" json:"Type,omitempty"`
	UserId               string   `protobuf:"bytes,5,opt,name=UserId,proto3" json:"UserId,omitempty"`
	FreezeIds            []string `protobuf:"bytes,6,rep,name=FreezeIds,proto3" json:"FreezeIds,omitempty"`
	Currency             string   `protobuf:"bytes,7,opt,name=Currency,proto3" json:"Currency,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BalanceDTO) Reset()         { *m = BalanceDTO{} }
func (m *BalanceDTO) String() string { return proto.CompactTextString(m) }
func (*BalanceDTO) ProtoMessage()    {}
func (*BalanceDTO) Descriptor() ([]byte, []int) {
	return fileDescriptor_afb05c4e66a226a6, []int{0}
}

func (m *BalanceDTO) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BalanceDTO.Unmarshal(m, b)
}
func (m *BalanceDTO) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BalanceDTO.Marshal(b, m, deterministic)
}
func (m *BalanceDTO) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BalanceDTO.Merge(m, src)
}
func (m *BalanceDTO) XXX_Size() int {
	return xxx_messageInfo_BalanceDTO.Size(m)
}
func (m *BalanceDTO) XXX_DiscardUnknown() {
	xxx_messageInfo_BalanceDTO.DiscardUnknown(m)
}

var xxx_messageInfo_BalanceDTO proto.InternalMessageInfo

func (m *BalanceDTO) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *BalanceDTO) GetAmountAvailable() int64 {
	if m != nil {
		return m.AmountAvailable
	}
	return 0
}

func (m *BalanceDTO) GetAmountFreeze() int64 {
	if m != nil {
		return m.AmountFreeze
	}
	return 0
}

func (m *BalanceDTO) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *BalanceDTO) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *BalanceDTO) GetFreezeIds() []string {
	if m != nil {
		return m.FreezeIds
	}
	return nil
}

func (m *BalanceDTO) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

func init() {
	proto.RegisterType((*BalanceDTO)(nil), "service_user.BalanceDTO")
}

func init() {
	proto.RegisterFile("service_user/balance.proto", fileDescriptor_afb05c4e66a226a6)
}

var fileDescriptor_afb05c4e66a226a6 = []byte{
	// 193 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x92, 0x2a, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0x8d, 0x2f, 0x05, 0x32, 0xf4, 0x93, 0x12, 0x73, 0x12, 0xf3, 0x92, 0x53, 0xf5,
	0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x78, 0x90, 0xe5, 0x94, 0x2e, 0x30, 0x72, 0x71, 0x39, 0x41,
	0xe4, 0x5d, 0x42, 0xfc, 0x85, 0xf8, 0xb8, 0x98, 0x3c, 0x53, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38,
	0x83, 0x80, 0x2c, 0x21, 0x0d, 0x2e, 0x7e, 0xc7, 0xdc, 0xfc, 0xd2, 0xbc, 0x12, 0xc7, 0xb2, 0xc4,
	0xcc, 0x9c, 0xc4, 0xa4, 0x9c, 0x54, 0x09, 0x26, 0xa0, 0x24, 0x73, 0x10, 0xba, 0xb0, 0x90, 0x12,
	0x17, 0x0f, 0x44, 0xc8, 0xad, 0x28, 0x35, 0xb5, 0x2a, 0x55, 0x82, 0x19, 0xac, 0x0c, 0x45, 0x4c,
	0x48, 0x88, 0x8b, 0x25, 0xa4, 0xb2, 0x20, 0x55, 0x82, 0x05, 0x6c, 0x3e, 0x98, 0x2d, 0x24, 0xc6,
	0xc5, 0x16, 0x0a, 0x74, 0x08, 0xd0, 0x56, 0x56, 0xb0, 0x28, 0x94, 0x27, 0x24, 0xc3, 0xc5, 0x09,
	0xd1, 0xe5, 0x99, 0x52, 0x2c, 0xc1, 0xa6, 0xc0, 0x0c, 0x94, 0x42, 0x08, 0x08, 0x49, 0x71, 0x71,
	0x38, 0x97, 0x16, 0x15, 0xa5, 0xe6, 0x25, 0x57, 0x4a, 0xb0, 0x83, 0xf5, 0xc1, 0xf9, 0x49, 0x6c,
	0x60, 0x7f, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x15, 0x85, 0x39, 0x28, 0x05, 0x01, 0x00,
	0x00,
}
