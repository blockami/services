// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/abi.proto

package abi

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/micro/v3/service/api"
	client "github.com/micro/micro/v3/service/client"
	server "github.com/micro/micro/v3/service/server"
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
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Abi service

func NewAbiEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Abi service

type AbiService interface {
	ContractAbi(ctx context.Context, in *AbiRequest, opts ...client.CallOption) (*AbiResponse, error)
}

type abiService struct {
	c    client.Client
	name string
}

func NewAbiService(name string, c client.Client) AbiService {
	return &abiService{
		c:    c,
		name: name,
	}
}

func (c *abiService) ContractAbi(ctx context.Context, in *AbiRequest, opts ...client.CallOption) (*AbiResponse, error) {
	req := c.c.NewRequest(c.name, "Abi.ContractAbi", in)
	out := new(AbiResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Abi service

type AbiHandler interface {
	ContractAbi(context.Context, *AbiRequest, *AbiResponse) error
}

func RegisterAbiHandler(s server.Server, hdlr AbiHandler, opts ...server.HandlerOption) error {
	type abi interface {
		ContractAbi(ctx context.Context, in *AbiRequest, out *AbiResponse) error
	}
	type Abi struct {
		abi
	}
	h := &abiHandler{hdlr}
	return s.Handle(s.NewHandler(&Abi{h}, opts...))
}

type abiHandler struct {
	AbiHandler
}

func (h *abiHandler) ContractAbi(ctx context.Context, in *AbiRequest, out *AbiResponse) error {
	return h.AbiHandler.ContractAbi(ctx, in, out)
}
