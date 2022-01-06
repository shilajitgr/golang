// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: PrimeDecomposeStream/primepb/prime.proto

package primeDecomposepb

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Number struct {
	Val                  int32    `protobuf:"varint,1,opt,name=val,proto3" json:"val,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Number) Reset()         { *m = Number{} }
func (m *Number) String() string { return proto.CompactTextString(m) }
func (*Number) ProtoMessage()    {}
func (*Number) Descriptor() ([]byte, []int) {
	return fileDescriptor_d713ea47f10bdb24, []int{0}
}
func (m *Number) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Number.Unmarshal(m, b)
}
func (m *Number) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Number.Marshal(b, m, deterministic)
}
func (m *Number) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Number.Merge(m, src)
}
func (m *Number) XXX_Size() int {
	return xxx_messageInfo_Number.Size(m)
}
func (m *Number) XXX_DiscardUnknown() {
	xxx_messageInfo_Number.DiscardUnknown(m)
}

var xxx_messageInfo_Number proto.InternalMessageInfo

func (m *Number) GetVal() int32 {
	if m != nil {
		return m.Val
	}
	return 0
}

type PrimeRequest struct {
	Input                *Number  `protobuf:"bytes,1,opt,name=input,proto3" json:"input,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrimeRequest) Reset()         { *m = PrimeRequest{} }
func (m *PrimeRequest) String() string { return proto.CompactTextString(m) }
func (*PrimeRequest) ProtoMessage()    {}
func (*PrimeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d713ea47f10bdb24, []int{1}
}
func (m *PrimeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrimeRequest.Unmarshal(m, b)
}
func (m *PrimeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrimeRequest.Marshal(b, m, deterministic)
}
func (m *PrimeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrimeRequest.Merge(m, src)
}
func (m *PrimeRequest) XXX_Size() int {
	return xxx_messageInfo_PrimeRequest.Size(m)
}
func (m *PrimeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_PrimeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_PrimeRequest proto.InternalMessageInfo

func (m *PrimeRequest) GetInput() *Number {
	if m != nil {
		return m.Input
	}
	return nil
}

type PrimeResponse struct {
	PrimeFactor          int32    `protobuf:"varint,1,opt,name=primeFactor,proto3" json:"primeFactor,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PrimeResponse) Reset()         { *m = PrimeResponse{} }
func (m *PrimeResponse) String() string { return proto.CompactTextString(m) }
func (*PrimeResponse) ProtoMessage()    {}
func (*PrimeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d713ea47f10bdb24, []int{2}
}
func (m *PrimeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PrimeResponse.Unmarshal(m, b)
}
func (m *PrimeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PrimeResponse.Marshal(b, m, deterministic)
}
func (m *PrimeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PrimeResponse.Merge(m, src)
}
func (m *PrimeResponse) XXX_Size() int {
	return xxx_messageInfo_PrimeResponse.Size(m)
}
func (m *PrimeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_PrimeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_PrimeResponse proto.InternalMessageInfo

func (m *PrimeResponse) GetPrimeFactor() int32 {
	if m != nil {
		return m.PrimeFactor
	}
	return 0
}

type AvgRequest struct {
	Input                *Number  `protobuf:"bytes,1,opt,name=input,proto3" json:"input,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AvgRequest) Reset()         { *m = AvgRequest{} }
func (m *AvgRequest) String() string { return proto.CompactTextString(m) }
func (*AvgRequest) ProtoMessage()    {}
func (*AvgRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d713ea47f10bdb24, []int{3}
}
func (m *AvgRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AvgRequest.Unmarshal(m, b)
}
func (m *AvgRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AvgRequest.Marshal(b, m, deterministic)
}
func (m *AvgRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AvgRequest.Merge(m, src)
}
func (m *AvgRequest) XXX_Size() int {
	return xxx_messageInfo_AvgRequest.Size(m)
}
func (m *AvgRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AvgRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AvgRequest proto.InternalMessageInfo

func (m *AvgRequest) GetInput() *Number {
	if m != nil {
		return m.Input
	}
	return nil
}

type AvgResponse struct {
	Avg                  float32  `protobuf:"fixed32,1,opt,name=avg,proto3" json:"avg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AvgResponse) Reset()         { *m = AvgResponse{} }
func (m *AvgResponse) String() string { return proto.CompactTextString(m) }
func (*AvgResponse) ProtoMessage()    {}
func (*AvgResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d713ea47f10bdb24, []int{4}
}
func (m *AvgResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AvgResponse.Unmarshal(m, b)
}
func (m *AvgResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AvgResponse.Marshal(b, m, deterministic)
}
func (m *AvgResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AvgResponse.Merge(m, src)
}
func (m *AvgResponse) XXX_Size() int {
	return xxx_messageInfo_AvgResponse.Size(m)
}
func (m *AvgResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AvgResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AvgResponse proto.InternalMessageInfo

func (m *AvgResponse) GetAvg() float32 {
	if m != nil {
		return m.Avg
	}
	return 0
}

type FindMaxRequest struct {
	Val                  int32    `protobuf:"varint,1,opt,name=val,proto3" json:"val,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindMaxRequest) Reset()         { *m = FindMaxRequest{} }
func (m *FindMaxRequest) String() string { return proto.CompactTextString(m) }
func (*FindMaxRequest) ProtoMessage()    {}
func (*FindMaxRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d713ea47f10bdb24, []int{5}
}
func (m *FindMaxRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindMaxRequest.Unmarshal(m, b)
}
func (m *FindMaxRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindMaxRequest.Marshal(b, m, deterministic)
}
func (m *FindMaxRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindMaxRequest.Merge(m, src)
}
func (m *FindMaxRequest) XXX_Size() int {
	return xxx_messageInfo_FindMaxRequest.Size(m)
}
func (m *FindMaxRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_FindMaxRequest.DiscardUnknown(m)
}

var xxx_messageInfo_FindMaxRequest proto.InternalMessageInfo

func (m *FindMaxRequest) GetVal() int32 {
	if m != nil {
		return m.Val
	}
	return 0
}

type FindMaxResponse struct {
	Max                  int32    `protobuf:"varint,1,opt,name=max,proto3" json:"max,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FindMaxResponse) Reset()         { *m = FindMaxResponse{} }
func (m *FindMaxResponse) String() string { return proto.CompactTextString(m) }
func (*FindMaxResponse) ProtoMessage()    {}
func (*FindMaxResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d713ea47f10bdb24, []int{6}
}
func (m *FindMaxResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FindMaxResponse.Unmarshal(m, b)
}
func (m *FindMaxResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FindMaxResponse.Marshal(b, m, deterministic)
}
func (m *FindMaxResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FindMaxResponse.Merge(m, src)
}
func (m *FindMaxResponse) XXX_Size() int {
	return xxx_messageInfo_FindMaxResponse.Size(m)
}
func (m *FindMaxResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_FindMaxResponse.DiscardUnknown(m)
}

var xxx_messageInfo_FindMaxResponse proto.InternalMessageInfo

func (m *FindMaxResponse) GetMax() int32 {
	if m != nil {
		return m.Max
	}
	return 0
}

type SqrtRequest struct {
	Val                  int32    `protobuf:"varint,1,opt,name=val,proto3" json:"val,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SqrtRequest) Reset()         { *m = SqrtRequest{} }
func (m *SqrtRequest) String() string { return proto.CompactTextString(m) }
func (*SqrtRequest) ProtoMessage()    {}
func (*SqrtRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d713ea47f10bdb24, []int{7}
}
func (m *SqrtRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SqrtRequest.Unmarshal(m, b)
}
func (m *SqrtRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SqrtRequest.Marshal(b, m, deterministic)
}
func (m *SqrtRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SqrtRequest.Merge(m, src)
}
func (m *SqrtRequest) XXX_Size() int {
	return xxx_messageInfo_SqrtRequest.Size(m)
}
func (m *SqrtRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SqrtRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SqrtRequest proto.InternalMessageInfo

func (m *SqrtRequest) GetVal() int32 {
	if m != nil {
		return m.Val
	}
	return 0
}

type SqrtResponse struct {
	Sqrt                 float32  `protobuf:"fixed32,1,opt,name=sqrt,proto3" json:"sqrt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SqrtResponse) Reset()         { *m = SqrtResponse{} }
func (m *SqrtResponse) String() string { return proto.CompactTextString(m) }
func (*SqrtResponse) ProtoMessage()    {}
func (*SqrtResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d713ea47f10bdb24, []int{8}
}
func (m *SqrtResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SqrtResponse.Unmarshal(m, b)
}
func (m *SqrtResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SqrtResponse.Marshal(b, m, deterministic)
}
func (m *SqrtResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SqrtResponse.Merge(m, src)
}
func (m *SqrtResponse) XXX_Size() int {
	return xxx_messageInfo_SqrtResponse.Size(m)
}
func (m *SqrtResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SqrtResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SqrtResponse proto.InternalMessageInfo

func (m *SqrtResponse) GetSqrt() float32 {
	if m != nil {
		return m.Sqrt
	}
	return 0
}

func init() {
	proto.RegisterType((*Number)(nil), "primeDecompose.Number")
	proto.RegisterType((*PrimeRequest)(nil), "primeDecompose.primeRequest")
	proto.RegisterType((*PrimeResponse)(nil), "primeDecompose.primeResponse")
	proto.RegisterType((*AvgRequest)(nil), "primeDecompose.avgRequest")
	proto.RegisterType((*AvgResponse)(nil), "primeDecompose.avgResponse")
	proto.RegisterType((*FindMaxRequest)(nil), "primeDecompose.findMaxRequest")
	proto.RegisterType((*FindMaxResponse)(nil), "primeDecompose.findMaxResponse")
	proto.RegisterType((*SqrtRequest)(nil), "primeDecompose.sqrtRequest")
	proto.RegisterType((*SqrtResponse)(nil), "primeDecompose.sqrtResponse")
}

func init() {
	proto.RegisterFile("PrimeDecomposeStream/primepb/prime.proto", fileDescriptor_d713ea47f10bdb24)
}

var fileDescriptor_d713ea47f10bdb24 = []byte{
	// 346 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x4f, 0x4f, 0xf2, 0x40,
	0x10, 0xc6, 0xdf, 0xc2, 0x0b, 0x26, 0x03, 0x22, 0xd9, 0x83, 0x31, 0x15, 0x85, 0xac, 0x17, 0x0e,
	0x06, 0x15, 0x6f, 0xc6, 0x8b, 0x7f, 0x42, 0xc2, 0x41, 0x43, 0xca, 0xcd, 0xdb, 0x16, 0x47, 0xd2,
	0xc4, 0x65, 0x97, 0xed, 0xb6, 0xe1, 0xeb, 0xf9, 0xcd, 0x4c, 0x77, 0xb7, 0x40, 0x81, 0x5e, 0x3c,
	0x75, 0x3a, 0xf3, 0x9b, 0xe7, 0x99, 0x9d, 0x5d, 0xe8, 0x4f, 0x54, 0xc4, 0xf1, 0x15, 0x67, 0x82,
	0x4b, 0x11, 0xe3, 0x54, 0x2b, 0x64, 0xfc, 0x46, 0x66, 0x49, 0x19, 0xda, 0xef, 0x40, 0x2a, 0xa1,
	0x05, 0x69, 0xc9, 0x02, 0x49, 0x7d, 0xa8, 0xbf, 0x27, 0x3c, 0x44, 0x45, 0xda, 0x50, 0x4d, 0xd9,
	0xf7, 0x99, 0xd7, 0xf3, 0xfa, 0xb5, 0x20, 0x0b, 0xe9, 0x23, 0x34, 0x0d, 0x1d, 0xe0, 0x32, 0xc1,
	0x58, 0x93, 0x6b, 0xa8, 0x45, 0x0b, 0x99, 0x68, 0xc3, 0x34, 0x86, 0xa7, 0x83, 0xa2, 0xd6, 0xc0,
	0x0a, 0x05, 0x16, 0xa2, 0x77, 0x70, 0xec, 0xba, 0x63, 0x29, 0x16, 0x31, 0x92, 0x1e, 0x34, 0x4c,
	0x62, 0xc4, 0x66, 0x5a, 0x28, 0x67, 0xb4, 0x9d, 0xa2, 0x0f, 0x00, 0x2c, 0x9d, 0xff, 0xcd, 0xae,
	0x0b, 0x0d, 0xd3, 0xeb, 0xcc, 0xda, 0x50, 0x65, 0xe9, 0xdc, 0xb4, 0x56, 0x82, 0x2c, 0xa4, 0x14,
	0x5a, 0x5f, 0xd1, 0xe2, 0xf3, 0x8d, 0xad, 0x72, 0x83, 0xfd, 0x13, 0x5f, 0xc1, 0xc9, 0x9a, 0xd9,
	0x08, 0x71, 0xb6, 0xca, 0x21, 0xce, 0x56, 0x99, 0x53, 0xbc, 0x54, 0xba, 0x5c, 0x85, 0x42, 0xd3,
	0x02, 0x4e, 0x82, 0xc0, 0xff, 0xec, 0xdf, 0x0d, 0x63, 0xe2, 0xe1, 0x4f, 0xc5, 0x2d, 0x77, 0x8a,
	0x2a, 0x8d, 0x66, 0x48, 0xc6, 0xeb, 0x8b, 0xe8, 0xec, 0x1e, 0x74, 0xfb, 0x12, 0xfc, 0x8b, 0x92,
	0xaa, 0xf5, 0xa2, 0xff, 0x6e, 0x3d, 0x32, 0x06, 0x78, 0x11, 0x5c, 0x26, 0x1a, 0x9f, 0xd2, 0x39,
	0xf1, 0x77, 0x1b, 0x36, 0x2b, 0xf6, 0xcf, 0x0f, 0xd6, 0x72, 0xa9, 0xbe, 0x47, 0x26, 0x70, 0x34,
	0xb2, 0x0b, 0x21, 0x97, 0xbb, 0x6c, 0x71, 0x9b, 0x7e, 0xb7, 0xb4, 0xbe, 0xd1, 0xb3, 0xc3, 0x4d,
	0x97, 0x09, 0x53, 0x18, 0x08, 0xa1, 0xc9, 0xde, 0x00, 0x5b, 0x9b, 0xf5, 0x3b, 0x87, 0x8b, 0xb9,
	0xdc, 0x33, 0xf9, 0x68, 0x17, 0x01, 0x19, 0x86, 0x75, 0xf3, 0xcc, 0xef, 0x7f, 0x03, 0x00, 0x00,
	0xff, 0xff, 0x53, 0xbb, 0xa8, 0xa8, 0x12, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PrimeServiceClient is the client API for PrimeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PrimeServiceClient interface {
	//server stream
	Number(ctx context.Context, in *PrimeRequest, opts ...grpc.CallOption) (PrimeService_NumberClient, error)
	// client Stream
	ComputeAvg(ctx context.Context, opts ...grpc.CallOption) (PrimeService_ComputeAvgClient, error)
	//bi-directional stream
	FindMax(ctx context.Context, opts ...grpc.CallOption) (PrimeService_FindMaxClient, error)
	//unary : find sqaure root
	SquareRoot(ctx context.Context, in *SqrtRequest, opts ...grpc.CallOption) (*SqrtResponse, error)
}

type primeServiceClient struct {
	cc *grpc.ClientConn
}

func NewPrimeServiceClient(cc *grpc.ClientConn) PrimeServiceClient {
	return &primeServiceClient{cc}
}

func (c *primeServiceClient) Number(ctx context.Context, in *PrimeRequest, opts ...grpc.CallOption) (PrimeService_NumberClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PrimeService_serviceDesc.Streams[0], "/primeDecompose.primeService/Number", opts...)
	if err != nil {
		return nil, err
	}
	x := &primeServiceNumberClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type PrimeService_NumberClient interface {
	Recv() (*PrimeResponse, error)
	grpc.ClientStream
}

type primeServiceNumberClient struct {
	grpc.ClientStream
}

func (x *primeServiceNumberClient) Recv() (*PrimeResponse, error) {
	m := new(PrimeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *primeServiceClient) ComputeAvg(ctx context.Context, opts ...grpc.CallOption) (PrimeService_ComputeAvgClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PrimeService_serviceDesc.Streams[1], "/primeDecompose.primeService/ComputeAvg", opts...)
	if err != nil {
		return nil, err
	}
	x := &primeServiceComputeAvgClient{stream}
	return x, nil
}

type PrimeService_ComputeAvgClient interface {
	Send(*AvgRequest) error
	CloseAndRecv() (*AvgResponse, error)
	grpc.ClientStream
}

type primeServiceComputeAvgClient struct {
	grpc.ClientStream
}

func (x *primeServiceComputeAvgClient) Send(m *AvgRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *primeServiceComputeAvgClient) CloseAndRecv() (*AvgResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(AvgResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *primeServiceClient) FindMax(ctx context.Context, opts ...grpc.CallOption) (PrimeService_FindMaxClient, error) {
	stream, err := c.cc.NewStream(ctx, &_PrimeService_serviceDesc.Streams[2], "/primeDecompose.primeService/FindMax", opts...)
	if err != nil {
		return nil, err
	}
	x := &primeServiceFindMaxClient{stream}
	return x, nil
}

type PrimeService_FindMaxClient interface {
	Send(*FindMaxRequest) error
	Recv() (*FindMaxResponse, error)
	grpc.ClientStream
}

type primeServiceFindMaxClient struct {
	grpc.ClientStream
}

func (x *primeServiceFindMaxClient) Send(m *FindMaxRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *primeServiceFindMaxClient) Recv() (*FindMaxResponse, error) {
	m := new(FindMaxResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *primeServiceClient) SquareRoot(ctx context.Context, in *SqrtRequest, opts ...grpc.CallOption) (*SqrtResponse, error) {
	out := new(SqrtResponse)
	err := c.cc.Invoke(ctx, "/primeDecompose.primeService/SquareRoot", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PrimeServiceServer is the server API for PrimeService service.
type PrimeServiceServer interface {
	//server stream
	Number(*PrimeRequest, PrimeService_NumberServer) error
	// client Stream
	ComputeAvg(PrimeService_ComputeAvgServer) error
	//bi-directional stream
	FindMax(PrimeService_FindMaxServer) error
	//unary : find sqaure root
	SquareRoot(context.Context, *SqrtRequest) (*SqrtResponse, error)
}

// UnimplementedPrimeServiceServer can be embedded to have forward compatible implementations.
type UnimplementedPrimeServiceServer struct {
}

func (*UnimplementedPrimeServiceServer) Number(req *PrimeRequest, srv PrimeService_NumberServer) error {
	return status.Errorf(codes.Unimplemented, "method Number not implemented")
}
func (*UnimplementedPrimeServiceServer) ComputeAvg(srv PrimeService_ComputeAvgServer) error {
	return status.Errorf(codes.Unimplemented, "method ComputeAvg not implemented")
}
func (*UnimplementedPrimeServiceServer) FindMax(srv PrimeService_FindMaxServer) error {
	return status.Errorf(codes.Unimplemented, "method FindMax not implemented")
}
func (*UnimplementedPrimeServiceServer) SquareRoot(ctx context.Context, req *SqrtRequest) (*SqrtResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SquareRoot not implemented")
}

func RegisterPrimeServiceServer(s *grpc.Server, srv PrimeServiceServer) {
	s.RegisterService(&_PrimeService_serviceDesc, srv)
}

func _PrimeService_Number_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PrimeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PrimeServiceServer).Number(m, &primeServiceNumberServer{stream})
}

type PrimeService_NumberServer interface {
	Send(*PrimeResponse) error
	grpc.ServerStream
}

type primeServiceNumberServer struct {
	grpc.ServerStream
}

func (x *primeServiceNumberServer) Send(m *PrimeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _PrimeService_ComputeAvg_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PrimeServiceServer).ComputeAvg(&primeServiceComputeAvgServer{stream})
}

type PrimeService_ComputeAvgServer interface {
	SendAndClose(*AvgResponse) error
	Recv() (*AvgRequest, error)
	grpc.ServerStream
}

type primeServiceComputeAvgServer struct {
	grpc.ServerStream
}

func (x *primeServiceComputeAvgServer) SendAndClose(m *AvgResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *primeServiceComputeAvgServer) Recv() (*AvgRequest, error) {
	m := new(AvgRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _PrimeService_FindMax_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PrimeServiceServer).FindMax(&primeServiceFindMaxServer{stream})
}

type PrimeService_FindMaxServer interface {
	Send(*FindMaxResponse) error
	Recv() (*FindMaxRequest, error)
	grpc.ServerStream
}

type primeServiceFindMaxServer struct {
	grpc.ServerStream
}

func (x *primeServiceFindMaxServer) Send(m *FindMaxResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *primeServiceFindMaxServer) Recv() (*FindMaxRequest, error) {
	m := new(FindMaxRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _PrimeService_SquareRoot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SqrtRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PrimeServiceServer).SquareRoot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/primeDecompose.primeService/SquareRoot",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PrimeServiceServer).SquareRoot(ctx, req.(*SqrtRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _PrimeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "primeDecompose.primeService",
	HandlerType: (*PrimeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SquareRoot",
			Handler:    _PrimeService_SquareRoot_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Number",
			Handler:       _PrimeService_Number_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ComputeAvg",
			Handler:       _PrimeService_ComputeAvg_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "FindMax",
			Handler:       _PrimeService_FindMax_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "PrimeDecomposeStream/primepb/prime.proto",
}
