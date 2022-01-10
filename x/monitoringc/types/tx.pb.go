// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: monitoringc/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	types "github.com/cosmos/ibc-go/modules/light-clients/07-tendermint/types"
	_ "github.com/gogo/protobuf/gogoproto"
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

type MsgCreateClient struct {
	Creator        string               `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	LaunchID       uint64               `protobuf:"varint,2,opt,name=launchID,proto3" json:"launchID,omitempty"`
	ClientState    types.ClientState    `protobuf:"bytes,3,opt,name=clientState,proto3" json:"clientState"`
	ConsensusState types.ConsensusState `protobuf:"bytes,4,opt,name=consensusState,proto3" json:"consensusState"`
}

func (m *MsgCreateClient) Reset()         { *m = MsgCreateClient{} }
func (m *MsgCreateClient) String() string { return proto.CompactTextString(m) }
func (*MsgCreateClient) ProtoMessage()    {}
func (*MsgCreateClient) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d32526277234083, []int{0}
}
func (m *MsgCreateClient) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateClient) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateClient.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateClient) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateClient.Merge(m, src)
}
func (m *MsgCreateClient) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateClient) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateClient.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateClient proto.InternalMessageInfo

func (m *MsgCreateClient) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgCreateClient) GetLaunchID() uint64 {
	if m != nil {
		return m.LaunchID
	}
	return 0
}

func (m *MsgCreateClient) GetClientState() types.ClientState {
	if m != nil {
		return m.ClientState
	}
	return types.ClientState{}
}

func (m *MsgCreateClient) GetConsensusState() types.ConsensusState {
	if m != nil {
		return m.ConsensusState
	}
	return types.ConsensusState{}
}

type MsgCreateClientResponse struct {
	ClientID string `protobuf:"bytes,1,opt,name=clientID,proto3" json:"clientID,omitempty"`
}

func (m *MsgCreateClientResponse) Reset()         { *m = MsgCreateClientResponse{} }
func (m *MsgCreateClientResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreateClientResponse) ProtoMessage()    {}
func (*MsgCreateClientResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6d32526277234083, []int{1}
}
func (m *MsgCreateClientResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateClientResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateClientResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateClientResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateClientResponse.Merge(m, src)
}
func (m *MsgCreateClientResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateClientResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateClientResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateClientResponse proto.InternalMessageInfo

func (m *MsgCreateClientResponse) GetClientID() string {
	if m != nil {
		return m.ClientID
	}
	return ""
}

func init() {
	proto.RegisterType((*MsgCreateClient)(nil), "tendermint.spn.monitoringc.MsgCreateClient")
	proto.RegisterType((*MsgCreateClientResponse)(nil), "tendermint.spn.monitoringc.MsgCreateClientResponse")
}

func init() { proto.RegisterFile("monitoringc/tx.proto", fileDescriptor_6d32526277234083) }

var fileDescriptor_6d32526277234083 = []byte{
	// 370 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xcd, 0x6a, 0xea, 0x40,
	0x14, 0xc7, 0x33, 0x57, 0xb9, 0xf7, 0x76, 0x2c, 0x2d, 0x04, 0xa1, 0x21, 0x8b, 0x54, 0x5c, 0x49,
	0xa5, 0x33, 0x44, 0xe9, 0x0b, 0xa8, 0x50, 0x5c, 0xb8, 0x89, 0xbb, 0xd2, 0x4d, 0x32, 0x1d, 0xc6,
	0x01, 0x33, 0x33, 0xcd, 0x4c, 0xac, 0xbe, 0x45, 0x1f, 0xcb, 0xa5, 0xcb, 0xae, 0x4a, 0xd1, 0xb7,
	0xe8, 0xaa, 0x24, 0xf1, 0x23, 0x0a, 0xa5, 0x74, 0x77, 0xfe, 0xc3, 0x39, 0xbf, 0xf9, 0x9f, 0x0f,
	0x58, 0x8f, 0xa5, 0xe0, 0x46, 0x26, 0x5c, 0x30, 0x82, 0xcd, 0x1c, 0xa9, 0x44, 0x1a, 0x69, 0xbb,
	0x86, 0x8a, 0x27, 0x9a, 0xc4, 0x5c, 0x18, 0xa4, 0x95, 0x40, 0xa5, 0x24, 0xf7, 0x86, 0x48, 0x1d,
	0x4b, 0x8d, 0xa3, 0x50, 0x53, 0xfc, 0x9c, 0xd2, 0x64, 0x81, 0x67, 0x7e, 0x44, 0x4d, 0xe8, 0x63,
	0x15, 0x32, 0x2e, 0x42, 0xc3, 0xa5, 0x28, 0x38, 0x2e, 0xe6, 0x11, 0xc1, 0x53, 0xce, 0x26, 0x86,
	0x4c, 0x39, 0x15, 0x46, 0xe3, 0x03, 0x18, 0xcf, 0xfc, 0x92, 0xda, 0x16, 0xd4, 0x99, 0x64, 0x32,
	0x0f, 0x71, 0x16, 0x15, 0xaf, 0xcd, 0x4f, 0x00, 0x2f, 0x47, 0x9a, 0xf5, 0x13, 0x1a, 0x1a, 0xda,
	0xcf, 0x49, 0xb6, 0x03, 0xff, 0x91, 0x4c, 0xcb, 0xc4, 0x01, 0x0d, 0xd0, 0x3a, 0x0b, 0x76, 0xd2,
	0x76, 0xe1, 0xff, 0x69, 0x98, 0x0a, 0x32, 0x19, 0x0e, 0x9c, 0x3f, 0x0d, 0xd0, 0xaa, 0x06, 0x7b,
	0x6d, 0x8f, 0x61, 0xad, 0x70, 0x32, 0x36, 0xa1, 0xa1, 0x4e, 0xa5, 0x01, 0x5a, 0xb5, 0x4e, 0x1b,
	0xf1, 0x88, 0xa0, 0xb2, 0x4d, 0x54, 0x32, 0x36, 0xf3, 0x51, 0xff, 0x50, 0xd2, 0xab, 0x2e, 0xdf,
	0xaf, 0xad, 0xa0, 0x4c, 0xb1, 0x1f, 0xe1, 0x05, 0x91, 0x42, 0x53, 0xa1, 0x53, 0x5d, 0x70, 0xab,
	0x39, 0x17, 0xfd, 0xc8, 0x3d, 0xaa, 0xda, 0xa2, 0x4f, 0x58, 0xcd, 0x3b, 0x78, 0x75, 0xd2, 0x7b,
	0x40, 0xb5, 0xca, 0x72, 0xb2, 0x4e, 0x0b, 0xf0, 0x70, 0xb0, 0x1d, 0xc2, 0x5e, 0x77, 0x5e, 0x60,
	0x65, 0xa4, 0x99, 0xad, 0xe0, 0xf9, 0xd1, 0xd8, 0xda, 0xe8, 0xfb, 0xd5, 0xa2, 0x93, 0x7f, 0xdc,
	0xee, 0x2f, 0x92, 0x77, 0xa6, 0x7a, 0xf7, 0xcb, 0xb5, 0x07, 0x56, 0x6b, 0x0f, 0x7c, 0xac, 0x3d,
	0xf0, 0xba, 0xf1, 0xac, 0xd5, 0xc6, 0xb3, 0xde, 0x36, 0x9e, 0xf5, 0x70, 0xcb, 0xb8, 0x99, 0xa4,
	0x11, 0x22, 0x32, 0x2e, 0xdf, 0x81, 0x56, 0x02, 0xcf, 0xf1, 0xd1, 0x1d, 0x2e, 0x14, 0xd5, 0xd1,
	0xdf, 0x7c, 0xf9, 0xdd, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9c, 0xb9, 0x2a, 0xce, 0xa3, 0x02,
	0x00, 0x00,
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
	CreateClient(ctx context.Context, in *MsgCreateClient, opts ...grpc.CallOption) (*MsgCreateClientResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreateClient(ctx context.Context, in *MsgCreateClient, opts ...grpc.CallOption) (*MsgCreateClientResponse, error) {
	out := new(MsgCreateClientResponse)
	err := c.cc.Invoke(ctx, "/tendermint.spn.monitoringc.Msg/CreateClient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	CreateClient(context.Context, *MsgCreateClient) (*MsgCreateClientResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) CreateClient(ctx context.Context, req *MsgCreateClient) (*MsgCreateClientResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateClient not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_CreateClient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateClient)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateClient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tendermint.spn.monitoringc.Msg/CreateClient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateClient(ctx, req.(*MsgCreateClient))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tendermint.spn.monitoringc.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateClient",
			Handler:    _Msg_CreateClient_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "monitoringc/tx.proto",
}

func (m *MsgCreateClient) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateClient) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateClient) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.ConsensusState.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size, err := m.ClientState.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if m.LaunchID != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.LaunchID))
		i--
		dAtA[i] = 0x10
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

func (m *MsgCreateClientResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateClientResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateClientResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ClientID) > 0 {
		i -= len(m.ClientID)
		copy(dAtA[i:], m.ClientID)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ClientID)))
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
func (m *MsgCreateClient) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.LaunchID != 0 {
		n += 1 + sovTx(uint64(m.LaunchID))
	}
	l = m.ClientState.Size()
	n += 1 + l + sovTx(uint64(l))
	l = m.ConsensusState.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgCreateClientResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ClientID)
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
func (m *MsgCreateClient) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgCreateClient: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateClient: illegal tag %d (wire type %d)", fieldNum, wire)
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LaunchID", wireType)
			}
			m.LaunchID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LaunchID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientState", wireType)
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
			if err := m.ClientState.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConsensusState", wireType)
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
			if err := m.ConsensusState.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *MsgCreateClientResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgCreateClientResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateClientResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientID", wireType)
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
			m.ClientID = string(dAtA[iNdEx:postIndex])
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
