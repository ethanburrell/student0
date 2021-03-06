// Code generated by protoc-gen-go. DO NOT EDIT.
// source: kv.proto

/*
Package api is a generated protocol buffer package.

It is generated from these files:
	kv.proto

It has these top-level messages:
	GetRequest
	PutRequest
	Response
*/
package api

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

type GetRequest struct {
	Key string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
}

func (m *GetRequest) Reset()                    { *m = GetRequest{} }
func (m *GetRequest) String() string            { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()               {}
func (*GetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type PutRequest struct {
	Key   string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value" json:"value,omitempty"`
}

func (m *PutRequest) Reset()                    { *m = PutRequest{} }
func (m *PutRequest) String() string            { return proto.CompactTextString(m) }
func (*PutRequest) ProtoMessage()               {}
func (*PutRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *PutRequest) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *PutRequest) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type Response struct {
	Value string `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Response) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*GetRequest)(nil), "api.GetRequest")
	proto.RegisterType((*PutRequest)(nil), "api.PutRequest")
	proto.RegisterType((*Response)(nil), "api.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for KeyValueAPI service

type KeyValueAPIClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Response, error)
	Put(ctx context.Context, in *PutRequest, opts ...grpc.CallOption) (*Response, error)
}

type keyValueAPIClient struct {
	cc *grpc.ClientConn
}

func NewKeyValueAPIClient(cc *grpc.ClientConn) KeyValueAPIClient {
	return &keyValueAPIClient{cc}
}

func (c *keyValueAPIClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/api.KeyValueAPI/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyValueAPIClient) Put(ctx context.Context, in *PutRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/api.KeyValueAPI/Put", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for KeyValueAPI service

type KeyValueAPIServer interface {
	Get(context.Context, *GetRequest) (*Response, error)
	Put(context.Context, *PutRequest) (*Response, error)
}

func RegisterKeyValueAPIServer(s *grpc.Server, srv KeyValueAPIServer) {
	s.RegisterService(&_KeyValueAPI_serviceDesc, srv)
}

func _KeyValueAPI_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyValueAPIServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.KeyValueAPI/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyValueAPIServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyValueAPI_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyValueAPIServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.KeyValueAPI/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyValueAPIServer).Put(ctx, req.(*PutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _KeyValueAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.KeyValueAPI",
	HandlerType: (*KeyValueAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _KeyValueAPI_Get_Handler,
		},
		{
			MethodName: "Put",
			Handler:    _KeyValueAPI_Put_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kv.proto",
}

func init() { proto.RegisterFile("kv.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 160 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xc8, 0x2e, 0xd3, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4e, 0x2c, 0xc8, 0x54, 0x92, 0xe3, 0xe2, 0x72, 0x4f, 0x2d,
	0x09, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x12, 0xe0, 0x62, 0xce, 0x4e, 0xad, 0x94, 0x60,
	0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x02, 0x31, 0x95, 0x4c, 0xb8, 0xb8, 0x02, 0x4a, 0x71, 0xcb, 0x0b,
	0x89, 0x70, 0xb1, 0x96, 0x25, 0xe6, 0x94, 0xa6, 0x4a, 0x30, 0x81, 0xc5, 0x20, 0x1c, 0x25, 0x05,
	0x2e, 0x8e, 0xa0, 0xd4, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x84, 0x0a, 0x46, 0x24, 0x15, 0x46,
	0xd1, 0x5c, 0xdc, 0xde, 0xa9, 0x95, 0x61, 0x20, 0xb6, 0x63, 0x80, 0xa7, 0x90, 0x2a, 0x17, 0xb3,
	0x7b, 0x6a, 0x89, 0x10, 0xbf, 0x5e, 0x62, 0x41, 0xa6, 0x1e, 0xc2, 0x41, 0x52, 0xbc, 0x60, 0x01,
	0xb8, 0x59, 0xaa, 0x5c, 0xcc, 0x01, 0xa5, 0x30, 0x65, 0x08, 0x77, 0xa1, 0x29, 0x4b, 0x62, 0x03,
	0x7b, 0xd0, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x92, 0x35, 0xd0, 0x35, 0xec, 0x00, 0x00, 0x00,
}
