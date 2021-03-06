// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: service_core/service.proto

package service_core

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Service service

type Service interface {
	FindWalletByID(ctx context.Context, in *FindWalletByIDRequest, opts ...client.CallOption) (*FindWalletByIDResponse, error)
	FindWalletByUserIDAndType(ctx context.Context, in *FindWalletByUserIDTypeRequest, opts ...client.CallOption) (*FindWalletByUserIDTypeResponse, error)
	FindOrCreateWalletByUserID(ctx context.Context, in *FindOrCreateByUserIDRequest, opts ...client.CallOption) (*FindOrCreateWalletByUserIDResponse, error)
	Freeze(ctx context.Context, in *FreezeRequest, opts ...client.CallOption) (*FreezeResponse, error)
	ReleaseFreeze(ctx context.Context, in *ReleaseFreezeRequest, opts ...client.CallOption) (*ReleaseFreezeResponse, error)
	TransferMoney(ctx context.Context, in *TransferRequest, opts ...client.CallOption) (*TransferMoneyResponse, error)
	FindWalletByUserID(ctx context.Context, in *FindWalletByUserIDRequest, opts ...client.CallOption) (*FindWalletByUserIDResponse, error)
}

type service struct {
	c    client.Client
	name string
}

func NewService(name string, c client.Client) Service {
	return &service{
		c:    c,
		name: name,
	}
}

func (c *service) FindWalletByID(ctx context.Context, in *FindWalletByIDRequest, opts ...client.CallOption) (*FindWalletByIDResponse, error) {
	req := c.c.NewRequest(c.name, "Service.FindWalletByID", in)
	out := new(FindWalletByIDResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *service) FindWalletByUserIDAndType(ctx context.Context, in *FindWalletByUserIDTypeRequest, opts ...client.CallOption) (*FindWalletByUserIDTypeResponse, error) {
	req := c.c.NewRequest(c.name, "Service.FindWalletByUserIDAndType", in)
	out := new(FindWalletByUserIDTypeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *service) FindOrCreateWalletByUserID(ctx context.Context, in *FindOrCreateByUserIDRequest, opts ...client.CallOption) (*FindOrCreateWalletByUserIDResponse, error) {
	req := c.c.NewRequest(c.name, "Service.FindOrCreateWalletByUserID", in)
	out := new(FindOrCreateWalletByUserIDResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *service) Freeze(ctx context.Context, in *FreezeRequest, opts ...client.CallOption) (*FreezeResponse, error) {
	req := c.c.NewRequest(c.name, "Service.Freeze", in)
	out := new(FreezeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *service) ReleaseFreeze(ctx context.Context, in *ReleaseFreezeRequest, opts ...client.CallOption) (*ReleaseFreezeResponse, error) {
	req := c.c.NewRequest(c.name, "Service.ReleaseFreeze", in)
	out := new(ReleaseFreezeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *service) TransferMoney(ctx context.Context, in *TransferRequest, opts ...client.CallOption) (*TransferMoneyResponse, error) {
	req := c.c.NewRequest(c.name, "Service.TransferMoney", in)
	out := new(TransferMoneyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *service) FindWalletByUserID(ctx context.Context, in *FindWalletByUserIDRequest, opts ...client.CallOption) (*FindWalletByUserIDResponse, error) {
	req := c.c.NewRequest(c.name, "Service.FindWalletByUserID", in)
	out := new(FindWalletByUserIDResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Service service

type ServiceHandler interface {
	FindWalletByID(context.Context, *FindWalletByIDRequest, *FindWalletByIDResponse) error
	FindWalletByUserIDAndType(context.Context, *FindWalletByUserIDTypeRequest, *FindWalletByUserIDTypeResponse) error
	FindOrCreateWalletByUserID(context.Context, *FindOrCreateByUserIDRequest, *FindOrCreateWalletByUserIDResponse) error
	Freeze(context.Context, *FreezeRequest, *FreezeResponse) error
	ReleaseFreeze(context.Context, *ReleaseFreezeRequest, *ReleaseFreezeResponse) error
	TransferMoney(context.Context, *TransferRequest, *TransferMoneyResponse) error
	FindWalletByUserID(context.Context, *FindWalletByUserIDRequest, *FindWalletByUserIDResponse) error
}

func RegisterServiceHandler(s server.Server, hdlr ServiceHandler, opts ...server.HandlerOption) error {
	type service interface {
		FindWalletByID(ctx context.Context, in *FindWalletByIDRequest, out *FindWalletByIDResponse) error
		FindWalletByUserIDAndType(ctx context.Context, in *FindWalletByUserIDTypeRequest, out *FindWalletByUserIDTypeResponse) error
		FindOrCreateWalletByUserID(ctx context.Context, in *FindOrCreateByUserIDRequest, out *FindOrCreateWalletByUserIDResponse) error
		Freeze(ctx context.Context, in *FreezeRequest, out *FreezeResponse) error
		ReleaseFreeze(ctx context.Context, in *ReleaseFreezeRequest, out *ReleaseFreezeResponse) error
		TransferMoney(ctx context.Context, in *TransferRequest, out *TransferMoneyResponse) error
		FindWalletByUserID(ctx context.Context, in *FindWalletByUserIDRequest, out *FindWalletByUserIDResponse) error
	}
	type Service struct {
		service
	}
	h := &serviceHandler{hdlr}
	return s.Handle(s.NewHandler(&Service{h}, opts...))
}

type serviceHandler struct {
	ServiceHandler
}

func (h *serviceHandler) FindWalletByID(ctx context.Context, in *FindWalletByIDRequest, out *FindWalletByIDResponse) error {
	return h.ServiceHandler.FindWalletByID(ctx, in, out)
}

func (h *serviceHandler) FindWalletByUserIDAndType(ctx context.Context, in *FindWalletByUserIDTypeRequest, out *FindWalletByUserIDTypeResponse) error {
	return h.ServiceHandler.FindWalletByUserIDAndType(ctx, in, out)
}

func (h *serviceHandler) FindOrCreateWalletByUserID(ctx context.Context, in *FindOrCreateByUserIDRequest, out *FindOrCreateWalletByUserIDResponse) error {
	return h.ServiceHandler.FindOrCreateWalletByUserID(ctx, in, out)
}

func (h *serviceHandler) Freeze(ctx context.Context, in *FreezeRequest, out *FreezeResponse) error {
	return h.ServiceHandler.Freeze(ctx, in, out)
}

func (h *serviceHandler) ReleaseFreeze(ctx context.Context, in *ReleaseFreezeRequest, out *ReleaseFreezeResponse) error {
	return h.ServiceHandler.ReleaseFreeze(ctx, in, out)
}

func (h *serviceHandler) TransferMoney(ctx context.Context, in *TransferRequest, out *TransferMoneyResponse) error {
	return h.ServiceHandler.TransferMoney(ctx, in, out)
}

func (h *serviceHandler) FindWalletByUserID(ctx context.Context, in *FindWalletByUserIDRequest, out *FindWalletByUserIDResponse) error {
	return h.ServiceHandler.FindWalletByUserID(ctx, in, out)
}
