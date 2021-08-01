// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: launch/request.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/codec/types"
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

type Request struct {
	ChainID   string     `protobuf:"bytes,1,opt,name=chainID,proto3" json:"chainID,omitempty"`
	RequestID uint64     `protobuf:"varint,2,opt,name=requestID,proto3" json:"requestID,omitempty"`
	Creator   string     `protobuf:"bytes,3,opt,name=creator,proto3" json:"creator,omitempty"`
	CreatedAt int64      `protobuf:"varint,4,opt,name=createdAt,proto3" json:"createdAt,omitempty"`
	Content   *types.Any `protobuf:"bytes,5,opt,name=content,proto3" json:"content,omitempty"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_028e4b0ce31bf039, []int{0}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Request.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return m.Size()
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetChainID() string {
	if m != nil {
		return m.ChainID
	}
	return ""
}

func (m *Request) GetRequestID() uint64 {
	if m != nil {
		return m.RequestID
	}
	return 0
}

func (m *Request) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *Request) GetCreatedAt() int64 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *Request) GetContent() *types.Any {
	if m != nil {
		return m.Content
	}
	return nil
}

type ContentRemoveAccount struct {
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *ContentRemoveAccount) Reset()         { *m = ContentRemoveAccount{} }
func (m *ContentRemoveAccount) String() string { return proto.CompactTextString(m) }
func (*ContentRemoveAccount) ProtoMessage()    {}
func (*ContentRemoveAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_028e4b0ce31bf039, []int{1}
}
func (m *ContentRemoveAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ContentRemoveAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ContentRemoveAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ContentRemoveAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContentRemoveAccount.Merge(m, src)
}
func (m *ContentRemoveAccount) XXX_Size() int {
	return m.Size()
}
func (m *ContentRemoveAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_ContentRemoveAccount.DiscardUnknown(m)
}

var xxx_messageInfo_ContentRemoveAccount proto.InternalMessageInfo

func (m *ContentRemoveAccount) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "tendermint.spn.launch.Request")
	proto.RegisterType((*ContentRemoveAccount)(nil), "tendermint.spn.launch.ContentRemoveAccount")
}

func init() { proto.RegisterFile("launch/request.proto", fileDescriptor_028e4b0ce31bf039) }

var fileDescriptor_028e4b0ce31bf039 = []byte{
	// 310 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x91, 0xb1, 0x4e, 0xf3, 0x30,
	0x14, 0x85, 0xeb, 0xbf, 0xfd, 0xa9, 0x6a, 0x24, 0x06, 0xab, 0x48, 0x6e, 0x85, 0xac, 0xa8, 0x53,
	0x16, 0x6c, 0x09, 0x36, 0xb6, 0x86, 0x2e, 0x5d, 0x33, 0xb2, 0xa0, 0xd4, 0xb9, 0xb4, 0x95, 0x1a,
	0x3b, 0xd8, 0x0e, 0xa2, 0x6f, 0xc1, 0xc3, 0xb0, 0xb2, 0x23, 0xa6, 0x8a, 0x89, 0x11, 0x35, 0x2f,
	0x82, 0xb0, 0x13, 0x65, 0x60, 0xf3, 0xb9, 0x3e, 0xdf, 0xd5, 0x39, 0x36, 0x1e, 0xef, 0xb2, 0x4a,
	0xc9, 0x8d, 0x30, 0xf0, 0x58, 0x81, 0x75, 0xbc, 0x34, 0xda, 0x69, 0x72, 0xee, 0x40, 0xe5, 0x60,
	0x8a, 0xad, 0x72, 0xdc, 0x96, 0x8a, 0x07, 0xd3, 0x74, 0xb2, 0xd6, 0x7a, 0xbd, 0x03, 0xe1, 0x4d,
	0xab, 0xea, 0x41, 0x64, 0x6a, 0x1f, 0x88, 0xe9, 0x44, 0x6a, 0x5b, 0x68, 0x7b, 0xef, 0x95, 0x08,
	0x22, 0x5c, 0xcd, 0xde, 0x10, 0x1e, 0xa6, 0x61, 0x3d, 0xa1, 0x78, 0x28, 0x37, 0xd9, 0x56, 0x2d,
	0x17, 0x14, 0x45, 0x28, 0x1e, 0xa5, 0xad, 0x24, 0x17, 0x78, 0xd4, 0x64, 0x58, 0x2e, 0xe8, 0xbf,
	0x08, 0xc5, 0x83, 0xb4, 0x1b, 0x78, 0xce, 0x40, 0xe6, 0xb4, 0xa1, 0xfd, 0x86, 0x0b, 0xf2, 0x97,
	0xf3, 0x47, 0xc8, 0xe7, 0x8e, 0x0e, 0x22, 0x14, 0xf7, 0xd3, 0x6e, 0x40, 0x12, 0x3c, 0x94, 0x5a,
	0x39, 0x50, 0x8e, 0xfe, 0x8f, 0x50, 0x7c, 0x7a, 0x35, 0xe6, 0xa1, 0x03, 0x6f, 0x3b, 0xf0, 0xb9,
	0xda, 0x27, 0xe4, 0xe3, 0xf5, 0xf2, 0xac, 0xc9, 0x78, 0x1b, 0xfc, 0x69, 0x0b, 0xce, 0x16, 0x78,
	0xdc, 0xce, 0xa0, 0xd0, 0x4f, 0x30, 0x97, 0x52, 0x57, 0xca, 0x77, 0xc9, 0xf2, 0xdc, 0x80, 0xb5,
	0x6d, 0x97, 0x46, 0xde, 0x90, 0xcf, 0x3f, 0xeb, 0x92, 0xe4, 0xfd, 0xc8, 0xd0, 0xe1, 0xc8, 0xd0,
	0xf7, 0x91, 0xa1, 0x97, 0x9a, 0xf5, 0x0e, 0x35, 0xeb, 0x7d, 0xd5, 0xac, 0x77, 0x17, 0xaf, 0xb7,
	0x6e, 0x53, 0xad, 0xb8, 0xd4, 0x85, 0xe8, 0xde, 0x5d, 0xd8, 0x52, 0x89, 0x67, 0xd1, 0x7c, 0x8f,
	0xdb, 0x97, 0x60, 0x57, 0x27, 0x3e, 0xf4, 0xf5, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x9a, 0x74,
	0xe0, 0xed, 0xb5, 0x01, 0x00, 0x00,
}

func (m *Request) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Request) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Request) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Content != nil {
		{
			size, err := m.Content.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRequest(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if m.CreatedAt != 0 {
		i = encodeVarintRequest(dAtA, i, uint64(m.CreatedAt))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintRequest(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x1a
	}
	if m.RequestID != 0 {
		i = encodeVarintRequest(dAtA, i, uint64(m.RequestID))
		i--
		dAtA[i] = 0x10
	}
	if len(m.ChainID) > 0 {
		i -= len(m.ChainID)
		copy(dAtA[i:], m.ChainID)
		i = encodeVarintRequest(dAtA, i, uint64(len(m.ChainID)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ContentRemoveAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ContentRemoveAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ContentRemoveAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintRequest(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintRequest(dAtA []byte, offset int, v uint64) int {
	offset -= sovRequest(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Request) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ChainID)
	if l > 0 {
		n += 1 + l + sovRequest(uint64(l))
	}
	if m.RequestID != 0 {
		n += 1 + sovRequest(uint64(m.RequestID))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovRequest(uint64(l))
	}
	if m.CreatedAt != 0 {
		n += 1 + sovRequest(uint64(m.CreatedAt))
	}
	if m.Content != nil {
		l = m.Content.Size()
		n += 1 + l + sovRequest(uint64(l))
	}
	return n
}

func (m *ContentRemoveAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovRequest(uint64(l))
	}
	return n
}

func sovRequest(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRequest(x uint64) (n int) {
	return sovRequest(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Request) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRequest
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
			return fmt.Errorf("proto: Request: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Request: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequest
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
				return ErrInvalidLengthRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestID", wireType)
			}
			m.RequestID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RequestID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequest
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
				return ErrInvalidLengthRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			m.CreatedAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequest
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CreatedAt |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Content", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequest
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
				return ErrInvalidLengthRequest
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Content == nil {
				m.Content = &types.Any{}
			}
			if err := m.Content.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRequest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRequest
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
func (m *ContentRemoveAccount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRequest
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
			return fmt.Errorf("proto: ContentRemoveAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ContentRemoveAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRequest
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
				return ErrInvalidLengthRequest
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRequest
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRequest(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRequest
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
func skipRequest(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRequest
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
					return 0, ErrIntOverflowRequest
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
					return 0, ErrIntOverflowRequest
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
				return 0, ErrInvalidLengthRequest
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRequest
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRequest
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRequest        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRequest          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRequest = fmt.Errorf("proto: unexpected end of group")
)
