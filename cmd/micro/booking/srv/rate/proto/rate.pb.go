// Code generated by protoc-gen-go. DO NOT EDIT.
// source: srv/rate/proto/rate.proto

package rate

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
	HotelIds             []string `protobuf:"bytes,1,rep,name=hotelIds,proto3" json:"hotelIds,omitempty"`
	InDate               string   `protobuf:"bytes,2,opt,name=inDate,proto3" json:"inDate,omitempty"`
	OutDate              string   `protobuf:"bytes,3,opt,name=outDate,proto3" json:"outDate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8fd8716f9769db8, []int{0}
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

func (m *Request) GetHotelIds() []string {
	if m != nil {
		return m.HotelIds
	}
	return nil
}

func (m *Request) GetInDate() string {
	if m != nil {
		return m.InDate
	}
	return ""
}

func (m *Request) GetOutDate() string {
	if m != nil {
		return m.OutDate
	}
	return ""
}

type Result struct {
	RatePlans            []*RatePlan `protobuf:"bytes,1,rep,name=ratePlans,proto3" json:"ratePlans,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *Result) Reset()         { *m = Result{} }
func (m *Result) String() string { return proto.CompactTextString(m) }
func (*Result) ProtoMessage()    {}
func (*Result) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8fd8716f9769db8, []int{1}
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

func (m *Result) GetRatePlans() []*RatePlan {
	if m != nil {
		return m.RatePlans
	}
	return nil
}

type RatePlan struct {
	HotelId              string    `protobuf:"bytes,1,opt,name=hotelId,proto3" json:"hotelId,omitempty"`
	Code                 string    `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	InDate               string    `protobuf:"bytes,3,opt,name=inDate,proto3" json:"inDate,omitempty"`
	OutDate              string    `protobuf:"bytes,4,opt,name=outDate,proto3" json:"outDate,omitempty"`
	RoomType             *RoomType `protobuf:"bytes,5,opt,name=roomType,proto3" json:"roomType,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *RatePlan) Reset()         { *m = RatePlan{} }
func (m *RatePlan) String() string { return proto.CompactTextString(m) }
func (*RatePlan) ProtoMessage()    {}
func (*RatePlan) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8fd8716f9769db8, []int{2}
}

func (m *RatePlan) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RatePlan.Unmarshal(m, b)
}
func (m *RatePlan) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RatePlan.Marshal(b, m, deterministic)
}
func (m *RatePlan) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RatePlan.Merge(m, src)
}
func (m *RatePlan) XXX_Size() int {
	return xxx_messageInfo_RatePlan.Size(m)
}
func (m *RatePlan) XXX_DiscardUnknown() {
	xxx_messageInfo_RatePlan.DiscardUnknown(m)
}

var xxx_messageInfo_RatePlan proto.InternalMessageInfo

func (m *RatePlan) GetHotelId() string {
	if m != nil {
		return m.HotelId
	}
	return ""
}

func (m *RatePlan) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *RatePlan) GetInDate() string {
	if m != nil {
		return m.InDate
	}
	return ""
}

func (m *RatePlan) GetOutDate() string {
	if m != nil {
		return m.OutDate
	}
	return ""
}

func (m *RatePlan) GetRoomType() *RoomType {
	if m != nil {
		return m.RoomType
	}
	return nil
}

