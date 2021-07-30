// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: launch/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
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

// this line is used by starport scaffolding # proto/tx/message
type MsgRequestRemoveValidator struct {
	ChainID          string `protobuf:"bytes,1,opt,name=chainID,proto3" json:"chainID,omitempty"`
	ValidatorAddress string `protobuf:"bytes,2,opt,name=validatorAddress,proto3" json:"validatorAddress,omitempty"`
}

func (m *MsgRequestRemoveValidator) Reset()         { *m = MsgRequestRemoveValidator{} }
func (m *MsgRequestRemoveValidator) String() string { return proto.CompactTextString(m) }
func (*MsgRequestRemoveValidator) ProtoMessage()    {}
func (*MsgRequestRemoveValidator) Descriptor() ([]byte, []int) {
	return fileDescriptor_6adab5ffa522f022, []int{0}
}
func (m *MsgRequestRemoveValidator) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRequestRemoveValidator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRequestRemoveValidator.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRequestRemoveValidator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRequestRemoveValidator.Merge(m, src)
}
func (m *MsgRequestRemoveValidator) XXX_Size() int {
	return m.Size()
}
func (m *MsgRequestRemoveValidator) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRequestRemoveValidator.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRequestRemoveValidator proto.InternalMessageInfo

func (m *MsgRequestRemoveValidator) GetChainID() string {
	if m != nil {
		return m.ChainID
	}
	return ""
}

func (m *MsgRequestRemoveValidator) GetValidatorAddress() string {
	if m != nil {
		return m.ValidatorAddress
	}
	return ""
}

type MsgRequestRemoveValidatorResponse struct {
	RequestID uint64 `protobuf:"varint,1,opt,name=requestID,proto3" json:"requestID,omitempty"`
}

func (m *MsgRequestRemoveValidatorResponse) Reset()         { *m = MsgRequestRemoveValidatorResponse{} }
func (m *MsgRequestRemoveValidatorResponse) String() string { return proto.CompactTextString(m) }
func (*MsgRequestRemoveValidatorResponse) ProtoMessage()    {}
func (*MsgRequestRemoveValidatorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6adab5ffa522f022, []int{1}
}
func (m *MsgRequestRemoveValidatorResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRequestRemoveValidatorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRequestRemoveValidatorResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRequestRemoveValidatorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRequestRemoveValidatorResponse.Merge(m, src)
}
func (m *MsgRequestRemoveValidatorResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgRequestRemoveValidatorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRequestRemoveValidatorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRequestRemoveValidatorResponse proto.InternalMessageInfo

func (m *MsgRequestRemoveValidatorResponse) GetRequestID() uint64 {
	if m != nil {
		return m.RequestID
	}
	return 0
}

type MsgCreateChain struct {
	Coordinator string `protobuf:"bytes,1,opt,name=coordinator,proto3" json:"coordinator,omitempty"`
	ChainName   string `protobuf:"bytes,2,opt,name=chainName,proto3" json:"chainName,omitempty"`
	SourceURL   string `protobuf:"bytes,3,opt,name=sourceURL,proto3" json:"sourceURL,omitempty"`
	SourceHash  string `protobuf:"bytes,4,opt,name=sourceHash,proto3" json:"sourceHash,omitempty"`
	GenesisURL  string `protobuf:"bytes,5,opt,name=genesisURL,proto3" json:"genesisURL,omitempty"`
	GenesisHash string `protobuf:"bytes,6,opt,name=genesisHash,proto3" json:"genesisHash,omitempty"`
}

func (m *MsgCreateChain) Reset()         { *m = MsgCreateChain{} }
func (m *MsgCreateChain) String() string { return proto.CompactTextString(m) }
func (*MsgCreateChain) ProtoMessage()    {}
func (*MsgCreateChain) Descriptor() ([]byte, []int) {
	return fileDescriptor_6adab5ffa522f022, []int{2}
}
func (m *MsgCreateChain) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateChain) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateChain.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateChain) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateChain.Merge(m, src)
}
func (m *MsgCreateChain) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateChain) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateChain.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateChain proto.InternalMessageInfo

