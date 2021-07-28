// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: launch/genesis.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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

// GenesisState defines the launch module's genesis state.
type GenesisState struct {
	// this line is used by starport scaffolding # genesis/proto/state
	ChainList            []*Chain            `protobuf:"bytes,1,rep,name=chainList,proto3" json:"chainList,omitempty"`
	GenesisAccountList   []*GenesisAccount   `protobuf:"bytes,2,rep,name=genesisAccountList,proto3" json:"genesisAccountList,omitempty"`
	VestedAccountList    []*VestedAccount    `protobuf:"bytes,3,rep,name=vestedAccountList,proto3" json:"vestedAccountList,omitempty"`
	GenesisValidatorList []*GenesisValidator `protobuf:"bytes,4,rep,name=genesisValidatorList,proto3" json:"genesisValidatorList,omitempty"`
	RequestList          []*Request          `protobuf:"bytes,10,rep,name=requestList,proto3" json:"requestList,omitempty"`
	RequestCountList     []*RequestCount     `protobuf:"bytes,11,rep,name=requestCountList,proto3" json:"requestCountList,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_02cd66d27edc51cd, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetChainList() []*Chain {
	if m != nil {
		return m.ChainList
	}
	return nil
}

func (m *GenesisState) GetGenesisAccountList() []*GenesisAccount {
	if m != nil {
		return m.GenesisAccountList
	}
	return nil
}

func (m *GenesisState) GetVestedAccountList() []*VestedAccount {
	if m != nil {
		return m.VestedAccountList
	}
	return nil
}

func (m *GenesisState) GetGenesisValidatorList() []*GenesisValidator {
	if m != nil {
		return m.GenesisValidatorList
	}
	return nil
}

func (m *GenesisState) GetRequestList() []*Request {
	if m != nil {
		return m.RequestList
	}
	return nil
}

func (m *GenesisState) GetRequestCountList() []*RequestCount {
	if m != nil {
		return m.RequestCountList
	}
	return nil
}

type RequestCount struct {
	ChainID string `protobuf:"bytes,1,opt,name=chainID,proto3" json:"chainID,omitempty"`
	Count   uint64 `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (m *RequestCount) Reset()         { *m = RequestCount{} }
func (m *RequestCount) String() string { return proto.CompactTextString(m) }
func (*RequestCount) ProtoMessage()    {}
func (*RequestCount) Descriptor() ([]byte, []int) {
	return fileDescriptor_02cd66d27edc51cd, []int{1}
}
func (m *RequestCount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RequestCount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RequestCount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RequestCount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestCount.Merge(m, src)
}
func (m *RequestCount) XXX_Size() int {
	return m.Size()
}
func (m *RequestCount) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestCount.DiscardUnknown(m)
}

var xxx_messageInfo_RequestCount proto.InternalMessageInfo

func (m *RequestCount) GetChainID() string {
	if m != nil {
		return m.ChainID
	}
	return ""
}

func (m *RequestCount) GetCount() uint64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "tendermint.spn.launch.GenesisState")
	proto.RegisterType((*RequestCount)(nil), "tendermint.spn.launch.RequestCount")
}

func init() { proto.RegisterFile("launch/genesis.proto", fileDescriptor_02cd66d27edc51cd) }

