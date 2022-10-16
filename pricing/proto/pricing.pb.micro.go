// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/pricing.proto

package pricing

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

// Api Endpoints for Pricing service

func NewPricingEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Pricing service

type PricingService interface {
	Call(ctx context.Context, in *PricingRequest, opts ...client.CallOption) (*PricingResponse, error)
}

type pricingService struct {
	c    client.Client
	name string
}

func NewPricingService(name string, c client.Client) PricingService {
	return &pricingService{
		c:    c,
		name: name,
	}
}

func (c *pricingService) Call(ctx context.Context, in *PricingRequest, opts ...client.CallOption) (*PricingResponse, error) {
	req := c.c.NewRequest(c.name, "Pricing.Call", in)
	out := new(PricingResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Pricing service

type PricingHandler interface {
	Call(context.Context, *PricingRequest, *PricingResponse) error
}

func RegisterPricingHandler(s server.Server, hdlr PricingHandler, opts ...server.HandlerOption) error {
	type pricing interface {
		Call(ctx context.Context, in *PricingRequest, out *PricingResponse) error
	}
	type Pricing struct {
		pricing
	}
	h := &pricingHandler{hdlr}
	return s.Handle(s.NewHandler(&Pricing{h}, opts...))
}

type pricingHandler struct {
	PricingHandler
}

func (h *pricingHandler) Call(ctx context.Context, in *PricingRequest, out *PricingResponse) error {
	return h.PricingHandler.Call(ctx, in, out)
}