// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: campaign/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

// this line is used by starport scaffolding # proto/tx/message
type MsgAddMainnetVestingAccount struct {
	Coordinator    string              `protobuf:"bytes,1,opt,name=coordinator,proto3" json:"coordinator,omitempty"`
	CampaignID     uint64              `protobuf:"varint,2,opt,name=campaignID,proto3" json:"campaignID,omitempty"`
	Address        string              `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	Shares         Shares              `protobuf:"bytes,4,rep,name=shares,proto3,casttype=github.com/cosmos/cosmos-sdk/types.Coin,castrepeated=Shares" json:"shares"`
	VestingOptions ShareVestingOptions `protobuf:"bytes,5,opt,name=vestingOptions,proto3" json:"vestingOptions"`
}

func (m *MsgAddMainnetVestingAccount) Reset()         { *m = MsgAddMainnetVestingAccount{} }
func (m *MsgAddMainnetVestingAccount) String() string { return proto.CompactTextString(m) }
func (*MsgAddMainnetVestingAccount) ProtoMessage()    {}
func (*MsgAddMainnetVestingAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb6bf904ffc53c1f, []int{0}
}
func (m *MsgAddMainnetVestingAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgAddMainnetVestingAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgAddMainnetVestingAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgAddMainnetVestingAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAddMainnetVestingAccount.Merge(m, src)
}
func (m *MsgAddMainnetVestingAccount) XXX_Size() int {
	return m.Size()
}
func (m *MsgAddMainnetVestingAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAddMainnetVestingAccount.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAddMainnetVestingAccount proto.InternalMessageInfo

func (m *MsgAddMainnetVestingAccount) GetCoordinator() string {
	if m != nil {
		return m.Coordinator
	}
	return ""
}

func (m *MsgAddMainnetVestingAccount) GetCampaignID() uint64 {
	if m != nil {
		return m.CampaignID
	}
	return 0
}

func (m *MsgAddMainnetVestingAccount) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *MsgAddMainnetVestingAccount) GetShares() Shares {
	if m != nil {
		return m.Shares
	}
	return nil
}

func (m *MsgAddMainnetVestingAccount) GetVestingOptions() ShareVestingOptions {
	if m != nil {
		return m.VestingOptions
	}
	return ShareVestingOptions{}
}

type MsgAddMainnetVestingAccountResponse struct {
}

func (m *MsgAddMainnetVestingAccountResponse) Reset()         { *m = MsgAddMainnetVestingAccountResponse{} }
func (m *MsgAddMainnetVestingAccountResponse) String() string { return proto.CompactTextString(m) }
func (*MsgAddMainnetVestingAccountResponse) ProtoMessage()    {}
func (*MsgAddMainnetVestingAccountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_fb6bf904ffc53c1f, []int{1}
}
func (m *MsgAddMainnetVestingAccountResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgAddMainnetVestingAccountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgAddMainnetVestingAccountResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgAddMainnetVestingAccountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgAddMainnetVestingAccountResponse.Merge(m, src)
}
func (m *MsgAddMainnetVestingAccountResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgAddMainnetVestingAccountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgAddMainnetVestingAccountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgAddMainnetVestingAccountResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgAddMainnetVestingAccount)(nil), "tendermint.spn.campaign.MsgAddMainnetVestingAccount")
	proto.RegisterType((*MsgAddMainnetVestingAccountResponse)(nil), "tendermint.spn.campaign.MsgAddMainnetVestingAccountResponse")
}

func init() { proto.RegisterFile("campaign/tx.proto", fileDescriptor_fb6bf904ffc53c1f) }

var fileDescriptor_fb6bf904ffc53c1f = []byte{
	// 412 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xcf, 0xca, 0xd3, 0x40,
	0x14, 0xc5, 0x33, 0x5f, 0x6b, 0xc5, 0x29, 0x08, 0x06, 0xc1, 0x58, 0x61, 0x1a, 0x2a, 0x6a, 0x10,
	0x9d, 0xa1, 0xd5, 0x8d, 0xe0, 0xa6, 0xb5, 0x1b, 0x17, 0x45, 0x88, 0xd0, 0x45, 0x37, 0x65, 0x92,
	0x0c, 0xe9, 0xa0, 0x99, 0x09, 0xb9, 0xd3, 0x52, 0x5f, 0x42, 0x5c, 0xf9, 0x10, 0xbe, 0x86, 0x9b,
	0x2e, 0xbb, 0x74, 0x55, 0xa5, 0x7d, 0x0b, 0x57, 0x92, 0x7f, 0xb6, 0x0a, 0xed, 0xe2, 0x5b, 0x25,
	0xb9, 0x39, 0xe7, 0x9e, 0x7b, 0x7f, 0x5c, 0x7c, 0x27, 0xe4, 0x49, 0xca, 0x65, 0xac, 0x98, 0x59,
	0xd3, 0x34, 0xd3, 0x46, 0xdb, 0xf7, 0x8c, 0x50, 0x91, 0xc8, 0x12, 0xa9, 0x0c, 0x85, 0x54, 0xd1,
	0x5a, 0xd1, 0xb9, 0x1b, 0xeb, 0x58, 0x17, 0x1a, 0x96, 0xbf, 0x95, 0xf2, 0x0e, 0x09, 0x35, 0x24,
	0x1a, 0x58, 0xc0, 0x41, 0xb0, 0x55, 0x3f, 0x10, 0x86, 0xf7, 0x59, 0xa8, 0xa5, 0xaa, 0xfe, 0x3f,
	0xfe, 0x9b, 0x90, 0x70, 0xa9, 0x94, 0x30, 0xf3, 0x95, 0x00, 0x23, 0x55, 0x3c, 0xe7, 0x61, 0xa8,
	0x97, 0xca, 0x94, 0xba, 0xde, 0xf7, 0x2b, 0xfc, 0x60, 0x02, 0xf1, 0x30, 0x8a, 0x26, 0xa5, 0x6e,
	0x5a, 0xca, 0x86, 0xa5, 0xca, 0x76, 0x71, 0x3b, 0xd4, 0x3a, 0x8b, 0xa4, 0xe2, 0x46, 0x67, 0x0e,
	0x72, 0x91, 0x77, 0xcb, 0x3f, 0x2d, 0xd9, 0x04, 0xe3, 0x3a, 0xeb, 0xed, 0xd8, 0xb9, 0x72, 0x91,
	0xd7, 0xf4, 0x4f, 0x2a, 0xb6, 0x83, 0x6f, 0xf2, 0x28, 0xca, 0x04, 0x80, 0xd3, 0x28, 0xdc, 0xf5,
	0xa7, 0xfd, 0x11, 0xb7, 0x60, 0xc1, 0x33, 0x01, 0x4e, 0xd3, 0x6d, 0x78, 0xed, 0xc1, 0x7d, 0x5a,
	0x2e, 0x45, 0xf3, 0xa5, 0x68, 0xb5, 0x14, 0x7d, 0xa3, 0xa5, 0x1a, 0xbd, 0xda, 0xec, 0xba, 0xd6,
	0xef, 0x5d, 0xf7, 0x49, 0x2c, 0xcd, 0x62, 0x19, 0xd0, 0x50, 0x27, 0xac, 0x22, 0x50, 0x3e, 0x9e,
	0x43, 0xf4, 0x81, 0x99, 0x4f, 0xa9, 0x80, 0xc2, 0xf0, 0xed, 0x67, 0xb7, 0xf5, 0xbe, 0xe8, 0xed,
	0x57, 0x19, 0xf6, 0x0c, 0xdf, 0xae, 0x10, 0xbc, 0x4b, 0x8d, 0xd4, 0x0a, 0x9c, 0x1b, 0x2e, 0xf2,
	0xda, 0x83, 0x67, 0xf4, 0x0c, 0x79, 0x5a, 0x34, 0x98, 0xfe, 0xe3, 0x19, 0x35, 0xf3, 0x41, 0xfc,
	0xff, 0x3a, 0xf5, 0x1e, 0xe1, 0x87, 0x17, 0x20, 0xfa, 0x02, 0x52, 0xad, 0x40, 0x0c, 0xbe, 0x22,
	0xdc, 0x98, 0x40, 0x6c, 0x7f, 0x46, 0xd8, 0x39, 0x4b, 0xfc, 0xe5, 0xd9, 0x79, 0x2e, 0x44, 0x74,
	0x5e, 0x5f, 0xc7, 0x55, 0x0f, 0x36, 0x1a, 0x6f, 0xf6, 0x04, 0x6d, 0xf7, 0x04, 0xfd, 0xda, 0x13,
	0xf4, 0xe5, 0x40, 0xac, 0xed, 0x81, 0x58, 0x3f, 0x0e, 0xc4, 0x9a, 0x3d, 0x3d, 0x01, 0x7e, 0x4c,
	0x60, 0x90, 0x2a, 0xb6, 0x66, 0xc7, 0x2b, 0xce, 0xc1, 0x07, 0xad, 0xe2, 0xa4, 0x5e, 0xfc, 0x09,
	0x00, 0x00, 0xff, 0xff, 0x06, 0x78, 0xd9, 0x3c, 0xde, 0x02, 0x00, 0x00,
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
	AddMainnetVestingAccount(ctx context.Context, in *MsgAddMainnetVestingAccount, opts ...grpc.CallOption) (*MsgAddMainnetVestingAccountResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) AddMainnetVestingAccount(ctx context.Context, in *MsgAddMainnetVestingAccount, opts ...grpc.CallOption) (*MsgAddMainnetVestingAccountResponse, error) {
	out := new(MsgAddMainnetVestingAccountResponse)
	err := c.cc.Invoke(ctx, "/tendermint.spn.campaign.Msg/AddMainnetVestingAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// this line is used by starport scaffolding # proto/tx/rpc
	AddMainnetVestingAccount(context.Context, *MsgAddMainnetVestingAccount) (*MsgAddMainnetVestingAccountResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) AddMainnetVestingAccount(ctx context.Context, req *MsgAddMainnetVestingAccount) (*MsgAddMainnetVestingAccountResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddMainnetVestingAccount not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_AddMainnetVestingAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgAddMainnetVestingAccount)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).AddMainnetVestingAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tendermint.spn.campaign.Msg/AddMainnetVestingAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).AddMainnetVestingAccount(ctx, req.(*MsgAddMainnetVestingAccount))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tendermint.spn.campaign.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddMainnetVestingAccount",
			Handler:    _Msg_AddMainnetVestingAccount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "campaign/tx.proto",
}

func (m *MsgAddMainnetVestingAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgAddMainnetVestingAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgAddMainnetVestingAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.VestingOptions.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if len(m.Shares) > 0 {
		for iNdEx := len(m.Shares) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Shares[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x1a
	}
	if m.CampaignID != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.CampaignID))
		i--
		dAtA[i] = 0x10
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

func (m *MsgAddMainnetVestingAccountResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgAddMainnetVestingAccountResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgAddMainnetVestingAccountResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
func (m *MsgAddMainnetVestingAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Coordinator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.CampaignID != 0 {
		n += 1 + sovTx(uint64(m.CampaignID))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.Shares) > 0 {
		for _, e := range m.Shares {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	l = m.VestingOptions.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgAddMainnetVestingAccountResponse) Size() (n int) {
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
func (m *MsgAddMainnetVestingAccount) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgAddMainnetVestingAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgAddMainnetVestingAccount: illegal tag %d (wire type %d)", fieldNum, wire)
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CampaignID", wireType)
			}
			m.CampaignID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CampaignID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
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
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Shares", wireType)
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
			m.Shares = append(m.Shares, github_com_cosmos_cosmos_sdk_types.Coin{})
			if err := m.Shares[len(m.Shares)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VestingOptions", wireType)
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
			if err := m.VestingOptions.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *MsgAddMainnetVestingAccountResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgAddMainnetVestingAccountResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgAddMainnetVestingAccountResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
