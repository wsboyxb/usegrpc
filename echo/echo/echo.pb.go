// Code generated by protoc-gen-go.
// source: echo.proto
// DO NOT EDIT!

/*
Package echo is a generated protocol buffer package.

It is generated from these files:
	echo.proto

It has these top-level messages:
	Request
	Reply
*/
package echo

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Request struct {
	UserName  string `protobuf:"bytes,1,opt,name=userName" json:"userName,omitempty"`
	Timestamp int64  `protobuf:"varint,2,opt,name=timestamp" json:"timestamp,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Request) GetUserName() string {
	if m != nil {
		return m.UserName
	}
	return ""
}

func (m *Request) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type Reply struct {
	Message   string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
	Timestamp int64  `protobuf:"varint,2,opt,name=timestamp" json:"timestamp,omitempty"`
}

func (m *Reply) Reset()                    { *m = Reply{} }
func (m *Reply) String() string            { return proto.CompactTextString(m) }
func (*Reply) ProtoMessage()               {}
func (*Reply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Reply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *Reply) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func init() {
	proto.RegisterType((*Request)(nil), "echo.Request")
	proto.RegisterType((*Reply)(nil), "echo.Reply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Echo service

type EchoClient interface {
	SayHi(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error)
}

type echoClient struct {
	cc *grpc.ClientConn
}

func NewEchoClient(cc *grpc.ClientConn) EchoClient {
	return &echoClient{cc}
}

func (c *echoClient) SayHi(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := grpc.Invoke(ctx, "/echo.Echo/SayHi", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Echo service

type EchoServer interface {
	SayHi(context.Context, *Request) (*Reply, error)
}

func RegisterEchoServer(s *grpc.Server, srv EchoServer) {
	s.RegisterService(&_Echo_serviceDesc, srv)
}

func _Echo_SayHi_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoServer).SayHi(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/echo.Echo/SayHi",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoServer).SayHi(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Echo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "echo.Echo",
	HandlerType: (*EchoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHi",
			Handler:    _Echo_SayHi_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "echo.proto",
}

func init() { proto.RegisterFile("echo.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 157 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x4d, 0xce, 0xc8,
	0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0x95, 0x9c, 0xb9, 0xd8, 0x83, 0x52,
	0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0xa4, 0xb8, 0x38, 0x4a, 0x8b, 0x53, 0x8b, 0xfc, 0x12, 0x73,
	0x53, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0xe0, 0x7c, 0x21, 0x19, 0x2e, 0xce, 0x92, 0xcc,
	0xdc, 0xd4, 0xe2, 0x92, 0xc4, 0xdc, 0x02, 0x09, 0x26, 0x05, 0x46, 0x0d, 0xe6, 0x20, 0x84, 0x80,
	0x92, 0x3d, 0x17, 0x6b, 0x50, 0x6a, 0x41, 0x4e, 0xa5, 0x90, 0x04, 0x17, 0x7b, 0x6e, 0x6a, 0x71,
	0x71, 0x62, 0x3a, 0xcc, 0x04, 0x18, 0x17, 0xbf, 0x01, 0x46, 0xba, 0x5c, 0x2c, 0xae, 0xc9, 0x19,
	0xf9, 0x42, 0xaa, 0x5c, 0xac, 0xc1, 0x89, 0x95, 0x1e, 0x99, 0x42, 0xbc, 0x7a, 0x60, 0x97, 0x42,
	0x9d, 0x26, 0xc5, 0x0d, 0xe3, 0x16, 0xe4, 0x54, 0x2a, 0x31, 0x24, 0xb1, 0x81, 0x7d, 0x60, 0x0c,
	0x08, 0x00, 0x00, 0xff, 0xff, 0xac, 0x4a, 0x83, 0xfd, 0xcf, 0x00, 0x00, 0x00,
}
