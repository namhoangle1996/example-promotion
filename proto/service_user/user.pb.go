// Code generated by protoc-gen-go. DO NOT EDIT.
// source: service_user/user.proto

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

type UserDTO struct {
	Id                           string   `protobuf:"bytes,1,opt,name=Id,proto3" json:"Id,omitempty"`
	IdGutina                     int64    `protobuf:"varint,2,opt,name=IdGutina,proto3" json:"IdGutina,omitempty"`
	PhoneNumber                  string   `protobuf:"bytes,3,opt,name=PhoneNumber,proto3" json:"PhoneNumber,omitempty"`
	Email                        string   `protobuf:"bytes,4,opt,name=Email,proto3" json:"Email,omitempty"`
	Name                         string   `protobuf:"bytes,5,opt,name=Name,proto3" json:"Name,omitempty"`
	Avatar                       string   `protobuf:"bytes,6,opt,name=Avatar,proto3" json:"Avatar,omitempty"`
	Gender                       string   `protobuf:"bytes,7,opt,name=Gender,proto3" json:"Gender,omitempty"`
	Status                       string   `protobuf:"bytes,8,opt,name=Status,proto3" json:"Status,omitempty"`
	Timezone                     string   `protobuf:"bytes,9,opt,name=Timezone,proto3" json:"Timezone,omitempty"`
	Language                     string   `protobuf:"bytes,10,opt,name=Language,proto3" json:"Language,omitempty"`
	Title                        string   `protobuf:"bytes,11,opt,name=Title,proto3" json:"Title,omitempty"`
	DateTime                     int64    `protobuf:"varint,12,opt,name=DateTime,proto3" json:"DateTime,omitempty"`
	Currency                     string   `protobuf:"bytes,13,opt,name=Currency,proto3" json:"Currency,omitempty"`
	Role                         string   `protobuf:"bytes,14,opt,name=Role,proto3" json:"Role,omitempty"`
	IsFacebookConnect            bool     `protobuf:"varint,15,opt,name=IsFacebookConnect,proto3" json:"IsFacebookConnect,omitempty"`
	EmailVerify                  bool     `protobuf:"varint,16,opt,name=EmailVerify,proto3" json:"EmailVerify,omitempty"`
	CanUpdateEmail               bool     `protobuf:"varint,17,opt,name=CanUpdateEmail,proto3" json:"CanUpdateEmail,omitempty"`
	CanAddBankAccount            bool     `protobuf:"varint,18,opt,name=CanAddBankAccount,proto3" json:"CanAddBankAccount,omitempty"`
	CanUpdatePhoneNumber         bool     `protobuf:"varint,19,opt,name=CanUpdatePhoneNumber,proto3" json:"CanUpdatePhoneNumber,omitempty"`
	PhoneNumberVerify            bool     `protobuf:"varint,20,opt,name=PhoneNumberVerify,proto3" json:"PhoneNumberVerify,omitempty"`
	AccessToken                  string   `protobuf:"bytes,21,opt,name=AccessToken,proto3" json:"AccessToken,omitempty"`
	IsExists                     bool     `protobuf:"varint,22,opt,name=IsExists,proto3" json:"IsExists,omitempty"`
	IdentityNumber               string   `protobuf:"bytes,23,opt,name=IdentityNumber,proto3" json:"IdentityNumber,omitempty"`
	BirthDay                     string   `protobuf:"bytes,24,opt,name=BirthDay,proto3" json:"BirthDay,omitempty"`
	ConnectedFacebook            bool     `protobuf:"varint,25,opt,name=ConnectedFacebook,proto3" json:"ConnectedFacebook,omitempty"`
	Address                      string   `protobuf:"bytes,26,opt,name=Address,proto3" json:"Address,omitempty"`
	MessageTopics                []string `protobuf:"bytes,27,rep,name=MessageTopics,proto3" json:"MessageTopics,omitempty"`
	DeviceIdLogin                string   `protobuf:"bytes,28,opt,name=DeviceIdLogin,proto3" json:"DeviceIdLogin,omitempty"`
	LastDeviceIdLogin            string   `protobuf:"bytes,29,opt,name=LastDeviceIdLogin,proto3" json:"LastDeviceIdLogin,omitempty"`
	Long                         float64  `protobuf:"fixed64,30,opt,name=Long,proto3" json:"Long,omitempty"`
	Lat                          float64  `protobuf:"fixed64,31,opt,name=Lat,proto3" json:"Lat,omitempty"`
	OriginDevice                 string   `protobuf:"bytes,32,opt,name=OriginDevice,proto3" json:"OriginDevice,omitempty"`
	Source                       string   `protobuf:"bytes,33,opt,name=Source,proto3" json:"Source,omitempty"`
	RandomString                 string   `protobuf:"bytes,34,opt,name=RandomString,proto3" json:"RandomString,omitempty"`
	LastAmountNetPrimaryWallet   int64    `protobuf:"varint,35,opt,name=LastAmountNetPrimaryWallet,proto3" json:"LastAmountNetPrimaryWallet,omitempty"`
	LastAmountPrimaryWallet      int64    `protobuf:"varint,36,opt,name=LastAmountPrimaryWallet,proto3" json:"LastAmountPrimaryWallet,omitempty"`
	IncrementAmountPrimaryWallet int64    `protobuf:"varint,37,opt,name=IncrementAmountPrimaryWallet,proto3" json:"IncrementAmountPrimaryWallet,omitempty"`
	DecrementAmountPrimaryWallet int64    `protobuf:"varint,38,opt,name=DecrementAmountPrimaryWallet,proto3" json:"DecrementAmountPrimaryWallet,omitempty"`
	LastReadAllMessage           int64    `protobuf:"varint,39,opt,name=LastReadAllMessage,proto3" json:"LastReadAllMessage,omitempty"`
	CreatedAt                    int64    `protobuf:"varint,40,opt,name=CreatedAt,proto3" json:"CreatedAt,omitempty"`
	UpdatedAt                    int64    `protobuf:"varint,41,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
	DeletedAt                    int64    `protobuf:"varint,42,opt,name=DeletedAt,proto3" json:"DeletedAt,omitempty"`
	XXX_NoUnkeyedLiteral         struct{} `json:"-"`
	XXX_unrecognized             []byte   `json:"-"`
	XXX_sizecache                int32    `json:"-"`
}

func (m *UserDTO) Reset()         { *m = UserDTO{} }
func (m *UserDTO) String() string { return proto.CompactTextString(m) }
func (*UserDTO) ProtoMessage()    {}
func (*UserDTO) Descriptor() ([]byte, []int) {
	return fileDescriptor_84ab6f6905cc578c, []int{0}
}

func (m *UserDTO) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserDTO.Unmarshal(m, b)
}
func (m *UserDTO) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserDTO.Marshal(b, m, deterministic)
}
func (m *UserDTO) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserDTO.Merge(m, src)
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

func (m *UserDTO) GetLastReadAllMessage() int64 {
	if m != nil {
		return m.LastReadAllMessage
	}
	return 0
}

func (m *UserDTO) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *UserDTO) GetUpdatedAt() int64 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

func (m *UserDTO) GetDeletedAt() int64 {
	if m != nil {
		return m.DeletedAt
	}
	return 0
}

func init() {
	proto.RegisterType((*UserDTO)(nil), "service_user.UserDTO")
}

func init() {
	proto.RegisterFile("service_user/user.proto", fileDescriptor_84ab6f6905cc578c)
}

var fileDescriptor_84ab6f6905cc578c = []byte{
	// 678 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x54, 0x6f, 0x4f, 0x13, 0x4f,
	0x10, 0x0e, 0x94, 0xbf, 0xcb, 0x9f, 0x1f, 0xec, 0x0f, 0x61, 0x44, 0xd4, 0x8a, 0x88, 0x68, 0x0c,
	0x26, 0xfa, 0xc6, 0x57, 0x26, 0xa5, 0x45, 0xd2, 0xa4, 0x02, 0x39, 0x8a, 0xbe, 0x34, 0xcb, 0xdd,
	0x5a, 0x36, 0xf4, 0x76, 0xc9, 0xde, 0x1e, 0xb1, 0x7e, 0x3b, 0xbf, 0x99, 0x3b, 0xb3, 0xed, 0x71,
	0xa5, 0xb5, 0x6f, 0x60, 0x9e, 0xe7, 0x99, 0x79, 0x66, 0x76, 0x6e, 0xb7, 0x6c, 0x2b, 0x93, 0xf6,
	0x4e, 0xc5, 0xf2, 0x47, 0xee, 0x83, 0xf7, 0xf8, 0xe7, 0xf0, 0xd6, 0x1a, 0x67, 0xf8, 0x72, 0x59,
	0xd8, 0xfd, 0xb3, 0xc4, 0xe6, 0x2f, 0x7d, 0xd0, 0x68, 0x9f, 0xf1, 0x55, 0x36, 0xdd, 0x4c, 0x60,
	0xaa, 0x3a, 0x75, 0xb0, 0x18, 0xf9, 0x88, 0x6f, 0xb3, 0x85, 0x66, 0x72, 0x92, 0x3b, 0xa5, 0x05,
	0x4c, 0x7b, 0xb6, 0x12, 0x15, 0x98, 0x57, 0xd9, 0xd2, 0xf9, 0xb5, 0xd1, 0xf2, 0x34, 0x4f, 0xaf,
	0xa4, 0x85, 0x0a, 0x15, 0x95, 0x29, 0xbe, 0xc1, 0x66, 0x8f, 0x53, 0xa1, 0xba, 0x30, 0x43, 0x5a,
	0x00, 0x9c, 0xb3, 0x99, 0x53, 0x91, 0x4a, 0x98, 0x25, 0x92, 0x62, 0xbe, 0xc9, 0xe6, 0x6a, 0x77,
	0xc2, 0x09, 0x0b, 0x73, 0xc4, 0xf6, 0x11, 0xf2, 0x27, 0x52, 0x27, 0xde, 0x7e, 0x3e, 0xf0, 0x01,
	0x21, 0x7f, 0xe1, 0x84, 0xcb, 0x33, 0x58, 0x08, 0x7c, 0x40, 0x38, 0x6f, 0x5b, 0xa5, 0xf2, 0xb7,
	0x9f, 0x01, 0x16, 0x49, 0x29, 0x30, 0x6a, 0x2d, 0xa1, 0x3b, 0xb9, 0xe8, 0x48, 0x60, 0x41, 0x1b,
	0x60, 0x9c, 0xb4, 0xad, 0x5c, 0x57, 0xc2, 0x52, 0x98, 0x94, 0x00, 0x56, 0x34, 0x84, 0x93, 0xe8,
	0x00, 0xcb, 0xe1, 0xf4, 0x03, 0x8c, 0x5a, 0x3d, 0xb7, 0x56, 0xea, 0xb8, 0x07, 0x2b, 0xc1, 0x6d,
	0x80, 0xf1, 0x84, 0x91, 0xf1, 0x66, 0xab, 0xe1, 0x84, 0x18, 0xf3, 0x77, 0x6c, 0xbd, 0x99, 0x7d,
	0x11, 0xb1, 0xbc, 0x32, 0xe6, 0xa6, 0x6e, 0xb4, 0x96, 0xb1, 0x83, 0xff, 0x7c, 0xc2, 0x42, 0x34,
	0x2a, 0xe0, 0x6e, 0x69, 0x59, 0xdf, 0xa4, 0x55, 0x3f, 0x7b, 0xb0, 0x46, 0x79, 0x65, 0x8a, 0xef,
	0xb3, 0xd5, 0xba, 0xd0, 0x97, 0xb7, 0x89, 0x1f, 0x28, 0x2c, 0x79, 0x9d, 0x92, 0x1e, 0xb0, 0xd8,
	0xd7, 0x33, 0xb5, 0x24, 0x39, 0x12, 0xfa, 0xa6, 0x16, 0xc7, 0x26, 0xd7, 0x0e, 0x78, 0xe8, 0x3b,
	0x22, 0xf0, 0x0f, 0x6c, 0xa3, 0xa8, 0x2f, 0x7f, 0xdc, 0xff, 0xa9, 0x60, 0xac, 0x86, 0x1d, 0x4a,
	0xb0, 0x3f, 0xf1, 0x46, 0xe8, 0x30, 0x22, 0xe0, 0xc9, 0x7c, 0x33, 0x99, 0x65, 0x6d, 0x73, 0x23,
	0x35, 0x3c, 0x0a, 0xb7, 0xa6, 0x44, 0xd1, 0x9d, 0xcb, 0x8e, 0x7f, 0xa9, 0xcc, 0x65, 0xb0, 0x49,
	0x36, 0x05, 0xc6, 0x53, 0x37, 0x13, 0xa9, 0x9d, 0x72, 0xbd, 0xfe, 0x64, 0x5b, 0x64, 0xf0, 0x80,
	0x45, 0x8f, 0x23, 0x65, 0xdd, 0x75, 0x43, 0xf4, 0x00, 0xc2, 0xd7, 0x19, 0x60, 0xda, 0x48, 0x58,
	0xb3, 0x4c, 0x06, 0x7b, 0x87, 0xc7, 0xfd, 0x8d, 0x3c, 0x14, 0x38, 0xb0, 0x79, 0xbf, 0x23, 0xeb,
	0xa7, 0x83, 0x6d, 0x32, 0x1a, 0x40, 0xbe, 0xc7, 0x56, 0xbe, 0xfa, 0xff, 0xfe, 0xfa, 0xb4, 0xcd,
	0xad, 0x8a, 0x33, 0x78, 0x52, 0xad, 0x78, 0x7d, 0x98, 0xc4, 0xac, 0x86, 0xc4, 0xc7, 0xd6, 0x4c,
	0x5a, 0xa6, 0xa3, 0x34, 0xec, 0x90, 0xcb, 0x30, 0x89, 0x33, 0xb5, 0x44, 0xe6, 0x86, 0x33, 0x9f,
	0x52, 0xe6, 0xa8, 0x80, 0xf7, 0xab, 0x65, 0x74, 0x07, 0x9e, 0xf9, 0x84, 0xa9, 0x88, 0x62, 0xbe,
	0xc6, 0x2a, 0x2d, 0xe1, 0xe0, 0x39, 0x51, 0x18, 0xf2, 0x5d, 0xb6, 0x7c, 0x66, 0x95, 0xcf, 0x0f,
	0xc5, 0x50, 0x25, 0xbb, 0x21, 0x8e, 0xde, 0x91, 0xc9, 0xad, 0x57, 0x5f, 0xf4, 0xdf, 0x11, 0x21,
	0xac, 0x8d, 0x84, 0x4e, 0x4c, 0x7a, 0xe1, 0xac, 0xf2, 0x9d, 0x76, 0x43, 0x6d, 0x99, 0xe3, 0x9f,
	0xd9, 0x36, 0x8e, 0x56, 0x4b, 0xf1, 0xe6, 0x9c, 0x4a, 0x77, 0x6e, 0x55, 0x2a, 0x6c, 0xef, 0xbb,
	0xe8, 0x76, 0xa5, 0x83, 0x97, 0xf4, 0x5e, 0x26, 0x64, 0xf0, 0x4f, 0x6c, 0xeb, 0x5e, 0x1d, 0x2e,
	0xde, 0xa3, 0xe2, 0x7f, 0xc9, 0xfc, 0x88, 0xed, 0x34, 0x75, 0x6c, 0x65, 0xea, 0xbf, 0xf9, 0xb8,
	0xf2, 0x57, 0x54, 0x3e, 0x31, 0x07, 0x3d, 0x1a, 0x72, 0x82, 0xc7, 0x7e, 0xf0, 0x98, 0x94, 0xc3,
	0x0f, 0x19, 0xc7, 0x11, 0x23, 0x29, 0x92, 0x5a, 0xb7, 0xdb, 0xff, 0xee, 0xf0, 0x9a, 0x2a, 0xc7,
	0x28, 0x7c, 0x87, 0x2d, 0xd6, 0xad, 0xf4, 0xcf, 0x27, 0xa9, 0x39, 0x38, 0xa0, 0xb4, 0x7b, 0x02,
	0xd5, 0xf0, 0xb8, 0x50, 0x7d, 0x13, 0xd4, 0x82, 0x40, 0xb5, 0x21, 0x7d, 0x53, 0x52, 0xdf, 0x06,
	0xb5, 0x20, 0xae, 0xe6, 0xe8, 0x87, 0xfd, 0xe3, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe8, 0x8d,
	0xb4, 0x63, 0xf3, 0x05, 0x00, 0x00,
}