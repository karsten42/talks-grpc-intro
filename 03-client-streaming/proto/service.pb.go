// Code generated by protoc-gen-go.
// source: service.proto
// DO NOT EDIT!

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	service.proto

It has these top-level messages:
	Request
	Response
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type Request struct {
	X float64 `protobuf:"fixed64,1,opt,name=x" json:"x,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto1.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Request) GetX() float64 {
	if m != nil {
		return m.X
	}
	return 0
}

type Response struct {
	Sum float64 `protobuf:"fixed64,1,opt,name=sum" json:"sum,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto1.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Response) GetSum() float64 {
	if m != nil {
		return m.Sum
	}
	return 0
}

func init() {
	proto1.RegisterType((*Request)(nil), "Request")
	proto1.RegisterType((*Response)(nil), "Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for SumBuilder service

type SumBuilderClient interface {
	Sum(ctx context.Context, opts ...grpc.CallOption) (SumBuilder_SumClient, error)
}

type sumBuilderClient struct {
	cc *grpc.ClientConn
}

func NewSumBuilderClient(cc *grpc.ClientConn) SumBuilderClient {
	return &sumBuilderClient{cc}
}

func (c *sumBuilderClient) Sum(ctx context.Context, opts ...grpc.CallOption) (SumBuilder_SumClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_SumBuilder_serviceDesc.Streams[0], c.cc, "/SumBuilder/Sum", opts...)
	if err != nil {
		return nil, err
	}
	x := &sumBuilderSumClient{stream}
	return x, nil
}

type SumBuilder_SumClient interface {
	Send(*Request) error
	CloseAndRecv() (*Response, error)
	grpc.ClientStream
}

type sumBuilderSumClient struct {
	grpc.ClientStream
}

func (x *sumBuilderSumClient) Send(m *Request) error {
	return x.ClientStream.SendMsg(m)
}

func (x *sumBuilderSumClient) CloseAndRecv() (*Response, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for SumBuilder service

type SumBuilderServer interface {
	Sum(SumBuilder_SumServer) error
}

func RegisterSumBuilderServer(s *grpc.Server, srv SumBuilderServer) {
	s.RegisterService(&_SumBuilder_serviceDesc, srv)
}

func _SumBuilder_Sum_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(SumBuilderServer).Sum(&sumBuilderSumServer{stream})
}

type SumBuilder_SumServer interface {
	SendAndClose(*Response) error
	Recv() (*Request, error)
	grpc.ServerStream
}

type sumBuilderSumServer struct {
	grpc.ServerStream
}

func (x *sumBuilderSumServer) SendAndClose(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *sumBuilderSumServer) Recv() (*Request, error) {
	m := new(Request)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _SumBuilder_serviceDesc = grpc.ServiceDesc{
	ServiceName: "SumBuilder",
	HandlerType: (*SumBuilderServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Sum",
			Handler:       _SumBuilder_Sum_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "service.proto",
}

func init() { proto1.RegisterFile("service.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 131 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x4e, 0x2d, 0x2a,
	0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x12, 0xe7, 0x62, 0x0f, 0x4a, 0x2d,
	0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0xe2, 0xe1, 0x62, 0xac, 0x90, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0c,
	0x62, 0xac, 0x50, 0x92, 0xe1, 0xe2, 0x08, 0x4a, 0x2d, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x15, 0x12,
	0xe0, 0x62, 0x2e, 0x2e, 0xcd, 0x85, 0xca, 0x81, 0x98, 0x46, 0x5a, 0x5c, 0x5c, 0xc1, 0xa5, 0xb9,
	0x4e, 0xa5, 0x99, 0x39, 0x29, 0xa9, 0x45, 0x42, 0x32, 0x5c, 0xcc, 0xc1, 0xa5, 0xb9, 0x42, 0x1c,
	0x7a, 0x50, 0xa3, 0xa4, 0x38, 0xf5, 0x60, 0x7a, 0x35, 0x18, 0x9d, 0xd8, 0xa3, 0x58, 0xc1, 0x76,
	0x25, 0xb1, 0x81, 0x29, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x76, 0x27, 0x23, 0xc0, 0x83,
	0x00, 0x00, 0x00,
}