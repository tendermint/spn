// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: participation/auction_used_allocations.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
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

// Allocations used by a user for a specific auction
type AuctionUsedAllocations struct {
	Address        string                                 `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	AuctionID      uint64                                 `protobuf:"varint,2,opt,name=auctionID,proto3" json:"auctionID,omitempty"`
	Withdrawn      bool                                   `protobuf:"varint,3,opt,name=withdrawn,proto3" json:"withdrawn,omitempty"`
	NumAllocations github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,4,opt,name=numAllocations,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"numAllocations"`
}

func (m *AuctionUsedAllocations) Reset()         { *m = AuctionUsedAllocations{} }
func (m *AuctionUsedAllocations) String() string { return proto.CompactTextString(m) }
func (*AuctionUsedAllocations) ProtoMessage()    {}
func (*AuctionUsedAllocations) Descriptor() ([]byte, []int) {
	return fileDescriptor_0dcb8de466150226, []int{0}
}
func (m *AuctionUsedAllocations) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AuctionUsedAllocations) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AuctionUsedAllocations.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AuctionUsedAllocations) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuctionUsedAllocations.Merge(m, src)
}
func (m *AuctionUsedAllocations) XXX_Size() int {
	return m.Size()
}
func (m *AuctionUsedAllocations) XXX_DiscardUnknown() {
	xxx_messageInfo_AuctionUsedAllocations.DiscardUnknown(m)
}

var xxx_messageInfo_AuctionUsedAllocations proto.InternalMessageInfo

func (m *AuctionUsedAllocations) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *AuctionUsedAllocations) GetAuctionID() uint64 {
	if m != nil {
		return m.AuctionID
	}
	return 0
}

func (m *AuctionUsedAllocations) GetWithdrawn() bool {
	if m != nil {
		return m.Withdrawn
	}
	return false
}

func init() {
	proto.RegisterType((*AuctionUsedAllocations)(nil), "tendermint.spn.participation.AuctionUsedAllocations")
}

func init() {
	proto.RegisterFile("participation/auction_used_allocations.proto", fileDescriptor_0dcb8de466150226)
}

var fileDescriptor_0dcb8de466150226 = []byte{
	// 299 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x29, 0x48, 0x2c, 0x2a,
	0xc9, 0x4c, 0xce, 0x2c, 0x48, 0x2c, 0xc9, 0xcc, 0xcf, 0xd3, 0x4f, 0x2c, 0x4d, 0x06, 0xd1, 0xf1,
	0xa5, 0xc5, 0xa9, 0x29, 0xf1, 0x89, 0x39, 0x39, 0xf9, 0xc9, 0x60, 0xf1, 0x62, 0xbd, 0x82, 0xa2,
	0xfc, 0x92, 0x7c, 0x21, 0x99, 0x92, 0xd4, 0xbc, 0x94, 0xd4, 0xa2, 0xdc, 0xcc, 0xbc, 0x12, 0xbd,
	0xe2, 0x82, 0x3c, 0x3d, 0x14, 0xcd, 0x52, 0x22, 0xe9, 0xf9, 0xe9, 0xf9, 0x60, 0x85, 0xfa, 0x20,
	0x16, 0x44, 0x8f, 0x94, 0x64, 0x72, 0x7e, 0x71, 0x6e, 0x7e, 0x71, 0x3c, 0x44, 0x02, 0xc2, 0x81,
	0x48, 0x29, 0x5d, 0x61, 0xe4, 0x12, 0x73, 0x84, 0xd8, 0x18, 0x5a, 0x9c, 0x9a, 0xe2, 0x88, 0xb0,
	0x4f, 0x48, 0x82, 0x8b, 0x3d, 0x31, 0x25, 0xa5, 0x28, 0xb5, 0xb8, 0x58, 0x82, 0x51, 0x81, 0x51,
	0x83, 0x33, 0x08, 0xc6, 0x15, 0x92, 0xe1, 0xe2, 0x84, 0xba, 0xd2, 0xd3, 0x45, 0x82, 0x49, 0x81,
	0x51, 0x83, 0x25, 0x08, 0x21, 0x00, 0x92, 0x2d, 0xcf, 0x2c, 0xc9, 0x48, 0x29, 0x4a, 0x2c, 0xcf,
	0x93, 0x60, 0x56, 0x60, 0xd4, 0xe0, 0x08, 0x42, 0x08, 0x08, 0xa5, 0x70, 0xf1, 0xe5, 0x95, 0xe6,
	0x22, 0xd9, 0x23, 0xc1, 0x02, 0x32, 0xdc, 0xc9, 0xe6, 0xc4, 0x3d, 0x79, 0x86, 0x5b, 0xf7, 0xe4,
	0xd5, 0xd2, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xa1, 0x2e, 0x85, 0x52, 0xba,
	0xc5, 0x29, 0xd9, 0xfa, 0x25, 0x95, 0x05, 0xa9, 0xc5, 0x7a, 0x9e, 0x79, 0x25, 0x97, 0xb6, 0xe8,
	0x72, 0x41, 0x3d, 0xe2, 0x99, 0x57, 0x12, 0x84, 0x66, 0xa6, 0x93, 0xe7, 0x89, 0x47, 0x72, 0x8c,
	0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72,
	0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44, 0xe9, 0x23, 0x99, 0x8f, 0x08, 0x4a, 0xfd, 0xe2, 0x82, 0x3c,
	0xfd, 0x0a, 0x7d, 0xd4, 0x98, 0x00, 0x5b, 0x96, 0xc4, 0x06, 0x0e, 0x28, 0x63, 0x40, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xa8, 0x8c, 0xd6, 0x8a, 0xa7, 0x01, 0x00, 0x00,
}

