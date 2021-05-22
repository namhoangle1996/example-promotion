// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service_user/merchant_detail.proto

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

type MerchantDetailDTO struct {
	User                 *UserDTO     `protobuf:"bytes,1,opt,name=User,proto3" json:"User,omitempty"`
	Merchant             *MerchantDTO `protobuf:"bytes,2,opt,name=Merchant,proto3" json:"Merchant,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *MerchantDetailDTO) Reset()         { *m = MerchantDetailDTO{} }
func (m *MerchantDetailDTO) String() string { return proto.CompactTextString(m) }
func (*MerchantDetailDTO) ProtoMessage()    {}
func (*MerchantDetailDTO) Descriptor() ([]byte, []int) {
	return fileDescriptor_2535d0bf6eb917fc, []int{0}
}

func (m *MerchantDetailDTO) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MerchantDetailDTO.Unmarshal(m, b)
}
func (m *MerchantDetailDTO) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MerchantDetailDTO.Marshal(b, m, deterministic)
}
func (m *MerchantDetailDTO) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MerchantDetailDTO.Merge(m, src)
}
func (m *MerchantDetailDTO) XXX_Size() int {
	return xxx_messageInfo_MerchantDetailDTO.Size(m)
}
func (m *MerchantDetailDTO) XXX_DiscardUnknown() {
	xxx_messageInfo_MerchantDetailDTO.DiscardUnknown(m)
}

var xxx_messageInfo_MerchantDetailDTO proto.InternalMessageInfo

func (m *MerchantDetailDTO) GetUser() *UserDTO {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *MerchantDetailDTO) GetMerchant() *MerchantDTO {
	if m != nil {
		return m.Merchant
	}
	return nil
}

func init() {
	proto.RegisterType((*MerchantDetailDTO)(nil), "service_user.MerchantDetailDTO")
}

func init() {
	proto.RegisterFile("service_user/merchant_detail.proto", fileDescriptor_2535d0bf6eb917fc)
}

var fileDescriptor_2535d0bf6eb917fc = []byte{
	// 140 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x52, 0x2a, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0x8d, 0x2f, 0x05, 0x32, 0xf4, 0x73, 0x53, 0x8b, 0x92, 0x33, 0x12, 0xf3, 0x4a,
	0xe2, 0x53, 0x52, 0x4b, 0x12, 0x33, 0x73, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x78, 0x90,
	0xd5, 0x48, 0x89, 0xa3, 0xe8, 0x00, 0x11, 0x10, 0x65, 0x52, 0xd2, 0x58, 0x8d, 0x82, 0x48, 0x2a,
	0x95, 0x72, 0x09, 0xfa, 0x42, 0x45, 0x5c, 0xc0, 0x66, 0xbb, 0x84, 0xf8, 0x0b, 0x69, 0x72, 0xb1,
	0x84, 0x02, 0xd5, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x1b, 0x89, 0xea, 0x21, 0x1b, 0xa0, 0x07,
	0x92, 0x01, 0x2a, 0x0a, 0x02, 0x2b, 0x11, 0x32, 0xe5, 0xe2, 0x80, 0xe9, 0x97, 0x60, 0x02, 0x2b,
	0x97, 0x44, 0x55, 0x0e, 0x37, 0x1d, 0xa8, 0x05, 0xae, 0x34, 0x89, 0x0d, 0x6c, 0xbb, 0x31, 0x20,
	0x00, 0x00, 0xff, 0xff, 0x43, 0x62, 0x45, 0xd2, 0xe7, 0x00, 0x00, 0x00,
}