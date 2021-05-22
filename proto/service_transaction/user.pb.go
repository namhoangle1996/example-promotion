// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service_transaction/user.proto

package service_transaction

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

type UserDTO struct {
	Id                           string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	IdGutina                     int64    `protobuf:"varint,2,opt,name=idGutina,proto3" json:"idGutina,omitempty"`
	PhoneNumber                  string   `protobuf:"bytes,3,opt,name=phoneNumber,proto3" json:"phoneNumber,omitempty"`
	Email                        string   `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Name                         string   `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	Avatar                       string   `protobuf:"bytes,6,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Gender                       string   `protobuf:"bytes,7,opt,name=gender,proto3" json:"gender,omitempty"`
	Status                       string   `protobuf:"bytes,8,opt,name=status,proto3" json:"status,omitempty"`
	Timezone                     string   `protobuf:"bytes,9,opt,name=timezone,proto3" json:"timezone,omitempty"`
	Language                     string   `protobuf:"bytes,10,opt,name=language,proto3" json:"language,omitempty"`
	Title                        string   `protobuf:"bytes,11,opt,name=title,proto3" json:"title,omitempty"`
	DateTime                     int64    `protobuf:"varint,12,opt,name=dateTime,proto3" json:"dateTime,omitempty"`
	Currency                     string   `protobuf:"bytes,13,opt,name=currency,proto3" json:"currency,omitempty"`
	Role                         string   `protobuf:"bytes,14,opt,name=role,proto3" json:"role,omitempty"`
	IsFacebookConnect            bool     `protobuf:"varint,15,opt,name=isFacebookConnect,proto3" json:"isFacebookConnect,omitempty"`
	EmailVerify                  bool     `protobuf:"varint,16,opt,name=emailVerify,proto3" json:"emailVerify,omitempty"`
	CanUpdateEmail               bool     `protobuf:"varint,17,opt,name=canUpdateEmail,proto3" json:"canUpdateEmail,omitempty"`
	CanAddBankAccount            bool     `protobuf:"varint,18,opt,name=canAddBankAccount,proto3" json:"canAddBankAccount,omitempty"`
	CanUpdatePhoneNumber         bool     `protobuf:"varint,19,opt,name=canUpdatePhoneNumber,proto3" json:"canUpdatePhoneNumber,omitempty"`
	PhoneNumberVerify            bool     `protobuf:"varint,20,opt,name=phoneNumberVerify,proto3" json:"phoneNumberVerify,omitempty"`
	AccessToken                  string   `protobuf:"bytes,21,opt,name=accessToken,proto3" json:"accessToken,omitempty"`
	IsExists                     bool     `protobuf:"varint,22,opt,name=isExists,proto3" json:"isExists,omitempty"`
	IdentityNumber               string   `protobuf:"bytes,23,opt,name=identityNumber,proto3" json:"identityNumber,omitempty"`
	BirthDay                     string   `protobuf:"bytes,24,opt,name=birthDay,proto3" json:"birthDay,omitempty"`
	ConnectedFacebook            bool     `protobuf:"varint,25,opt,name=connectedFacebook,proto3" json:"connectedFacebook,omitempty"`
	Address                      string   `protobuf:"bytes,26,opt,name=address,proto3" json:"address,omitempty"`
	MessageTopics                []string `protobuf:"bytes,27,rep,name=messageTopics,proto3" json:"messageTopics,omitempty"`
	DeviceIdLogin                string   `protobuf:"bytes,28,opt,name=deviceIdLogin,proto3" json:"deviceIdLogin,omitempty"`
	LastDeviceIdLogin            string   `protobuf:"bytes,29,opt,name=lastDeviceIdLogin,proto3" json:"lastDeviceIdLogin,omitempty"`
	Password                     string   `protobuf:"bytes,30,opt,name=password,proto3" json:"password,omitempty"`
	Long                         float64  `protobuf:"fixed64,31,opt,name=long,proto3" json:"long,omitempty"`
	Lat                          float64  `protobuf:"fixed64,32,opt,name=lat,proto3" json:"lat,omitempty"`
	OriginDevice                 string   `protobuf:"bytes,33,opt,name=originDevice,proto3" json:"originDevice,omitempty"`
	Source                       string   `protobuf:"bytes,34,opt,name=source,proto3" json:"source,omitempty"`
	RandomString                 string   `protobuf:"bytes,35,opt,name=randomString,proto3" json:"randomString,omitempty"`
	LastAmountNetPrimaryWallet   int64    `protobuf:"varint,36,opt,name=lastAmountNetPrimaryWallet,proto3" json:"lastAmountNetPrimaryWallet,omitempty"`
	LastAmountPrimaryWallet      int64    `protobuf:"varint,37,opt,name=lastAmountPrimaryWallet,proto3" json:"lastAmountPrimaryWallet,omitempty"`
	IncrementAmountPrimaryWallet int64    `protobuf:"varint,38,opt,name=incrementAmountPrimaryWallet,proto3" json:"incrementAmountPrimaryWallet,omitempty"`
	DecrementAmountPrimaryWallet int64    `protobuf:"varint,39,opt,name=decrementAmountPrimaryWallet,proto3" json:"decrementAmountPrimaryWallet,omitempty"`
	XXX_NoUnkeyedLiteral         struct{} `json:"-"`
	XXX_unrecognized             []byte   `json:"-"`
	XXX_sizecache                int32    `json:"-"`
}

func (m *UserDTO) Reset()         { *m = UserDTO{} }
func (m *UserDTO) String() string { return proto.CompactTextString(m) }
func (*UserDTO) ProtoMessage()    {}
func (*UserDTO) Descriptor() ([]byte, []int) {
	return fileDescriptor_user_9ea8d2c55d42c72c, []int{0}
}
func (m *UserDTO) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserDTO.Unmarshal(m, b)
}
func (m *UserDTO) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserDTO.Marshal(b, m, deterministic)
}
func (dst *UserDTO) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserDTO.Merge(dst, src)
}
func (m *UserDTO) XXX_Size() int {
	return xxx_messageInfo_UserDTO.Size(m)
}
func (m *UserDTO) XXX_DiscardUnknown() {
	xxx_messageInfo_UserDTO.DiscardUnknown(m)
}

var xxx_messageInfo_UserDTO proto.InternalMessageInfo

func (m *UserDTO) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *UserDTO) GetIdGutina() int64 {
	if m != nil {
		return m.IdGutina
	}
	return 0
}

func (m *UserDTO) GetPhoneNumber() string {
	if m != nil {
		return m.PhoneNumber
	}
	return ""
}

func (m *UserDTO) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *UserDTO) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UserDTO) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

func (m *UserDTO) GetGender() string {
	if m != nil {
		return m.Gender
	}
	return ""
}

func (m *UserDTO) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *UserDTO) GetTimezone() string {
	if m != nil {
		return m.Timezone
	}
	return ""
}

func (m *UserDTO) GetLanguage() string {
	if m != nil {
		return m.Language
	}
	return ""
}

func (m *UserDTO) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *UserDTO) GetDateTime() int64 {
	if m != nil {
		return m.DateTime
	}
	return 0
}

func (m *UserDTO) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

func (m *UserDTO) GetRole() string {
	if m != nil {
		return m.Role
	}
	return ""
}

func (m *UserDTO) GetIsFacebookConnect() bool {
	if m != nil {
		return m.IsFacebookConnect
	}
	return false
}

func (m *UserDTO) GetEmailVerify() bool {
	if m != nil {
		return m.EmailVerify
	}
	return false
}

func (m *UserDTO) GetCanUpdateEmail() bool {
	if m != nil {
		return m.CanUpdateEmail
	}
	return false
}

func (m *UserDTO) GetCanAddBankAccount() bool {
	if m != nil {
		return m.CanAddBankAccount
	}
	return false
}

func (m *UserDTO) GetCanUpdatePhoneNumber() bool {
	if m != nil {
		return m.CanUpdatePhoneNumber
	}
	return false
}

func (m *UserDTO) GetPhoneNumberVerify() bool {
	if m != nil {
		return m.PhoneNumberVerify
	}
	return false
}

func (m *UserDTO) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *UserDTO) GetIsExists() bool {
	if m != nil {
		return m.IsExists
	}
	return false
}

func (m *UserDTO) GetIdentityNumber() string {
	if m != nil {
		return m.IdentityNumber
	}
	return ""
}

func (m *UserDTO) GetBirthDay() string {
	if m != nil {
		return m.BirthDay
	}
	return ""
}

func (m *UserDTO) GetConnectedFacebook() bool {
	if m != nil {
		return m.ConnectedFacebook
	}
	return false
}

func (m *UserDTO) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *UserDTO) GetMessageTopics() []string {
	if m != nil {
		return m.MessageTopics
	}
	return nil
}

func (m *UserDTO) GetDeviceIdLogin() string {
	if m != nil {
		return m.DeviceIdLogin
	}
	return ""
}

func (m *UserDTO) GetLastDeviceIdLogin() string {
	if m != nil {
		return m.LastDeviceIdLogin
	}
	return ""
}

func (m *UserDTO) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *UserDTO) GetLong() float64 {
	if m != nil {
		return m.Long
	}
	return 0
}

func (m *UserDTO) GetLat() float64 {
	if m != nil {
		return m.Lat
	}
	return 0
}

func (m *UserDTO) GetOriginDevice() string {
	if m != nil {
		return m.OriginDevice
	}
	return ""
}

func (m *UserDTO) GetSource() string {
	if m != nil {
		return m.Source
	}
	return ""
}

func (m *UserDTO) GetRandomString() string {
	if m != nil {
		return m.RandomString
	}
	return ""
}

func (m *UserDTO) GetLastAmountNetPrimaryWallet() int64 {
	if m != nil {
		return m.LastAmountNetPrimaryWallet
	}
	return 0
}

func (m *UserDTO) GetLastAmountPrimaryWallet() int64 {
	if m != nil {
		return m.LastAmountPrimaryWallet
	}
	return 0
}

func (m *UserDTO) GetIncrementAmountPrimaryWallet() int64 {
	if m != nil {
		return m.IncrementAmountPrimaryWallet
	}
	return 0
}

func (m *UserDTO) GetDecrementAmountPrimaryWallet() int64 {
	if m != nil {
		return m.DecrementAmountPrimaryWallet
	}
	return 0
}

func init() {
	proto.RegisterType((*UserDTO)(nil), "service_transaction.UserDTO")
}

func init() {
	proto.RegisterFile("service_transaction/user.proto", fileDescriptor_user_9ea8d2c55d42c72c)
}

var fileDescriptor_user_9ea8d2c55d42c72c = []byte{
	// 644 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x54, 0xdd, 0x6e, 0x13, 0x3d,
	0x10, 0x55, 0x9a, 0xfe, 0xba, 0x3f, 0x5f, 0xeb, 0xf6, 0x6b, 0x87, 0x52, 0xca, 0x52, 0x4a, 0xc9,
	0x05, 0x02, 0x09, 0x6e, 0xb8, 0x42, 0x6a, 0x69, 0x41, 0x48, 0xa8, 0x54, 0x21, 0x85, 0x4b, 0xe4,
	0xd8, 0xc3, 0xd6, 0xea, 0xae, 0x1d, 0xd9, 0xde, 0x42, 0x78, 0x4d, 0x5e, 0x08, 0x79, 0x9c, 0x0d,
	0x49, 0x53, 0x72, 0xe7, 0x73, 0xce, 0xcc, 0xf1, 0xcc, 0xac, 0x67, 0xd9, 0xbe, 0x47, 0x77, 0xa3,
	0x25, 0x7e, 0x0b, 0x4e, 0x18, 0x2f, 0x64, 0xd0, 0xd6, 0xbc, 0xa8, 0x3c, 0xba, 0xe7, 0x3d, 0x67,
	0x83, 0xe5, 0x9b, 0x77, 0xe8, 0x07, 0xbf, 0x19, 0x5b, 0xb8, 0xf4, 0xe8, 0x4e, 0x3b, 0x9f, 0xf8,
	0x1a, 0x9b, 0xd1, 0x0a, 0x1a, 0x59, 0xa3, 0xb5, 0xd4, 0x9e, 0xd1, 0x8a, 0xef, 0xb2, 0x45, 0xad,
	0xde, 0x57, 0x41, 0x1b, 0x01, 0x33, 0x59, 0xa3, 0xd5, 0x6c, 0x0f, 0x31, 0xcf, 0xd8, 0x72, 0xef,
	0xca, 0x1a, 0x3c, 0xaf, 0xca, 0x2e, 0x3a, 0x68, 0x52, 0xd2, 0x28, 0xc5, 0xb7, 0xd8, 0x1c, 0x96,
	0x42, 0x17, 0x30, 0x4b, 0x5a, 0x02, 0x9c, 0xb3, 0x59, 0x23, 0x4a, 0x84, 0x39, 0x22, 0xe9, 0xcc,
	0xb7, 0xd9, 0xbc, 0xb8, 0x11, 0x41, 0x38, 0x98, 0x27, 0x76, 0x80, 0x22, 0x9f, 0xa3, 0x51, 0xe8,
	0x60, 0x21, 0xf1, 0x09, 0x45, 0xde, 0x07, 0x11, 0x2a, 0x0f, 0x8b, 0x89, 0x4f, 0x28, 0xd6, 0x1b,
	0x74, 0x89, 0xbf, 0xac, 0x41, 0x58, 0x22, 0x65, 0x88, 0xa3, 0x56, 0x08, 0x93, 0x57, 0x22, 0x47,
	0x60, 0x49, 0xab, 0x71, 0xac, 0x34, 0xe8, 0x50, 0x20, 0x2c, 0xa7, 0x4a, 0x09, 0xc4, 0x0c, 0x25,
	0x02, 0x76, 0x74, 0x89, 0xb0, 0x92, 0xba, 0xaf, 0x71, 0xd4, 0x64, 0xe5, 0x1c, 0x1a, 0xd9, 0x87,
	0xd5, 0xe4, 0x56, 0xe3, 0xd8, 0xa1, 0xb3, 0x05, 0xc2, 0x5a, 0xea, 0x30, 0x9e, 0xf9, 0x33, 0xb6,
	0xa1, 0xfd, 0x3b, 0x21, 0xb1, 0x6b, 0xed, 0xf5, 0x5b, 0x6b, 0x0c, 0xca, 0x00, 0xff, 0x65, 0x8d,
	0xd6, 0x62, 0x7b, 0x52, 0x88, 0xb3, 0xa5, 0x61, 0x7d, 0x41, 0xa7, 0xbf, 0xf7, 0x61, 0x9d, 0xe2,
	0x46, 0x29, 0x7e, 0xc4, 0xd6, 0xa4, 0x30, 0x97, 0xbd, 0x58, 0xd0, 0x19, 0x0d, 0x79, 0x83, 0x82,
	0x6e, 0xb1, 0xf1, 0x5e, 0x29, 0xcc, 0xb1, 0x52, 0x27, 0xc2, 0x5c, 0x1f, 0x4b, 0x69, 0x2b, 0x13,
	0x80, 0xa7, 0x7b, 0x27, 0x04, 0xfe, 0x92, 0x6d, 0x0d, 0xf3, 0x2f, 0x46, 0x3e, 0xee, 0x26, 0x25,
	0xdc, 0xa9, 0xc5, 0x1b, 0x46, 0x3e, 0xfa, 0xa0, 0xe2, 0xad, 0x74, 0xc3, 0x84, 0x10, 0x3b, 0x13,
	0x52, 0xa2, 0xf7, 0x1d, 0x7b, 0x8d, 0x06, 0xfe, 0x4f, 0xaf, 0x66, 0x84, 0xa2, 0x37, 0xe7, 0xcf,
	0x7e, 0x6a, 0x1f, 0x3c, 0x6c, 0x93, 0xcd, 0x10, 0xc7, 0xae, 0xb5, 0x42, 0x13, 0x74, 0xe8, 0x0f,
	0x2a, 0xdb, 0x21, 0x83, 0x5b, 0x6c, 0xf4, 0xe8, 0x6a, 0x17, 0xae, 0x4e, 0x45, 0x1f, 0x20, 0x7d,
	0x9d, 0x1a, 0xd3, 0x44, 0xd2, 0x98, 0x51, 0xd5, 0x73, 0x87, 0x7b, 0x83, 0x89, 0xdc, 0x16, 0x38,
	0xb0, 0x05, 0xa1, 0x94, 0x43, 0xef, 0x61, 0x97, 0x8c, 0x6a, 0xc8, 0x0f, 0xd9, 0x6a, 0x89, 0xde,
	0x8b, 0x1c, 0x3b, 0xb6, 0xa7, 0xa5, 0x87, 0xfb, 0x59, 0xb3, 0xb5, 0xd4, 0x1e, 0x27, 0x63, 0x94,
	0xc2, 0xb8, 0x73, 0x1f, 0xd4, 0x47, 0x9b, 0x6b, 0x03, 0x7b, 0xe4, 0x32, 0x4e, 0xc6, 0x9a, 0x0a,
	0xe1, 0xc3, 0xe9, 0x58, 0xe4, 0x03, 0x8a, 0x9c, 0x14, 0x62, 0x77, 0x3d, 0xe1, 0xfd, 0x0f, 0xeb,
	0x14, 0xec, 0xa7, 0xee, 0x6a, 0x1c, 0xdf, 0x5e, 0x61, 0x4d, 0x0e, 0x0f, 0xb3, 0x46, 0xab, 0xd1,
	0xa6, 0x33, 0x5f, 0x67, 0xcd, 0x42, 0x04, 0xc8, 0x88, 0x8a, 0x47, 0x7e, 0xc0, 0x56, 0xac, 0xd3,
	0xb9, 0x36, 0xc9, 0x18, 0x1e, 0x91, 0xcb, 0x18, 0x47, 0x3b, 0x66, 0x2b, 0x27, 0x11, 0x0e, 0x06,
	0x3b, 0x46, 0x28, 0xe6, 0x3a, 0x61, 0x94, 0x2d, 0x3f, 0x07, 0xa7, 0x4d, 0x0e, 0x8f, 0x53, 0xee,
	0x28, 0xc7, 0xdf, 0xb0, 0xdd, 0x58, 0xf6, 0x71, 0x19, 0x5f, 0xd5, 0x39, 0x86, 0x0b, 0xa7, 0x4b,
	0xe1, 0xfa, 0x5f, 0x45, 0x51, 0x60, 0x80, 0x43, 0xda, 0xa5, 0x29, 0x11, 0xfc, 0x35, 0xdb, 0xf9,
	0xab, 0x8e, 0x27, 0x3f, 0xa1, 0xe4, 0x7f, 0xc9, 0xfc, 0x84, 0xed, 0x69, 0x23, 0x1d, 0x96, 0x68,
	0xee, 0x4c, 0x3f, 0xa2, 0xf4, 0xa9, 0x31, 0xd1, 0x43, 0xe1, 0x14, 0x8f, 0xa7, 0xc9, 0x63, 0x5a,
	0x4c, 0x77, 0x9e, 0xfe, 0xb8, 0xaf, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0x96, 0xeb, 0x69, 0xcd,
	0x93, 0x05, 0x00, 0x00,
}
