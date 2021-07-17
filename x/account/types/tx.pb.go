// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: account/tx.proto

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
	Creator  string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Address  string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Identity string `protobuf:"bytes,3,opt,name=identity,proto3" json:"identity,omitempty"`
	Website  string `protobuf:"bytes,4,opt,name=website,proto3" json:"website,omitempty"`
	Details  string `protobuf:"bytes,5,opt,name=details,proto3" json:"details,omitempty"`
}

func (m *MsgCreateCoordinator) Reset()         { *m = MsgCreateCoordinator{} }
func (m *MsgCreateCoordinator) String() string { return proto.CompactTextString(m) }
func (*MsgCreateCoordinator) ProtoMessage()    {}
func (*MsgCreateCoordinator) Descriptor() ([]byte, []int) {
	return fileDescriptor_dff9bd8bfd5c5f3b, []int{0}
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

func (m *MsgCreateCoordinator) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgCreateCoordinator) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *MsgCreateCoordinator) GetIdentity() string {
	if m != nil {
		return m.Identity
	}
	return ""
}

func (m *MsgCreateCoordinator) GetWebsite() string {
	if m != nil {
		return m.Website
	}
	return ""
}

func (m *MsgCreateCoordinator) GetDetails() string {
	if m != nil {
		return m.Details
	}
	return ""
}

type MsgCreateCoordinatorResponse struct {
	CoordinatorId uint64 `protobuf:"varint,1,opt,name=coordinatorId,proto3" json:"coordinatorId,omitempty"`
}

func (m *MsgCreateCoordinatorResponse) Reset()         { *m = MsgCreateCoordinatorResponse{} }
func (m *MsgCreateCoordinatorResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreateCoordinatorResponse) ProtoMessage()    {}
func (*MsgCreateCoordinatorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dff9bd8bfd5c5f3b, []int{1}
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
	proto.RegisterType((*MsgCreateCoordinator)(nil), "tendermint.spn.account.MsgCreateCoordinator")
	proto.RegisterType((*MsgCreateCoordinatorResponse)(nil), "tendermint.spn.account.MsgCreateCoordinatorResponse")
}

func init() { proto.RegisterFile("account/tx.proto", fileDescriptor_dff9bd8bfd5c5f3b) }

