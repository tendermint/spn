// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: types/monitoring.proto

package types

import (
	fmt "fmt"
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

// MonitoringPacket is the packet sent over IBC that contains all the signature counts
type MonitoringPacket struct {
	BlockHeight     uint64          `protobuf:"varint,1,opt,name=blockHeight,proto3" json:"blockHeight,omitempty"`
	SignatureCounts SignatureCounts `protobuf:"bytes,2,opt,name=signatureCounts,proto3" json:"signatureCounts"`
}

func (m *MonitoringPacket) Reset()         { *m = MonitoringPacket{} }
func (m *MonitoringPacket) String() string { return proto.CompactTextString(m) }
func (*MonitoringPacket) ProtoMessage()    {}
func (*MonitoringPacket) Descriptor() ([]byte, []int) {
	return fileDescriptor_4a0d1b50e3af2385, []int{0}
}
func (m *MonitoringPacket) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MonitoringPacket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MonitoringPacket.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MonitoringPacket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MonitoringPacket.Merge(m, src)
}
func (m *MonitoringPacket) XXX_Size() int {
	return m.Size()
}
func (m *MonitoringPacket) XXX_DiscardUnknown() {
	xxx_messageInfo_MonitoringPacket.DiscardUnknown(m)
}

var xxx_messageInfo_MonitoringPacket proto.InternalMessageInfo

func (m *MonitoringPacket) GetBlockHeight() uint64 {
	if m != nil {
		return m.BlockHeight
	}
	return 0
}

func (m *MonitoringPacket) GetSignatureCounts() SignatureCounts {
	if m != nil {
		return m.SignatureCounts
	}
	return SignatureCounts{}
}

// SignatureCounts contains information about signature reporting for a number of blocks
type SignatureCounts struct {
	BlockCount uint64           `protobuf:"varint,1,opt,name=blockCount,proto3" json:"blockCount,omitempty"`
	Counts     []SignatureCount `protobuf:"bytes,2,rep,name=counts,proto3" json:"counts"`
}

func (m *SignatureCounts) Reset()         { *m = SignatureCounts{} }
func (m *SignatureCounts) String() string { return proto.CompactTextString(m) }
func (*SignatureCounts) ProtoMessage()    {}
func (*SignatureCounts) Descriptor() ([]byte, []int) {
	return fileDescriptor_4a0d1b50e3af2385, []int{1}
}
func (m *SignatureCounts) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SignatureCounts) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SignatureCounts.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SignatureCounts) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignatureCounts.Merge(m, src)
}
func (m *SignatureCounts) XXX_Size() int {
	return m.Size()
}
func (m *SignatureCounts) XXX_DiscardUnknown() {
	xxx_messageInfo_SignatureCounts.DiscardUnknown(m)
}

var xxx_messageInfo_SignatureCounts proto.InternalMessageInfo

func (m *SignatureCounts) GetBlockCount() uint64 {
	if m != nil {
		return m.BlockCount
	}
	return 0
}

func (m *SignatureCounts) GetCounts() []SignatureCount {
	if m != nil {
		return m.Counts
	}
	return nil
}

