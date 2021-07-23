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
type MsgCreateCoordinator struct {
	Address     string                  `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Description *CoordinatorDescription `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (m *MsgCreateCoordinator) Reset()         { *m = MsgCreateCoordinator{} }
func (m *MsgCreateCoordinator) String() string { return proto.CompactTextString(m) }
func (*MsgCreateCoordinator) ProtoMessage()    {}
func (*MsgCreateCoordinator) Descriptor() ([]byte, []int) {
	return fileDescriptor_a471fea62152592e, []int{0}
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
	return fileDescriptor_a471fea62152592e, []int{1}
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

type MsgDeleteCoordinator struct {
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *MsgDeleteCoordinator) Reset()         { *m = MsgDeleteCoordinator{} }
func (m *MsgDeleteCoordinator) String() string { return proto.CompactTextString(m) }
func (*MsgDeleteCoordinator) ProtoMessage()    {}
func (*MsgDeleteCoordinator) Descriptor() ([]byte, []int) {
	return fileDescriptor_a471fea62152592e, []int{2}
}
func (m *MsgDeleteCoordinator) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDeleteCoordinator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDeleteCoordinator.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDeleteCoordinator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDeleteCoordinator.Merge(m, src)
}
func (m *MsgDeleteCoordinator) XXX_Size() int {
	return m.Size()
}
func (m *MsgDeleteCoordinator) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDeleteCoordinator.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDeleteCoordinator proto.InternalMessageInfo

func (m *MsgDeleteCoordinator) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type MsgDeleteCoordinatorResponse struct {
}

func (m *MsgDeleteCoordinatorResponse) Reset()         { *m = MsgDeleteCoordinatorResponse{} }
func (m *MsgDeleteCoordinatorResponse) String() string { return proto.CompactTextString(m) }
func (*MsgDeleteCoordinatorResponse) ProtoMessage()    {}
func (*MsgDeleteCoordinatorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a471fea62152592e, []int{3}
}
func (m *MsgDeleteCoordinatorResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDeleteCoordinatorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDeleteCoordinatorResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDeleteCoordinatorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDeleteCoordinatorResponse.Merge(m, src)
}
func (m *MsgDeleteCoordinatorResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgDeleteCoordinatorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDeleteCoordinatorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDeleteCoordinatorResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgCreateCoordinator)(nil), "tendermint.spn.profile.MsgCreateCoordinator")
	proto.RegisterType((*MsgCreateCoordinatorResponse)(nil), "tendermint.spn.profile.MsgCreateCoordinatorResponse")
	proto.RegisterType((*MsgDeleteCoordinator)(nil), "tendermint.spn.profile.MsgDeleteCoordinator")
	proto.RegisterType((*MsgDeleteCoordinatorResponse)(nil), "tendermint.spn.profile.MsgDeleteCoordinatorResponse")
}

func init() { proto.RegisterFile("profile/tx.proto", fileDescriptor_a471fea62152592e) }

