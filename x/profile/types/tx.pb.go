// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: profile/tx.proto

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
type MsgDeleteValidator struct {
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *MsgDeleteValidator) Reset()         { *m = MsgDeleteValidator{} }
func (m *MsgDeleteValidator) String() string { return proto.CompactTextString(m) }
func (*MsgDeleteValidator) ProtoMessage()    {}
func (*MsgDeleteValidator) Descriptor() ([]byte, []int) {
	return fileDescriptor_a471fea62152592e, []int{0}
}
func (m *MsgDeleteValidator) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDeleteValidator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDeleteValidator.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDeleteValidator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDeleteValidator.Merge(m, src)
}
func (m *MsgDeleteValidator) XXX_Size() int {
	return m.Size()
}
func (m *MsgDeleteValidator) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDeleteValidator.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDeleteValidator proto.InternalMessageInfo

func (m *MsgDeleteValidator) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type MsgDeleteValidatorResponse struct {
}

func (m *MsgDeleteValidatorResponse) Reset()         { *m = MsgDeleteValidatorResponse{} }
func (m *MsgDeleteValidatorResponse) String() string { return proto.CompactTextString(m) }
func (*MsgDeleteValidatorResponse) ProtoMessage()    {}
func (*MsgDeleteValidatorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a471fea62152592e, []int{1}
}
func (m *MsgDeleteValidatorResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDeleteValidatorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDeleteValidatorResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDeleteValidatorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDeleteValidatorResponse.Merge(m, src)
}
func (m *MsgDeleteValidatorResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgDeleteValidatorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDeleteValidatorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDeleteValidatorResponse proto.InternalMessageInfo

type MsgCreateCoordinator struct {
	Address     string                  `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Description *CoordinatorDescription `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (m *MsgCreateCoordinator) Reset()         { *m = MsgCreateCoordinator{} }
func (m *MsgCreateCoordinator) String() string { return proto.CompactTextString(m) }
func (*MsgCreateCoordinator) ProtoMessage()    {}
func (*MsgCreateCoordinator) Descriptor() ([]byte, []int) {
	return fileDescriptor_a471fea62152592e, []int{2}
}
func (m *MsgCreateCoordinator) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateCoordinator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateCoordinator.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateCoordinator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateCoordinator.Merge(m, src)
}
func (m *MsgCreateCoordinator) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateCoordinator) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateCoordinator.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateCoordinator proto.InternalMessageInfo

func (m *MsgCreateCoordinator) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *MsgCreateCoordinator) GetDescription() *CoordinatorDescription {
	if m != nil {
		return m.Description
	}
	return nil
}

type MsgCreateCoordinatorResponse struct {
	CoordinatorId uint64 `protobuf:"varint,1,opt,name=coordinatorId,proto3" json:"coordinatorId,omitempty"`
}