// SignatureCount contains information of signature reporting for one specific validator with consensus address
// RelativeSignatures is the sum of all signatures relative to the validator set size
type SignatureCount struct {
	ConsAddress        string                                 `protobuf:"bytes,1,opt,name=consAddress,proto3" json:"consAddress,omitempty"`
	RelativeSignatures github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=RelativeSignatures,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"RelativeSignatures"`
}

func (m *SignatureCount) Reset()         { *m = SignatureCount{} }
func (m *SignatureCount) String() string { return proto.CompactTextString(m) }
func (*SignatureCount) ProtoMessage()    {}
func (*SignatureCount) Descriptor() ([]byte, []int) {
	return fileDescriptor_4a0d1b50e3af2385, []int{2}
}
func (m *SignatureCount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SignatureCount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SignatureCount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SignatureCount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SignatureCount.Merge(m, src)
}
func (m *SignatureCount) XXX_Size() int {
	return m.Size()
}
func (m *SignatureCount) XXX_DiscardUnknown() {
	xxx_messageInfo_SignatureCount.DiscardUnknown(m)
}

var xxx_messageInfo_SignatureCount proto.InternalMessageInfo

func (m *SignatureCount) GetConsAddress() string {
	if m != nil {
		return m.ConsAddress
	}
	return ""
}

func init() {
	proto.RegisterType((*MonitoringPacket)(nil), "tendermint.spn.types.MonitoringPacket")
	proto.RegisterType((*SignatureCounts)(nil), "tendermint.spn.types.SignatureCounts")
	proto.RegisterType((*SignatureCount)(nil), "tendermint.spn.types.SignatureCount")
}

func init() { proto.RegisterFile("types/monitoring.proto", fileDescriptor_4a0d1b50e3af2385) }

var fileDescriptor_4a0d1b50e3af2385 = []byte{
	// 332 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xcd, 0x4a, 0x3b, 0x31,
	0x14, 0xc5, 0x27, 0xff, 0x7f, 0x29, 0x34, 0x05, 0x2b, 0xa1, 0x48, 0x71, 0x91, 0x96, 0xfa, 0x41,
	0x37, 0x66, 0xa0, 0xae, 0x5d, 0x38, 0xba, 0x70, 0x23, 0xc8, 0x88, 0x1b, 0x17, 0x42, 0x9b, 0x09,
	0x69, 0x98, 0x4e, 0xee, 0x30, 0xc9, 0x08, 0x3e, 0x83, 0x1b, 0xf1, 0xa9, 0xba, 0xec, 0x52, 0x5c,
	0x14, 0x69, 0x5f, 0x44, 0xe6, 0xc3, 0x3a, 0x2d, 0x5d, 0xb8, 0x9a, 0xc9, 0xe1, 0xdc, 0xdf, 0x39,
	0xb9, 0xc1, 0x07, 0xf6, 0x25, 0x16, 0xc6, 0x8d, 0x40, 0x2b, 0x0b, 0x89, 0xd2, 0x92, 0xc5, 0x09,
	0x58, 0x20, 0x6d, 0x2b, 0x74, 0x20, 0x92, 0x48, 0x69, 0xcb, 0x4c, 0xac, 0x59, 0x6e, 0x3b, 0x6c,
	0x4b, 0x90, 0x90, 0x1b, 0xdc, 0xec, 0xaf, 0xf0, 0xf6, 0x5f, 0x11, 0xde, 0xbf, 0x5d, 0x03, 0xee,
	0x46, 0x3c, 0x14, 0x96, 0xf4, 0x70, 0x73, 0x3c, 0x05, 0x1e, 0xde, 0x08, 0x25, 0x27, 0xb6, 0x83,
	0x7a, 0x68, 0x50, 0xf3, 0xab, 0x12, 0x79, 0xc0, 0x2d, 0xa3, 0xa4, 0x1e, 0xd9, 0x34, 0x11, 0x57,
	0x90, 0x6a, 0x6b, 0x3a, 0xff, 0x7a, 0x68, 0xd0, 0x1c, 0x9e, 0xb0, 0x5d, 0xe1, 0xec, 0x7e, 0xd3,
	0xec, 0xd5, 0x66, 0x8b, 0xae, 0xe3, 0x6f, 0x33, 0xfa, 0x29, 0x6e, 0x6d, 0x39, 0x09, 0xc5, 0x38,
	0x0f, 0xce, 0x8f, 0x65, 0x95, 0x8a, 0x42, 0x3c, 0x5c, 0xe7, 0x3f, 0x05, 0xfe, 0x0f, 0x9a, 0xc3,
	0xe3, 0xbf, 0x14, 0x28, 0xf3, 0xcb, 0xc9, 0xfe, 0x3b, 0xc2, 0x7b, 0x9b, 0x86, 0x6c, 0x05, 0x1c,
	0xb4, 0xb9, 0x0c, 0x82, 0x44, 0x18, 0x93, 0xe7, 0x36, 0xfc, 0xaa, 0x44, 0x9e, 0x30, 0xf1, 0xc5,
	0x74, 0x64, 0xd5, 0xb3, 0x58, 0xcf, 0x16, 0x5b, 0x68, 0x78, 0x2c, 0xc3, 0x7f, 0x2e, 0xba, 0xa7,
	0x52, 0xd9, 0x49, 0x3a, 0x66, 0x1c, 0x22, 0x97, 0x83, 0x89, 0xc0, 0x94, 0x9f, 0x33, 0x13, 0x84,
	0x6e, 0xd1, 0xec, 0x5a, 0x70, 0x7f, 0x07, 0xc9, 0xbb, 0x98, 0x2d, 0x29, 0x9a, 0x2f, 0x29, 0xfa,
	0x5a, 0x52, 0xf4, 0xb6, 0xa2, 0xce, 0x7c, 0x45, 0x9d, 0x8f, 0x15, 0x75, 0x1e, 0x8f, 0x2a, 0xd4,
	0xdf, 0xcb, 0xba, 0x26, 0xd6, 0x6e, 0x1c, 0xca, 0x02, 0x3b, 0xae, 0xe7, 0xef, 0x7b, 0xfe, 0x1d,
	0x00, 0x00, 0xff, 0xff, 0x01, 0x7e, 0xa8, 0x06, 0x25, 0x02, 0x00, 0x00,
}

func (m *MonitoringPacket) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MonitoringPacket) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MonitoringPacket) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.SignatureCounts.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintMonitoring(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.BlockHeight != 0 {
		i = encodeVarintMonitoring(dAtA, i, uint64(m.BlockHeight))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *SignatureCounts) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SignatureCounts) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SignatureCounts) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Counts) > 0 {
		for iNdEx := len(m.Counts) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Counts[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintMonitoring(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if m.BlockCount != 0 {
		i = encodeVarintMonitoring(dAtA, i, uint64(m.BlockCount))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *SignatureCount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SignatureCount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SignatureCount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.RelativeSignatures.Size()
		i -= size
		if _, err := m.RelativeSignatures.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMonitoring(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.ConsAddress) > 0 {
		i -= len(m.ConsAddress)
		copy(dAtA[i:], m.ConsAddress)
		i = encodeVarintMonitoring(dAtA, i, uint64(len(m.ConsAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMonitoring(dAtA []byte, offset int, v uint64) int {
	offset -= sovMonitoring(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MonitoringPacket) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BlockHeight != 0 {
		n += 1 + sovMonitoring(uint64(m.BlockHeight))
	}
	l = m.SignatureCounts.Size()
	n += 1 + l + sovMonitoring(uint64(l))
	return n
}

func (m *SignatureCounts) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BlockCount != 0 {
		n += 1 + sovMonitoring(uint64(m.BlockCount))
	}
	if len(m.Counts) > 0 {
		for _, e := range m.Counts {
			l = e.Size()
			n += 1 + l + sovMonitoring(uint64(l))
		}
	}
	return n
}

func (m *SignatureCount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ConsAddress)
	if l > 0 {
		n += 1 + l + sovMonitoring(uint64(l))
	}
	l = m.RelativeSignatures.Size()
	n += 1 + l + sovMonitoring(uint64(l))
	return n
}

func sovMonitoring(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMonitoring(x uint64) (n int) {
	return sovMonitoring(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MonitoringPacket) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMonitoring
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
			return fmt.Errorf("proto: MonitoringPacket: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MonitoringPacket: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockHeight", wireType)
			}
			m.BlockHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMonitoring
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockHeight |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SignatureCounts", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMonitoring
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
				return ErrInvalidLengthMonitoring
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMonitoring
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.SignatureCounts.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMonitoring(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMonitoring
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
func (m *SignatureCounts) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMonitoring
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
			return fmt.Errorf("proto: SignatureCounts: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SignatureCounts: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BlockCount", wireType)
			}
			m.BlockCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMonitoring
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BlockCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Counts", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMonitoring
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
				return ErrInvalidLengthMonitoring
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMonitoring
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Counts = append(m.Counts, SignatureCount{})
			if err := m.Counts[len(m.Counts)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMonitoring(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMonitoring
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
func (m *SignatureCount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMonitoring
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
			return fmt.Errorf("proto: SignatureCount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SignatureCount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConsAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMonitoring
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
				return ErrInvalidLengthMonitoring
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMonitoring
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ConsAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RelativeSignatures", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMonitoring
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
				return ErrInvalidLengthMonitoring
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMonitoring
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.RelativeSignatures.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMonitoring(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMonitoring
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
func skipMonitoring(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMonitoring
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
					return 0, ErrIntOverflowMonitoring
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
					return 0, ErrIntOverflowMonitoring
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
				return 0, ErrInvalidLengthMonitoring
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMonitoring
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMonitoring
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMonitoring        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMonitoring          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMonitoring = fmt.Errorf("proto: unexpected end of group")
)