func (m *AuctionUsedAllocations) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AuctionUsedAllocations) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AuctionUsedAllocations) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.NumAllocations.Size()
		i -= size
		if _, err := m.NumAllocations.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintAuctionUsedAllocations(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if m.Withdrawn {
		i--
		if m.Withdrawn {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if m.AuctionID != 0 {
		i = encodeVarintAuctionUsedAllocations(dAtA, i, uint64(m.AuctionID))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintAuctionUsedAllocations(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintAuctionUsedAllocations(dAtA []byte, offset int, v uint64) int {
	offset -= sovAuctionUsedAllocations(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AuctionUsedAllocations) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovAuctionUsedAllocations(uint64(l))
	}
	if m.AuctionID != 0 {
		n += 1 + sovAuctionUsedAllocations(uint64(m.AuctionID))
	}
	if m.Withdrawn {
		n += 2
	}
	l = m.NumAllocations.Size()
	n += 1 + l + sovAuctionUsedAllocations(uint64(l))
	return n
}

func sovAuctionUsedAllocations(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAuctionUsedAllocations(x uint64) (n int) {
	return sovAuctionUsedAllocations(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AuctionUsedAllocations) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuctionUsedAllocations
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
			return fmt.Errorf("proto: AuctionUsedAllocations: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuctionUsedAllocations: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuctionUsedAllocations
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
				return ErrInvalidLengthAuctionUsedAllocations
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAuctionUsedAllocations
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuctionID", wireType)
			}
			m.AuctionID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuctionUsedAllocations
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AuctionID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Withdrawn", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuctionUsedAllocations
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Withdrawn = bool(v != 0)
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumAllocations", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuctionUsedAllocations
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
				return ErrInvalidLengthAuctionUsedAllocations
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAuctionUsedAllocations
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.NumAllocations.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAuctionUsedAllocations(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAuctionUsedAllocations
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
func skipAuctionUsedAllocations(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAuctionUsedAllocations
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
					return 0, ErrIntOverflowAuctionUsedAllocations
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
					return 0, ErrIntOverflowAuctionUsedAllocations
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
				return 0, ErrInvalidLengthAuctionUsedAllocations
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAuctionUsedAllocations
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAuctionUsedAllocations
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAuctionUsedAllocations        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAuctionUsedAllocations          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAuctionUsedAllocations = fmt.Errorf("proto: unexpected end of group")
)
