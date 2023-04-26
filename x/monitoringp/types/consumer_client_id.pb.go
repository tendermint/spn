// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: spn/monitoringp/consumer_client_id.proto

package types

import (
	fmt "fmt"
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

type ConsumerClientID struct {
	ClientID string `protobuf:"bytes,1,opt,name=clientID,proto3" json:"clientID,omitempty"`
}

func (m *ConsumerClientID) Reset()         { *m = ConsumerClientID{} }
func (m *ConsumerClientID) String() string { return proto.CompactTextString(m) }
func (*ConsumerClientID) ProtoMessage()    {}
func (*ConsumerClientID) Descriptor() ([]byte, []int) {
	return fileDescriptor_8a2c45e5d19e6417, []int{0}
}
func (m *ConsumerClientID) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ConsumerClientID) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ConsumerClientID.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ConsumerClientID) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConsumerClientID.Merge(m, src)
}
func (m *ConsumerClientID) XXX_Size() int {
	return m.Size()
}
func (m *ConsumerClientID) XXX_DiscardUnknown() {
	xxx_messageInfo_ConsumerClientID.DiscardUnknown(m)
}

var xxx_messageInfo_ConsumerClientID proto.InternalMessageInfo

func (m *ConsumerClientID) GetClientID() string {
	if m != nil {
		return m.ClientID
	}
	return ""
}

func init() {
	proto.RegisterType((*ConsumerClientID)(nil), "spn.monitoringp.ConsumerClientID")
}

func init() {
	proto.RegisterFile("spn/monitoringp/consumer_client_id.proto", fileDescriptor_8a2c45e5d19e6417)
}

var fileDescriptor_8a2c45e5d19e6417 = []byte{
	// 169 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x28, 0x2e, 0xc8, 0xd3,
	0xcf, 0xcd, 0xcf, 0xcb, 0x2c, 0xc9, 0x2f, 0xca, 0xcc, 0x4b, 0x2f, 0xd0, 0x4f, 0xce, 0xcf, 0x2b,
	0x2e, 0xcd, 0x4d, 0x2d, 0x8a, 0x4f, 0xce, 0xc9, 0x4c, 0xcd, 0x2b, 0x89, 0xcf, 0x4c, 0xd1, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x2f, 0x2e, 0xc8, 0xd3, 0x43, 0x52, 0xa9, 0xa4, 0xc7, 0x25,
	0xe0, 0x0c, 0x55, 0xec, 0x0c, 0x56, 0xeb, 0xe9, 0x22, 0x24, 0xc5, 0xc5, 0x91, 0x0c, 0x65, 0x4b,
	0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0xc1, 0xf9, 0x4e, 0xee, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78,
	0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc,
	0x78, 0x2c, 0xc7, 0x10, 0xa5, 0x9b, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab,
	0x5f, 0x92, 0x9a, 0x97, 0x92, 0x5a, 0x94, 0x9b, 0x99, 0x57, 0xa2, 0x0f, 0x72, 0x5a, 0x05, 0x8a,
	0xe3, 0x4a, 0x2a, 0x0b, 0x52, 0x8b, 0x93, 0xd8, 0xc0, 0x0e, 0x32, 0x06, 0x04, 0x00, 0x00, 0xff,
	0xff, 0x6c, 0x85, 0x25, 0x9f, 0xbc, 0x00, 0x00, 0x00,
}

func (m *ConsumerClientID) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ConsumerClientID) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ConsumerClientID) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ClientID) > 0 {
		i -= len(m.ClientID)
		copy(dAtA[i:], m.ClientID)
		i = encodeVarintConsumerClientId(dAtA, i, uint64(len(m.ClientID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintConsumerClientId(dAtA []byte, offset int, v uint64) int {
	offset -= sovConsumerClientId(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ConsumerClientID) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ClientID)
	if l > 0 {
		n += 1 + l + sovConsumerClientId(uint64(l))
	}
	return n
}

func sovConsumerClientId(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozConsumerClientId(x uint64) (n int) {
	return sovConsumerClientId(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ConsumerClientID) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConsumerClientId
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
			return fmt.Errorf("proto: ConsumerClientID: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ConsumerClientID: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConsumerClientId
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
				return ErrInvalidLengthConsumerClientId
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConsumerClientId
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConsumerClientId(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthConsumerClientId
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
func skipConsumerClientId(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowConsumerClientId
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
					return 0, ErrIntOverflowConsumerClientId
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
					return 0, ErrIntOverflowConsumerClientId
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
				return 0, ErrInvalidLengthConsumerClientId
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupConsumerClientId
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthConsumerClientId
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthConsumerClientId        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowConsumerClientId          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupConsumerClientId = fmt.Errorf("proto: unexpected end of group")
)
