// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/run-controller-service/run-controller-service.proto

package go_micro_service_runcontrollerservice

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

// Api Endpoints for RunControllerService service

func NewRunControllerServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for RunControllerService service

type RunControllerService interface {
	Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (RunControllerService_StreamService, error)
	PingPong(ctx context.Context, opts ...client.CallOption) (RunControllerService_PingPongService, error)
}

type runControllerService struct {
	c    client.Client
	name string
}

func NewRunControllerService(name string, c client.Client) RunControllerService {
	return &runControllerService{
		c:    c,
		name: name,
	}
}

func (c *runControllerService) Call(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "RunControllerService.Call", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *runControllerService) Stream(ctx context.Context, in *StreamingRequest, opts ...client.CallOption) (RunControllerService_StreamService, error) {
	req := c.c.NewRequest(c.name, "RunControllerService.Stream", &StreamingRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &runControllerServiceStream{stream}, nil
}

type RunControllerService_StreamService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*StreamingResponse, error)
}

type runControllerServiceStream struct {
	stream client.Stream
}

func (x *runControllerServiceStream) Close() error {
	return x.stream.Close()
}

func (x *runControllerServiceStream) Context() context.Context {
	return x.stream.Context()
}

func (x *runControllerServiceStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *runControllerServiceStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *runControllerServiceStream) Recv() (*StreamingResponse, error) {
	m := new(StreamingResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (c *runControllerService) PingPong(ctx context.Context, opts ...client.CallOption) (RunControllerService_PingPongService, error) {
	req := c.c.NewRequest(c.name, "RunControllerService.PingPong", &Ping{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	return &runControllerServicePingPong{stream}, nil
}

type RunControllerService_PingPongService interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Ping) error
	Recv() (*Pong, error)
}

type runControllerServicePingPong struct {
	stream client.Stream
}

func (x *runControllerServicePingPong) Close() error {
	return x.stream.Close()
}

func (x *runControllerServicePingPong) Context() context.Context {
	return x.stream.Context()
}

func (x *runControllerServicePingPong) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *runControllerServicePingPong) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *runControllerServicePingPong) Send(m *Ping) error {
	return x.stream.Send(m)
}

func (x *runControllerServicePingPong) Recv() (*Pong, error) {
	m := new(Pong)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for RunControllerService service

type RunControllerServiceHandler interface {
	Call(context.Context, *Request, *Response) error
	Stream(context.Context, *StreamingRequest, RunControllerService_StreamStream) error
	PingPong(context.Context, RunControllerService_PingPongStream) error
}

func RegisterRunControllerServiceHandler(s server.Server, hdlr RunControllerServiceHandler, opts ...server.HandlerOption) error {
	type runControllerService interface {
		Call(ctx context.Context, in *Request, out *Response) error
		Stream(ctx context.Context, stream server.Stream) error
		PingPong(ctx context.Context, stream server.Stream) error
	}
	type RunControllerService struct {
		runControllerService
	}
	h := &runControllerServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&RunControllerService{h}, opts...))
}

type runControllerServiceHandler struct {
	RunControllerServiceHandler
}

func (h *runControllerServiceHandler) Call(ctx context.Context, in *Request, out *Response) error {
	return h.RunControllerServiceHandler.Call(ctx, in, out)
}

func (h *runControllerServiceHandler) Stream(ctx context.Context, stream server.Stream) error {
	m := new(StreamingRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.RunControllerServiceHandler.Stream(ctx, m, &runControllerServiceStreamStream{stream})
}

type RunControllerService_StreamStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*StreamingResponse) error
}

type runControllerServiceStreamStream struct {
	stream server.Stream
}

func (x *runControllerServiceStreamStream) Close() error {
	return x.stream.Close()
}

func (x *runControllerServiceStreamStream) Context() context.Context {
	return x.stream.Context()
}

func (x *runControllerServiceStreamStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *runControllerServiceStreamStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *runControllerServiceStreamStream) Send(m *StreamingResponse) error {
	return x.stream.Send(m)
}

func (h *runControllerServiceHandler) PingPong(ctx context.Context, stream server.Stream) error {
	return h.RunControllerServiceHandler.PingPong(ctx, &runControllerServicePingPongStream{stream})
}

type RunControllerService_PingPongStream interface {
	Context() context.Context
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*Pong) error
	Recv() (*Ping, error)
}

type runControllerServicePingPongStream struct {
	stream server.Stream
}

func (x *runControllerServicePingPongStream) Close() error {
	return x.stream.Close()
}

func (x *runControllerServicePingPongStream) Context() context.Context {
	return x.stream.Context()
}

func (x *runControllerServicePingPongStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *runControllerServicePingPongStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *runControllerServicePingPongStream) Send(m *Pong) error {
	return x.stream.Send(m)
}

func (x *runControllerServicePingPongStream) Recv() (*Ping, error) {
	m := new(Ping)
	if err := x.stream.Recv(m); err != nil {
		return nil, err
	}
	return m, nil
}
