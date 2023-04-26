// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: spn/project/project_chains.proto

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

type ProjectChains struct {
	ProjectID uint64   `protobuf:"varint,1,opt,name=projectID,proto3" json:"projectID,omitempty"`
	Chains    []uint64 `protobuf:"varint,2,rep,packed,name=chains,proto3" json:"chains,omitempty"`
}

func (m *ProjectChains) Reset()         { *m = ProjectChains{} }
func (m *ProjectChains) String() string { return proto.CompactTextString(m) }
func (*ProjectChains) ProtoMessage()    {}
func (*ProjectChains) Descriptor() ([]byte, []int) {
	return fileDescriptor_954119d7be6667b6, []int{0}
}
func (m *ProjectChains) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProjectChains) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProjectChains.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProjectChains) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProjectChains.Merge(m, src)
}
func (m *ProjectChains) XXX_Size() int {
	return m.Size()
}
func (m *ProjectChains) XXX_DiscardUnknown() {
	xxx_messageInfo_ProjectChains.DiscardUnknown(m)
}

var xxx_messageInfo_ProjectChains proto.InternalMessageInfo

func (m *ProjectChains) GetProjectID() uint64 {
	if m != nil {
		return m.ProjectID
	}
	return 0
}

func (m *ProjectChains) GetChains() []uint64 {
	if m != nil {
		return m.Chains
	}
	return nil
}

func init() {
	proto.RegisterType((*ProjectChains)(nil), "spn.project.ProjectChains")
}

func init() { proto.RegisterFile("spn/project/project_chains.proto", fileDescriptor_954119d7be6667b6) }

var fileDescriptor_954119d7be6667b6 = []byte{
	// 168 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x28, 0x2e, 0xc8, 0xd3,
	0x2f, 0x28, 0xca, 0xcf, 0x4a, 0x4d, 0x2e, 0x81, 0xd1, 0xf1, 0xc9, 0x19, 0x89, 0x99, 0x79, 0xc5,
	0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0xdc, 0xc5, 0x05, 0x79, 0x7a, 0x50, 0x19, 0x25, 0x57,
	0x2e, 0xde, 0x00, 0x08, 0xd3, 0x19, 0xac, 0x46, 0x48, 0x86, 0x8b, 0x13, 0x2a, 0xe7, 0xe9, 0x22,
	0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x12, 0x84, 0x10, 0x10, 0x12, 0xe3, 0x62, 0x83, 0x98, 0x25, 0xc1,
	0xa4, 0xc0, 0xac, 0xc1, 0x12, 0x04, 0xe5, 0x39, 0x39, 0x9f, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91,
	0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3,
	0xb1, 0x1c, 0x43, 0x94, 0x66, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0x7e,
	0x49, 0x6a, 0x5e, 0x4a, 0x6a, 0x51, 0x6e, 0x66, 0x5e, 0x89, 0x3e, 0xc8, 0x95, 0x15, 0x70, 0x77,
	0x96, 0x54, 0x16, 0xa4, 0x16, 0x27, 0xb1, 0x81, 0xdd, 0x67, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff,
	0xe7, 0x9c, 0x16, 0x84, 0xc3, 0x00, 0x00, 0x00,
}

func (m *ProjectChains) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProjectChains) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProjectChains) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Chains) > 0 {
		dAtA2 := make([]byte, len(m.Chains)*10)
		var j1 int
		for _, num := range m.Chains {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintProjectChains(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0x12
	}
	if m.ProjectID != 0 {
		i = encodeVarintProjectChains(dAtA, i, uint64(m.ProjectID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintProjectChains(dAtA []byte, offset int, v uint64) int {
	offset -= sovProjectChains(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ProjectChains) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ProjectID != 0 {
		n += 1 + sovProjectChains(uint64(m.ProjectID))
	}
	if len(m.Chains) > 0 {
		l = 0
		for _, e := range m.Chains {
			l += sovProjectChains(uint64(e))
		}
		n += 1 + sovProjectChains(uint64(l)) + l
	}
	return n
}

func sovProjectChains(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProjectChains(x uint64) (n int) {
	return sovProjectChains(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ProjectChains) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProjectChains
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
			return fmt.Errorf("proto: ProjectChains: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProjectChains: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProjectID", wireType)
			}
			m.ProjectID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProjectChains
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ProjectID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowProjectChains
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Chains = append(m.Chains, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowProjectChains
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthProjectChains
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthProjectChains
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.Chains) == 0 {
					m.Chains = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowProjectChains
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Chains = append(m.Chains, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Chains", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipProjectChains(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProjectChains
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
func skipProjectChains(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProjectChains
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
					return 0, ErrIntOverflowProjectChains
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
					return 0, ErrIntOverflowProjectChains
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
				return 0, ErrInvalidLengthProjectChains
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProjectChains
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProjectChains
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProjectChains        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProjectChains          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProjectChains = fmt.Errorf("proto: unexpected end of group")
)