var fileDescriptor_a471fea62152592e = []byte{
	// 303 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x28, 0x28, 0xca, 0x4f,
	0xcb, 0xcc, 0x49, 0xd5, 0x2f, 0xa9, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x2b, 0x49,
	0xcd, 0x4b, 0x49, 0x2d, 0xca, 0xcd, 0xcc, 0x2b, 0xd1, 0x2b, 0x2e, 0xc8, 0xd3, 0x83, 0x2a, 0x90,
	0x92, 0x84, 0xa9, 0x4c, 0xce, 0xcf, 0x2f, 0x4a, 0xc9, 0xcc, 0x4b, 0x2c, 0xc9, 0x2f, 0x82, 0x68,
	0x51, 0x6a, 0x62, 0xe4, 0x12, 0xf1, 0x2d, 0x4e, 0x77, 0x2e, 0x4a, 0x4d, 0x2c, 0x49, 0x75, 0x46,
	0x48, 0x0b, 0x49, 0x70, 0xb1, 0x27, 0xa6, 0xa4, 0x14, 0xa5, 0x16, 0x17, 0x4b, 0x30, 0x2a, 0x30,
	0x6a, 0x70, 0x06, 0xc1, 0xb8, 0x42, 0x01, 0x5c, 0xdc, 0x29, 0xa9, 0xc5, 0xc9, 0x45, 0x99, 0x05,
	0x25, 0x99, 0xf9, 0x79, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0xdc, 0x46, 0x7a, 0x7a, 0xd8, 0xed, 0xd6,
	0x43, 0x32, 0xd3, 0x05, 0xa1, 0x2b, 0x08, 0xd9, 0x08, 0x25, 0x17, 0x2e, 0x19, 0x6c, 0x6e, 0x08,
	0x4a, 0x2d, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x15, 0x52, 0xe1, 0xe2, 0x45, 0x72, 0xb9, 0x67, 0x0a,
	0xd8, 0x45, 0x2c, 0x41, 0xa8, 0x82, 0x4a, 0x06, 0x60, 0x9f, 0xb8, 0xa4, 0xe6, 0xa4, 0x12, 0xe9,
	0x13, 0x25, 0x39, 0xb0, 0xbd, 0x18, 0x3a, 0x60, 0xf6, 0x1a, 0x7d, 0x67, 0xe4, 0x62, 0xf6, 0x2d,
	0x4e, 0x17, 0x2a, 0xe7, 0x12, 0xc4, 0x0c, 0x20, 0x1d, 0x5c, 0x3e, 0xc6, 0xe6, 0x15, 0x29, 0x13,
	0x52, 0x54, 0xc3, 0x3d, 0x5e, 0xce, 0x25, 0x88, 0xe9, 0x1f, 0x7c, 0x16, 0x63, 0xa8, 0xc6, 0x6b,
	0x31, 0x4e, 0x9f, 0x3b, 0x39, 0x9f, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47,
	0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c, 0x43, 0x94,
	0x66, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0x3e, 0xc2, 0x64, 0xfd, 0xe2,
	0x82, 0x3c, 0xfd, 0x0a, 0x7d, 0x78, 0x8a, 0xac, 0x2c, 0x48, 0x2d, 0x4e, 0x62, 0x03, 0x27, 0x31,
	0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0c, 0xdb, 0xee, 0x90, 0xa9, 0x02, 0x00, 0x00,
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
	CreateCoordinator(ctx context.Context, in *MsgCreateCoordinator, opts ...grpc.CallOption) (*MsgCreateCoordinatorResponse, error)
	DeleteCoordinator(ctx context.Context, in *MsgDeleteCoordinator, opts ...grpc.CallOption) (*MsgDeleteCoordinatorResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreateCoordinator(ctx context.Context, in *MsgCreateCoordinator, opts ...grpc.CallOption) (*MsgCreateCoordinatorResponse, error) {
	out := new(MsgCreateCoordinatorResponse)
	err := c.cc.Invoke(ctx, "/tendermint.spn.profile.Msg/CreateCoordinator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) DeleteCoordinator(ctx context.Context, in *MsgDeleteCoordinator, opts ...grpc.CallOption) (*MsgDeleteCoordinatorResponse, error) {
	out := new(MsgDeleteCoordinatorResponse)
	err := c.cc.Invoke(ctx, "/tendermint.spn.profile.Msg/DeleteCoordinator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// this line is used by starport scaffolding # proto/tx/rpc
	CreateCoordinator(context.Context, *MsgCreateCoordinator) (*MsgCreateCoordinatorResponse, error)
	DeleteCoordinator(context.Context, *MsgDeleteCoordinator) (*MsgDeleteCoordinatorResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) CreateCoordinator(ctx context.Context, req *MsgCreateCoordinator) (*MsgCreateCoordinatorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCoordinator not implemented")
}
func (*UnimplementedMsgServer) DeleteCoordinator(ctx context.Context, req *MsgDeleteCoordinator) (*MsgDeleteCoordinatorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCoordinator not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
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

func _Msg_DeleteCoordinator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDeleteCoordinator)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DeleteCoordinator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tendermint.spn.profile.Msg/DeleteCoordinator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DeleteCoordinator(ctx, req.(*MsgDeleteCoordinator))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tendermint.spn.profile.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCoordinator",
			Handler:    _Msg_CreateCoordinator_Handler,
		},
		{
			MethodName: "DeleteCoordinator",
			Handler:    _Msg_DeleteCoordinator_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "profile/tx.proto",
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

func (m *MsgDeleteCoordinator) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDeleteCoordinator) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDeleteCoordinator) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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

func (m *MsgDeleteCoordinatorResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDeleteCoordinatorResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDeleteCoordinatorResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
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

func (m *MsgDeleteCoordinator) Size() (n int) {
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

func (m *MsgDeleteCoordinatorResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
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
func (m *MsgDeleteCoordinator) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgDeleteCoordinator: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDeleteCoordinator: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *MsgDeleteCoordinatorResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgDeleteCoordinatorResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDeleteCoordinatorResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
