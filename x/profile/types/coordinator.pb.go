// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: spn/profile/coordinator.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
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

type Coordinator struct {
	CoordinatorID uint64                 `protobuf:"varint,1,opt,name=coordinatorID,proto3" json:"coordinatorID,omitempty"`
	Address       string                 `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Description   CoordinatorDescription `protobuf:"bytes,3,opt,name=description,proto3" json:"description"`
	Active        bool                   `protobuf:"varint,4,opt,name=active,proto3" json:"active,omitempty"`
}

func (m *Coordinator) Reset()         { *m = Coordinator{} }
func (m *Coordinator) String() string { return proto.CompactTextString(m) }
func (*Coordinator) ProtoMessage()    {}
func (*Coordinator) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7929f77698b58d3, []int{0}
}
func (m *Coordinator) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Coordinator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Coordinator.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Coordinator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Coordinator.Merge(m, src)
}
func (m *Coordinator) XXX_Size() int {
	return m.Size()
}
func (m *Coordinator) XXX_DiscardUnknown() {
	xxx_messageInfo_Coordinator.DiscardUnknown(m)
}

var xxx_messageInfo_Coordinator proto.InternalMessageInfo

func (m *Coordinator) GetCoordinatorID() uint64 {
	if m != nil {
		return m.CoordinatorID
	}
	return 0
}

func (m *Coordinator) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Coordinator) GetDescription() CoordinatorDescription {
	if m != nil {
		return m.Description
	}
	return CoordinatorDescription{}
}

func (m *Coordinator) GetActive() bool {
	if m != nil {
		return m.Active
	}
	return false
}

type CoordinatorDescription struct {
	Identity string `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"`
	Website  string `protobuf:"bytes,2,opt,name=website,proto3" json:"website,omitempty"`
	Details  string `protobuf:"bytes,3,opt,name=details,proto3" json:"details,omitempty"`
}

func (m *CoordinatorDescription) Reset()         { *m = CoordinatorDescription{} }
func (m *CoordinatorDescription) String() string { return proto.CompactTextString(m) }
func (*CoordinatorDescription) ProtoMessage()    {}
func (*CoordinatorDescription) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7929f77698b58d3, []int{1}
}
func (m *CoordinatorDescription) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CoordinatorDescription) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CoordinatorDescription.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CoordinatorDescription) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CoordinatorDescription.Merge(m, src)
}
func (m *CoordinatorDescription) XXX_Size() int {
	return m.Size()
}
func (m *CoordinatorDescription) XXX_DiscardUnknown() {
	xxx_messageInfo_CoordinatorDescription.DiscardUnknown(m)
}

var xxx_messageInfo_CoordinatorDescription proto.InternalMessageInfo

func (m *CoordinatorDescription) GetIdentity() string {
	if m != nil {
		return m.Identity
	}
	return ""
}

func (m *CoordinatorDescription) GetWebsite() string {
	if m != nil {
		return m.Website
	}
	return ""
}

func (m *CoordinatorDescription) GetDetails() string {
	if m != nil {
		return m.Details
	}
	return ""
}

type CoordinatorByAddress struct {
	Address       string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	CoordinatorID uint64 `protobuf:"varint,2,opt,name=coordinatorID,proto3" json:"coordinatorID,omitempty"`
}

func (m *CoordinatorByAddress) Reset()         { *m = CoordinatorByAddress{} }
func (m *CoordinatorByAddress) String() string { return proto.CompactTextString(m) }
func (*CoordinatorByAddress) ProtoMessage()    {}
func (*CoordinatorByAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_f7929f77698b58d3, []int{2}
}
func (m *CoordinatorByAddress) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CoordinatorByAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CoordinatorByAddress.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CoordinatorByAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CoordinatorByAddress.Merge(m, src)
}
func (m *CoordinatorByAddress) XXX_Size() int {
	return m.Size()
}
func (m *CoordinatorByAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_CoordinatorByAddress.DiscardUnknown(m)
}

var xxx_messageInfo_CoordinatorByAddress proto.InternalMessageInfo

func (m *CoordinatorByAddress) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *CoordinatorByAddress) GetCoordinatorID() uint64 {
	if m != nil {
		return m.CoordinatorID
	}
	return 0
}

func init() {
	proto.RegisterType((*Coordinator)(nil), "spn.profile.Coordinator")
	proto.RegisterType((*CoordinatorDescription)(nil), "spn.profile.CoordinatorDescription")
	proto.RegisterType((*CoordinatorByAddress)(nil), "spn.profile.CoordinatorByAddress")
}

