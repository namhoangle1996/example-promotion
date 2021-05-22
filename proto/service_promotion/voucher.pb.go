// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service_promotion/voucher.proto

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

type VoucherDTO struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Thumbnail            string   `protobuf:"bytes,3,opt,name=thumbnail,proto3" json:"thumbnail,omitempty"`
	Rule                 string   `protobuf:"bytes,4,opt,name=rule,proto3" json:"rule,omitempty"`
	Description          string   `protobuf:"bytes,5,opt,name=description,proto3" json:"description,omitempty"`
	ShortDescription     string   `protobuf:"bytes,6,opt,name=shortDescription,proto3" json:"shortDescription,omitempty"`
	Code                 string   `protobuf:"bytes,7,opt,name=code,proto3" json:"code,omitempty"`
	Amount               int64    `protobuf:"varint,8,opt,name=amount,proto3" json:"amount,omitempty"`
	Type                 string   `protobuf:"bytes,9,opt,name=type,proto3" json:"type,omitempty"`
	Total                int64    `protobuf:"varint,10,opt,name=total,proto3" json:"total,omitempty"`
	TotalFreeze          int64    `protobuf:"varint,11,opt,name=totalFreeze,proto3" json:"totalFreeze,omitempty"`
	FreezeIds            []string `protobuf:"bytes,12,rep,name=freezeIds,proto3" json:"freezeIds,omitempty"`
	TotalCanBuy          int64    `protobuf:"varint,13,opt,name=totalCanBuy,proto3" json:"totalCanBuy,omitempty"`
	TotalCanBuyPerDay    int64    `protobuf:"varint,14,opt,name=totalCanBuyPerDay,proto3" json:"totalCanBuyPerDay,omitempty"`
	TotalCanUsePerDay    int64    `protobuf:"varint,15,opt,name=totalCanUsePerDay,proto3" json:"totalCanUsePerDay,omitempty"`
	TotalCanUsePerWeek   int64    `protobuf:"varint,16,opt,name=totalCanUsePerWeek,proto3" json:"totalCanUsePerWeek,omitempty"`
	TotalCanUse          int64    `protobuf:"varint,17,opt,name=totalCanUse,proto3" json:"totalCanUse,omitempty"`
	DiscountAmount       uint64   `protobuf:"varint,18,opt,name=discountAmount,proto3" json:"discountAmount,omitempty"`
	DiscountPercent      float64  `protobuf:"fixed64,19,opt,name=discountPercent,proto3" json:"discountPercent,omitempty"`
	MerchantID           string   `protobuf:"bytes,20,opt,name=merchantID,proto3" json:"merchantID,omitempty"`
	MaxAmount            uint64   `protobuf:"varint,21,opt,name=maxAmount,proto3" json:"maxAmount,omitempty"`
	Status               string   `protobuf:"bytes,22,opt,name=status,proto3" json:"status,omitempty"`
	Show                 bool     `protobuf:"varint,23,opt,name=show,proto3" json:"show,omitempty"`
	AllowAll             bool     `protobuf:"varint,24,opt,name=allowAll,proto3" json:"allowAll,omitempty"`
	ShowInMyVoucher      bool     `protobuf:"varint,25,opt,name=showInMyVoucher,proto3" json:"showInMyVoucher,omitempty"`
	StartedAt            int64    `protobuf:"varint,26,opt,name=startedAt,proto3" json:"startedAt,omitempty"`
	ExpiredAt            int64    `protobuf:"varint,27,opt,name=expiredAt,proto3" json:"expiredAt,omitempty"`
	MinTransValue        int64    `protobuf:"varint,28,opt,name=min_trans_value,json=minTransValue,proto3" json:"min_trans_value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VoucherDTO) Reset()         { *m = VoucherDTO{} }
func (m *VoucherDTO) String() string { return proto.CompactTextString(m) }
func (*VoucherDTO) ProtoMessage()    {}
func (*VoucherDTO) Descriptor() ([]byte, []int) {
	return fileDescriptor_voucher_9c100c72a26b05c2, []int{0}
}
func (m *VoucherDTO) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VoucherDTO.Unmarshal(m, b)
}
func (m *VoucherDTO) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VoucherDTO.Marshal(b, m, deterministic)
}
func (dst *VoucherDTO) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VoucherDTO.Merge(dst, src)
}
func (m *VoucherDTO) XXX_Size() int {
	return xxx_messageInfo_VoucherDTO.Size(m)
}
func (m *VoucherDTO) XXX_DiscardUnknown() {
	xxx_messageInfo_VoucherDTO.DiscardUnknown(m)
}

var xxx_messageInfo_VoucherDTO proto.InternalMessageInfo

func (m *VoucherDTO) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *VoucherDTO) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *VoucherDTO) GetThumbnail() string {
	if m != nil {
		return m.Thumbnail
	}
	return ""
}

func (m *VoucherDTO) GetRule() string {
	if m != nil {
		return m.Rule
	}
	return ""
}

func (m *VoucherDTO) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *VoucherDTO) GetShortDescription() string {
	if m != nil {
		return m.ShortDescription
	}
	return ""
}

func (m *VoucherDTO) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *VoucherDTO) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *VoucherDTO) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *VoucherDTO) GetTotal() int64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *VoucherDTO) GetTotalFreeze() int64 {
	if m != nil {
		return m.TotalFreeze
	}
	return 0
}

func (m *VoucherDTO) GetFreezeIds() []string {
	if m != nil {
		return m.FreezeIds
	}
	return nil
}

func (m *VoucherDTO) GetTotalCanBuy() int64 {
	if m != nil {
		return m.TotalCanBuy
	}
	return 0
}

func (m *VoucherDTO) GetTotalCanBuyPerDay() int64 {
	if m != nil {
		return m.TotalCanBuyPerDay
	}
	return 0
}

func (m *VoucherDTO) GetTotalCanUsePerDay() int64 {
	if m != nil {
		return m.TotalCanUsePerDay
	}
	return 0
}

func (m *VoucherDTO) GetTotalCanUsePerWeek() int64 {
	if m != nil {
		return m.TotalCanUsePerWeek
	}
	return 0
}

func (m *VoucherDTO) GetTotalCanUse() int64 {
	if m != nil {
		return m.TotalCanUse
	}
	return 0
}

func (m *VoucherDTO) GetDiscountAmount() uint64 {
	if m != nil {
		return m.DiscountAmount
	}
	return 0
}

func (m *VoucherDTO) GetDiscountPercent() float64 {
	if m != nil {
		return m.DiscountPercent
	}
	return 0
}

func (m *VoucherDTO) GetMerchantID() string {
	if m != nil {
		return m.MerchantID
	}
	return ""
}

func (m *VoucherDTO) GetMaxAmount() uint64 {
	if m != nil {
		return m.MaxAmount
	}
	return 0
}

func (m *VoucherDTO) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *VoucherDTO) GetShow() bool {
	if m != nil {
		return m.Show
	}
	return false
}

func (m *VoucherDTO) GetAllowAll() bool {
	if m != nil {
		return m.AllowAll
	}
	return false
}

func (m *VoucherDTO) GetShowInMyVoucher() bool {
	if m != nil {
		return m.ShowInMyVoucher
	}
	return false
}

func (m *VoucherDTO) GetStartedAt() int64 {
	if m != nil {
		return m.StartedAt
	}
	return 0
}

func (m *VoucherDTO) GetExpiredAt() int64 {
	if m != nil {
		return m.ExpiredAt
	}
	return 0
}

func (m *VoucherDTO) GetMinTransValue() int64 {
	if m != nil {
		return m.MinTransValue
	}
	return 0
}

func init() {
	proto.RegisterType((*VoucherDTO)(nil), "service_promotion.VoucherDTO")
}

func init() {
	proto.RegisterFile("service_promotion/voucher.proto", fileDescriptor_voucher_9c100c72a26b05c2)
}

var fileDescriptor_voucher_9c100c72a26b05c2 = []byte{
	// 477 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x93, 0x51, 0x6f, 0xd3, 0x30,
	0x10, 0xc7, 0x95, 0xb6, 0x2b, 0xad, 0xc7, 0xda, 0xd5, 0x8c, 0x71, 0x8c, 0x09, 0x22, 0x1e, 0xa6,
	0x08, 0xa1, 0xf1, 0xc0, 0x27, 0x28, 0x54, 0x48, 0x7d, 0x40, 0x4c, 0xd5, 0x3a, 0x1e, 0x2b, 0x2f,
	0x39, 0x14, 0x8b, 0xc4, 0x8e, 0x6c, 0xa7, 0x5b, 0xf9, 0xf0, 0x08, 0xf9, 0x12, 0xd2, 0xb4, 0xdd,
	0xdb, 0xdd, 0xef, 0xff, 0x8b, 0xee, 0x6c, 0xb7, 0xec, 0x9d, 0x45, 0xb3, 0x96, 0x31, 0xae, 0x0a,
	0xa3, 0x73, 0xed, 0xa4, 0x56, 0x9f, 0xd6, 0xba, 0x8c, 0x53, 0x34, 0xd7, 0x85, 0xd1, 0x4e, 0xf3,
	0xc9, 0x81, 0xf0, 0xfe, 0x6f, 0x9f, 0xb1, 0xbb, 0x4a, 0x9a, 0xdd, 0xfe, 0xe0, 0x23, 0xd6, 0x91,
	0x09, 0x04, 0x61, 0x10, 0x0d, 0x17, 0x1d, 0x99, 0xf0, 0x33, 0x76, 0xe4, 0xa4, 0xcb, 0x10, 0x3a,
	0x84, 0xaa, 0x86, 0x5f, 0xb2, 0xa1, 0x4b, 0xcb, 0xfc, 0x5e, 0x09, 0x99, 0x41, 0x97, 0x92, 0x2d,
	0xe0, 0x9c, 0xf5, 0x4c, 0x99, 0x21, 0xf4, 0x28, 0xa0, 0x9a, 0x87, 0xec, 0x38, 0x41, 0x1b, 0x1b,
	0x59, 0xf8, 0xa9, 0x70, 0x44, 0x51, 0x1b, 0xf1, 0x0f, 0xec, 0xd4, 0xa6, 0xda, 0xb8, 0x59, 0x4b,
	0xeb, 0x93, 0x76, 0xc0, 0xfd, 0x84, 0x58, 0x27, 0x08, 0xcf, 0xaa, 0x09, 0xbe, 0xe6, 0xe7, 0xac,
	0x2f, 0x72, 0x5d, 0x2a, 0x07, 0x83, 0x30, 0x88, 0xba, 0x8b, 0xba, 0xf3, 0xae, 0xdb, 0x14, 0x08,
	0xc3, 0xca, 0xf5, 0x35, 0x9d, 0x4a, 0x3b, 0x91, 0x01, 0x23, 0xb5, 0x6a, 0xfc, 0x8e, 0x54, 0x7c,
	0x33, 0x88, 0x7f, 0x10, 0x8e, 0x29, 0x6b, 0x23, 0x7f, 0xee, 0x5f, 0x54, 0xcd, 0x13, 0x0b, 0xcf,
	0xc3, 0xae, 0x3f, 0x77, 0x03, 0x9a, 0xef, 0xbf, 0x0a, 0xf5, 0xa5, 0xdc, 0xc0, 0x49, 0xeb, 0xfb,
	0x0a, 0xf1, 0x8f, 0x6c, 0xd2, 0x6a, 0x6f, 0xd0, 0xcc, 0xc4, 0x06, 0x46, 0xe4, 0x1d, 0x06, 0x6d,
	0x7b, 0x69, 0xb1, 0xb6, 0xc7, 0xbb, 0x76, 0x13, 0xf0, 0x6b, 0xc6, 0x77, 0xe1, 0x4f, 0xc4, 0xdf,
	0x70, 0x4a, 0xfa, 0x13, 0x49, 0x7b, 0xdb, 0xa5, 0x45, 0x98, 0xec, 0x6e, 0xbb, 0xb4, 0xc8, 0xaf,
	0xd8, 0x28, 0x91, 0x36, 0xf6, 0xb7, 0x38, 0xad, 0x6e, 0x96, 0x87, 0x41, 0xd4, 0x5b, 0xec, 0x51,
	0x1e, 0xb1, 0xf1, 0x7f, 0x72, 0x83, 0x26, 0x46, 0xe5, 0xe0, 0x45, 0x18, 0x44, 0xc1, 0x62, 0x1f,
	0xf3, 0xb7, 0x8c, 0xe5, 0x68, 0xe2, 0x54, 0x28, 0x37, 0x9f, 0xc1, 0x19, 0xbd, 0x48, 0x8b, 0xf8,
	0xfb, 0xcd, 0xc5, 0x63, 0x3d, 0xec, 0x25, 0x0d, 0xdb, 0x02, 0xff, 0xc2, 0xd6, 0x09, 0x57, 0x5a,
	0x38, 0xa7, 0x2f, 0xeb, 0xce, 0xbf, 0xb0, 0x4d, 0xf5, 0x03, 0xbc, 0x0a, 0x83, 0x68, 0xb0, 0xa0,
	0x9a, 0x5f, 0xb0, 0x81, 0xc8, 0x32, 0xfd, 0x30, 0xcd, 0x32, 0x00, 0xe2, 0x4d, 0xef, 0xf7, 0xf5,
	0xce, 0x5c, 0x7d, 0xdf, 0xd4, 0xbf, 0x7c, 0x78, 0x4d, 0xca, 0x3e, 0xf6, 0xfb, 0x58, 0x27, 0x8c,
	0xc3, 0x64, 0xea, 0xe0, 0x82, 0x6e, 0x68, 0x0b, 0x7c, 0x8a, 0x8f, 0x85, 0x34, 0x94, 0xbe, 0xa9,
	0xd2, 0x06, 0xf0, 0x2b, 0x36, 0xce, 0xa5, 0x5a, 0x39, 0x23, 0x94, 0x5d, 0xad, 0x45, 0x56, 0x22,
	0x5c, 0x92, 0x73, 0x92, 0x4b, 0x75, 0xeb, 0xe9, 0x9d, 0x87, 0xf7, 0x7d, 0xfa, 0x6b, 0x7e, 0xfe,
	0x17, 0x00, 0x00, 0xff, 0xff, 0x25, 0xf7, 0x23, 0xf5, 0xbd, 0x03, 0x00, 0x00,
}