type RoomType struct {
	BookableRate         float64  `protobuf:"fixed64,1,opt,name=bookableRate,proto3" json:"bookableRate,omitempty"`
	TotalRate            float64  `protobuf:"fixed64,2,opt,name=totalRate,proto3" json:"totalRate,omitempty"`
	TotalRateInclusive   float64  `protobuf:"fixed64,3,opt,name=totalRateInclusive,proto3" json:"totalRateInclusive,omitempty"`
	Code                 string   `protobuf:"bytes,4,opt,name=code,proto3" json:"code,omitempty"`
	Currency             string   `protobuf:"bytes,5,opt,name=currency,proto3" json:"currency,omitempty"`
	RoomDescription      string   `protobuf:"bytes,6,opt,name=roomDescription,proto3" json:"roomDescription,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoomType) Reset()         { *m = RoomType{} }
func (m *RoomType) String() string { return proto.CompactTextString(m) }
func (*RoomType) ProtoMessage()    {}
func (*RoomType) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8fd8716f9769db8, []int{3}
}

func (m *RoomType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoomType.Unmarshal(m, b)
}
func (m *RoomType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoomType.Marshal(b, m, deterministic)
}
func (m *RoomType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoomType.Merge(m, src)
}
func (m *RoomType) XXX_Size() int {
	return xxx_messageInfo_RoomType.Size(m)
}
func (m *RoomType) XXX_DiscardUnknown() {
	xxx_messageInfo_RoomType.DiscardUnknown(m)
}

var xxx_messageInfo_RoomType proto.InternalMessageInfo

func (m *RoomType) GetBookableRate() float64 {
	if m != nil {
		return m.BookableRate
	}
	return 0
}

func (m *RoomType) GetTotalRate() float64 {
	if m != nil {
		return m.TotalRate
	}
	return 0
}

func (m *RoomType) GetTotalRateInclusive() float64 {
	if m != nil {
		return m.TotalRateInclusive
	}
	return 0
}

func (m *RoomType) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *RoomType) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

func (m *RoomType) GetRoomDescription() string {
	if m != nil {
		return m.RoomDescription
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "rate.Request")
	proto.RegisterType((*Result)(nil), "rate.Result")
	proto.RegisterType((*RatePlan)(nil), "rate.RatePlan")
	proto.RegisterType((*RoomType)(nil), "rate.RoomType")
}

func init() { proto.RegisterFile("srv/rate/proto/rate.proto", fileDescriptor_b8fd8716f9769db8) }

var fileDescriptor_b8fd8716f9769db8 = []byte{
	// 324 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xdd, 0x4a, 0xc3, 0x30,
	0x14, 0xc7, 0x89, 0xab, 0x5d, 0x7b, 0x9c, 0x0a, 0xe7, 0x42, 0xe2, 0xf0, 0x62, 0xf4, 0xc6, 0x21,
	0xb2, 0xc1, 0x04, 0x9f, 0x60, 0x20, 0xbb, 0x93, 0x83, 0xe0, 0x75, 0xd7, 0x05, 0x2c, 0xc6, 0x66,
	0x26, 0xe9, 0x60, 0x2f, 0xe2, 0xa3, 0xf9, 0x3c, 0x92, 0x34, 0xed, 0x3a, 0xd1, 0xbb, 0xff, 0x47,
	0x38, 0xfd, 0x25, 0x3d, 0x70, 0x6d, 0xf4, 0x6e, 0xae, 0x73, 0x2b, 0xe6, 0x5b, 0xad, 0xac, 0xf2,
	0x72, 0xe6, 0x25, 0x46, 0x4e, 0x67, 0xaf, 0x30, 0x24, 0xf1, 0x59, 0x0b, 0x63, 0x71, 0x0c, 0xc9,
	0x9b, 0xb2, 0x42, 0xae, 0x36, 0x86, 0xb3, 0xc9, 0x60, 0x9a, 0x52, 0xe7, 0xf1, 0x0a, 0xe2, 0xb2,
	0x5a, 0xe6, 0x56, 0xf0, 0x93, 0x09, 0x9b, 0xa6, 0x14, 0x1c, 0x72, 0x18, 0xaa, 0xda, 0xfa, 0x62,
	0xe0, 0x8b, 0xd6, 0x66, 0x8f, 0x10, 0x93, 0x30, 0xb5, 0xb4, 0x78, 0x0f, 0xa9, 0xfb, 0xd4, 0xb3,
	0xcc, 0xab, 0x66, 0xf0, 0xd9, 0xe2, 0x62, 0xe6, 0x41, 0x28, 0xc4, 0x74, 0x38, 0x90, 0x7d, 0x31,
	0x48, 0xda, 0xdc, 0x8d, 0x0f, 0x08, 0x9c, 0x35, 0xe3, 0x83, 0x45, 0x84, 0xa8, 0x50, 0x9b, 0x16,
	0xc7, 0xeb, 0x1e, 0xe4, 0xe0, 0x3f, 0xc8, 0xe8, 0x08, 0x12, 0xef, 0x20, 0xd1, 0x4a, 0x7d, 0xbc,
	0xec, 0xb7, 0x82, 0x9f, 0x4e, 0x58, 0x8f, 0x2c, 0xa4, 0xd4, 0xf5, 0xd9, 0xb7, 0x03, 0x0b, 0x06,
	0x33, 0x18, 0xad, 0x95, 0x7a, 0xcf, 0xd7, 0x52, 0x38, 0x58, 0x4f, 0xc7, 0xe8, 0x28, 0xc3, 0x1b,
	0x48, 0xad, 0xb2, 0xb9, 0xa4, 0xf6, 0xd9, 0x18, 0x1d, 0x02, 0x9c, 0x01, 0x76, 0x66, 0x55, 0x15,
	0xb2, 0x36, 0xe5, 0xae, 0x01, 0x67, 0xf4, 0x47, 0xd3, 0x5d, 0x38, 0xea, 0x5d, 0x78, 0x0c, 0x49,
	0x51, 0x6b, 0x2d, 0xaa, 0x62, 0xef, 0xf1, 0x53, 0xea, 0x3c, 0x4e, 0xe1, 0xd2, 0xa1, 0x2f, 0x85,
	0x29, 0x74, 0xb9, 0xb5, 0xa5, 0xaa, 0x78, 0xec, 0x8f, 0xfc, 0x8e, 0x17, 0x73, 0x88, 0x3c, 0xd1,
	0x2d, 0x24, 0x4f, 0xc2, 0x3a, 0x69, 0xf0, 0x3c, 0x3c, 0x43, 0xb3, 0x1a, 0xe3, 0x51, 0x6b, 0xdd,
	0x0f, 0x5d, 0xc7, 0x7e, 0x81, 0x1e, 0x7e, 0x02, 0x00, 0x00, 0xff, 0xff, 0xdc, 0xf4, 0x9f, 0xd8,
	0x5d, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RateClient is the client API for Rate service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RateClient interface {
	// GetRates returns rate codes for hotels for a given date range
	GetRates(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Result, error)
}

type rateClient struct {
	cc *grpc.ClientConn
}

func NewRateClient(cc *grpc.ClientConn) RateClient {
	return &rateClient{cc}
}

func (c *rateClient) GetRates(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/rate.Rate/GetRates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RateServer is the server API for Rate service.
type RateServer interface {
	// GetRates returns rate codes for hotels for a given date range
	GetRates(context.Context, *Request) (*Result, error)
}

// UnimplementedRateServer can be embedded to have forward compatible implementations.
type UnimplementedRateServer struct {
}

func (*UnimplementedRateServer) GetRates(ctx context.Context, req *Request) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRates not implemented")
}

func RegisterRateServer(s *grpc.Server, srv RateServer) {
	s.RegisterService(&_Rate_serviceDesc, srv)
}

func _Rate_GetRates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RateServer).GetRates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rate.Rate/GetRates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RateServer).GetRates(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _Rate_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rate.Rate",
	HandlerType: (*RateServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRates",
			Handler:    _Rate_GetRates_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "srv/rate/proto/rate.proto",
}
