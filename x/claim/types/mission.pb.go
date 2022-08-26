// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: claim/mission.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
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

type Mission struct {
	MissionID   uint64                                 `protobuf:"varint,1,opt,name=missionID,proto3" json:"missionID,omitempty"`
	Description string                                 `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Weight      github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,3,opt,name=weight,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"weight"`
}

func (m *Mission) Reset()         { *m = Mission{} }
func (m *Mission) String() string { return proto.CompactTextString(m) }
func (*Mission) ProtoMessage()    {}
func (*Mission) Descriptor() ([]byte, []int) {
	return fileDescriptor_5a9d5f8cba75736e, []int{0}
}
func (m *Mission) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Mission) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Mission.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Mission) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Mission.Merge(m, src)
}
func (m *Mission) XXX_Size() int {
	return m.Size()
}
func (m *Mission) XXX_DiscardUnknown() {
	xxx_messageInfo_Mission.DiscardUnknown(m)
}

var xxx_messageInfo_Mission proto.InternalMessageInfo

func (m *Mission) GetMissionID() uint64 {
	if m != nil {
		return m.MissionID
	}
	return 0
}

func (m *Mission) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func init() {
	proto.RegisterType((*Mission)(nil), "tendermint.spn.claim.Mission")
}

func init() { proto.RegisterFile("claim/mission.proto", fileDescriptor_5a9d5f8cba75736e) }

var fileDescriptor_5a9d5f8cba75736e = []byte{
	// 256 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4e, 0xce, 0x49, 0xcc,
	0xcc, 0xd5, 0xcf, 0xcd, 0x2c, 0x2e, 0xce, 0xcc, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x12, 0x29, 0x49, 0xcd, 0x4b, 0x49, 0x2d, 0xca, 0xcd, 0xcc, 0x2b, 0xd1, 0x2b, 0x2e, 0xc8, 0xd3,
	0x03, 0xab, 0x91, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07, 0x2b, 0xd0, 0x07, 0xb1, 0x20, 0x6a, 0xa5,
	0x24, 0x93, 0xf3, 0x8b, 0x73, 0xf3, 0x8b, 0xe3, 0x21, 0x12, 0x10, 0x0e, 0x44, 0x4a, 0x69, 0x3e,
	0x23, 0x17, 0xbb, 0x2f, 0xc4, 0x60, 0x21, 0x19, 0x2e, 0x4e, 0xa8, 0x1d, 0x9e, 0x2e, 0x12, 0x8c,
	0x0a, 0x8c, 0x1a, 0x2c, 0x41, 0x08, 0x01, 0x21, 0x05, 0x2e, 0xee, 0x94, 0xd4, 0xe2, 0xe4, 0xa2,
	0xcc, 0x82, 0x92, 0xcc, 0xfc, 0x3c, 0x09, 0x26, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x64, 0x21, 0xa1,
	0x10, 0x2e, 0xb6, 0xf2, 0xd4, 0xcc, 0xf4, 0x8c, 0x12, 0x09, 0x66, 0x90, 0xa4, 0x93, 0xcd, 0x89,
	0x7b, 0xf2, 0x0c, 0xb7, 0xee, 0xc9, 0xab, 0xa5, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7,
	0xe7, 0x42, 0x2d, 0x87, 0x52, 0xba, 0xc5, 0x29, 0xd9, 0xfa, 0x25, 0x95, 0x05, 0xa9, 0xc5, 0x7a,
	0x2e, 0xa9, 0xc9, 0x97, 0xb6, 0xe8, 0x72, 0x41, 0xdd, 0xe6, 0x92, 0x9a, 0x1c, 0x04, 0x35, 0xcb,
	0xc9, 0xf1, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0,
	0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xa2, 0xd4, 0x91, 0xcc, 0x45,
	0x84, 0x86, 0x7e, 0x71, 0x41, 0x9e, 0x7e, 0x85, 0x3e, 0x24, 0xcc, 0xc0, 0x86, 0x27, 0xb1, 0x81,
	0xfd, 0x6a, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x40, 0x3c, 0x6a, 0x41, 0x49, 0x01, 0x00, 0x00,
}

func (m *Mission) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Mission) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Mission) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Weight.Size()
		i -= size
		if _, err := m.Weight.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMission(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintMission(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if m.MissionID != 0 {
		i = encodeVarintMission(dAtA, i, uint64(m.MissionID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintMission(dAtA []byte, offset int, v uint64) int {
	offset -= sovMission(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Mission) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MissionID != 0 {
		n += 1 + sovMission(uint64(m.MissionID))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovMission(uint64(l))
	}
	l = m.Weight.Size()
	n += 1 + l + sovMission(uint64(l))
	return n
}

func sovMission(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMission(x uint64) (n int) {
	return sovMission(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Mission) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMission
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
			return fmt.Errorf("proto: Mission: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Mission: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MissionID", wireType)
			}
			m.MissionID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMission
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MissionID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMission
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
				return ErrInvalidLengthMission
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMission
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Weight", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMission
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
				return ErrInvalidLengthMission
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMission
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Weight.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMission(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMission
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
func skipMission(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMission
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
					return 0, ErrIntOverflowMission
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
					return 0, ErrIntOverflowMission
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
				return 0, ErrInvalidLengthMission
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMission
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMission
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMission        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMission          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMission = fmt.Errorf("proto: unexpected end of group")
)