func (m *MsgCreateCoordinatorResponse) Reset()         { *m = MsgCreateCoordinatorResponse{} }
func (m *MsgCreateCoordinatorResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreateCoordinatorResponse) ProtoMessage()    {}
func (*MsgCreateCoordinatorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a471fea62152592e, []int{3}
}
func (m *MsgCreateCoordinatorResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateCoordinatorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateCoordinatorResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateCoordinatorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateCoordinatorResponse.Merge(m, src)
}
func (m *MsgCreateCoordinatorResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateCoordinatorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateCoordinatorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateCoordinatorResponse proto.InternalMessageInfo

func (m *MsgCreateCoordinatorResponse) GetCoordinatorId() uint64 {
	if m != nil {
		return m.CoordinatorId
	}
	return 0
}

func init() {
	proto.RegisterType((*MsgDeleteValidator)(nil), "tendermint.spn.profile.MsgDeleteValidator")
	proto.RegisterType((*MsgDeleteValidatorResponse)(nil), "tendermint.spn.profile.MsgDeleteValidatorResponse")
	proto.RegisterType((*MsgCreateCoordinator)(nil), "tendermint.spn.profile.MsgCreateCoordinator")
	proto.RegisterType((*MsgCreateCoordinatorResponse)(nil), "tendermint.spn.profile.MsgCreateCoordinatorResponse")
}

func init() { proto.RegisterFile("profile/tx.proto", fileDescriptor_a471fea62152592e) }

var fileDescriptor_a471fea62152592e = []byte{
	// 317 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x28, 0x28, 0xca, 0x4f,
	0xcb, 0xcc, 0x49, 0xd5, 0x2f, 0xa9, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x2b, 0x49,
	0xcd, 0x4b, 0x49, 0x2d, 0xca, 0xcd, 0xcc, 0x2b, 0xd1, 0x2b, 0x2e, 0xc8, 0xd3, 0x83, 0x2a, 0x90,
	0x92, 0x84, 0xa9, 0x4c, 0xce, 0xcf, 0x2f, 0x4a, 0xc9, 0xcc, 0x4b, 0x2c, 0xc9, 0x2f, 0x82, 0x68,
	0x51, 0xd2, 0xe3, 0x12, 0xf2, 0x2d, 0x4e, 0x77, 0x49, 0xcd, 0x49, 0x2d, 0x49, 0x0d, 0x4b, 0xcc,
	0xc9, 0x4c, 0x01, 0xc9, 0x09, 0x49, 0x70, 0xb1, 0x27, 0xa6, 0xa4, 0x14, 0xa5, 0x16, 0x17, 0x4b,
	0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0xc1, 0xb8, 0x4a, 0x32, 0x5c, 0x52, 0x98, 0xea, 0x83, 0x52,
	0x8b, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x95, 0x9a, 0x18, 0xb9, 0x44, 0x7c, 0x8b, 0xd3, 0x9d, 0x8b,
	0x52, 0x13, 0x4b, 0x52, 0x9d, 0x11, 0x96, 0xe1, 0x36, 0x50, 0x28, 0x80, 0x8b, 0x3b, 0x25, 0xb5,
	0x38, 0xb9, 0x28, 0xb3, 0xa0, 0x24, 0x33, 0x3f, 0x4f, 0x82, 0x49, 0x81, 0x51, 0x83, 0xdb, 0x48,
	0x4f, 0x0f, 0xbb, 0x4f, 0xf4, 0x90, 0xcc, 0x74, 0x41, 0xe8, 0x0a, 0x42, 0x36, 0x42, 0xc9, 0x85,
	0x4b, 0x06, 0x9b, 0x1b, 0x60, 0x8e, 0x14, 0x52, 0xe1, 0xe2, 0x45, 0x0a, 0x07, 0xcf, 0x14, 0xb0,
	0x8b, 0x58, 0x82, 0x50, 0x05, 0x8d, 0x3e, 0x32, 0x72, 0x31, 0xfb, 0x16, 0xa7, 0x0b, 0x15, 0x72,
	0xf1, 0xa3, 0x87, 0x8e, 0x16, 0x2e, 0xd7, 0x61, 0x86, 0x8c, 0x94, 0x11, 0xf1, 0x6a, 0xe1, 0x0e,
	0x2c, 0xe7, 0x12, 0xc4, 0x0c, 0x41, 0x1d, 0x3c, 0x06, 0x61, 0xa8, 0x96, 0x32, 0x21, 0x45, 0x35,
	0xcc, 0x62, 0x27, 0xe7, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e,
	0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88, 0xd2, 0x4c,
	0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x47, 0x98, 0xac, 0x5f, 0x5c, 0x90,
	0xa7, 0x5f, 0xa1, 0x0f, 0x4f, 0x87, 0x95, 0x05, 0xa9, 0xc5, 0x49, 0x6c, 0xe0, 0x84, 0x65, 0x0c,
	0x08, 0x00, 0x00, 0xff, 0xff, 0x2e, 0xad, 0xef, 0x6c, 0x9f, 0x02, 0x00, 0x00,
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
	DeleteValidator(ctx context.Context, in *MsgDeleteValidator, opts ...grpc.CallOption) (*MsgDeleteValidatorResponse, error)
	CreateCoordinator(ctx context.Context, in *MsgCreateCoordinator, opts ...grpc.CallOption) (*MsgCreateCoordinatorResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) DeleteValidator(ctx context.Context, in *MsgDeleteValidator, opts ...grpc.CallOption) (*MsgDeleteValidatorResponse, error) {
	out := new(MsgDeleteValidatorResponse)
	err := c.cc.Invoke(ctx, "/tendermint.spn.profile.Msg/DeleteValidator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) CreateCoordinator(ctx context.Context, in *MsgCreateCoordinator, opts ...grpc.CallOption) (*MsgCreateCoordinatorResponse, error) {
	out := new(MsgCreateCoordinatorResponse)
	err := c.cc.Invoke(ctx, "/tendermint.spn.profile.Msg/CreateCoordinator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// this line is used by starport scaffolding # proto/tx/rpc
	DeleteValidator(context.Context, *MsgDeleteValidator) (*MsgDeleteValidatorResponse, error)
	CreateCoordinator(context.Context, *MsgCreateCoordinator) (*MsgCreateCoordinatorResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) DeleteValidator(ctx context.Context, req *MsgDeleteValidator) (*MsgDeleteValidatorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteValidator not implemented")
}
func (*UnimplementedMsgServer) CreateCoordinator(ctx context.Context, req *MsgCreateCoordinator) (*MsgCreateCoordinatorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCoordinator not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_DeleteValidator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDeleteValidator)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DeleteValidator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tendermint.spn.profile.Msg/DeleteValidator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DeleteValidator(ctx, req.(*MsgDeleteValidator))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_CreateCoordinator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateCoordinator)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateCoordinator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tendermint.spn.profile.Msg/CreateCoordinator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateCoordinator(ctx, req.(*MsgCreateCoordinator))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tendermint.spn.profile.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DeleteValidator",
			Handler:    _Msg_DeleteValidator_Handler,
		},
		{
			MethodName: "CreateCoordinator",
			Handler:    _Msg_CreateCoordinator_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "profile/tx.proto",
}

func (m *MsgDeleteValidator) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDeleteValidator) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDeleteValidator) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgDeleteValidatorResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDeleteValidatorResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDeleteValidatorResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgCreateCoordinator) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateCoordinator) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateCoordinator) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Description != nil {
		{
			size, err := m.Description.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCreateCoordinatorResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateCoordinatorResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateCoordinatorResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CoordinatorId != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.CoordinatorId))
		i--
		dAtA[i] = 0x8
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
func (m *MsgDeleteValidator) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgDeleteValidatorResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgCreateCoordinator) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Description != nil {
		l = m.Description.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgCreateCoordinatorResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CoordinatorId != 0 {
		n += 1 + sovTx(uint64(m.CoordinatorId))
	}
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgDeleteValidator) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgDeleteValidator: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDeleteValidator: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
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
			m.Address = string(dAtA[iNdEx:postIndex])
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
func (m *MsgDeleteValidatorResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgDeleteValidatorResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDeleteValidatorResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
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
func (m *MsgCreateCoordinator) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgCreateCoordinator: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateCoordinator: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
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
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Description == nil {
				m.Description = &CoordinatorDescription{}
			}
			if err := m.Description.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
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
func (m *MsgCreateCoordinatorResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgCreateCoordinatorResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateCoordinatorResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CoordinatorId", wireType)
			}
			m.CoordinatorId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CoordinatorId |= uint64(b&0x7F) << shift
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
