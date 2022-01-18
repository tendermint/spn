// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: profile/validator.proto

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

type Validator struct {
	Address          string               `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	ConsensusAddress string               `protobuf:"bytes,2,opt,name=consensusAddress,proto3" json:"consensusAddress,omitempty"`
	Description      ValidatorDescription `protobuf:"bytes,3,opt,name=description,proto3" json:"description"`
}

func (m *Validator) Reset()         { *m = Validator{} }
func (m *Validator) String() string { return proto.CompactTextString(m) }
func (*Validator) ProtoMessage()    {}
func (*Validator) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e2276f43ab77aa3, []int{0}
}
func (m *Validator) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Validator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Validator.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Validator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Validator.Merge(m, src)
}
func (m *Validator) XXX_Size() int {
	return m.Size()
}
func (m *Validator) XXX_DiscardUnknown() {
	xxx_messageInfo_Validator.DiscardUnknown(m)
}

var xxx_messageInfo_Validator proto.InternalMessageInfo

func (m *Validator) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Validator) GetConsensusAddress() string {
	if m != nil {
		return m.ConsensusAddress
	}
	return ""
}

func (m *Validator) GetDescription() ValidatorDescription {
	if m != nil {
		return m.Description
	}
	return ValidatorDescription{}
}

type ValidatorDescription struct {
	Identity        string `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"`
	Moniker         string `protobuf:"bytes,2,opt,name=moniker,proto3" json:"moniker,omitempty"`
	Website         string `protobuf:"bytes,3,opt,name=website,proto3" json:"website,omitempty"`
	SecurityContact string `protobuf:"bytes,4,opt,name=securityContact,proto3" json:"securityContact,omitempty"`
	Details         string `protobuf:"bytes,5,opt,name=details,proto3" json:"details,omitempty"`
}

func (m *ValidatorDescription) Reset()         { *m = ValidatorDescription{} }
func (m *ValidatorDescription) String() string { return proto.CompactTextString(m) }
func (*ValidatorDescription) ProtoMessage()    {}
func (*ValidatorDescription) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e2276f43ab77aa3, []int{1}
}
func (m *ValidatorDescription) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ValidatorDescription) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ValidatorDescription.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ValidatorDescription) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidatorDescription.Merge(m, src)
}
func (m *ValidatorDescription) XXX_Size() int {
	return m.Size()
}
func (m *ValidatorDescription) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidatorDescription.DiscardUnknown(m)
}

var xxx_messageInfo_ValidatorDescription proto.InternalMessageInfo

func (m *ValidatorDescription) GetIdentity() string {
	if m != nil {
		return m.Identity
	}
	return ""
}

func (m *ValidatorDescription) GetMoniker() string {
	if m != nil {
		return m.Moniker
	}
	return ""
}

func (m *ValidatorDescription) GetWebsite() string {
	if m != nil {
		return m.Website
	}
	return ""
}

func (m *ValidatorDescription) GetSecurityContact() string {
	if m != nil {
		return m.SecurityContact
	}
	return ""
}

func (m *ValidatorDescription) GetDetails() string {
	if m != nil {
		return m.Details
	}
	return ""
}

type ValidatorByConsAddress struct {
	ConsensusAddress string `protobuf:"bytes,1,opt,name=consensusAddress,proto3" json:"consensusAddress,omitempty"`
	ValidatorAddress string `protobuf:"bytes,2,opt,name=validatorAddress,proto3" json:"validatorAddress,omitempty"`
}

func (m *ValidatorByConsAddress) Reset()         { *m = ValidatorByConsAddress{} }
func (m *ValidatorByConsAddress) String() string { return proto.CompactTextString(m) }
func (*ValidatorByConsAddress) ProtoMessage()    {}
func (*ValidatorByConsAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_8e2276f43ab77aa3, []int{2}
}
func (m *ValidatorByConsAddress) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ValidatorByConsAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ValidatorByConsAddress.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ValidatorByConsAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidatorByConsAddress.Merge(m, src)
}
func (m *ValidatorByConsAddress) XXX_Size() int {
	return m.Size()
}
func (m *ValidatorByConsAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidatorByConsAddress.DiscardUnknown(m)
}

var xxx_messageInfo_ValidatorByConsAddress proto.InternalMessageInfo

func (m *ValidatorByConsAddress) GetConsensusAddress() string {
	if m != nil {
		return m.ConsensusAddress
	}
	return ""
}

func (m *ValidatorByConsAddress) GetValidatorAddress() string {
	if m != nil {
		return m.ValidatorAddress
	}
	return ""
}

func init() {
	proto.RegisterType((*Validator)(nil), "tendermint.spn.profile.Validator")
	proto.RegisterType((*ValidatorDescription)(nil), "tendermint.spn.profile.ValidatorDescription")
	proto.RegisterType((*ValidatorByConsAddress)(nil), "tendermint.spn.profile.ValidatorByConsAddress")
}

