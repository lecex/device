// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/cashier/cashier.proto

package cashier

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

// Client API for Cashiers service

type CashiersService interface {
	// 用过 code 查询收银员是否存在
	Exist(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	// 获取所有收银员信息
	All(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	// 获取设备列表
	List(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	// 获取单条设备信息
	Get(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	// 创建或更新商品设备信息
	Create(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	// 更新数据
	Update(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	// 删除商品设备信息
	Delete(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type cashiersService struct {
	c    client.Client
	name string
}

func NewCashiersService(name string, c client.Client) CashiersService {
	return &cashiersService{
		c:    c,
		name: name,
	}
}

func (c *cashiersService) Exist(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Cashiers.Exist", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cashiersService) All(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Cashiers.All", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cashiersService) List(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Cashiers.List", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cashiersService) Get(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Cashiers.Get", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cashiersService) Create(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Cashiers.Create", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cashiersService) Update(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Cashiers.Update", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cashiersService) Delete(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Cashiers.Delete", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Cashiers service

type CashiersHandler interface {
	// 用过 code 查询收银员是否存在
	Exist(context.Context, *Request, *Response) error
	// 获取所有收银员信息
	All(context.Context, *Request, *Response) error
	// 获取设备列表
	List(context.Context, *Request, *Response) error
	// 获取单条设备信息
	Get(context.Context, *Request, *Response) error
	// 创建或更新商品设备信息
	Create(context.Context, *Request, *Response) error
	// 更新数据
	Update(context.Context, *Request, *Response) error
	// 删除商品设备信息
	Delete(context.Context, *Request, *Response) error
}

func RegisterCashiersHandler(s server.Server, hdlr CashiersHandler, opts ...server.HandlerOption) error {
	type cashiers interface {
		Exist(ctx context.Context, in *Request, out *Response) error
		All(ctx context.Context, in *Request, out *Response) error
		List(ctx context.Context, in *Request, out *Response) error
		Get(ctx context.Context, in *Request, out *Response) error
		Create(ctx context.Context, in *Request, out *Response) error
		Update(ctx context.Context, in *Request, out *Response) error
		Delete(ctx context.Context, in *Request, out *Response) error
	}
	type Cashiers struct {
		cashiers
	}
	h := &cashiersHandler{hdlr}
	return s.Handle(s.NewHandler(&Cashiers{h}, opts...))
}

type cashiersHandler struct {
	CashiersHandler
}

func (h *cashiersHandler) Exist(ctx context.Context, in *Request, out *Response) error {
	return h.CashiersHandler.Exist(ctx, in, out)
}

func (h *cashiersHandler) All(ctx context.Context, in *Request, out *Response) error {
	return h.CashiersHandler.All(ctx, in, out)
}

func (h *cashiersHandler) List(ctx context.Context, in *Request, out *Response) error {
	return h.CashiersHandler.List(ctx, in, out)
}

func (h *cashiersHandler) Get(ctx context.Context, in *Request, out *Response) error {
	return h.CashiersHandler.Get(ctx, in, out)
}

func (h *cashiersHandler) Create(ctx context.Context, in *Request, out *Response) error {
	return h.CashiersHandler.Create(ctx, in, out)
}

func (h *cashiersHandler) Update(ctx context.Context, in *Request, out *Response) error {
	return h.CashiersHandler.Update(ctx, in, out)
}

func (h *cashiersHandler) Delete(ctx context.Context, in *Request, out *Response) error {
	return h.CashiersHandler.Delete(ctx, in, out)
}