var fileDescriptor_02cd66d27edc51cd = []byte{
	// 369 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x4f, 0x4b, 0xc3, 0x30,
	0x18, 0x87, 0xd7, 0x6d, 0x2a, 0xcb, 0x76, 0xd0, 0x30, 0xa1, 0xcc, 0x11, 0xc6, 0x54, 0xec, 0xa9,
	0x05, 0xbd, 0x79, 0x10, 0xdd, 0x84, 0x21, 0x08, 0x42, 0xc4, 0x1d, 0xf4, 0x20, 0x5d, 0x17, 0xb6,
	0xc2, 0x96, 0xd6, 0x26, 0x1d, 0xfa, 0x2d, 0xfc, 0x58, 0x1e, 0x77, 0xf4, 0x28, 0x1b, 0x7e, 0x0f,
	0xf1, 0x4d, 0xba, 0x75, 0x7f, 0x8f, 0xc9, 0xef, 0x79, 0x9f, 0xbc, 0x79, 0x13, 0x54, 0x1e, 0xb8,
	0x31, 0xf7, 0xfa, 0x4e, 0x8f, 0x71, 0x26, 0x7c, 0x61, 0x87, 0x51, 0x20, 0x03, 0x7c, 0x28, 0x19,
	0xef, 0xb2, 0x68, 0xe8, 0x73, 0x69, 0x8b, 0x90, 0xdb, 0x0a, 0xaa, 0x24, 0x70, 0xc4, 0xde, 0x62,
	0x26, 0xa4, 0x82, 0x2b, 0x47, 0x7a, 0x77, 0xc4, 0x84, 0x64, 0xdd, 0x57, 0xd7, 0xf3, 0x82, 0x98,
	0x27, 0x61, 0x75, 0xd1, 0xbf, 0x94, 0x92, 0xa5, 0x74, 0xe4, 0x0e, 0xfc, 0xae, 0x2b, 0x83, 0x48,
	0xe7, 0x58, 0xe7, 0x5e, 0xdf, 0xf5, 0xb9, 0xda, 0xab, 0xff, 0xe6, 0x50, 0xa9, 0xa5, 0xf8, 0x47,
	0xe9, 0x4a, 0x86, 0x2f, 0x51, 0x01, 0xf2, 0x7b, 0x5f, 0x48, 0xd3, 0xa8, 0xe5, 0xac, 0xe2, 0x79,
	0xd5, 0x5e, 0x7b, 0x01, 0xbb, 0xf9, 0xcf, 0xd1, 0x39, 0x8e, 0x9f, 0x10, 0xd6, 0x67, 0xdf, 0xa8,
	0xc6, 0x40, 0x92, 0x05, 0xc9, 0xe9, 0x06, 0x49, 0x6b, 0xa1, 0x80, 0xae, 0x11, 0x60, 0x8a, 0x0e,
	0xd4, 0x34, 0xd2, 0xd6, 0x1c, 0x58, 0x4f, 0x36, 0x58, 0xdb, 0x69, 0x9e, 0xae, 0x96, 0xe3, 0x17,
	0x54, 0xd6, 0x27, 0xb5, 0x93, 0x29, 0x81, 0x36, 0x0f, 0xda, 0xb3, 0xed, 0xcd, 0xce, 0x4a, 0xe8,
	0x5a, 0x09, 0xbe, 0x46, 0x45, 0xfd, 0xa8, 0xe0, 0x44, 0xe0, 0x24, 0x1b, 0x9c, 0x54, 0x91, 0x34,
	0x5d, 0x82, 0x1f, 0xd0, 0xbe, 0x5e, 0x36, 0x67, 0x37, 0x2e, 0x82, 0xe6, 0x78, 0xbb, 0x06, 0x70,
	0xba, 0x52, 0x5c, 0xbf, 0x42, 0xa5, 0x34, 0x81, 0x4d, 0xb4, 0x07, 0xef, 0x76, 0x77, 0x6b, 0x1a,
	0x35, 0xc3, 0x2a, 0xd0, 0x64, 0x89, 0xcb, 0x68, 0x07, 0xc6, 0x64, 0x66, 0x6b, 0x86, 0x95, 0xa7,
	0x6a, 0xd1, 0x68, 0x7c, 0x4d, 0x88, 0x31, 0x9e, 0x10, 0xe3, 0x67, 0x42, 0x8c, 0xcf, 0x29, 0xc9,
	0x8c, 0xa7, 0x24, 0xf3, 0x3d, 0x25, 0x99, 0x67, 0xab, 0xe7, 0xcb, 0x7e, 0xdc, 0xb1, 0xbd, 0x60,
	0xe8, 0xcc, 0x5b, 0x73, 0x44, 0xc8, 0x9d, 0x77, 0x47, 0xff, 0x38, 0xf9, 0x11, 0x32, 0xd1, 0xd9,
	0x85, 0x2f, 0x77, 0xf1, 0x17, 0x00, 0x00, 0xff, 0xff, 0x98, 0x5e, 0x30, 0x05, 0x26, 0x03, 0x00,
	0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.RequestCountList) > 0 {
		for iNdEx := len(m.RequestCountList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.RequestCountList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x5a
		}
	}
	if len(m.RequestList) > 0 {
		for iNdEx := len(m.RequestList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.RequestList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x52
		}
	}
	if len(m.GenesisValidatorList) > 0 {
		for iNdEx := len(m.GenesisValidatorList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.GenesisValidatorList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.VestedAccountList) > 0 {
		for iNdEx := len(m.VestedAccountList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.VestedAccountList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.GenesisAccountList) > 0 {
		for iNdEx := len(m.GenesisAccountList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.GenesisAccountList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.ChainList) > 0 {
		for iNdEx := len(m.ChainList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ChainList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *RequestCount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RequestCount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RequestCount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Count != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.Count))
		i--
		dAtA[i] = 0x10
	}
	if len(m.ChainID) > 0 {
		i -= len(m.ChainID)
		copy(dAtA[i:], m.ChainID)
		i = encodeVarintGenesis(dAtA, i, uint64(len(m.ChainID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.ChainList) > 0 {
		for _, e := range m.ChainList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.GenesisAccountList) > 0 {
		for _, e := range m.GenesisAccountList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.VestedAccountList) > 0 {
		for _, e := range m.VestedAccountList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.GenesisValidatorList) > 0 {
		for _, e := range m.GenesisValidatorList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.RequestList) > 0 {
		for _, e := range m.RequestList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.RequestCountList) > 0 {
		for _, e := range m.RequestCountList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *RequestCount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ChainID)
	if l > 0 {
		n += 1 + l + sovGenesis(uint64(l))
	}
	if m.Count != 0 {
		n += 1 + sovGenesis(uint64(m.Count))
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainList = append(m.ChainList, &Chain{})
			if err := m.ChainList[len(m.ChainList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GenesisAccountList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GenesisAccountList = append(m.GenesisAccountList, &GenesisAccount{})
			if err := m.GenesisAccountList[len(m.GenesisAccountList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VestedAccountList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VestedAccountList = append(m.VestedAccountList, &VestedAccount{})
			if err := m.VestedAccountList[len(m.VestedAccountList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GenesisValidatorList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GenesisValidatorList = append(m.GenesisValidatorList, &GenesisValidator{})
			if err := m.GenesisValidatorList[len(m.GenesisValidatorList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RequestList = append(m.RequestList, &Request{})
			if err := m.RequestList[len(m.RequestList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestCountList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RequestCountList = append(m.RequestCountList, &RequestCount{})
			if err := m.RequestCountList[len(m.RequestCountList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *RequestCount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: RequestCount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RequestCount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Count", wireType)
			}
			m.Count = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Count |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
