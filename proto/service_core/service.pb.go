// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service_core/service.proto

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

type ReleaseFreezeRequest struct {
	TransactionId        string   `protobuf:"bytes,1,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReleaseFreezeRequest) Reset()         { *m = ReleaseFreezeRequest{} }
func (m *ReleaseFreezeRequest) String() string { return proto.CompactTextString(m) }
func (*ReleaseFreezeRequest) ProtoMessage()    {}
func (*ReleaseFreezeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_aa1d687600079dba, []int{0}
}

func (m *ReleaseFreezeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReleaseFreezeRequest.Unmarshal(m, b)
}
func (m *ReleaseFreezeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReleaseFreezeRequest.Marshal(b, m, deterministic)
}
func (m *ReleaseFreezeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReleaseFreezeRequest.Merge(m, src)
}
func (m *ReleaseFreezeRequest) XXX_Size() int {
	return xxx_messageInfo_ReleaseFreezeRequest.Size(m)
}
func (m *ReleaseFreezeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ReleaseFreezeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ReleaseFreezeRequest proto.InternalMessageInfo

func (m *ReleaseFreezeRequest) GetTransactionId() string {
	if m != nil {
		return m.TransactionId
	}
	return ""
}

type FindWalletByIDRequest struct {
	WalletId             string   `protobuf:"bytes,1,opt,name=wallet_id,json=walletId,proto3" json:"wallet_id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindWalletByIDRequest) Reset()         { *m = FindWalletByIDRequest{} }
func (m *FindWalletByIDRequest) String() string { return proto.CompactTextString(m) }
func (*FindWalletByIDRequest) ProtoMessage()    {}
func (*FindWalletByIDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_aa1d687600079dba, []int{1}
}

func (m *FindWalletByIDRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindWalletByIDRequest.Unmarshal(m, b)
}
func (m *FindWalletByIDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindWalletByIDRequest.Marshal(b, m, deterministic)
}
func (m *FindWalletByIDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindWalletByIDRequest.Merge(m, src)
}
func (m *FindWalletByIDRequest) XXX_Size() int {
	return xxx_messageInfo_FindWalletByIDRequest.Size(m)
}
func (m *FindWalletByIDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindWalletByIDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindWalletByIDRequest proto.InternalMessageInfo

func (m *FindWalletByIDRequest) GetWalletId() string {
	if m != nil {
		return m.WalletId
	}
	return ""
}

type FindOrCreateByUserIDRequest struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id"`
	Currency             string   `protobuf:"bytes,2,opt,name=currency,proto3" json:"currency"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindOrCreateByUserIDRequest) Reset()         { *m = FindOrCreateByUserIDRequest{} }
func (m *FindOrCreateByUserIDRequest) String() string { return proto.CompactTextString(m) }
func (*FindOrCreateByUserIDRequest) ProtoMessage()    {}
func (*FindOrCreateByUserIDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_aa1d687600079dba, []int{2}
}

func (m *FindOrCreateByUserIDRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindOrCreateByUserIDRequest.Unmarshal(m, b)
}
func (m *FindOrCreateByUserIDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindOrCreateByUserIDRequest.Marshal(b, m, deterministic)
}
func (m *FindOrCreateByUserIDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindOrCreateByUserIDRequest.Merge(m, src)
}
func (m *FindOrCreateByUserIDRequest) XXX_Size() int {
	return xxx_messageInfo_FindOrCreateByUserIDRequest.Size(m)
}
func (m *FindOrCreateByUserIDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindOrCreateByUserIDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindOrCreateByUserIDRequest proto.InternalMessageInfo

func (m *FindOrCreateByUserIDRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *FindOrCreateByUserIDRequest) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

type FindWalletByUserIDRequest struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindWalletByUserIDRequest) Reset()         { *m = FindWalletByUserIDRequest{} }
func (m *FindWalletByUserIDRequest) String() string { return proto.CompactTextString(m) }
func (*FindWalletByUserIDRequest) ProtoMessage()    {}
func (*FindWalletByUserIDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_aa1d687600079dba, []int{3}
}

func (m *FindWalletByUserIDRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindWalletByUserIDRequest.Unmarshal(m, b)
}
func (m *FindWalletByUserIDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindWalletByUserIDRequest.Marshal(b, m, deterministic)
}
func (m *FindWalletByUserIDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindWalletByUserIDRequest.Merge(m, src)
}
func (m *FindWalletByUserIDRequest) XXX_Size() int {
	return xxx_messageInfo_FindWalletByUserIDRequest.Size(m)
}
func (m *FindWalletByUserIDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindWalletByUserIDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindWalletByUserIDRequest proto.InternalMessageInfo

func (m *FindWalletByUserIDRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type FreezeRequest struct {
	WalletId             string   `protobuf:"bytes,1,opt,name=wallet_id,json=walletId,proto3" json:"wallet_id"`
	Amount               uint64   `protobuf:"varint,2,opt,name=amount,proto3" json:"amount"`
	TransactionId        string   `protobuf:"bytes,3,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FreezeRequest) Reset()         { *m = FreezeRequest{} }
func (m *FreezeRequest) String() string { return proto.CompactTextString(m) }
func (*FreezeRequest) ProtoMessage()    {}
func (*FreezeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_aa1d687600079dba, []int{4}
}

func (m *FreezeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FreezeRequest.Unmarshal(m, b)
}
func (m *FreezeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FreezeRequest.Marshal(b, m, deterministic)
}
func (m *FreezeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FreezeRequest.Merge(m, src)
}
func (m *FreezeRequest) XXX_Size() int {
	return xxx_messageInfo_FreezeRequest.Size(m)
}
func (m *FreezeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FreezeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FreezeRequest proto.InternalMessageInfo

func (m *FreezeRequest) GetWalletId() string {
	if m != nil {
		return m.WalletId
	}
	return ""
}

func (m *FreezeRequest) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *FreezeRequest) GetTransactionId() string {
	if m != nil {
		return m.TransactionId
	}
	return ""
}

type TransferRequest struct {
	WalletIdSrc          string   `protobuf:"bytes,1,opt,name=wallet_id_src,json=walletIdSrc,proto3" json:"wallet_id_src"`
	WalletIdDest         string   `protobuf:"bytes,2,opt,name=wallet_id_dest,json=walletIdDest,proto3" json:"wallet_id_dest"`
	Amount               uint64   `protobuf:"varint,3,opt,name=amount,proto3" json:"amount"`
	TransactionId        string   `protobuf:"bytes,4,opt,name=transaction_id,json=transactionId,proto3" json:"transaction_id"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransferRequest) Reset()         { *m = TransferRequest{} }
func (m *TransferRequest) String() string { return proto.CompactTextString(m) }
func (*TransferRequest) ProtoMessage()    {}
func (*TransferRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_aa1d687600079dba, []int{5}
}

func (m *TransferRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransferRequest.Unmarshal(m, b)
}
func (m *TransferRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransferRequest.Marshal(b, m, deterministic)
}
func (m *TransferRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransferRequest.Merge(m, src)
}
func (m *TransferRequest) XXX_Size() int {
	return xxx_messageInfo_TransferRequest.Size(m)
}
func (m *TransferRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TransferRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TransferRequest proto.InternalMessageInfo

func (m *TransferRequest) GetWalletIdSrc() string {
	if m != nil {
		return m.WalletIdSrc
	}
	return ""
}

func (m *TransferRequest) GetWalletIdDest() string {
	if m != nil {
		return m.WalletIdDest
	}
	return ""
}

func (m *TransferRequest) GetAmount() uint64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *TransferRequest) GetTransactionId() string {
	if m != nil {
		return m.TransactionId
	}
	return ""
}

type FindWalletByUserIDTypeRequest struct {
	UserId               string   `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id"`
	TypeWallet           string   `protobuf:"bytes,2,opt,name=type_wallet,json=typeWallet,proto3" json:"type_wallet"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindWalletByUserIDTypeRequest) Reset()         { *m = FindWalletByUserIDTypeRequest{} }
func (m *FindWalletByUserIDTypeRequest) String() string { return proto.CompactTextString(m) }
func (*FindWalletByUserIDTypeRequest) ProtoMessage()    {}
func (*FindWalletByUserIDTypeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_aa1d687600079dba, []int{6}
}

func (m *FindWalletByUserIDTypeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindWalletByUserIDTypeRequest.Unmarshal(m, b)
}
func (m *FindWalletByUserIDTypeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindWalletByUserIDTypeRequest.Marshal(b, m, deterministic)
}
func (m *FindWalletByUserIDTypeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindWalletByUserIDTypeRequest.Merge(m, src)
}
func (m *FindWalletByUserIDTypeRequest) XXX_Size() int {
	return xxx_messageInfo_FindWalletByUserIDTypeRequest.Size(m)
}
func (m *FindWalletByUserIDTypeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindWalletByUserIDTypeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindWalletByUserIDTypeRequest proto.InternalMessageInfo

func (m *FindWalletByUserIDTypeRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *FindWalletByUserIDTypeRequest) GetTypeWallet() string {
	if m != nil {
		return m.TypeWallet
	}
	return ""
}

type FindWalletByUserIDTypeResponse struct {
	Balance              *BalanceDTO `protobuf:"bytes,1,opt,name=balance,proto3" json:"balance"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *FindWalletByUserIDTypeResponse) Reset()         { *m = FindWalletByUserIDTypeResponse{} }
func (m *FindWalletByUserIDTypeResponse) String() string { return proto.CompactTextString(m) }
func (*FindWalletByUserIDTypeResponse) ProtoMessage()    {}
func (*FindWalletByUserIDTypeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_aa1d687600079dba, []int{7}
}

func (m *FindWalletByUserIDTypeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindWalletByUserIDTypeResponse.Unmarshal(m, b)
}
func (m *FindWalletByUserIDTypeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindWalletByUserIDTypeResponse.Marshal(b, m, deterministic)
}
func (m *FindWalletByUserIDTypeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindWalletByUserIDTypeResponse.Merge(m, src)
}
func (m *FindWalletByUserIDTypeResponse) XXX_Size() int {
	return xxx_messageInfo_FindWalletByUserIDTypeResponse.Size(m)
}
func (m *FindWalletByUserIDTypeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FindWalletByUserIDTypeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FindWalletByUserIDTypeResponse proto.InternalMessageInfo

func (m *FindWalletByUserIDTypeResponse) GetBalance() *BalanceDTO {
	if m != nil {
		return m.Balance
	}
	return nil
}

type FindWalletByIDResponse struct {
	Balance              *BalanceDTO `protobuf:"bytes,1,opt,name=balance,proto3" json:"balance"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *FindWalletByIDResponse) Reset()         { *m = FindWalletByIDResponse{} }
func (m *FindWalletByIDResponse) String() string { return proto.CompactTextString(m) }
func (*FindWalletByIDResponse) ProtoMessage()    {}
func (*FindWalletByIDResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_aa1d687600079dba, []int{8}
}

func (m *FindWalletByIDResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindWalletByIDResponse.Unmarshal(m, b)
}
func (m *FindWalletByIDResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindWalletByIDResponse.Marshal(b, m, deterministic)
}
func (m *FindWalletByIDResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindWalletByIDResponse.Merge(m, src)
}
func (m *FindWalletByIDResponse) XXX_Size() int {
	return xxx_messageInfo_FindWalletByIDResponse.Size(m)
}
func (m *FindWalletByIDResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FindWalletByIDResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FindWalletByIDResponse proto.InternalMessageInfo

func (m *FindWalletByIDResponse) GetBalance() *BalanceDTO {
	if m != nil {
		return m.Balance
	}
	return nil
}

type FindOrCreateWalletByUserIDResponse struct {
	Balances             []*BalanceDTO `protobuf:"bytes,1,rep,name=balances,proto3" json:"balances"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *FindOrCreateWalletByUserIDResponse) Reset()         { *m = FindOrCreateWalletByUserIDResponse{} }
func (m *FindOrCreateWalletByUserIDResponse) String() string { return proto.CompactTextString(m) }
func (*FindOrCreateWalletByUserIDResponse) ProtoMessage()    {}
func (*FindOrCreateWalletByUserIDResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_aa1d687600079dba, []int{9}
}

func (m *FindOrCreateWalletByUserIDResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindOrCreateWalletByUserIDResponse.Unmarshal(m, b)
}
func (m *FindOrCreateWalletByUserIDResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindOrCreateWalletByUserIDResponse.Marshal(b, m, deterministic)
}
func (m *FindOrCreateWalletByUserIDResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindOrCreateWalletByUserIDResponse.Merge(m, src)
}
func (m *FindOrCreateWalletByUserIDResponse) XXX_Size() int {
	return xxx_messageInfo_FindOrCreateWalletByUserIDResponse.Size(m)
}
func (m *FindOrCreateWalletByUserIDResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FindOrCreateWalletByUserIDResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FindOrCreateWalletByUserIDResponse proto.InternalMessageInfo

func (m *FindOrCreateWalletByUserIDResponse) GetBalances() []*BalanceDTO {
	if m != nil {
		return m.Balances
	}
	return nil
}

type FindWalletByUserIDResponse struct {
	Balances             []*BalanceDTO `protobuf:"bytes,1,rep,name=balances,proto3" json:"balances"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *FindWalletByUserIDResponse) Reset()         { *m = FindWalletByUserIDResponse{} }
func (m *FindWalletByUserIDResponse) String() string { return proto.CompactTextString(m) }
func (*FindWalletByUserIDResponse) ProtoMessage()    {}
func (*FindWalletByUserIDResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_aa1d687600079dba, []int{10}
}

func (m *FindWalletByUserIDResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindWalletByUserIDResponse.Unmarshal(m, b)
}
func (m *FindWalletByUserIDResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindWalletByUserIDResponse.Marshal(b, m, deterministic)
}
func (m *FindWalletByUserIDResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindWalletByUserIDResponse.Merge(m, src)
}
func (m *FindWalletByUserIDResponse) XXX_Size() int {
	return xxx_messageInfo_FindWalletByUserIDResponse.Size(m)
}
func (m *FindWalletByUserIDResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FindWalletByUserIDResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FindWalletByUserIDResponse proto.InternalMessageInfo

func (m *FindWalletByUserIDResponse) GetBalances() []*BalanceDTO {
	if m != nil {
		return m.Balances
	}
	return nil
}

type FreezeResponse struct {
	Done                 bool     `protobuf:"varint,1,opt,name=done,proto3" json:"done"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FreezeResponse) Reset()         { *m = FreezeResponse{} }
func (m *FreezeResponse) String() string { return proto.CompactTextString(m) }
func (*FreezeResponse) ProtoMessage()    {}
func (*FreezeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_aa1d687600079dba, []int{11}
}

func (m *FreezeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FreezeResponse.Unmarshal(m, b)
}
func (m *FreezeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FreezeResponse.Marshal(b, m, deterministic)
}
func (m *FreezeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FreezeResponse.Merge(m, src)
}
func (m *FreezeResponse) XXX_Size() int {
	return xxx_messageInfo_FreezeResponse.Size(m)
}
func (m *FreezeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FreezeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FreezeResponse proto.InternalMessageInfo

func (m *FreezeResponse) GetDone() bool {
	if m != nil {
		return m.Done
	}
	return false
}

type ReleaseFreezeResponse struct {
	Done                 bool     `protobuf:"varint,1,opt,name=done,proto3" json:"done"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReleaseFreezeResponse) Reset()         { *m = ReleaseFreezeResponse{} }
func (m *ReleaseFreezeResponse) String() string { return proto.CompactTextString(m) }
func (*ReleaseFreezeResponse) ProtoMessage()    {}
func (*ReleaseFreezeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_aa1d687600079dba, []int{12}
}

func (m *ReleaseFreezeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReleaseFreezeResponse.Unmarshal(m, b)
}
func (m *ReleaseFreezeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReleaseFreezeResponse.Marshal(b, m, deterministic)
}
func (m *ReleaseFreezeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReleaseFreezeResponse.Merge(m, src)
}
func (m *ReleaseFreezeResponse) XXX_Size() int {
	return xxx_messageInfo_ReleaseFreezeResponse.Size(m)
}
func (m *ReleaseFreezeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ReleaseFreezeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ReleaseFreezeResponse proto.InternalMessageInfo

func (m *ReleaseFreezeResponse) GetDone() bool {
	if m != nil {
		return m.Done
	}
	return false
}

type TransferMoneyResponse struct {
	Done                 bool     `protobuf:"varint,1,opt,name=done,proto3" json:"done"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TransferMoneyResponse) Reset()         { *m = TransferMoneyResponse{} }
func (m *TransferMoneyResponse) String() string { return proto.CompactTextString(m) }
func (*TransferMoneyResponse) ProtoMessage()    {}
func (*TransferMoneyResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_aa1d687600079dba, []int{13}
}

func (m *TransferMoneyResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TransferMoneyResponse.Unmarshal(m, b)
}
func (m *TransferMoneyResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TransferMoneyResponse.Marshal(b, m, deterministic)
}
func (m *TransferMoneyResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransferMoneyResponse.Merge(m, src)
}
func (m *TransferMoneyResponse) XXX_Size() int {
	return xxx_messageInfo_TransferMoneyResponse.Size(m)
}
func (m *TransferMoneyResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TransferMoneyResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TransferMoneyResponse proto.InternalMessageInfo

func (m *TransferMoneyResponse) GetDone() bool {
	if m != nil {
		return m.Done
	}
	return false
}

func init() {
	proto.RegisterType((*ReleaseFreezeRequest)(nil), "service_core.ReleaseFreezeRequest")
	proto.RegisterType((*FindWalletByIDRequest)(nil), "service_core.FindWalletByIDRequest")
	proto.RegisterType((*FindOrCreateByUserIDRequest)(nil), "service_core.FindOrCreateByUserIDRequest")
	proto.RegisterType((*FindWalletByUserIDRequest)(nil), "service_core.FindWalletByUserIDRequest")
	proto.RegisterType((*FreezeRequest)(nil), "service_core.FreezeRequest")
	proto.RegisterType((*TransferRequest)(nil), "service_core.TransferRequest")
	proto.RegisterType((*FindWalletByUserIDTypeRequest)(nil), "service_core.FindWalletByUserIDTypeRequest")
	proto.RegisterType((*FindWalletByUserIDTypeResponse)(nil), "service_core.FindWalletByUserIDTypeResponse")
	proto.RegisterType((*FindWalletByIDResponse)(nil), "service_core.FindWalletByIDResponse")
	proto.RegisterType((*FindOrCreateWalletByUserIDResponse)(nil), "service_core.FindOrCreateWalletByUserIDResponse")
	proto.RegisterType((*FindWalletByUserIDResponse)(nil), "service_core.FindWalletByUserIDResponse")
	proto.RegisterType((*FreezeResponse)(nil), "service_core.FreezeResponse")
	proto.RegisterType((*ReleaseFreezeResponse)(nil), "service_core.ReleaseFreezeResponse")
	proto.RegisterType((*TransferMoneyResponse)(nil), "service_core.TransferMoneyResponse")
}

func init() {
	proto.RegisterFile("service_core/service.proto", fileDescriptor_aa1d687600079dba)
}

var fileDescriptor_aa1d687600079dba = []byte{
	// 566 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0xdd, 0x6e, 0x12, 0x41,
	0x14, 0x06, 0x41, 0xa0, 0x87, 0x82, 0xc9, 0xc4, 0x56, 0x5c, 0xac, 0x36, 0x53, 0x8c, 0x35, 0x35,
	0x68, 0x90, 0x5b, 0x2f, 0x44, 0x6c, 0x42, 0xa2, 0x69, 0xb2, 0xa5, 0x31, 0x36, 0x31, 0x64, 0xbb,
	0x7b, 0x4c, 0x88, 0x38, 0x8b, 0x33, 0x8b, 0xba, 0xfa, 0x26, 0x3e, 0x8f, 0x0f, 0xd6, 0xec, 0xce,
	0xce, 0xfe, 0x31, 0x2c, 0x4d, 0x7a, 0xc7, 0xcc, 0x7c, 0x3f, 0x67, 0xe6, 0x7c, 0x7b, 0x00, 0x43,
	0x20, 0xff, 0x39, 0xb7, 0x71, 0x66, 0xbb, 0x1c, 0x5f, 0x46, 0x8b, 0xfe, 0x92, 0xbb, 0x9e, 0x4b,
	0x76, 0xd3, 0x67, 0x46, 0x16, 0x79, 0x65, 0x2d, 0x2c, 0xa6, 0x90, 0xf4, 0x0d, 0xdc, 0x37, 0x71,
	0x81, 0x96, 0xc0, 0x53, 0x8e, 0xf8, 0x07, 0x4d, 0xfc, 0xb1, 0x42, 0xe1, 0x91, 0xa7, 0xd0, 0xf6,
	0xb8, 0xc5, 0x84, 0x65, 0x7b, 0x73, 0x97, 0xcd, 0xe6, 0x4e, 0xa7, 0x7c, 0x58, 0x3e, 0xde, 0x31,
	0x5b, 0xa9, 0xdd, 0x89, 0x43, 0x87, 0xb0, 0x77, 0x3a, 0x67, 0xce, 0x27, 0x6b, 0xb1, 0x40, 0x6f,
	0xe4, 0x4f, 0xc6, 0x8a, 0xdf, 0x85, 0x9d, 0x5f, 0xe1, 0x66, 0x42, 0x6d, 0xc8, 0x8d, 0x89, 0x43,
	0x4d, 0xe8, 0x06, 0xac, 0x33, 0xfe, 0x8e, 0xa3, 0xe5, 0xe1, 0xc8, 0xbf, 0x10, 0xc8, 0x13, 0xee,
	0x03, 0xa8, 0xaf, 0x04, 0xf2, 0x84, 0x59, 0x0b, 0x96, 0x13, 0x87, 0x18, 0xd0, 0xb0, 0x57, 0x9c,
	0x23, 0xb3, 0xfd, 0xce, 0x1d, 0xa9, 0xa9, 0xd6, 0x74, 0x08, 0x0f, 0xd3, 0x95, 0xdc, 0x4c, 0x91,
	0x7e, 0x83, 0x56, 0xf6, 0xde, 0x45, 0x75, 0x93, 0x7d, 0xa8, 0x59, 0xdf, 0xdd, 0x15, 0xf3, 0x42,
	0xf7, 0xaa, 0x19, 0xad, 0x34, 0x8f, 0x55, 0xd1, 0x3d, 0xd6, 0xbf, 0x32, 0xdc, 0x9b, 0x06, 0x3b,
	0x5f, 0x91, 0x2b, 0x3f, 0x0a, 0xad, 0xd8, 0x6f, 0x26, 0xb8, 0x1d, 0x79, 0x36, 0x95, 0xe7, 0x39,
	0xb7, 0x49, 0x0f, 0xda, 0x09, 0xc6, 0x41, 0xe1, 0x45, 0x97, 0xdf, 0x55, 0xa0, 0x71, 0xa0, 0x94,
	0x14, 0x57, 0xd9, 0x52, 0x5c, 0x55, 0x57, 0xdc, 0x67, 0x38, 0x58, 0x7f, 0xbf, 0xa9, 0xbf, 0xc4,
	0xad, 0x5d, 0x79, 0x02, 0x4d, 0xcf, 0x5f, 0xe2, 0x4c, 0x56, 0x13, 0xd5, 0x06, 0xc1, 0x96, 0x14,
	0xa3, 0x53, 0x78, 0xbc, 0x49, 0x5a, 0x2c, 0x5d, 0x26, 0x90, 0x0c, 0xa0, 0x1e, 0xc5, 0x32, 0xd4,
	0x6e, 0x0e, 0x3a, 0xfd, 0x74, 0x66, 0xfb, 0x23, 0x79, 0x38, 0x9e, 0x9e, 0x99, 0x0a, 0x48, 0x3f,
	0xc0, 0x7e, 0x3e, 0x7a, 0xb7, 0x50, 0xbb, 0x04, 0x9a, 0x8e, 0x64, 0x3e, 0x46, 0x91, 0xf2, 0x10,
	0x1a, 0x11, 0x41, 0x74, 0xca, 0x87, 0x95, 0x42, 0xe9, 0x18, 0x49, 0x4d, 0x30, 0x74, 0xd1, 0xbc,
	0x95, 0x66, 0x0f, 0xda, 0x2a, 0xb8, 0x91, 0x0e, 0x81, 0xaa, 0xe3, 0x32, 0x79, 0xe5, 0x86, 0x19,
	0xfe, 0xa6, 0x27, 0xb0, 0x97, 0xfb, 0xba, 0x8b, 0xc1, 0x2a, 0x9d, 0x1f, 0x5d, 0x86, 0x7e, 0x11,
	0x78, 0xf0, 0xff, 0x2e, 0xd4, 0xcf, 0x65, 0x95, 0xe4, 0x0b, 0xb4, 0xb3, 0x9d, 0x20, 0x47, 0xd9,
	0x1b, 0x68, 0x47, 0x84, 0xd1, 0x2b, 0x06, 0x49, 0x73, 0x5a, 0x22, 0xbf, 0x75, 0x5f, 0xf6, 0x5b,
	0xe6, 0x04, 0x09, 0x22, 0x27, 0x9b, 0x45, 0xd6, 0x22, 0x6c, 0xbc, 0xb8, 0x19, 0x38, 0x76, 0xfe,
	0x2b, 0x1b, 0xa7, 0x0f, 0x05, 0x79, 0xbe, 0xae, 0xb6, 0x61, 0xa2, 0x19, 0xaf, 0x36, 0x43, 0xf5,
	0xa9, 0xa0, 0x25, 0xf2, 0x1e, 0x6a, 0xb2, 0x69, 0xa4, 0x9b, 0x63, 0xa7, 0x07, 0x96, 0xf1, 0x48,
	0x7f, 0x18, 0xcb, 0x5c, 0x42, 0x2b, 0x13, 0x01, 0x42, 0xb3, 0x04, 0xdd, 0xf4, 0x37, 0x8e, 0x0a,
	0x31, 0xb1, 0xf6, 0x05, 0xb4, 0x32, 0x89, 0x21, 0x07, 0x59, 0x5e, 0x6e, 0xd8, 0xe5, 0x65, 0xb5,
	0x69, 0xa3, 0x25, 0x32, 0x07, 0xb2, 0xde, 0x1a, 0xf2, 0x6c, 0x5b, 0xf3, 0x94, 0xcb, 0xf1, 0x76,
	0xa0, 0xb2, 0xba, 0xaa, 0x85, 0xff, 0x82, 0xaf, 0xaf, 0x03, 0x00, 0x00, 0xff, 0xff, 0x14, 0x53,
	0xa2, 0x34, 0x4d, 0x07, 0x00, 0x00,
}