func init() { proto.RegisterFile("profile/validator.proto", fileDescriptor_8e2276f43ab77aa3) }

var fileDescriptor_8e2276f43ab77aa3 = []byte{
	// 341 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x4f, 0x4e, 0xc2, 0x40,
	0x14, 0xc6, 0x3b, 0x8a, 0x7f, 0x18, 0x16, 0x92, 0x86, 0x60, 0xc3, 0xa2, 0x12, 0x56, 0x48, 0x4c,
	0x9b, 0xe8, 0x09, 0x04, 0x4f, 0x40, 0x8c, 0x0b, 0x77, 0xa5, 0xf3, 0xc4, 0x89, 0x30, 0x33, 0x99,
	0xf7, 0x50, 0xb9, 0x85, 0xb7, 0x30, 0xf1, 0x24, 0x2c, 0x59, 0xba, 0x32, 0x06, 0x2e, 0x62, 0x5a,
	0xda, 0x62, 0xb0, 0xee, 0xfa, 0xbd, 0xfe, 0xfa, 0xf5, 0x9b, 0x6f, 0x1e, 0x3f, 0x35, 0x56, 0x3f,
	0xc8, 0x09, 0x84, 0xcf, 0xd1, 0x44, 0x8a, 0x88, 0xb4, 0x0d, 0x8c, 0xd5, 0xa4, 0xdd, 0x26, 0x81,
	0x12, 0x60, 0xa7, 0x52, 0x51, 0x80, 0x46, 0x05, 0x19, 0xd7, 0x6a, 0x8c, 0xf5, 0x58, 0xa7, 0x48,
	0x98, 0x3c, 0x6d, 0xe8, 0xce, 0x3b, 0xe3, 0xd5, 0xbb, 0xdc, 0xc1, 0xf5, 0xf8, 0x51, 0x24, 0x84,
	0x05, 0x44, 0x8f, 0xb5, 0x59, 0xb7, 0x3a, 0xcc, 0xa5, 0xdb, 0xe3, 0xf5, 0x58, 0x2b, 0x04, 0x85,
	0x33, 0xbc, 0xce, 0x90, 0xbd, 0x14, 0xf9, 0x33, 0x77, 0x6f, 0x79, 0x4d, 0x00, 0xc6, 0x56, 0x1a,
	0x92, 0x5a, 0x79, 0xfb, 0x6d, 0xd6, 0xad, 0x5d, 0x5e, 0x04, 0xe5, 0xb9, 0x82, 0xe2, 0xef, 0x37,
	0xdb, 0x6f, 0xfa, 0x95, 0xc5, 0xd7, 0x99, 0x33, 0xfc, 0x6d, 0xd3, 0xf9, 0x60, 0xbc, 0x51, 0xc6,
	0xba, 0x2d, 0x7e, 0x2c, 0x05, 0x28, 0x92, 0x34, 0xcf, 0x52, 0x17, 0x3a, 0x39, 0xd0, 0x54, 0x2b,
	0xf9, 0x04, 0x36, 0x4b, 0x9b, 0xcb, 0xe4, 0xcd, 0x0b, 0x8c, 0x50, 0x12, 0xa4, 0x01, 0xab, 0xc3,
	0x5c, 0xba, 0x5d, 0x7e, 0x82, 0x10, 0xcf, 0xac, 0xa4, 0xf9, 0x40, 0x2b, 0x8a, 0x62, 0xf2, 0x2a,
	0x29, 0xb1, 0x3b, 0x4e, 0x3c, 0x04, 0x50, 0x24, 0x27, 0xe8, 0x1d, 0x6c, 0x3c, 0x32, 0xd9, 0x31,
	0xbc, 0x59, 0x64, 0xed, 0x27, 0x7c, 0x51, 0x4e, 0x59, 0x91, 0xec, 0x9f, 0x22, 0x7b, 0xbc, 0x5e,
	0xdc, 0xee, 0x4e, 0xe9, 0xbb, 0xf3, 0xfe, 0x60, 0xb1, 0xf2, 0xd9, 0x72, 0xe5, 0xb3, 0xef, 0x95,
	0xcf, 0xde, 0xd6, 0xbe, 0xb3, 0x5c, 0xfb, 0xce, 0xe7, 0xda, 0x77, 0xee, 0xcf, 0xc7, 0x92, 0x1e,
	0x67, 0xa3, 0x20, 0xd6, 0xd3, 0x70, 0x7b, 0x07, 0x21, 0x1a, 0x15, 0xbe, 0x86, 0xf9, 0x16, 0xd1,
	0xdc, 0x00, 0x8e, 0x0e, 0xd3, 0xa5, 0xb8, 0xfa, 0x09, 0x00, 0x00, 0xff, 0xff, 0xe4, 0x32, 0xe2,
	0xe9, 0x5d, 0x02, 0x00, 0x00,
}

