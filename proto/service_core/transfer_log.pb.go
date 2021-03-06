// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service_core/transfer_log.proto

package service_core

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

type TransferMoneyLogDTO struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
	WalletIdSrc          string   `protobuf:"bytes,2,opt,name=walletIdSrc,proto3" json:"walletIdSrc"`
	WalletIdDest         string   `protobuf:"bytes,3,opt,name=walletIdDest,proto3" json:"walletIdDest"`
	TransactionId        string   `protobuf:"bytes,4,opt,name=transactionId,proto3" json:"transactionId"`
	Amount               uint64   `protobuf:"varint,5,opt,name=amount,proto3" json:"amount"`
	State                string   `protobuf:"bytes,6,opt,name=state,proto3" json:"state"`
	CreatedAt            int64    `protobuf:"varint,7,opt,name=createdAt,proto3" json:"createdAt"`
	UpdatedAt            int64    `protobuf:"varint,8,opt,name=updatedAt,proto3" json:"updatedAt"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransferMoneyLogDTO) Reset()         { *m = TransferMoneyLogDTO{} }
func (m *TransferMoneyLogDTO) String() string { return proto.CompactTextString(m) }
func (*TransferMoneyLogDTO) ProtoMessage()    {}
func (*TransferMoneyLogDTO) Descriptor() ([]byte, []int) {
	return fileDescriptor_c183c052743f0ff5, []int{0}
}

func (m *TransferMoneyLogDTO) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransferMoneyLogDTO.Unmarshal(m, b)
}
func (m *TransferMoneyLogDTO) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransferMoneyLogDTO.Marshal(b, m, deterministic)
}
func (m *TransferMoneyLogDTO) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransferMoneyLogDTO.Merge(m, src)
}
func (m *TransferMoneyLogDTO) XXX_Size() int {
	return xxx_messageInfo_TransferMoneyLogDTO.Size(m)
}
func (m *TransferMoneyLogDTO) XXX_DiscardUnknown() {
	xxx_messageInfo_TransferMoneyLogDTO.DiscardUnknown(m)
}

var xxx_messageInfo_TransferMoneyLogDTO proto.InternalMessageInfo

func (m *TransferMoneyLogDTO) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *TransferMoneyLogDTO) GetWalletIdSrc() string {
	if m != nil {
		return m.WalletIdSrc
	}
	return ""
}

func (m *TransferMoneyLogDTO) GetWalletIdDest() string {
	if m != nil {
		return m.WalletIdDest
	}
	return ""
}

func (m *TransferMoneyLogDTO) GetTransactionId() string {
	if m != nil {
		return m.TransactionId
	}
	return ""
}

func (m *TransferMoneyLogDTO) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *TransferMoneyLogDTO) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *TransferMoneyLogDTO) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *TransferMoneyLogDTO) GetUpdatedAt() int64 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

func init() {
	proto.RegisterType((*TransferMoneyLogDTO)(nil), "service_core.TransferMoneyLogDTO")
}

func init() {
	proto.RegisterFile("service_core/transfer_log.proto", fileDescriptor_c183c052743f0ff5)
}

var fileDescriptor_c183c052743f0ff5 = []byte{
	// 221 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0x41, 0x4b, 0x03, 0x31,
	0x10, 0x46, 0xc9, 0xb6, 0x5d, 0xed, 0x58, 0x3d, 0x8c, 0x22, 0x39, 0x08, 0x86, 0xe2, 0x61, 0x4f,
	0x7a, 0xf0, 0x17, 0x08, 0xbd, 0x14, 0x14, 0x61, 0xed, 0xbd, 0xc4, 0x64, 0x2c, 0x81, 0x35, 0x53,
	0x92, 0xa9, 0xe2, 0x6f, 0xf7, 0x22, 0xa6, 0xab, 0x6d, 0x8f, 0xdf, 0x7b, 0x2f, 0x10, 0x06, 0xae,
	0x33, 0xa5, 0x8f, 0xe0, 0x68, 0xe9, 0x38, 0xd1, 0x9d, 0x24, 0x1b, 0xf3, 0x1b, 0xa5, 0x65, 0xc7,
	0xab, 0xdb, 0x75, 0x62, 0x61, 0x9c, 0xec, 0x07, 0xd3, 0x6f, 0x05, 0xe7, 0x8b, 0x3e, 0x7a, 0xe2,
	0x48, 0x5f, 0x8f, 0xbc, 0x9a, 0x2d, 0x9e, 0xf1, 0x0c, 0xaa, 0xe0, 0xb5, 0x32, 0xaa, 0x19, 0xb7,
	0x55, 0xf0, 0x68, 0xe0, 0xe4, 0xd3, 0x76, 0x1d, 0xc9, 0xdc, 0xbf, 0x24, 0xa7, 0xab, 0x22, 0xf6,
	0x11, 0x4e, 0x61, 0xf2, 0x37, 0x67, 0x94, 0x45, 0x0f, 0x4a, 0x72, 0xc0, 0xf0, 0x06, 0x4e, 0xcb,
	0x8f, 0xac, 0x93, 0xc0, 0x71, 0xee, 0xf5, 0xb0, 0x44, 0x87, 0x10, 0x2f, 0xa1, 0xb6, 0xef, 0xbc,
	0x89, 0xa2, 0x47, 0x46, 0x35, 0xc3, 0xb6, 0x5f, 0x78, 0x01, 0xa3, 0x2c, 0x56, 0x48, 0xd7, 0xe5,
	0xd5, 0x76, 0xe0, 0x15, 0x8c, 0x5d, 0x22, 0x2b, 0xe4, 0x1f, 0x44, 0x1f, 0x19, 0xd5, 0x0c, 0xda,
	0x1d, 0xf8, 0xb5, 0x9b, 0xb5, 0xef, 0xed, 0xf1, 0xd6, 0xfe, 0x83, 0xd7, 0xba, 0x9c, 0xe4, 0xfe,
	0x27, 0x00, 0x00, 0xff, 0xff, 0x78, 0xcf, 0x54, 0xf2, 0x35, 0x01, 0x00, 0x00,
}