var fileDescriptor_dff9bd8bfd5c5f3b = []byte{
	// 283 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xcd, 0x4a, 0x03, 0x31,
	0x10, 0xc7, 0x1b, 0x5b, 0xbf, 0x02, 0x82, 0x2e, 0x22, 0xa1, 0x48, 0x90, 0xe2, 0x41, 0x41, 0xb2,
	0xa0, 0x3e, 0x81, 0xf5, 0xe2, 0xa1, 0x97, 0x3d, 0x7a, 0xdb, 0xdd, 0x0c, 0x6b, 0xc0, 0x26, 0x4b,
	0x66, 0x4a, 0xdb, 0x8b, 0xcf, 0xe0, 0xd9, 0x27, 0xf2, 0xd8, 0xa3, 0x47, 0xd9, 0x7d, 0x11, 0xd9,
	0x8f, 0xb6, 0x88, 0x7b, 0xf1, 0xf8, 0x9b, 0xdf, 0x4c, 0xf2, 0x4f, 0x86, 0x1f, 0xc7, 0x69, 0xea,
	0x66, 0x96, 0x42, 0x5a, 0xa8, 0xdc, 0x3b, 0x72, 0xc1, 0x19, 0x81, 0xd5, 0xe0, 0xa7, 0xc6, 0x92,
	0xc2, 0xdc, 0xaa, 0xb6, 0x61, 0xf4, 0xc1, 0xf8, 0xe9, 0x04, 0xb3, 0xb1, 0x87, 0x98, 0x60, 0xec,
	0x9c, 0xd7, 0xc6, 0xc6, 0xe4, 0x7c, 0x20, 0xf8, 0x7e, 0x5a, 0x15, 0x9d, 0x17, 0xec, 0x82, 0x5d,
	0x1d, 0x46, 0x6b, 0xac, 0x4c, 0xac, 0xb5, 0x07, 0x44, 0xb1, 0xd3, 0x98, 0x16, 0x83, 0x21, 0x3f,
	0x30, 0x1a, 0x2c, 0x19, 0x5a, 0x8a, 0x7e, 0xad, 0x36, 0x5c, 0x4d, 0xcd, 0x21, 0x41, 0x43, 0x20,
	0x06, 0xcd, 0x54, 0x8b, 0x95, 0xd1, 0x40, 0xb1, 0x79, 0x45, 0xb1, 0xdb, 0x98, 0x16, 0x47, 0x8f,
	0xfc, 0xbc, 0x2b, 0x5b, 0x04, 0x98, 0x3b, 0x8b, 0x10, 0x5c, 0xf2, 0xa3, 0x74, 0x5b, 0x7e, 0xd2,
	0x75, 0xd2, 0x41, 0xf4, 0xbb, 0x78, 0xfb, 0xc6, 0xfb, 0x13, 0xcc, 0x82, 0x39, 0x3f, 0xf9, 0xfb,
	0xca, 0x1b, 0xd5, 0xfd, 0x2f, 0xaa, 0xeb, 0xde, 0xe1, 0xfd, 0x7f, 0xba, 0xd7, 0x29, 0x1f, 0xc6,
	0x9f, 0x85, 0x64, 0xab, 0x42, 0xb2, 0xef, 0x42, 0xb2, 0xf7, 0x52, 0xf6, 0x56, 0xa5, 0xec, 0x7d,
	0x95, 0xb2, 0xf7, 0x7c, 0x9d, 0x19, 0x7a, 0x99, 0x25, 0x2a, 0x75, 0xd3, 0x70, 0x7b, 0x72, 0x88,
	0xb9, 0x0d, 0x17, 0xe1, 0x66, 0x85, 0xcb, 0x1c, 0x30, 0xd9, 0xab, 0xd7, 0x78, 0xf7, 0x13, 0x00,
	0x00, 0xff, 0xff, 0xbe, 0xff, 0x34, 0x33, 0xda, 0x01, 0x00, 0x00,
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
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreateCoordinator(ctx context.Context, in *MsgCreateCoordinator, opts ...grpc.CallOption) (*MsgCreateCoordinatorResponse, error) {
	out := new(MsgCreateCoordinatorResponse)
	err := c.cc.Invoke(ctx, "/tendermint.spn.account.Msg/CreateCoordinator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// this line is used by starport scaffolding # proto/tx/rpc
	CreateCoordinator(context.Context, *MsgCreateCoordinator) (*MsgCreateCoordinatorResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) CreateCoordinator(ctx context.Context, req *MsgCreateCoordinator) (*MsgCreateCoordinatorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCoordinator not implemented")
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
		FullMethod: "/tendermint.spn.account.Msg/CreateCoordinator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateCoordinator(ctx, req.(*MsgCreateCoordinator))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tendermint.spn.account.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateCoordinator",
			Handler:    _Msg_CreateCoordinator_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "account/tx.proto",
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
	if len(m.Details) > 0 {
		i -= len(m.Details)
		copy(dAtA[i:], m.Details)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Details)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Website) > 0 {
		i -= len(m.Website)
		copy(dAtA[i:], m.Website)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Website)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Identity) > 0 {
		i -= len(m.Identity)
		copy(dAtA[i:], m.Identity)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Identity)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
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
func (m *MsgCreateCoordinator) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Identity)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Website)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Details)
	if l > 0 {
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
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
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
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
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
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Identity", wireType)
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
			m.Identity = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Website", wireType)
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
			m.Website = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Details", wireType)
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
			m.Details = string(dAtA[iNdEx:postIndex])
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