func (m *Validator) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Validator) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Validator) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Description.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintValidator(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.ConsensusAddress) > 0 {
		i -= len(m.ConsensusAddress)
		copy(dAtA[i:], m.ConsensusAddress)
		i = encodeVarintValidator(dAtA, i, uint64(len(m.ConsensusAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintValidator(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ValidatorDescription) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ValidatorDescription) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ValidatorDescription) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Details) > 0 {
		i -= len(m.Details)
		copy(dAtA[i:], m.Details)
		i = encodeVarintValidator(dAtA, i, uint64(len(m.Details)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.SecurityContact) > 0 {
		i -= len(m.SecurityContact)
		copy(dAtA[i:], m.SecurityContact)
		i = encodeVarintValidator(dAtA, i, uint64(len(m.SecurityContact)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Website) > 0 {
		i -= len(m.Website)
		copy(dAtA[i:], m.Website)
		i = encodeVarintValidator(dAtA, i, uint64(len(m.Website)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Moniker) > 0 {
		i -= len(m.Moniker)
		copy(dAtA[i:], m.Moniker)
		i = encodeVarintValidator(dAtA, i, uint64(len(m.Moniker)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Identity) > 0 {
		i -= len(m.Identity)
		copy(dAtA[i:], m.Identity)
		i = encodeVarintValidator(dAtA, i, uint64(len(m.Identity)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ValidatorByConsAddress) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ValidatorByConsAddress) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ValidatorByConsAddress) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ValidatorAddress) > 0 {
		i -= len(m.ValidatorAddress)
		copy(dAtA[i:], m.ValidatorAddress)
		i = encodeVarintValidator(dAtA, i, uint64(len(m.ValidatorAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ConsensusAddress) > 0 {
		i -= len(m.ConsensusAddress)
		copy(dAtA[i:], m.ConsensusAddress)
		i = encodeVarintValidator(dAtA, i, uint64(len(m.ConsensusAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintValidator(dAtA []byte, offset int, v uint64) int {
	offset -= sovValidator(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Validator) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovValidator(uint64(l))
	}
	l = len(m.ConsensusAddress)
	if l > 0 {
		n += 1 + l + sovValidator(uint64(l))
	}
	l = m.Description.Size()
	n += 1 + l + sovValidator(uint64(l))
	return n
}

func (m *ValidatorDescription) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Identity)
	if l > 0 {
		n += 1 + l + sovValidator(uint64(l))
	}
	l = len(m.Moniker)
	if l > 0 {
		n += 1 + l + sovValidator(uint64(l))
	}
	l = len(m.Website)
	if l > 0 {
		n += 1 + l + sovValidator(uint64(l))
	}
	l = len(m.SecurityContact)
	if l > 0 {
		n += 1 + l + sovValidator(uint64(l))
	}
	l = len(m.Details)
	if l > 0 {
		n += 1 + l + sovValidator(uint64(l))
	}
	return n
}

func (m *ValidatorByConsAddress) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ConsensusAddress)
	if l > 0 {
		n += 1 + l + sovValidator(uint64(l))
	}
	l = len(m.ValidatorAddress)
	if l > 0 {
		n += 1 + l + sovValidator(uint64(l))
	}
	return n
}

func sovValidator(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozValidator(x uint64) (n int) {
	return sovValidator(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Validator) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowValidator
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
			return fmt.Errorf("proto: Validator: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Validator: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
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
				return ErrInvalidLengthValidator
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConsensusAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
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
				return ErrInvalidLengthValidator
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ConsensusAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
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
				return ErrInvalidLengthValidator
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Description.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipValidator(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthValidator
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
func (m *ValidatorDescription) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowValidator
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
			return fmt.Errorf("proto: ValidatorDescription: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ValidatorDescription: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Identity", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
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
				return ErrInvalidLengthValidator
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Identity = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Moniker", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
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
				return ErrInvalidLengthValidator
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Moniker = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Website", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
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
				return ErrInvalidLengthValidator
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Website = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SecurityContact", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
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
				return ErrInvalidLengthValidator
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SecurityContact = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Details", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
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
				return ErrInvalidLengthValidator
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Details = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipValidator(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthValidator
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
func (m *ValidatorByConsAddress) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowValidator
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
			return fmt.Errorf("proto: ValidatorByConsAddress: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ValidatorByConsAddress: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConsensusAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
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
				return ErrInvalidLengthValidator
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ConsensusAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowValidator
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
				return ErrInvalidLengthValidator
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthValidator
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ValidatorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipValidator(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthValidator
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
func skipValidator(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowValidator
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
					return 0, ErrIntOverflowValidator
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
					return 0, ErrIntOverflowValidator
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
				return 0, ErrInvalidLengthValidator
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupValidator
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthValidator
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthValidator        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowValidator          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupValidator = fmt.Errorf("proto: unexpected end of group")
)
