// Code generated by protoc-gen-go. DO NOT EDIT.
// source: secret.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	secret.proto

It has these top-level messages:
	Id
	Secret
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

type Id struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *Id) Reset()                    { *m = Id{} }
func (m *Id) String() string            { return proto1.CompactTextString(m) }
func (*Id) ProtoMessage()               {}
func (*Id) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Id) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type Secret struct {
	Secret string `protobuf:"bytes,1,opt,name=secret" json:"secret,omitempty"`
}

func (m *Secret) Reset()                    { *m = Secret{} }
func (m *Secret) String() string            { return proto1.CompactTextString(m) }
func (*Secret) ProtoMessage()               {}
func (*Secret) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Secret) GetSecret() string {
	if m != nil {
		return m.Secret
	}
	return ""
}

func init() {
	proto1.RegisterType((*Id)(nil), "proto.Id")
	proto1.RegisterType((*Secret)(nil), "proto.Secret")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for SecretService service

type SecretServiceClient interface {
	PostSecret(ctx context.Context, in *Secret, opts ...grpc.CallOption) (*Id, error)
	ReadSecret(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Secret, error)
}

type secretServiceClient struct {
	cc *grpc.ClientConn
}

func NewSecretServiceClient(cc *grpc.ClientConn) SecretServiceClient {
	return &secretServiceClient{cc}
}

func (c *secretServiceClient) PostSecret(ctx context.Context, in *Secret, opts ...grpc.CallOption) (*Id, error) {
	out := new(Id)
	err := grpc.Invoke(ctx, "/proto.SecretService/PostSecret", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *secretServiceClient) ReadSecret(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Secret, error) {
	out := new(Secret)
	err := grpc.Invoke(ctx, "/proto.SecretService/ReadSecret", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SecretService service

type SecretServiceServer interface {
	PostSecret(context.Context, *Secret) (*Id, error)
	ReadSecret(context.Context, *Id) (*Secret, error)
}

func RegisterSecretServiceServer(s *grpc.Server, srv SecretServiceServer) {
	s.RegisterService(&_SecretService_serviceDesc, srv)
}

func _SecretService_PostSecret_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Secret)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretServiceServer).PostSecret(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SecretService/PostSecret",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretServiceServer).PostSecret(ctx, req.(*Secret))
	}
	return interceptor(ctx, in, info, handler)
}

func _SecretService_ReadSecret_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SecretServiceServer).ReadSecret(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.SecretService/ReadSecret",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SecretServiceServer).ReadSecret(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

var _SecretService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.SecretService",
	HandlerType: (*SecretServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PostSecret",
			Handler:    _SecretService_PostSecret_Handler,
		},
		{
			MethodName: "ReadSecret",
			Handler:    _SecretService_ReadSecret_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "secret.proto",
}

func init() { proto1.RegisterFile("secret.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 137 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0x4e, 0x4d, 0x2e,
	0x4a, 0x2d, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x4a, 0x12, 0x5c, 0x4c,
	0x9e, 0x29, 0x42, 0x42, 0x5c, 0x2c, 0x79, 0x89, 0xb9, 0xa9, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c,
	0x41, 0x60, 0xb6, 0x92, 0x02, 0x17, 0x5b, 0x30, 0x58, 0x83, 0x90, 0x18, 0x17, 0x1b, 0x44, 0x2b,
	0x54, 0x1e, 0xca, 0x33, 0x4a, 0xe6, 0xe2, 0x85, 0xa8, 0x08, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e,
	0x15, 0xd2, 0xe0, 0xe2, 0x0a, 0xc8, 0x2f, 0x2e, 0x81, 0x6a, 0xe3, 0x85, 0xd8, 0xa4, 0x07, 0xe1,
	0x4a, 0x71, 0x42, 0xb9, 0x9e, 0x29, 0x4a, 0x0c, 0x20, 0x95, 0x41, 0xa9, 0x89, 0x29, 0x50, 0x95,
	0x08, 0x29, 0x29, 0x54, 0x4d, 0x4a, 0x0c, 0x49, 0x6c, 0x60, 0xbe, 0x31, 0x20, 0x00, 0x00, 0xff,
	0xff, 0x5e, 0xce, 0x82, 0x0a, 0xbe, 0x00, 0x00, 0x00,
}
