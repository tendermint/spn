// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: launch/params.proto

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

// Params defines the parameters for the staking module.
type Params struct {
	LaunchTimeRange LaunchTimeRange `protobuf:"bytes,1,opt,name=launchTimeRange,proto3" json:"launchTimeRange"`
	RevertDelay     int64           `protobuf:"varint,2,opt,name=revertDelay,proto3" json:"revertDelay,omitempty"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8f73d6645a211b2, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetLaunchTimeRange() LaunchTimeRange {
	if m != nil {
		return m.LaunchTimeRange
	}
	return LaunchTimeRange{}
}

func (m *Params) GetRevertDelay() int64 {
	if m != nil {
		return m.RevertDelay
	}
	return 0
}

type LaunchTimeRange struct {
	MinLaunchTime uint64 `protobuf:"varint,1,opt,name=minLaunchTime,proto3" json:"minLaunchTime,omitempty"`
	MaxLaunchTime uint64 `protobuf:"varint,2,opt,name=maxLaunchTime,proto3" json:"maxLaunchTime,omitempty"`
}

func (m *LaunchTimeRange) Reset()         { *m = LaunchTimeRange{} }
func (m *LaunchTimeRange) String() string { return proto.CompactTextString(m) }
func (*LaunchTimeRange) ProtoMessage()    {}
func (*LaunchTimeRange) Descriptor() ([]byte, []int) {
	return fileDescriptor_b8f73d6645a211b2, []int{1}
}
func (m *LaunchTimeRange) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LaunchTimeRange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LaunchTimeRange.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LaunchTimeRange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LaunchTimeRange.Merge(m, src)
}
func (m *LaunchTimeRange) XXX_Size() int {
	return m.Size()
}
func (m *LaunchTimeRange) XXX_DiscardUnknown() {
	xxx_messageInfo_LaunchTimeRange.DiscardUnknown(m)
}

var xxx_messageInfo_LaunchTimeRange proto.InternalMessageInfo

func (m *LaunchTimeRange) GetMinLaunchTime() uint64 {
	if m != nil {
		return m.MinLaunchTime
	}
	return 0
}

func (m *LaunchTimeRange) GetMaxLaunchTime() uint64 {
	if m != nil {
		return m.MaxLaunchTime
	}
	return 0
}

func init() {
	proto.RegisterType((*Params)(nil), "tendermint.spn.launch.Params")
	proto.RegisterType((*LaunchTimeRange)(nil), "tendermint.spn.launch.LaunchTimeRange")
}

func init() { proto.RegisterFile("launch/params.proto", fileDescriptor_b8f73d6645a211b2) }

var fileDescriptor_b8f73d6645a211b2 = []byte{
	// 250 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xce, 0x49, 0x2c, 0xcd,
	0x4b, 0xce, 0xd0, 0x2f, 0x48, 0x2c, 0x4a, 0xcc, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x12, 0x2d, 0x49, 0xcd, 0x4b, 0x49, 0x2d, 0xca, 0xcd, 0xcc, 0x2b, 0xd1, 0x2b, 0x2e, 0xc8, 0xd3,
	0x83, 0xa8, 0x91, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07, 0xab, 0xd0, 0x07, 0xb1, 0x20, 0x8a, 0x95,
	0x3a, 0x18, 0xb9, 0xd8, 0x02, 0xc0, 0xba, 0x85, 0xc2, 0xb8, 0xf8, 0x21, 0x4a, 0x43, 0x32, 0x73,
	0x53, 0x83, 0x12, 0xf3, 0xd2, 0x53, 0x25, 0x18, 0x15, 0x18, 0x35, 0xb8, 0x8d, 0xd4, 0xf4, 0xb0,
	0x9a, 0xa8, 0xe7, 0x83, 0xaa, 0xda, 0x89, 0xe5, 0xc4, 0x3d, 0x79, 0x86, 0x20, 0x74, 0x43, 0x84,
	0x14, 0xb8, 0xb8, 0x8b, 0x52, 0xcb, 0x52, 0x8b, 0x4a, 0x5c, 0x52, 0x73, 0x12, 0x2b, 0x25, 0x98,
	0x14, 0x18, 0x35, 0x98, 0x83, 0x90, 0x85, 0xac, 0x58, 0x66, 0x2c, 0x90, 0x67, 0x50, 0x8a, 0xe5,
	0xe2, 0x47, 0x33, 0x51, 0x48, 0x85, 0x8b, 0x37, 0x37, 0x33, 0x0f, 0x21, 0x0a, 0x76, 0x10, 0x4b,
	0x10, 0xaa, 0x20, 0x58, 0x55, 0x62, 0x05, 0x92, 0x2a, 0x26, 0xa8, 0x2a, 0x64, 0x41, 0x27, 0xa7,
	0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39,
	0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88, 0xd2, 0x48, 0xcf, 0x2c, 0xc9, 0x28,
	0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x47, 0xf8, 0x54, 0xbf, 0xb8, 0x20, 0x4f, 0xbf, 0x42, 0x1f,
	0x1a, 0xc2, 0x25, 0x95, 0x05, 0xa9, 0xc5, 0x49, 0x6c, 0xe0, 0x40, 0x33, 0x06, 0x04, 0x00, 0x00,
	0xff, 0xff, 0x3e, 0x22, 0x6c, 0xca, 0x78, 0x01, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.RevertDelay != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.RevertDelay))
		i--
		dAtA[i] = 0x10
	}
	{
		size, err := m.LaunchTimeRange.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *LaunchTimeRange) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LaunchTimeRange) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LaunchTimeRange) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MaxLaunchTime != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MaxLaunchTime))
		i--
		dAtA[i] = 0x10
	}
	if m.MinLaunchTime != 0 {
		i = encodeVarintParams(dAtA, i, uint64(m.MinLaunchTime))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.LaunchTimeRange.Size()
	n += 1 + l + sovParams(uint64(l))
	if m.RevertDelay != 0 {
		n += 1 + sovParams(uint64(m.RevertDelay))
	}
	return n
}

func (m *LaunchTimeRange) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MinLaunchTime != 0 {
		n += 1 + sovParams(uint64(m.MinLaunchTime))
	}
	if m.MaxLaunchTime != 0 {
		n += 1 + sovParams(uint64(m.MaxLaunchTime))
	}
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LaunchTimeRange", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LaunchTimeRange.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RevertDelay", wireType)
			}
			m.RevertDelay = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RevertDelay |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func (m *LaunchTimeRange) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: LaunchTimeRange: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LaunchTimeRange: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinLaunchTime", wireType)
			}
			m.MinLaunchTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MinLaunchTime |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxLaunchTime", wireType)
			}
			m.MaxLaunchTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxLaunchTime |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
