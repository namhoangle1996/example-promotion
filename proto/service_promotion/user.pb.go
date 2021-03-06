// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service_promotion/user.proto

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
	return fileDescriptor_user_8869b8553dae6b27, []int{0}
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
	proto.RegisterType((*UserDTO)(nil), "service_promotion.UserDTO")
}

func init() { proto.RegisterFile("service_promotion/user.proto", fileDescriptor_user_8869b8553dae6b27) }

var fileDescriptor_user_8869b8553dae6b27 = []byte{
	// 642 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x54, 0x5b, 0x6f, 0x13, 0x3d,
	0x10, 0x55, 0x9a, 0x5e, 0xdd, 0xcb, 0xd7, 0xfa, 0x2b, 0xed, 0x50, 0x0a, 0x2c, 0xa5, 0x94, 0x3c,
	0x20, 0x90, 0xe0, 0x85, 0x27, 0xa4, 0x96, 0x16, 0x84, 0x84, 0x4a, 0x15, 0x52, 0x78, 0x44, 0x8e,
	0x3d, 0x6c, 0xad, 0xee, 0xda, 0x91, 0xed, 0x2d, 0x84, 0xbf, 0xc9, 0x1f, 0x42, 0x1e, 0x67, 0x43,
	0x2e, 0x25, 0x6f, 0x3e, 0xe7, 0xcc, 0x1c, 0xcf, 0xcc, 0x7a, 0x96, 0xed, 0x7b, 0x74, 0x37, 0x5a,
	0xe2, 0xb7, 0x9e, 0xb3, 0xa5, 0x0d, 0xda, 0x9a, 0x17, 0x95, 0x47, 0xf7, 0xbc, 0xe7, 0x6c, 0xb0,
	0x7c, 0x6b, 0x4a, 0x3d, 0xf8, 0xcd, 0xd8, 0xd2, 0xa5, 0x47, 0x77, 0xda, 0xf9, 0xc4, 0x37, 0xd8,
	0x9c, 0x56, 0xd0, 0xc8, 0x1a, 0xad, 0x95, 0xf6, 0x9c, 0x56, 0x7c, 0x8f, 0x2d, 0x6b, 0xf5, 0xbe,
	0x0a, 0xda, 0x08, 0x98, 0xcb, 0x1a, 0xad, 0x66, 0x7b, 0x88, 0x79, 0xc6, 0x56, 0x7b, 0x57, 0xd6,
	0xe0, 0x79, 0x55, 0x76, 0xd1, 0x41, 0x93, 0x92, 0x46, 0x29, 0xbe, 0xcd, 0x16, 0xb0, 0x14, 0xba,
	0x80, 0x79, 0xd2, 0x12, 0xe0, 0x9c, 0xcd, 0x1b, 0x51, 0x22, 0x2c, 0x10, 0x49, 0x67, 0xbe, 0xc3,
	0x16, 0xc5, 0x8d, 0x08, 0xc2, 0xc1, 0x22, 0xb1, 0x03, 0x14, 0xf9, 0x1c, 0x8d, 0x42, 0x07, 0x4b,
	0x89, 0x4f, 0x28, 0xf2, 0x3e, 0x88, 0x50, 0x79, 0x58, 0x4e, 0x7c, 0x42, 0xb1, 0xde, 0xa0, 0x4b,
	0xfc, 0x65, 0x0d, 0xc2, 0x0a, 0x29, 0x43, 0x1c, 0xb5, 0x42, 0x98, 0xbc, 0x12, 0x39, 0x02, 0x4b,
	0x5a, 0x8d, 0x63, 0xa5, 0x41, 0x87, 0x02, 0x61, 0x35, 0x55, 0x4a, 0x20, 0x66, 0x28, 0x11, 0xb0,
	0xa3, 0x4b, 0x84, 0xb5, 0xd4, 0x7d, 0x8d, 0xa3, 0x26, 0x2b, 0xe7, 0xd0, 0xc8, 0x3e, 0xac, 0x27,
	0xb7, 0x1a, 0xc7, 0x0e, 0x9d, 0x2d, 0x10, 0x36, 0x52, 0x87, 0xf1, 0xcc, 0x9f, 0xb1, 0x2d, 0xed,
	0xdf, 0x09, 0x89, 0x5d, 0x6b, 0xaf, 0xdf, 0x5a, 0x63, 0x50, 0x06, 0xf8, 0x2f, 0x6b, 0xb4, 0x96,
	0xdb, 0xd3, 0x42, 0x9c, 0x2d, 0x0d, 0xeb, 0x0b, 0x3a, 0xfd, 0xbd, 0x0f, 0x9b, 0x14, 0x37, 0x4a,
	0xf1, 0x23, 0xb6, 0x21, 0x85, 0xb9, 0xec, 0xc5, 0x82, 0xce, 0x68, 0xc8, 0x5b, 0x14, 0x34, 0xc1,
	0xc6, 0x7b, 0xa5, 0x30, 0xc7, 0x4a, 0x9d, 0x08, 0x73, 0x7d, 0x2c, 0xa5, 0xad, 0x4c, 0x00, 0x9e,
	0xee, 0x9d, 0x12, 0xf8, 0x4b, 0xb6, 0x3d, 0xcc, 0xbf, 0x18, 0xf9, 0xb8, 0xff, 0x53, 0xc2, 0xad,
	0x5a, 0xbc, 0x61, 0xe4, 0xa3, 0x0f, 0x2a, 0xde, 0x4e, 0x37, 0x4c, 0x09, 0xb1, 0x33, 0x21, 0x25,
	0x7a, 0xdf, 0xb1, 0xd7, 0x68, 0xe0, 0x4e, 0x7a, 0x35, 0x23, 0x14, 0xbd, 0x39, 0x7f, 0xf6, 0x53,
	0xfb, 0xe0, 0x61, 0x87, 0x6c, 0x86, 0x38, 0x76, 0xad, 0x15, 0x9a, 0xa0, 0x43, 0x7f, 0x50, 0xd9,
	0x2e, 0x19, 0x4c, 0xb0, 0xd1, 0xa3, 0xab, 0x5d, 0xb8, 0x3a, 0x15, 0x7d, 0x80, 0xf4, 0x75, 0x6a,
	0x4c, 0x13, 0x49, 0x63, 0x46, 0x55, 0xcf, 0x1d, 0xee, 0x0e, 0x26, 0x32, 0x29, 0x70, 0x60, 0x4b,
	0x42, 0x29, 0x87, 0xde, 0xc3, 0x1e, 0x19, 0xd5, 0x90, 0x1f, 0xb2, 0xf5, 0x12, 0xbd, 0x17, 0x39,
	0x76, 0x6c, 0x4f, 0x4b, 0x0f, 0xf7, 0xb2, 0x66, 0x6b, 0xa5, 0x3d, 0x4e, 0xc6, 0x28, 0x85, 0x71,
	0xe3, 0x3e, 0xa8, 0x8f, 0x36, 0xd7, 0x06, 0xf6, 0xc9, 0x65, 0x9c, 0x8c, 0x35, 0x15, 0xc2, 0x87,
	0xd3, 0xb1, 0xc8, 0xfb, 0x14, 0x39, 0x2d, 0xc4, 0xee, 0x7a, 0xc2, 0xfb, 0x1f, 0xd6, 0x29, 0x78,
	0x90, 0xba, 0xab, 0x71, 0x7c, 0x7b, 0x85, 0x35, 0x39, 0x3c, 0xcc, 0x1a, 0xad, 0x46, 0x9b, 0xce,
	0x7c, 0x93, 0x35, 0x0b, 0x11, 0x20, 0x23, 0x2a, 0x1e, 0xf9, 0x01, 0x5b, 0xb3, 0x4e, 0xe7, 0xda,
	0x24, 0x63, 0x78, 0x44, 0x2e, 0x63, 0x1c, 0xed, 0x98, 0xad, 0x9c, 0x44, 0x38, 0x18, 0xec, 0x18,
	0xa1, 0x98, 0xeb, 0x84, 0x51, 0xb6, 0xfc, 0x1c, 0x9c, 0x36, 0x39, 0x3c, 0x4e, 0xb9, 0xa3, 0x1c,
	0x7f, 0xc3, 0xf6, 0x62, 0xd9, 0xc7, 0x65, 0x7c, 0x55, 0xe7, 0x18, 0x2e, 0x9c, 0x2e, 0x85, 0xeb,
	0x7f, 0x15, 0x45, 0x81, 0x01, 0x0e, 0x69, 0x97, 0x66, 0x44, 0xf0, 0xd7, 0x6c, 0xf7, 0xaf, 0x3a,
	0x9e, 0xfc, 0x84, 0x92, 0xff, 0x25, 0xf3, 0x13, 0xb6, 0xaf, 0x8d, 0x74, 0x58, 0xa2, 0xb9, 0x35,
	0xfd, 0x88, 0xd2, 0x67, 0xc6, 0x44, 0x0f, 0x85, 0x33, 0x3c, 0x9e, 0x26, 0x8f, 0x59, 0x31, 0xdd,
	0x45, 0xfa, 0xdf, 0xbe, 0xfa, 0x13, 0x00, 0x00, 0xff, 0xff, 0xab, 0xca, 0x6b, 0xa0, 0x8f, 0x05,
	0x00, 0x00,
}