func (m *MsgCreateChain) GetCoordinator() string {
	if m != nil {
		return m.Coordinator
	}
	return ""
}

func (m *MsgCreateChain) GetChainName() string {
	if m != nil {
		return m.ChainName
	}
	return ""
}

func (m *MsgCreateChain) GetSourceURL() string {
	if m != nil {
		return m.SourceURL
	}
	return ""
}

func (m *MsgCreateChain) GetSourceHash() string {
	if m != nil {
		return m.SourceHash
	}
	return ""
}

func (m *MsgCreateChain) GetGenesisURL() string {
	if m != nil {
		return m.GenesisURL
	}
	return ""
}

func (m *MsgCreateChain) GetGenesisHash() string {
	if m != nil {
		return m.GenesisHash
	}
	return ""
}

type MsgCreateChainResponse struct {
	ChainID string `protobuf:"bytes,1,opt,name=chainID,proto3" json:"chainID,omitempty"`
}

func (m *MsgCreateChainResponse) Reset()         { *m = MsgCreateChainResponse{} }
func (m *MsgCreateChainResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreateChainResponse) ProtoMessage()    {}
func (*MsgCreateChainResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6adab5ffa522f022, []int{3}
}
func (m *MsgCreateChainResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateChainResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateChainResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateChainResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateChainResponse.Merge(m, src)
}
func (m *MsgCreateChainResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateChainResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateChainResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateChainResponse proto.InternalMessageInfo

func (m *MsgCreateChainResponse) GetChainID() string {
	if m != nil {
		return m.ChainID
	}
	return ""
}

func init() {
	proto.RegisterType((*MsgRequestRemoveValidator)(nil), "tendermint.spn.launch.MsgRequestRemoveValidator")
	proto.RegisterType((*MsgRequestRemoveValidatorResponse)(nil), "tendermint.spn.launch.MsgRequestRemoveValidatorResponse")
	proto.RegisterType((*MsgCreateChain)(nil), "tendermint.spn.launch.MsgCreateChain")
	proto.RegisterType((*MsgCreateChainResponse)(nil), "tendermint.spn.launch.MsgCreateChainResponse")
}

func init() { proto.RegisterFile("launch/tx.proto", fileDescriptor_6adab5ffa522f022) }

var fileDescriptor_6adab5ffa522f022 = []byte{
	// 377 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xcf, 0x4a, 0xeb, 0x40,
	0x14, 0xc6, 0x3b, 0xb7, 0xbd, 0xbd, 0x74, 0x0a, 0x57, 0x09, 0x58, 0x62, 0x91, 0x50, 0x03, 0x42,
	0x11, 0x4c, 0xa4, 0x6e, 0xdc, 0xb6, 0x75, 0xa1, 0x60, 0x5d, 0x04, 0x74, 0xe1, 0x6e, 0x9a, 0x1c,
	0x92, 0x40, 0x33, 0x13, 0x67, 0x26, 0xa5, 0xee, 0x7d, 0x00, 0x1f, 0xcb, 0x65, 0x97, 0x2e, 0xa5,
	0x7d, 0x02, 0xdf, 0x40, 0xf2, 0xaf, 0x49, 0xb1, 0x01, 0x5d, 0x9e, 0xef, 0x3b, 0xe7, 0x9b, 0x5f,
	0x4e, 0x66, 0xf0, 0xde, 0x8c, 0x44, 0xd4, 0xf6, 0x4c, 0xb9, 0x30, 0x42, 0xce, 0x24, 0x53, 0x0e,
	0x24, 0x50, 0x07, 0x78, 0xe0, 0x53, 0x69, 0x88, 0x90, 0x1a, 0xa9, 0xaf, 0x13, 0x7c, 0x38, 0x11,
	0xae, 0x05, 0x4f, 0x11, 0x08, 0x69, 0x41, 0xc0, 0xe6, 0xf0, 0x40, 0x66, 0xbe, 0x43, 0x24, 0xe3,
	0x8a, 0x8a, 0xff, 0xd9, 0x1e, 0xf1, 0xe9, 0xcd, 0x95, 0x8a, 0x7a, 0xa8, 0xdf, 0xb2, 0xf2, 0x52,
	0x39, 0xc5, 0xfb, 0xf3, 0xbc, 0x6d, 0xe8, 0x38, 0x1c, 0x84, 0x50, 0xff, 0x24, 0x2d, 0xdf, 0x74,
	0x7d, 0x88, 0x8f, 0x2b, 0x8f, 0xb0, 0x40, 0x84, 0x8c, 0x0a, 0x50, 0x8e, 0x70, 0x8b, 0xa7, 0x1d,
	0xd9, 0x61, 0x0d, 0xab, 0x10, 0xf4, 0x25, 0xc2, 0xff, 0x27, 0xc2, 0x1d, 0x73, 0x20, 0x12, 0xc6,
	0x31, 0x83, 0xd2, 0xc3, 0x6d, 0x9b, 0x31, 0xee, 0xf8, 0x34, 0xce, 0xc9, 0xf8, 0xca, 0x52, 0x1c,
	0x99, 0xe0, 0xde, 0x91, 0x00, 0x32, 0xb8, 0x42, 0x88, 0x5d, 0xc1, 0x22, 0x6e, 0xc3, 0xbd, 0x75,
	0xab, 0xd6, 0x53, 0x77, 0x23, 0x28, 0x1a, 0xc6, 0x69, 0x71, 0x4d, 0x84, 0xa7, 0x36, 0x12, 0xbb,
	0xa4, 0xc4, 0xbe, 0x0b, 0x14, 0x84, 0x2f, 0xe2, 0xf1, 0xbf, 0xa9, 0x5f, 0x28, 0x31, 0x5d, 0x56,
	0x25, 0x01, 0xcd, 0x94, 0xae, 0x24, 0xe9, 0x03, 0xdc, 0xd9, 0xfe, 0xa2, 0xcd, 0x2a, 0x2a, 0xb7,
	0x3e, 0xf8, 0x44, 0xb8, 0x3e, 0x11, 0xae, 0x62, 0xe3, 0x76, 0x79, 0x15, 0x27, 0xc6, 0xce, 0x7f,
	0x6b, 0x6c, 0xe7, 0x77, 0xcf, 0x7e, 0xd4, 0xb6, 0xc1, 0x78, 0x41, 0xb8, 0x53, 0x71, 0x2f, 0xce,
	0xab, 0x93, 0x76, 0x4f, 0x74, 0x2f, 0x7f, 0x3b, 0x91, 0x63, 0x8c, 0x46, 0x6f, 0x2b, 0x0d, 0x2d,
	0x57, 0x1a, 0xfa, 0x58, 0x69, 0xe8, 0x75, 0xad, 0xd5, 0x96, 0x6b, 0xad, 0xf6, 0xbe, 0xd6, 0x6a,
	0x8f, 0x7d, 0xd7, 0x97, 0x5e, 0x34, 0x35, 0x6c, 0x16, 0x98, 0x45, 0xba, 0x29, 0x42, 0x6a, 0x2e,
	0xcc, 0xfc, 0xfa, 0x3f, 0x87, 0x20, 0xa6, 0xcd, 0xe4, 0x09, 0x5c, 0x7c, 0x05, 0x00, 0x00, 0xff,
	0xff, 0xdb, 0x97, 0x45, 0xc6, 0x15, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// this line is used by starport scaffolding # proto/tx/rpc
	CreateChain(ctx context.Context, in *MsgCreateChain, opts ...grpc.CallOption) (*MsgCreateChainResponse, error)
	RequestRemoveValidator(ctx context.Context, in *MsgRequestRemoveValidator, opts ...grpc.CallOption) (*MsgRequestRemoveValidatorResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreateChain(ctx context.Context, in *MsgCreateChain, opts ...grpc.CallOption) (*MsgCreateChainResponse, error) {
	out := new(MsgCreateChainResponse)
	err := c.cc.Invoke(ctx, "/tendermint.spn.launch.Msg/CreateChain", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) RequestRemoveValidator(ctx context.Context, in *MsgRequestRemoveValidator, opts ...grpc.CallOption) (*MsgRequestRemoveValidatorResponse, error) {
	out := new(MsgRequestRemoveValidatorResponse)
	err := c.cc.Invoke(ctx, "/tendermint.spn.launch.Msg/RequestRemoveValidator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// this line is used by starport scaffolding # proto/tx/rpc
	CreateChain(context.Context, *MsgCreateChain) (*MsgCreateChainResponse, error)
	RequestRemoveValidator(context.Context, *MsgRequestRemoveValidator) (*MsgRequestRemoveValidatorResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) CreateChain(ctx context.Context, req *MsgCreateChain) (*MsgCreateChainResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateChain not implemented")
}
func (*UnimplementedMsgServer) RequestRemoveValidator(ctx context.Context, req *MsgRequestRemoveValidator) (*MsgRequestRemoveValidatorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestRemoveValidator not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_CreateChain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateChain)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateChain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tendermint.spn.launch.Msg/CreateChain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateChain(ctx, req.(*MsgCreateChain))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_RequestRemoveValidator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRequestRemoveValidator)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RequestRemoveValidator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tendermint.spn.launch.Msg/RequestRemoveValidator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RequestRemoveValidator(ctx, req.(*MsgRequestRemoveValidator))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tendermint.spn.launch.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateChain",
			Handler:    _Msg_CreateChain_Handler,
		},
		{
			MethodName: "RequestRemoveValidator",
			Handler:    _Msg_RequestRemoveValidator_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "launch/tx.proto",
}

func (m *MsgRequestRemoveValidator) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRequestRemoveValidator) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRequestRemoveValidator) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ValidatorAddress) > 0 {
		i -= len(m.ValidatorAddress)
		copy(dAtA[i:], m.ValidatorAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ValidatorAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ChainID) > 0 {
		i -= len(m.ChainID)
		copy(dAtA[i:], m.ChainID)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ChainID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgRequestRemoveValidatorResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRequestRemoveValidatorResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRequestRemoveValidatorResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.RequestID != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.RequestID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MsgCreateChain) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateChain) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateChain) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.GenesisHash) > 0 {
		i -= len(m.GenesisHash)
		copy(dAtA[i:], m.GenesisHash)
		i = encodeVarintTx(dAtA, i, uint64(len(m.GenesisHash)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.GenesisURL) > 0 {
		i -= len(m.GenesisURL)
		copy(dAtA[i:], m.GenesisURL)
		i = encodeVarintTx(dAtA, i, uint64(len(m.GenesisURL)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.SourceHash) > 0 {
		i -= len(m.SourceHash)
		copy(dAtA[i:], m.SourceHash)
		i = encodeVarintTx(dAtA, i, uint64(len(m.SourceHash)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.SourceURL) > 0 {
		i -= len(m.SourceURL)
		copy(dAtA[i:], m.SourceURL)
		i = encodeVarintTx(dAtA, i, uint64(len(m.SourceURL)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.ChainName) > 0 {
		i -= len(m.ChainName)
		copy(dAtA[i:], m.ChainName)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ChainName)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Coordinator) > 0 {
		i -= len(m.Coordinator)
		copy(dAtA[i:], m.Coordinator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Coordinator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCreateChainResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateChainResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateChainResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ChainID) > 0 {
		i -= len(m.ChainID)
		copy(dAtA[i:], m.ChainID)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ChainID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgRequestRemoveValidator) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ChainID)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.ValidatorAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgRequestRemoveValidatorResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.RequestID != 0 {
		n += 1 + sovTx(uint64(m.RequestID))
	}
	return n
}

func (m *MsgCreateChain) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Coordinator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.ChainName)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.SourceURL)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.SourceHash)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.GenesisURL)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.GenesisHash)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgCreateChainResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ChainID)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgRequestRemoveValidator) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgRequestRemoveValidator: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRequestRemoveValidator: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgRequestRemoveValidatorResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgRequestRemoveValidatorResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRequestRemoveValidatorResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestID", wireType)
			}
			m.RequestID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RequestID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgCreateChain) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgCreateChain: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateChain: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Coordinator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Coordinator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceURL", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SourceURL = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SourceHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GenesisURL", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GenesisURL = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GenesisHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GenesisHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgCreateChainResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgCreateChainResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateChainResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
