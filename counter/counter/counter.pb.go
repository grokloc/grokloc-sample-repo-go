// Code generated by protoc-gen-go. DO NOT EDIT.
// source: counter.proto

package counter

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

// Value contains a countername and its present value.
type Value struct {
	Countername          string   `protobuf:"bytes,1,opt,name=countername,proto3" json:"countername,omitempty"`
	Value                int32    `protobuf:"varint,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Value) Reset()         { *m = Value{} }
func (m *Value) String() string { return proto.CompactTextString(m) }
func (*Value) ProtoMessage()    {}
func (*Value) Descriptor() ([]byte, []int) {
	return fileDescriptor_75dcd656fce7132f, []int{0}
}

func (m *Value) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Value.Unmarshal(m, b)
}
func (m *Value) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Value.Marshal(b, m, deterministic)
}
func (m *Value) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Value.Merge(m, src)
}
func (m *Value) XXX_Size() int {
	return xxx_messageInfo_Value.Size(m)
}
func (m *Value) XXX_DiscardUnknown() {
	xxx_messageInfo_Value.DiscardUnknown(m)
}

var xxx_messageInfo_Value proto.InternalMessageInfo

func (m *Value) GetCountername() string {
	if m != nil {
		return m.Countername
	}
	return ""
}

func (m *Value) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

// Increment contains a countername and a value to increment it by.
// Assume a counter not defined will be incremented from 0 if not defined.
type Increment struct {
	Countername          string   `protobuf:"bytes,1,opt,name=countername,proto3" json:"countername,omitempty"`
	Increment            int32    `protobuf:"varint,2,opt,name=increment,proto3" json:"increment,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Increment) Reset()         { *m = Increment{} }
func (m *Increment) String() string { return proto.CompactTextString(m) }
func (*Increment) ProtoMessage()    {}
func (*Increment) Descriptor() ([]byte, []int) {
	return fileDescriptor_75dcd656fce7132f, []int{1}
}

func (m *Increment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Increment.Unmarshal(m, b)
}
func (m *Increment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Increment.Marshal(b, m, deterministic)
}
func (m *Increment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Increment.Merge(m, src)
}
func (m *Increment) XXX_Size() int {
	return xxx_messageInfo_Increment.Size(m)
}
func (m *Increment) XXX_DiscardUnknown() {
	xxx_messageInfo_Increment.DiscardUnknown(m)
}

var xxx_messageInfo_Increment proto.InternalMessageInfo

func (m *Increment) GetCountername() string {
	if m != nil {
		return m.Countername
	}
	return ""
}

func (m *Increment) GetIncrement() int32 {
	if m != nil {
		return m.Increment
	}
	return 0
}

// Read represents a request for a Value for countername.
type Read struct {
	Countername          string   `protobuf:"bytes,1,opt,name=countername,proto3" json:"countername,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Read) Reset()         { *m = Read{} }
func (m *Read) String() string { return proto.CompactTextString(m) }
func (*Read) ProtoMessage()    {}
func (*Read) Descriptor() ([]byte, []int) {
	return fileDescriptor_75dcd656fce7132f, []int{2}
}

func (m *Read) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Read.Unmarshal(m, b)
}
func (m *Read) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Read.Marshal(b, m, deterministic)
}
func (m *Read) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Read.Merge(m, src)
}
func (m *Read) XXX_Size() int {
	return xxx_messageInfo_Read.Size(m)
}
func (m *Read) XXX_DiscardUnknown() {
	xxx_messageInfo_Read.DiscardUnknown(m)
}

var xxx_messageInfo_Read proto.InternalMessageInfo

func (m *Read) GetCountername() string {
	if m != nil {
		return m.Countername
	}
	return ""
}

// ValueResponse can contain either a CounterValue or an error in case
// the counter is not defined.
type ValueResponse struct {
	Countervalue         *Value   `protobuf:"bytes,1,opt,name=countervalue,proto3" json:"countervalue,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ValueResponse) Reset()         { *m = ValueResponse{} }
func (m *ValueResponse) String() string { return proto.CompactTextString(m) }
func (*ValueResponse) ProtoMessage()    {}
func (*ValueResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_75dcd656fce7132f, []int{3}
}

func (m *ValueResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ValueResponse.Unmarshal(m, b)
}
func (m *ValueResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ValueResponse.Marshal(b, m, deterministic)
}
func (m *ValueResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValueResponse.Merge(m, src)
}
func (m *ValueResponse) XXX_Size() int {
	return xxx_messageInfo_ValueResponse.Size(m)
}
func (m *ValueResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ValueResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ValueResponse proto.InternalMessageInfo

func (m *ValueResponse) GetCountervalue() *Value {
	if m != nil {
		return m.Countervalue
	}
	return nil
}

func init() {
	proto.RegisterType((*Value)(nil), "counter.Value")
	proto.RegisterType((*Increment)(nil), "counter.Increment")
	proto.RegisterType((*Read)(nil), "counter.Read")
	proto.RegisterType((*ValueResponse)(nil), "counter.ValueResponse")
}

func init() { proto.RegisterFile("counter.proto", fileDescriptor_75dcd656fce7132f) }

var fileDescriptor_75dcd656fce7132f = []byte{
	// 207 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0xce, 0x2f, 0xcd,
	0x2b, 0x49, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x72, 0x95, 0xec, 0xb9,
	0x58, 0xc3, 0x12, 0x73, 0x4a, 0x53, 0x85, 0x14, 0xb8, 0xb8, 0xa1, 0x62, 0x79, 0x89, 0xb9, 0xa9,
	0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0xc8, 0x42, 0x42, 0x22, 0x5c, 0xac, 0x65, 0x20, 0xa5,
	0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0xac, 0x41, 0x10, 0x8e, 0x92, 0x37, 0x17, 0xa7, 0x67, 0x5e, 0x72,
	0x51, 0x6a, 0x6e, 0x6a, 0x5e, 0x09, 0x11, 0x86, 0xc8, 0x70, 0x71, 0x66, 0xc2, 0x94, 0x43, 0x0d,
	0x42, 0x08, 0x28, 0x69, 0x70, 0xb1, 0x04, 0xa5, 0x26, 0xa6, 0x10, 0x36, 0x47, 0xc9, 0x99, 0x8b,
	0x17, 0xec, 0xee, 0xa0, 0xd4, 0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x21, 0x23, 0x2e, 0x1e, 0xa8,
	0x3c, 0xc4, 0x91, 0x20, 0x3d, 0xdc, 0x46, 0x7c, 0x7a, 0x30, 0x7f, 0x43, 0x54, 0xa3, 0xa8, 0x31,
	0x6a, 0x66, 0xe4, 0x62, 0x77, 0x86, 0x08, 0x08, 0x39, 0x70, 0x09, 0xc0, 0xfd, 0x01, 0x13, 0x13,
	0x82, 0xeb, 0x86, 0x4b, 0x49, 0x89, 0xa1, 0x99, 0x08, 0xb5, 0x5f, 0x89, 0x41, 0xc8, 0x8c, 0x8b,
	0x1b, 0xe4, 0x78, 0x98, 0x66, 0x5e, 0xb8, 0x42, 0x90, 0x28, 0x6e, 0x7d, 0x49, 0x6c, 0xe0, 0x28,
	0x31, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x2a, 0x71, 0x56, 0x3c, 0xa3, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CounterClient is the client API for Counter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CounterClient interface {
	IncrementCounter(ctx context.Context, in *Increment, opts ...grpc.CallOption) (*ValueResponse, error)
	ReadCounter(ctx context.Context, in *Read, opts ...grpc.CallOption) (*ValueResponse, error)
}

type counterClient struct {
	cc *grpc.ClientConn
}

func NewCounterClient(cc *grpc.ClientConn) CounterClient {
	return &counterClient{cc}
}

func (c *counterClient) IncrementCounter(ctx context.Context, in *Increment, opts ...grpc.CallOption) (*ValueResponse, error) {
	out := new(ValueResponse)
	err := c.cc.Invoke(ctx, "/counter.Counter/IncrementCounter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *counterClient) ReadCounter(ctx context.Context, in *Read, opts ...grpc.CallOption) (*ValueResponse, error) {
	out := new(ValueResponse)
	err := c.cc.Invoke(ctx, "/counter.Counter/ReadCounter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CounterServer is the server API for Counter service.
type CounterServer interface {
	IncrementCounter(context.Context, *Increment) (*ValueResponse, error)
	ReadCounter(context.Context, *Read) (*ValueResponse, error)
}

// UnimplementedCounterServer can be embedded to have forward compatible implementations.
type UnimplementedCounterServer struct {
}

func (*UnimplementedCounterServer) IncrementCounter(ctx context.Context, req *Increment) (*ValueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IncrementCounter not implemented")
}
func (*UnimplementedCounterServer) ReadCounter(ctx context.Context, req *Read) (*ValueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadCounter not implemented")
}

func RegisterCounterServer(s *grpc.Server, srv CounterServer) {
	s.RegisterService(&_Counter_serviceDesc, srv)
}

func _Counter_IncrementCounter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Increment)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CounterServer).IncrementCounter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/counter.Counter/IncrementCounter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CounterServer).IncrementCounter(ctx, req.(*Increment))
	}
	return interceptor(ctx, in, info, handler)
}

func _Counter_ReadCounter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Read)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CounterServer).ReadCounter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/counter.Counter/ReadCounter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CounterServer).ReadCounter(ctx, req.(*Read))
	}
	return interceptor(ctx, in, info, handler)
}

var _Counter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "counter.Counter",
	HandlerType: (*CounterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IncrementCounter",
			Handler:    _Counter_IncrementCounter_Handler,
		},
		{
			MethodName: "ReadCounter",
			Handler:    _Counter_ReadCounter_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "counter.proto",
}
