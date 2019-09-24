// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/zrma/1d1c/cmd/micro/booking/srv/auth/proto/auth.proto

package auth

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Request struct {
	AuthToken            string   `protobuf:"bytes,1,opt,name=authToken,proto3" json:"authToken,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_28d32e5d2aa06eb0, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetAuthToken() string {
	if m != nil {
		return m.AuthToken
	}
	return ""
}

type Result struct {
	Customer             *Customer `protobuf:"bytes,1,opt,name=customer,proto3" json:"customer,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *Result) Reset()         { *m = Result{} }
func (m *Result) String() string { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()    {}
func (*Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_28d32e5d2aa06eb0, []int{1}
}

func (m *Result) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Result.Unmarshal(m, b)
}
func (m *Result) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Result.Marshal(b, m, deterministic)
}
func (m *Result) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Result.Merge(m, src)
}
func (m *Result) XXX_Size() int {
	return xxx_messageInfo_Result.Size(m)
}
func (m *Result) XXX_DiscardUnknown() {
	xxx_messageInfo_Result.DiscardUnknown(m)
}

var xxx_messageInfo_Result proto.InternalMessageInfo

func (m *Result) GetCustomer() *Customer {
	if m != nil {
		return m.Customer
	}
	return nil
}

type Customer struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	AuthToken            string   `protobuf:"bytes,2,opt,name=authToken,proto3" json:"authToken,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Customer) Reset()         { *m = Customer{} }
func (m *Customer) String() string { return proto.CompactTextString(m) }
func (*Customer) ProtoMessage()    {}
func (*Customer) Descriptor() ([]byte, []int) {
	return fileDescriptor_28d32e5d2aa06eb0, []int{2}
}

func (m *Customer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Customer.Unmarshal(m, b)
}
func (m *Customer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Customer.Marshal(b, m, deterministic)
}
func (m *Customer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Customer.Merge(m, src)
}
func (m *Customer) XXX_Size() int {
	return xxx_messageInfo_Customer.Size(m)
}
func (m *Customer) XXX_DiscardUnknown() {
	xxx_messageInfo_Customer.DiscardUnknown(m)
}

var xxx_messageInfo_Customer proto.InternalMessageInfo

func (m *Customer) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Customer) GetAuthToken() string {
	if m != nil {
		return m.AuthToken
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "auth.Request")
	proto.RegisterType((*Result)(nil), "auth.Result")
	proto.RegisterType((*Customer)(nil), "auth.Customer")
}

func init() {
	proto.RegisterFile("github.com/zrma/1d1c/cmd/micro/booking/srv/auth/proto/auth.proto", fileDescriptor_28d32e5d2aa06eb0)
}

var fileDescriptor_28d32e5d2aa06eb0 = []byte{
	// 217 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x8f, 0x4f, 0x4b, 0xc4, 0x30,
	0x10, 0xc5, 0xdd, 0xb2, 0xae, 0xbb, 0xb3, 0xba, 0x87, 0x9c, 0x44, 0x3c, 0x48, 0x2e, 0x8a, 0x48,
	0x87, 0xd6, 0x1e, 0x3c, 0x2a, 0x7e, 0x83, 0x20, 0xde, 0xdb, 0x24, 0xb6, 0xa1, 0xa6, 0xa3, 0xf9,
	0x23, 0xe8, 0xa7, 0x97, 0xc6, 0x56, 0x71, 0x6f, 0x6f, 0x7e, 0xf3, 0xe7, 0xbd, 0x81, 0xfb, 0xd6,
	0x84, 0x2e, 0x36, 0xb9, 0x24, 0x8b, 0x5f, 0xce, 0xd6, 0x58, 0xa8, 0x42, 0xa2, 0xb4, 0x0a, 0xad,
	0x91, 0x8e, 0xb0, 0x21, 0xea, 0xcd, 0xd0, 0xa2, 0x77, 0x1f, 0x58, 0xc7, 0xd0, 0xe1, 0x9b, 0xa3,
	0x40, 0x49, 0xe6, 0x49, 0xb2, 0xe5, 0xa8, 0xf9, 0x25, 0x1c, 0x09, 0xfd, 0x1e, 0xb5, 0x0f, 0xec,
	0x1c, 0x36, 0x23, 0x7a, 0xa2, 0x5e, 0x0f, 0xa7, 0x8b, 0x8b, 0xc5, 0xd5, 0x46, 0xfc, 0x01, 0x5e,
	0xc1, 0x4a, 0x68, 0x1f, 0x5f, 0x03, 0xbb, 0x86, 0xb5, 0x8c, 0x3e, 0x90, 0xd5, 0x2e, 0x8d, 0x6d,
	0xcb, 0x5d, 0x9e, 0xee, 0x3e, 0x4e, 0x54, 0xfc, 0xf6, 0xf9, 0x1d, 0xac, 0x67, 0xca, 0x76, 0x90,
	0x19, 0x95, 0x36, 0x0e, 0x45, 0x66, 0xd4, 0x7f, 0xbf, 0x6c, 0xcf, 0xaf, 0xac, 0x60, 0xf9, 0x10,
	0x43, 0xc7, 0x6e, 0x60, 0xfb, 0xac, 0x9d, 0x79, 0xf9, 0x4c, 0x98, 0x9d, 0xfc, 0x58, 0x4d, 0x99,
	0xcf, 0x8e, 0xe7, 0x72, 0x4c, 0xc6, 0x0f, 0x9a, 0x55, 0xfa, 0xed, 0xf6, 0x3b, 0x00, 0x00, 0xff,
	0xff, 0x43, 0x9c, 0xe6, 0xcf, 0x1f, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuthClient is the client API for Auth service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthClient interface {
	VerifyToken(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Result, error)
}

type authClient struct {
	cc *grpc.ClientConn
}

func NewAuthClient(cc *grpc.ClientConn) AuthClient {
	return &authClient{cc}
}

func (c *authClient) VerifyToken(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/auth.Auth/VerifyToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServer is the server API for Auth service.
type AuthServer interface {
	VerifyToken(context.Context, *Request) (*Result, error)
}

// UnimplementedAuthServer can be embedded to have forward compatible implementations.
type UnimplementedAuthServer struct {
}

func (*UnimplementedAuthServer) VerifyToken(ctx context.Context, req *Request) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyToken not implemented")
}

func RegisterAuthServer(s *grpc.Server, srv AuthServer) {
	s.RegisterService(&_Auth_serviceDesc, srv)
}

func _Auth_VerifyToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServer).VerifyToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/auth.Auth/VerifyToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServer).VerifyToken(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Auth_serviceDesc = grpc.ServiceDesc{
	ServiceName: "auth.Auth",
	HandlerType: (*AuthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "VerifyToken",
			Handler:    _Auth_VerifyToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/zrma/1d1c/cmd/micro/booking/srv/auth/proto/auth.proto",
}
