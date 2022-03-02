// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: profile/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
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

// GenesisState defines the profile module's genesis state.
type GenesisState struct {
	ValidatorList              []Validator              `protobuf:"bytes,1,rep,name=validatorList,proto3" json:"validatorList"`
	ValidatorByConsAddressList []ValidatorByConsAddress `protobuf:"bytes,2,rep,name=validatorByConsAddressList,proto3" json:"validatorByConsAddressList"`
	CoordinatorList            []Coordinator            `protobuf:"bytes,3,rep,name=coordinatorList,proto3" json:"coordinatorList"`
	CoordinatorCounter         uint64                   `protobuf:"varint,4,opt,name=coordinatorCounter,proto3" json:"coordinatorCounter,omitempty"`
	CoordinatorByAddressList   []CoordinatorByAddress   `protobuf:"bytes,5,rep,name=coordinatorByAddressList,proto3" json:"coordinatorByAddressList"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_db4bc1562021cf42, []int{0}
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

func (m *GenesisState) GetValidatorList() []Validator {
	if m != nil {
		return m.ValidatorList
	}
	return nil
}

func (m *GenesisState) GetValidatorByConsAddressList() []ValidatorByConsAddress {
	if m != nil {
		return m.ValidatorByConsAddressList
	}
	return nil
}

func (m *GenesisState) GetCoordinatorList() []Coordinator {
	if m != nil {
		return m.CoordinatorList
	}
	return nil
}

func (m *GenesisState) GetCoordinatorCounter() uint64 {
	if m != nil {
		return m.CoordinatorCounter
	}
	return 0
}

func (m *GenesisState) GetCoordinatorByAddressList() []CoordinatorByAddress {
	if m != nil {
		return m.CoordinatorByAddressList
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "tendermint.spn.profile.GenesisState")
}

func init() { proto.RegisterFile("profile/genesis.proto", fileDescriptor_db4bc1562021cf42) }

var fileDescriptor_db4bc1562021cf42 = []byte{
	// 316 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xbd, 0x4e, 0xc3, 0x30,
	0x14, 0x85, 0x13, 0x5a, 0x18, 0x0c, 0x08, 0xc9, 0xe2, 0xa7, 0x64, 0x30, 0x05, 0x96, 0x22, 0x21,
	0x5b, 0x82, 0x27, 0x20, 0x19, 0x58, 0x60, 0x69, 0x25, 0x06, 0xb6, 0xb4, 0x31, 0xc1, 0x52, 0x6b,
	0x47, 0xf6, 0x2d, 0x22, 0x6f, 0xc1, 0x63, 0x65, 0xec, 0xc8, 0x84, 0x50, 0xf2, 0x22, 0x08, 0x37,
	0x69, 0x52, 0xd4, 0x02, 0x9b, 0x75, 0xcf, 0xb9, 0xdf, 0xb9, 0x47, 0x32, 0x3a, 0x48, 0xb4, 0x7a,
	0x12, 0x63, 0xce, 0x62, 0x2e, 0xb9, 0x11, 0x86, 0x26, 0x5a, 0x81, 0xc2, 0x87, 0xc0, 0x65, 0xc4,
	0xf5, 0x44, 0x48, 0xa0, 0x26, 0x91, 0xb4, 0x74, 0x79, 0xfb, 0xb1, 0x8a, 0x95, 0xb5, 0xb0, 0xef,
	0xd7, 0xdc, 0xed, 0x1d, 0x55, 0x90, 0x97, 0x70, 0x2c, 0xa2, 0x10, 0x94, 0x2e, 0x85, 0xe3, 0x4a,
	0x18, 0x29, 0xa5, 0x23, 0x21, 0x6b, 0xe9, 0x2c, 0x6b, 0xa1, 0x9d, 0xdb, 0x79, 0xe6, 0x00, 0x42,
	0xe0, 0xf8, 0x1e, 0xed, 0x2e, 0xd6, 0xef, 0x84, 0x81, 0x8e, 0xdb, 0x6d, 0xf5, 0xb6, 0xaf, 0x4e,
	0xe9, 0xea, 0x53, 0xe8, 0x43, 0x65, 0xf6, 0xdb, 0xd9, 0xc7, 0x89, 0xd3, 0x5f, 0xde, 0xc6, 0x80,
	0xbc, 0xc5, 0xc0, 0x4f, 0x03, 0x25, 0xcd, 0x4d, 0x14, 0x69, 0x6e, 0x8c, 0x65, 0x6f, 0x58, 0x36,
	0xfd, 0x9b, 0xdd, 0xdc, 0x2c, 0x83, 0x7e, 0xe1, 0xe2, 0x01, 0xda, 0x6b, 0x54, 0xb5, 0x51, 0x2d,
	0x1b, 0x75, 0xbe, 0x2e, 0x2a, 0xa8, 0xed, 0x25, 0xff, 0x27, 0x01, 0x53, 0x84, 0x1b, 0xa3, 0x40,
	0x4d, 0x25, 0x70, 0xdd, 0x69, 0x77, 0xdd, 0x5e, 0xbb, 0xbf, 0x42, 0xc1, 0x12, 0x75, 0x1a, 0x53,
	0x3f, 0x6d, 0x16, 0xdf, 0xb4, 0xd7, 0x5c, 0xfe, 0xe7, 0x9a, 0x74, 0xb9, 0xf6, 0x5a, 0xa6, 0x1f,
	0x64, 0x39, 0x71, 0x67, 0x39, 0x71, 0x3f, 0x73, 0xe2, 0xbe, 0x15, 0xc4, 0x99, 0x15, 0xc4, 0x79,
	0x2f, 0x88, 0xf3, 0x78, 0x11, 0x0b, 0x78, 0x9e, 0x0e, 0xe9, 0x48, 0x4d, 0x58, 0x9d, 0xc8, 0x4c,
	0x22, 0xd9, 0x2b, 0xab, 0xfe, 0x06, 0xa4, 0x09, 0x37, 0xc3, 0x2d, 0xfb, 0x2d, 0xae, 0xbf, 0x02,
	0x00, 0x00, 0xff, 0xff, 0x6e, 0xf9, 0xcd, 0x1a, 0x91, 0x02, 0x00, 0x00,
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
	if len(m.CoordinatorByAddressList) > 0 {
		for iNdEx := len(m.CoordinatorByAddressList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CoordinatorByAddressList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if m.CoordinatorCounter != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.CoordinatorCounter))
		i--
		dAtA[i] = 0x20
	}
	if len(m.CoordinatorList) > 0 {
		for iNdEx := len(m.CoordinatorList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CoordinatorList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.ValidatorByConsAddressList) > 0 {
		for iNdEx := len(m.ValidatorByConsAddressList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ValidatorByConsAddressList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.ValidatorList) > 0 {
		for iNdEx := len(m.ValidatorList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ValidatorList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.ValidatorList) > 0 {
		for _, e := range m.ValidatorList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.ValidatorByConsAddressList) > 0 {
		for _, e := range m.ValidatorByConsAddressList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.CoordinatorList) > 0 {
		for _, e := range m.CoordinatorList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.CoordinatorCounter != 0 {
		n += 1 + sovGenesis(uint64(m.CoordinatorCounter))
	}
	if len(m.CoordinatorByAddressList) > 0 {
		for _, e := range m.CoordinatorByAddressList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
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
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorList", wireType)
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
			m.ValidatorList = append(m.ValidatorList, Validator{})
			if err := m.ValidatorList[len(m.ValidatorList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorByConsAddressList", wireType)
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
			m.ValidatorByConsAddressList = append(m.ValidatorByConsAddressList, ValidatorByConsAddress{})
			if err := m.ValidatorByConsAddressList[len(m.ValidatorByConsAddressList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CoordinatorList", wireType)
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
			m.CoordinatorList = append(m.CoordinatorList, Coordinator{})
			if err := m.CoordinatorList[len(m.CoordinatorList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CoordinatorCounter", wireType)
			}
			m.CoordinatorCounter = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CoordinatorCounter |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CoordinatorByAddressList", wireType)
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
			m.CoordinatorByAddressList = append(m.CoordinatorByAddressList, CoordinatorByAddress{})
			if err := m.CoordinatorByAddressList[len(m.CoordinatorByAddressList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
