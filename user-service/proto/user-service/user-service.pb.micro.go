// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/user-service/user-service.proto

package go_micro_service_userservice

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
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
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for UserService service

func NewUserServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for UserService service

type UserService interface {
	Create(ctx context.Context, in *User, opts ...client.CallOption) (*EmptyResponse, error)
	GetProfile(ctx context.Context, in *EmptyRequest, opts ...client.CallOption) (*AuthResponse, error)
	Authenticate(ctx context.Context, in *User, opts ...client.CallOption) (*AuthResponse, error)
	ValidateToken(ctx context.Context, in *Token, opts ...client.CallOption) (*TokenValidationResponse, error)
}

type userService struct {
	c    client.Client
	name string
}

func NewUserService(name string, c client.Client) UserService {
	return &userService{
		c:    c,
		name: name,
	}
}

func (c *userService) Create(ctx context.Context, in *User, opts ...client.CallOption) (*EmptyResponse, error) {
	req := c.c.NewRequest(c.name, "UserService.Create", in)
	out := new(EmptyResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) GetProfile(ctx context.Context, in *EmptyRequest, opts ...client.CallOption) (*AuthResponse, error) {
	req := c.c.NewRequest(c.name, "UserService.GetProfile", in)
	out := new(AuthResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) Authenticate(ctx context.Context, in *User, opts ...client.CallOption) (*AuthResponse, error) {
	req := c.c.NewRequest(c.name, "UserService.Authenticate", in)
	out := new(AuthResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userService) ValidateToken(ctx context.Context, in *Token, opts ...client.CallOption) (*TokenValidationResponse, error) {
	req := c.c.NewRequest(c.name, "UserService.ValidateToken", in)
	out := new(TokenValidationResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserService service

type UserServiceHandler interface {
	Create(context.Context, *User, *EmptyResponse) error
	GetProfile(context.Context, *EmptyRequest, *AuthResponse) error
	Authenticate(context.Context, *User, *AuthResponse) error
	ValidateToken(context.Context, *Token, *TokenValidationResponse) error
}

func RegisterUserServiceHandler(s server.Server, hdlr UserServiceHandler, opts ...server.HandlerOption) error {
	type userService interface {
		Create(ctx context.Context, in *User, out *EmptyResponse) error
		GetProfile(ctx context.Context, in *EmptyRequest, out *AuthResponse) error
		Authenticate(ctx context.Context, in *User, out *AuthResponse) error
		ValidateToken(ctx context.Context, in *Token, out *TokenValidationResponse) error
	}
	type UserService struct {
		userService
	}
	h := &userServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&UserService{h}, opts...))
}

type userServiceHandler struct {
	UserServiceHandler
}

func (h *userServiceHandler) Create(ctx context.Context, in *User, out *EmptyResponse) error {
	return h.UserServiceHandler.Create(ctx, in, out)
}

func (h *userServiceHandler) GetProfile(ctx context.Context, in *EmptyRequest, out *AuthResponse) error {
	return h.UserServiceHandler.GetProfile(ctx, in, out)
}

func (h *userServiceHandler) Authenticate(ctx context.Context, in *User, out *AuthResponse) error {
	return h.UserServiceHandler.Authenticate(ctx, in, out)
}

func (h *userServiceHandler) ValidateToken(ctx context.Context, in *Token, out *TokenValidationResponse) error {
	return h.UserServiceHandler.ValidateToken(ctx, in, out)
}
