// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: Captcha.proto

package Captcha

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

// Api Endpoints for Captcha service

func NewCaptchaEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Captcha service

type CaptchaService interface {
	GetCaptcha(ctx context.Context, in *EmptyRequest, opts ...client.CallOption) (*Img, error)
}

type captchaService struct {
	c    client.Client
	name string
}

func NewCaptchaService(name string, c client.Client) CaptchaService {
	return &captchaService{
		c:    c,
		name: name,
	}
}

func (c *captchaService) GetCaptcha(ctx context.Context, in *EmptyRequest, opts ...client.CallOption) (*Img, error) {
	req := c.c.NewRequest(c.name, "Captcha.GetCaptcha", in)
	out := new(Img)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Captcha service

type CaptchaHandler interface {
	GetCaptcha(context.Context, *EmptyRequest, *Img) error
}

func RegisterCaptchaHandler(s server.Server, hdlr CaptchaHandler, opts ...server.HandlerOption) error {
	type captcha interface {
		GetCaptcha(ctx context.Context, in *EmptyRequest, out *Img) error
	}
	type Captcha struct {
		captcha
	}
	h := &captchaHandler{hdlr}
	return s.Handle(s.NewHandler(&Captcha{h}, opts...))
}

type captchaHandler struct {
	CaptchaHandler
}

func (h *captchaHandler) GetCaptcha(ctx context.Context, in *EmptyRequest, out *Img) error {
	return h.CaptchaHandler.GetCaptcha(ctx, in, out)
}