func init() { proto.RegisterFile("spn/profile/coordinator.proto", fileDescriptor_f7929f77698b58d3) }

var fileDescriptor_f7929f77698b58d3 = []byte{
	// 350 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xbb, 0x4e, 0xc3, 0x30,
	0x14, 0x8d, 0x4b, 0xd5, 0x87, 0x23, 0x16, 0xab, 0xaa, 0x42, 0x25, 0x42, 0x54, 0x18, 0xc2, 0x40,
	0x22, 0x95, 0x2f, 0x20, 0xed, 0x82, 0xd8, 0xcc, 0xc6, 0x82, 0xd2, 0xd8, 0xa4, 0x96, 0x5a, 0xdb,
	0xb2, 0xcd, 0xa3, 0x7f, 0xc1, 0xc7, 0xf0, 0x11, 0x1d, 0x18, 0x2a, 0x26, 0x26, 0x84, 0xda, 0x1f,
	0x41, 0xcd, 0xa3, 0x0d, 0xa2, 0x03, 0x5b, 0xce, 0x3d, 0x37, 0xf7, 0x9e, 0x73, 0x7c, 0xe1, 0xb1,
	0x96, 0x3c, 0x94, 0x4a, 0x3c, 0xb0, 0x29, 0x0d, 0x13, 0x21, 0x14, 0x61, 0x3c, 0x36, 0x42, 0x05,
	0x52, 0x09, 0x23, 0x90, 0xad, 0x25, 0x0f, 0x0a, 0xba, 0xd7, 0x49, 0x45, 0x2a, 0xb2, 0x7a, 0xb8,
	0xf9, 0xca, 0x5b, 0x7a, 0x47, 0x89, 0xd0, 0x33, 0xa1, 0xef, 0x73, 0x22, 0x07, 0x39, 0xd5, 0x7f,
	0x07, 0xd0, 0x1e, 0xee, 0x66, 0xa2, 0x33, 0x78, 0x58, 0x59, 0x71, 0x3d, 0x72, 0x80, 0x07, 0xfc,
	0x3a, 0xfe, 0x5d, 0x44, 0x03, 0xd8, 0x8c, 0x09, 0x51, 0x54, 0x6b, 0xa7, 0xe6, 0x01, 0xbf, 0x1d,
	0x39, 0x1f, 0x6f, 0x17, 0x9d, 0x62, 0xf0, 0x55, 0xce, 0xdc, 0x1a, 0xc5, 0x78, 0x8a, 0xcb, 0x46,
	0x74, 0x03, 0x6d, 0x42, 0x75, 0xa2, 0x98, 0x34, 0x4c, 0x70, 0xe7, 0xc0, 0x03, 0xbe, 0x3d, 0x38,
	0x0d, 0x2a, 0xea, 0x83, 0x8a, 0x90, 0xd1, 0xae, 0x35, 0xaa, 0x2f, 0xbe, 0x4e, 0x2c, 0x5c, 0xfd,
	0x1b, 0x75, 0x61, 0x23, 0x4e, 0x0c, 0x7b, 0xa2, 0x4e, 0xdd, 0x03, 0x7e, 0x0b, 0x17, 0xa8, 0x3f,
	0x81, 0xdd, 0xfd, 0x43, 0x50, 0x0f, 0xb6, 0x18, 0xa1, 0xdc, 0x30, 0x33, 0xcf, 0x3c, 0xb5, 0xf1,
	0x16, 0x23, 0x07, 0x36, 0x9f, 0xe9, 0x58, 0x33, 0x43, 0x73, 0x3b, 0xb8, 0x84, 0x1b, 0x86, 0x50,
	0x13, 0xb3, 0xa9, 0xce, 0x04, 0xb7, 0x71, 0x09, 0xfb, 0x12, 0x76, 0x2a, 0x9b, 0xa2, 0x79, 0xe1,
	0xba, 0x1a, 0x0d, 0xf8, 0x6f, 0x34, 0x7f, 0x42, 0xaf, 0xed, 0x09, 0x3d, 0x1a, 0x2e, 0x56, 0x2e,
	0x58, 0xae, 0x5c, 0xf0, 0xbd, 0x72, 0xc1, 0xeb, 0xda, 0xb5, 0x96, 0x6b, 0xd7, 0xfa, 0x5c, 0xbb,
	0xd6, 0xdd, 0x79, 0xca, 0xcc, 0xe4, 0x71, 0x1c, 0x24, 0x62, 0x16, 0x1a, 0xca, 0x09, 0x55, 0x33,
	0xc6, 0x4d, 0xb8, 0xb9, 0x9b, 0x97, 0xed, 0xe5, 0x98, 0xb9, 0xa4, 0x7a, 0xdc, 0xc8, 0x9e, 0xfd,
	0xf2, 0x27, 0x00, 0x00, 0xff, 0xff, 0x4b, 0xec, 0xd7, 0x45, 0x55, 0x02, 0x00, 0x00,
}

