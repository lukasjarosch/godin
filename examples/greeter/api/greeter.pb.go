// Code generated by protoc-gen-go. DO NOT EDIT.
// source: greeter.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
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

type HelloRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloRequest) Reset()         { *m = HelloRequest{} }
func (m *HelloRequest) String() string { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()    {}
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_e585294ab3f34af5, []int{0}
}

func (m *HelloRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloRequest.Unmarshal(m, b)
}
func (m *HelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloRequest.Marshal(b, m, deterministic)
}
func (m *HelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRequest.Merge(m, src)
}
func (m *HelloRequest) XXX_Size() int {
	return xxx_messageInfo_HelloRequest.Size(m)
}
func (m *HelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRequest proto.InternalMessageInfo

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type HelloResponse struct {
	Greeting             string   `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloResponse) Reset()         { *m = HelloResponse{} }
func (m *HelloResponse) String() string { return proto.CompactTextString(m) }
func (*HelloResponse) ProtoMessage()    {}
func (*HelloResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e585294ab3f34af5, []int{1}
}

func (m *HelloResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloResponse.Unmarshal(m, b)
}
func (m *HelloResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloResponse.Marshal(b, m, deterministic)
}
func (m *HelloResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloResponse.Merge(m, src)
}
func (m *HelloResponse) XXX_Size() int {
	return xxx_messageInfo_HelloResponse.Size(m)
}
func (m *HelloResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HelloResponse proto.InternalMessageInfo

func (m *HelloResponse) GetGreeting() string {
	if m != nil {
		return m.Greeting
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "godin.greeter.v1beta1.HelloRequest")
	proto.RegisterType((*HelloResponse)(nil), "godin.greeter.v1beta1.HelloResponse")
}

func init() { proto.RegisterFile("greeter.proto", fileDescriptor_e585294ab3f34af5) }

var fileDescriptor_e585294ab3f34af5 = []byte{
	// 165 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0x2f, 0x4a, 0x4d,
	0x2d, 0x49, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x4d, 0xcf, 0x4f, 0xc9, 0xcc,
	0xd3, 0x83, 0x09, 0x96, 0x19, 0x26, 0xa5, 0x96, 0x24, 0x1a, 0x2a, 0x29, 0x71, 0xf1, 0x78, 0xa4,
	0xe6, 0xe4, 0xe4, 0x07, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0x09, 0x71, 0xb1, 0xe4, 0x25,
	0xe6, 0xa6, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81, 0xd9, 0x4a, 0xda, 0x5c, 0xbc, 0x50,
	0x35, 0xc5, 0x05, 0xf9, 0x79, 0xc5, 0xa9, 0x42, 0x52, 0x5c, 0x1c, 0x60, 0x73, 0x32, 0xf3, 0xd2,
	0xa1, 0x0a, 0xe1, 0x7c, 0xa3, 0x14, 0x2e, 0x3e, 0x77, 0x88, 0x1d, 0xc1, 0xa9, 0x45, 0x65, 0x99,
	0xc9, 0xa9, 0x42, 0x41, 0x5c, 0xac, 0x60, 0xed, 0x42, 0xca, 0x7a, 0x58, 0xdd, 0xa0, 0x87, 0xec,
	0x00, 0x29, 0x15, 0xfc, 0x8a, 0x20, 0x2e, 0x70, 0xe2, 0x8c, 0x62, 0xaf, 0xac, 0xac, 0xd4, 0x4f,
	0x2c, 0xc8, 0x4c, 0x62, 0x03, 0xfb, 0xcf, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xb2, 0x24, 0xca,
	0x59, 0xf0, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GreeterServiceClient is the client API for GreeterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GreeterServiceClient interface {
	Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
}

type greeterServiceClient struct {
	cc *grpc.ClientConn
}

func NewGreeterServiceClient(cc *grpc.ClientConn) GreeterServiceClient {
	return &greeterServiceClient{cc}
}

func (c *greeterServiceClient) Hello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, "/godin.greeter.v1beta1.GreeterService/Hello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreeterServiceServer is the server API for GreeterService service.
type GreeterServiceServer interface {
	Hello(context.Context, *HelloRequest) (*HelloResponse, error)
}

func RegisterGreeterServiceServer(s *grpc.Server, srv GreeterServiceServer) {
	s.RegisterService(&_GreeterService_serviceDesc, srv)
}

func _GreeterService_Hello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServiceServer).Hello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/godin.greeter.v1beta1.GreeterService/Hello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServiceServer).Hello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GreeterService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "godin.greeter.v1beta1.GreeterService",
	HandlerType: (*GreeterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Hello",
			Handler:    _GreeterService_Hello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "greeter.proto",
}