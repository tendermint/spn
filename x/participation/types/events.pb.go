// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: spn/participation/events.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

type EventAllocationsUsed struct {
	Participant    string                                 `protobuf:"bytes,1,opt,name=participant,proto3" json:"participant,omitempty"`
	AuctionID      uint64                                 `protobuf:"varint,2,opt,name=auctionID,proto3" json:"auctionID,omitempty"`
	NumAllocations github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,3,opt,name=numAllocations,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"numAllocations"`
}

func (m *EventAllocationsUsed) Reset()         { *m = EventAllocationsUsed{} }
func (m *EventAllocationsUsed) String() string { return proto.CompactTextString(m) }
func (*EventAllocationsUsed) ProtoMessage()    {}
func (*EventAllocationsUsed) Descriptor() ([]byte, []int) {
	return fileDescriptor_9fdf41644bb3267b, []int{0}
}
func (m *EventAllocationsUsed) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventAllocationsUsed) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventAllocationsUsed.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventAllocationsUsed) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventAllocationsUsed.Merge(m, src)
}
func (m *EventAllocationsUsed) XXX_Size() int {
	return m.Size()
}
func (m *EventAllocationsUsed) XXX_DiscardUnknown() {
	xxx_messageInfo_EventAllocationsUsed.DiscardUnknown(m)
}

var xxx_messageInfo_EventAllocationsUsed proto.InternalMessageInfo

func (m *EventAllocationsUsed) GetParticipant() string {
	if m != nil {
		return m.Participant
	}
	return ""
}

func (m *EventAllocationsUsed) GetAuctionID() uint64 {
	if m != nil {
		return m.AuctionID
	}
	return 0
}

type EventAllocationsWithdrawn struct {
	Participant string `protobuf:"bytes,1,opt,name=participant,proto3" json:"participant,omitempty"`
	AuctionID   uint64 `protobuf:"varint,2,opt,name=auctionID,proto3" json:"auctionID,omitempty"`
}

func (m *EventAllocationsWithdrawn) Reset()         { *m = EventAllocationsWithdrawn{} }
func (m *EventAllocationsWithdrawn) String() string { return proto.CompactTextString(m) }
func (*EventAllocationsWithdrawn) ProtoMessage()    {}
func (*EventAllocationsWithdrawn) Descriptor() ([]byte, []int) {
	return fileDescriptor_9fdf41644bb3267b, []int{1}
}
func (m *EventAllocationsWithdrawn) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventAllocationsWithdrawn) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventAllocationsWithdrawn.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventAllocationsWithdrawn) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventAllocationsWithdrawn.Merge(m, src)
}
func (m *EventAllocationsWithdrawn) XXX_Size() int {
	return m.Size()
}
func (m *EventAllocationsWithdrawn) XXX_DiscardUnknown() {
	xxx_messageInfo_EventAllocationsWithdrawn.DiscardUnknown(m)
}

var xxx_messageInfo_EventAllocationsWithdrawn proto.InternalMessageInfo

func (m *EventAllocationsWithdrawn) GetParticipant() string {
	if m != nil {
		return m.Participant
	}
	return ""
}

func (m *EventAllocationsWithdrawn) GetAuctionID() uint64 {
	if m != nil {
		return m.AuctionID
	}
	return 0
}

func init() {
	proto.RegisterType((*EventAllocationsUsed)(nil), "spn.participation.EventAllocationsUsed")
	proto.RegisterType((*EventAllocationsWithdrawn)(nil), "spn.participation.EventAllocationsWithdrawn")
}

func init() { proto.RegisterFile("spn/participation/events.proto", fileDescriptor_9fdf41644bb3267b) }

var fileDescriptor_9fdf41644bb3267b = []byte{
	// 287 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2b, 0x2e, 0xc8, 0xd3,
	0x2f, 0x48, 0x2c, 0x2a, 0xc9, 0x4c, 0xce, 0x2c, 0x48, 0x2c, 0xc9, 0xcc, 0xcf, 0xd3, 0x4f, 0x2d,
	0x4b, 0xcd, 0x2b, 0x29, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x2c, 0x2e, 0xc8, 0xd3,
	0x43, 0x91, 0x97, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07, 0xcb, 0xea, 0x83, 0x58, 0x10, 0x85, 0x52,
	0x92, 0xc9, 0xf9, 0xc5, 0xb9, 0xf9, 0xc5, 0xf1, 0x10, 0x09, 0x08, 0x07, 0x22, 0xa5, 0xb4, 0x87,
	0x91, 0x4b, 0xc4, 0x15, 0x64, 0xa8, 0x63, 0x4e, 0x4e, 0x7e, 0x32, 0xd8, 0x90, 0xe2, 0xd0, 0xe2,
	0xd4, 0x14, 0x21, 0x05, 0x2e, 0x6e, 0xb8, 0xd1, 0x79, 0x25, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c,
	0x41, 0xc8, 0x42, 0x42, 0x32, 0x5c, 0x9c, 0x89, 0xa5, 0xc9, 0x20, 0x1d, 0x9e, 0x2e, 0x12, 0x4c,
	0x0a, 0x8c, 0x1a, 0x2c, 0x41, 0x08, 0x01, 0xa1, 0x14, 0x2e, 0xbe, 0xbc, 0xd2, 0x5c, 0x24, 0x53,
	0x25, 0x98, 0x41, 0x46, 0x38, 0xd9, 0x9c, 0xb8, 0x27, 0xcf, 0x70, 0xeb, 0x9e, 0xbc, 0x5a, 0x7a,
	0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0x2e, 0xd4, 0x45, 0x50, 0x4a, 0xb7, 0x38, 0x25,
	0x5b, 0xbf, 0xa4, 0xb2, 0x20, 0xb5, 0x58, 0xcf, 0x33, 0xaf, 0xe4, 0xd2, 0x16, 0x5d, 0x2e, 0xa8,
	0x83, 0x3d, 0xf3, 0x4a, 0x82, 0xd0, 0xcc, 0x54, 0x8a, 0xe6, 0x92, 0x44, 0x77, 0x7d, 0x78, 0x66,
	0x49, 0x46, 0x4a, 0x51, 0x62, 0x79, 0x1e, 0xa5, 0x5e, 0x70, 0xf2, 0x3c, 0xf1, 0x48, 0x8e, 0xf1,
	0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e,
	0xe1, 0xc6, 0x63, 0x39, 0x86, 0x28, 0x7d, 0x24, 0xc7, 0x97, 0xa4, 0xe6, 0xa5, 0xa4, 0x16, 0xe5,
	0x66, 0xe6, 0x95, 0xe8, 0x83, 0xe2, 0xab, 0x02, 0x2d, 0xc6, 0xc0, 0x3e, 0x49, 0x62, 0x03, 0x87,
	0xb6, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x85, 0x27, 0xeb, 0xd9, 0xd3, 0x01, 0x00, 0x00,
}

func (m *EventAllocationsUsed) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventAllocationsUsed) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventAllocationsUsed) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
		i = encodeVarintEvents(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if m.AuctionID != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.AuctionID))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Participant) > 0 {
		i -= len(m.Participant)
		copy(dAtA[i:], m.Participant)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Participant)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventAllocationsWithdrawn) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventAllocationsWithdrawn) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventAllocationsWithdrawn) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.AuctionID != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.AuctionID))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Participant) > 0 {
		i -= len(m.Participant)
		copy(dAtA[i:], m.Participant)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Participant)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvents(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvents(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EventAllocationsUsed) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Participant)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.AuctionID != 0 {
		n += 1 + sovEvents(uint64(m.AuctionID))
	}
	l = m.NumAllocations.Size()
	n += 1 + l + sovEvents(uint64(l))
	return n
}

func (m *EventAllocationsWithdrawn) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Participant)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.AuctionID != 0 {
		n += 1 + sovEvents(uint64(m.AuctionID))
	}
	return n
}

func sovEvents(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvents(x uint64) (n int) {
	return sovEvents(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EventAllocationsUsed) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventAllocationsUsed: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventAllocationsUsed: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Participant", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Participant = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuctionID", wireType)
			}
			m.AuctionID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumAllocations", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
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
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventAllocationsWithdrawn) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventAllocationsWithdrawn: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventAllocationsWithdrawn: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Participant", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Participant = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuctionID", wireType)
			}
			m.AuctionID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func skipEvents(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
				return 0, ErrInvalidLengthEvents
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvents
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvents
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvents        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvents          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvents = fmt.Errorf("proto: unexpected end of group")
)
