// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service_promotion/voucher_detail.proto

package service_promotion

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

type VoucherDetail struct {
	Voucher              *VoucherDTO  `protobuf:"bytes,1,opt,name=voucher,proto3" json:"voucher,omitempty"`
	Merchant             *MerchantDTO `protobuf:"bytes,2,opt,name=merchant,proto3" json:"merchant,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *VoucherDetail) Reset()         { *m = VoucherDetail{} }
func (m *VoucherDetail) String() string { return proto.CompactTextString(m) }
func (*VoucherDetail) ProtoMessage()    {}
func (*VoucherDetail) Descriptor() ([]byte, []int) {
	return fileDescriptor_voucher_detail_2bb69ade8e759ca2, []int{0}
}
func (m *VoucherDetail) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VoucherDetail.Unmarshal(m, b)
}
func (m *VoucherDetail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VoucherDetail.Marshal(b, m, deterministic)
}
func (dst *VoucherDetail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VoucherDetail.Merge(dst, src)
}
func (m *VoucherDetail) XXX_Size() int {
	return xxx_messageInfo_VoucherDetail.Size(m)
}
func (m *VoucherDetail) XXX_DiscardUnknown() {
	xxx_messageInfo_VoucherDetail.DiscardUnknown(m)
}

var xxx_messageInfo_VoucherDetail proto.InternalMessageInfo

func (m *VoucherDetail) GetVoucher() *VoucherDTO {
	if m != nil {
		return m.Voucher
	}
	return nil
}

func (m *VoucherDetail) GetMerchant() *MerchantDTO {
	if m != nil {
		return m.Merchant
	}
	return nil
}

func init() {
	proto.RegisterType((*VoucherDetail)(nil), "service_promotion.VoucherDetail")
}

func init() {
	proto.RegisterFile("service_promotion/voucher_detail.proto", fileDescriptor_voucher_detail_2bb69ade8e759ca2)
}

var fileDescriptor_voucher_detail_2bb69ade8e759ca2 = []byte{
	// 151 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2b, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0x8d, 0x2f, 0x28, 0xca, 0xcf, 0xcd, 0x2f, 0xc9, 0xcc, 0xcf, 0xd3, 0x2f, 0xcb,
	0x2f, 0x4d, 0xce, 0x48, 0x2d, 0x8a, 0x4f, 0x49, 0x2d, 0x49, 0xcc, 0xcc, 0xd1, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0x12, 0xc4, 0x50, 0x27, 0x25, 0x8f, 0x53, 0x2b, 0x44, 0x8f, 0x94, 0x02, 0xa6,
	0x82, 0xdc, 0xd4, 0xa2, 0xe4, 0x8c, 0xc4, 0xbc, 0x12, 0x88, 0x0a, 0xa5, 0x16, 0x46, 0x2e, 0xde,
	0x30, 0x88, 0x1e, 0x17, 0xb0, 0x6d, 0x42, 0xe6, 0x5c, 0xec, 0x50, 0x43, 0x24, 0x18, 0x15, 0x18,
	0x35, 0xb8, 0x8d, 0x64, 0xf5, 0x30, 0x4c, 0xd1, 0x83, 0x69, 0x09, 0xf1, 0x0f, 0x82, 0xa9, 0x16,
	0xb2, 0xe2, 0xe2, 0x80, 0x19, 0x2e, 0xc1, 0x04, 0xd6, 0x29, 0x87, 0x45, 0xa7, 0x2f, 0x54, 0x09,
	0x48, 0x2b, 0x5c, 0x7d, 0x12, 0x1b, 0xd8, 0x35, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xb7,
	0x57, 0x62, 0xfe, 0x0d, 0x01, 0x00, 0x00,
}