func (m *Coordinator) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Coordinator) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Coordinator) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Active {
		i--
		if m.Active {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	{
		size, err := m.Description.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintCoordinator(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintCoordinator(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if m.CoordinatorID != 0 {
		i = encodeVarintCoordinator(dAtA, i, uint64(m.CoordinatorID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *CoordinatorDescription) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CoordinatorDescription) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CoordinatorDescription) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Details) > 0 {
		i -= len(m.Details)
		copy(dAtA[i:], m.Details)
		i = encodeVarintCoordinator(dAtA, i, uint64(len(m.Details)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Website) > 0 {
		i -= len(m.Website)
		copy(dAtA[i:], m.Website)
		i = encodeVarintCoordinator(dAtA, i, uint64(len(m.Website)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Identity) > 0 {
		i -= len(m.Identity)
		copy(dAtA[i:], m.Identity)
		i = encodeVarintCoordinator(dAtA, i, uint64(len(m.Identity)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *CoordinatorByAddress) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CoordinatorByAddress) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CoordinatorByAddress) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CoordinatorID != 0 {
		i = encodeVarintCoordinator(dAtA, i, uint64(m.CoordinatorID))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintCoordinator(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintCoordinator(dAtA []byte, offset int, v uint64) int {
	offset -= sovCoordinator(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Coordinator) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CoordinatorID != 0 {
		n += 1 + sovCoordinator(uint64(m.CoordinatorID))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovCoordinator(uint64(l))
	}
	l = m.Description.Size()
	n += 1 + l + sovCoordinator(uint64(l))
	if m.Active {
		n += 2
	}
	return n
}

func (m *CoordinatorDescription) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Identity)
	if l > 0 {
		n += 1 + l + sovCoordinator(uint64(l))
	}
	l = len(m.Website)
	if l > 0 {
		n += 1 + l + sovCoordinator(uint64(l))
	}
	l = len(m.Details)
	if l > 0 {
		n += 1 + l + sovCoordinator(uint64(l))
	}
	return n
}

func (m *CoordinatorByAddress) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovCoordinator(uint64(l))
	}
	if m.CoordinatorID != 0 {
		n += 1 + sovCoordinator(uint64(m.CoordinatorID))
	}
	return n
}

func sovCoordinator(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCoordinator(x uint64) (n int) {
	return sovCoordinator(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Coordinator) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCoordinator
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
			return fmt.Errorf("proto: Coordinator: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Coordinator: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CoordinatorID", wireType)
			}
			m.CoordinatorID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoordinator
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CoordinatorID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoordinator
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
				return ErrInvalidLengthCoordinator
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCoordinator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoordinator
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
				return ErrInvalidLengthCoordinator
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCoordinator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Description.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Active", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoordinator
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
			m.Active = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipCoordinator(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCoordinator
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
func (m *CoordinatorDescription) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCoordinator
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
			return fmt.Errorf("proto: CoordinatorDescription: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CoordinatorDescription: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Identity", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoordinator
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
				return ErrInvalidLengthCoordinator
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCoordinator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Identity = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Website", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoordinator
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
				return ErrInvalidLengthCoordinator
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCoordinator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Website = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Details", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoordinator
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
				return ErrInvalidLengthCoordinator
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCoordinator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Details = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCoordinator(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCoordinator
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
func (m *CoordinatorByAddress) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCoordinator
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
			return fmt.Errorf("proto: CoordinatorByAddress: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CoordinatorByAddress: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoordinator
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
				return ErrInvalidLengthCoordinator
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCoordinator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CoordinatorID", wireType)
			}
			m.CoordinatorID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCoordinator
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CoordinatorID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipCoordinator(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCoordinator
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
func skipCoordinator(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCoordinator
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
					return 0, ErrIntOverflowCoordinator
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
					return 0, ErrIntOverflowCoordinator
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
				return 0, ErrInvalidLengthCoordinator
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCoordinator
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCoordinator
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCoordinator        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCoordinator          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCoordinator = fmt.Errorf("proto: unexpected end of group")
)
