// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: chat/v1beta/messages.proto

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

type MsgCreateChannel struct {
	Creator     github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,1,opt,name=creator,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"creator,omitempty"`
	Title       string                                        `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description string                                        `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Payload     []byte                                        `protobuf:"bytes,4,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (m *MsgCreateChannel) Reset()         { *m = MsgCreateChannel{} }
func (m *MsgCreateChannel) String() string { return proto.CompactTextString(m) }
func (*MsgCreateChannel) ProtoMessage()    {}
func (*MsgCreateChannel) Descriptor() ([]byte, []int) {
	return fileDescriptor_3395017881b3c0cc, []int{0}
}
func (m *MsgCreateChannel) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateChannel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateChannel.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateChannel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateChannel.Merge(m, src)
}
func (m *MsgCreateChannel) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateChannel) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateChannel.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateChannel proto.InternalMessageInfo

func (m *MsgCreateChannel) GetCreator() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Creator
	}
	return nil
}

func (m *MsgCreateChannel) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *MsgCreateChannel) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *MsgCreateChannel) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

type MsgSendMessage struct {
	ChannelID   int32                                         `protobuf:"varint,1,opt,name=channelID,proto3" json:"channelID,omitempty"`
	Creator     github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,2,opt,name=creator,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"creator,omitempty"`
	Content     string                                        `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	Tags        []string                                      `protobuf:"bytes,4,rep,name=tags,proto3" json:"tags,omitempty"`
	PollOptions []string                                      `protobuf:"bytes,5,rep,name=pollOptions,proto3" json:"pollOptions,omitempty"`
	Payload     []byte                                        `protobuf:"bytes,6,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (m *MsgSendMessage) Reset()         { *m = MsgSendMessage{} }
func (m *MsgSendMessage) String() string { return proto.CompactTextString(m) }
func (*MsgSendMessage) ProtoMessage()    {}
func (*MsgSendMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_3395017881b3c0cc, []int{1}
}
func (m *MsgSendMessage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgSendMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgSendMessage.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgSendMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSendMessage.Merge(m, src)
}
func (m *MsgSendMessage) XXX_Size() int {
	return m.Size()
}
func (m *MsgSendMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSendMessage.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSendMessage proto.InternalMessageInfo

func (m *MsgSendMessage) GetChannelID() int32 {
	if m != nil {
		return m.ChannelID
	}
	return 0
}

func (m *MsgSendMessage) GetCreator() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Creator
	}
	return nil
}

func (m *MsgSendMessage) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *MsgSendMessage) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *MsgSendMessage) GetPollOptions() []string {
	if m != nil {
		return m.PollOptions
	}
	return nil
}

func (m *MsgSendMessage) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

type MsgVotePoll struct {
	ChannelID    int32                                         `protobuf:"varint,1,opt,name=channelID,proto3" json:"channelID,omitempty"`
	MessageIndex int32                                         `protobuf:"varint,2,opt,name=messageIndex,proto3" json:"messageIndex,omitempty"`
	Creator      github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,3,opt,name=creator,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"creator,omitempty"`
	Value        int32                                         `protobuf:"varint,4,opt,name=value,proto3" json:"value,omitempty"`
	Payload      []byte                                        `protobuf:"bytes,5,opt,name=payload,proto3" json:"payload,omitempty"`
}

func (m *MsgVotePoll) Reset()         { *m = MsgVotePoll{} }
func (m *MsgVotePoll) String() string { return proto.CompactTextString(m) }
func (*MsgVotePoll) ProtoMessage()    {}
func (*MsgVotePoll) Descriptor() ([]byte, []int) {
	return fileDescriptor_3395017881b3c0cc, []int{2}
}
func (m *MsgVotePoll) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgVotePoll) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgVotePoll.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgVotePoll) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgVotePoll.Merge(m, src)
}
func (m *MsgVotePoll) XXX_Size() int {
	return m.Size()
}
func (m *MsgVotePoll) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgVotePoll.DiscardUnknown(m)
}

var xxx_messageInfo_MsgVotePoll proto.InternalMessageInfo

func (m *MsgVotePoll) GetChannelID() int32 {
	if m != nil {
		return m.ChannelID
	}
	return 0
}

func (m *MsgVotePoll) GetMessageIndex() int32 {
	if m != nil {
		return m.MessageIndex
	}
	return 0
}

func (m *MsgVotePoll) GetCreator() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Creator
	}
	return nil
}

func (m *MsgVotePoll) GetValue() int32 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *MsgVotePoll) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func init() {
	proto.RegisterType((*MsgCreateChannel)(nil), "spn.chat.v1beta1.MsgCreateChannel")
	proto.RegisterType((*MsgSendMessage)(nil), "spn.chat.v1beta1.MsgSendMessage")
	proto.RegisterType((*MsgVotePoll)(nil), "spn.chat.v1beta1.MsgVotePoll")
}

func init() { proto.RegisterFile("chat/v1beta/messages.proto", fileDescriptor_3395017881b3c0cc) }

var fileDescriptor_3395017881b3c0cc = []byte{
	// 406 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0x41, 0x6a, 0xdb, 0x40,
	0x14, 0x86, 0x3d, 0xb1, 0x15, 0xe3, 0x49, 0x28, 0x61, 0xc8, 0x62, 0x08, 0x45, 0x35, 0x5e, 0x14,
	0x6f, 0xa2, 0xc1, 0xf4, 0x02, 0x4d, 0xd2, 0x4d, 0x28, 0xa2, 0x45, 0x85, 0x2e, 0xba, 0x1b, 0xcf,
	0x3c, 0xc6, 0xa2, 0xa3, 0x19, 0xa1, 0x37, 0x09, 0xc9, 0x2d, 0x7a, 0x95, 0xde, 0xa2, 0x8b, 0x2e,
	0xbc, 0xec, 0xaa, 0x14, 0x9b, 0x5e, 0xa2, 0xab, 0xa2, 0x91, 0x4d, 0xe5, 0x55, 0x17, 0xed, 0x4a,
	0xef, 0xff, 0x7f, 0x21, 0xfe, 0x4f, 0xef, 0xd1, 0x0b, 0xb5, 0x92, 0x41, 0xdc, 0x2f, 0x96, 0x10,
	0xa4, 0xa8, 0x00, 0x51, 0x1a, 0xc0, 0xac, 0x6e, 0x7c, 0xf0, 0xec, 0x0c, 0x6b, 0x97, 0xb5, 0x79,
	0xd6, 0xe5, 0x8b, 0x8b, 0x73, 0xe3, 0x8d, 0x8f, 0xa1, 0x68, 0xa7, 0xee, 0xbd, 0xd9, 0x67, 0x42,
	0xcf, 0x72, 0x34, 0x37, 0x0d, 0xc8, 0x00, 0x37, 0x2b, 0xe9, 0x1c, 0x58, 0xf6, 0x9a, 0x8e, 0x55,
	0x6b, 0xf8, 0x86, 0x93, 0x29, 0x99, 0x9f, 0x5e, 0x2f, 0x7e, 0x7d, 0x7f, 0x76, 0x69, 0xca, 0xb0,
	0xba, 0x5b, 0x66, 0xca, 0x57, 0x42, 0x79, 0xac, 0x3c, 0xee, 0x1e, 0x97, 0xa8, 0x3f, 0x8a, 0xf0,
	0x58, 0x03, 0x66, 0x57, 0x4a, 0x5d, 0x69, 0xdd, 0x00, 0x62, 0xb1, 0xff, 0x02, 0x3b, 0xa7, 0x49,
	0x28, 0x83, 0x05, 0x7e, 0x34, 0x25, 0xf3, 0x49, 0xd1, 0x09, 0x36, 0xa5, 0x27, 0x1a, 0x50, 0x35,
	0x65, 0x1d, 0x4a, 0xef, 0xf8, 0x30, 0x66, 0x7d, 0x8b, 0x71, 0x3a, 0xae, 0xe5, 0xa3, 0xf5, 0x52,
	0xf3, 0x51, 0x5b, 0xa2, 0xd8, 0xcb, 0xd9, 0x4f, 0x42, 0x9f, 0xe4, 0x68, 0xde, 0x81, 0xd3, 0x79,
	0x47, 0xcd, 0x9e, 0xd2, 0x89, 0xea, 0xca, 0xdf, 0xbe, 0x8a, 0x9d, 0x93, 0xe2, 0x8f, 0xd1, 0xe7,
	0x39, 0xfa, 0x67, 0x1e, 0x4e, 0xc7, 0xca, 0xbb, 0x00, 0x2e, 0xec, 0x5a, 0xef, 0x25, 0x63, 0x74,
	0x14, 0xa4, 0x41, 0x3e, 0x9a, 0x0e, 0xe7, 0x93, 0x22, 0xce, 0x2d, 0x67, 0xed, 0xad, 0x7d, 0x13,
	0x99, 0x90, 0x27, 0x31, 0xea, 0x5b, 0x7d, 0xce, 0xe3, 0x43, 0xce, 0xaf, 0x84, 0x9e, 0xe4, 0x68,
	0xde, 0xfb, 0x00, 0x6f, 0xbd, 0xb5, 0x7f, 0x81, 0x9c, 0xd1, 0xd3, 0xdd, 0x0d, 0xdc, 0x3a, 0x0d,
	0x0f, 0x91, 0x34, 0x29, 0x0e, 0xbc, 0xfe, 0x8f, 0x18, 0xfe, 0x8f, 0xc5, 0xde, 0x4b, 0x7b, 0x07,
	0x71, 0x3d, 0x49, 0xd1, 0x89, 0x3e, 0x4e, 0x72, 0x80, 0x73, 0xfd, 0xf2, 0xcb, 0x26, 0x25, 0xeb,
	0x4d, 0x4a, 0x7e, 0x6c, 0x52, 0xf2, 0x69, 0x9b, 0x0e, 0xd6, 0xdb, 0x74, 0xf0, 0x6d, 0x9b, 0x0e,
	0x3e, 0x3c, 0xef, 0x35, 0x08, 0xe0, 0x34, 0x34, 0x55, 0xe9, 0x82, 0xc0, 0xda, 0x89, 0x07, 0x11,
	0x8f, 0x3c, 0xb6, 0x58, 0x1e, 0xc7, 0x9b, 0x7d, 0xf1, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x20, 0xda,
	0xed, 0x7f, 0xf9, 0x02, 0x00, 0x00,
}

func (m *MsgCreateChannel) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateChannel) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateChannel) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Payload) > 0 {
		i -= len(m.Payload)
		copy(dAtA[i:], m.Payload)
		i = encodeVarintMessages(dAtA, i, uint64(len(m.Payload)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintMessages(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintMessages(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintMessages(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgSendMessage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgSendMessage) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgSendMessage) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Payload) > 0 {
		i -= len(m.Payload)
		copy(dAtA[i:], m.Payload)
		i = encodeVarintMessages(dAtA, i, uint64(len(m.Payload)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.PollOptions) > 0 {
		for iNdEx := len(m.PollOptions) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.PollOptions[iNdEx])
			copy(dAtA[i:], m.PollOptions[iNdEx])
			i = encodeVarintMessages(dAtA, i, uint64(len(m.PollOptions[iNdEx])))
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.Tags) > 0 {
		for iNdEx := len(m.Tags) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Tags[iNdEx])
			copy(dAtA[i:], m.Tags[iNdEx])
			i = encodeVarintMessages(dAtA, i, uint64(len(m.Tags[iNdEx])))
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Content) > 0 {
		i -= len(m.Content)
		copy(dAtA[i:], m.Content)
		i = encodeVarintMessages(dAtA, i, uint64(len(m.Content)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintMessages(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x12
	}
	if m.ChannelID != 0 {
		i = encodeVarintMessages(dAtA, i, uint64(m.ChannelID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MsgVotePoll) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgVotePoll) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgVotePoll) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Payload) > 0 {
		i -= len(m.Payload)
		copy(dAtA[i:], m.Payload)
		i = encodeVarintMessages(dAtA, i, uint64(len(m.Payload)))
		i--
		dAtA[i] = 0x2a
	}
	if m.Value != 0 {
		i = encodeVarintMessages(dAtA, i, uint64(m.Value))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintMessages(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x1a
	}
	if m.MessageIndex != 0 {
		i = encodeVarintMessages(dAtA, i, uint64(m.MessageIndex))
		i--
		dAtA[i] = 0x10
	}
	if m.ChannelID != 0 {
		i = encodeVarintMessages(dAtA, i, uint64(m.ChannelID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintMessages(dAtA []byte, offset int, v uint64) int {
	offset -= sovMessages(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgCreateChannel) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovMessages(uint64(l))
	}
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovMessages(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovMessages(uint64(l))
	}
	l = len(m.Payload)
	if l > 0 {
		n += 1 + l + sovMessages(uint64(l))
	}
	return n
}

func (m *MsgSendMessage) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ChannelID != 0 {
		n += 1 + sovMessages(uint64(m.ChannelID))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovMessages(uint64(l))
	}
	l = len(m.Content)
	if l > 0 {
		n += 1 + l + sovMessages(uint64(l))
	}
	if len(m.Tags) > 0 {
		for _, s := range m.Tags {
			l = len(s)
			n += 1 + l + sovMessages(uint64(l))
		}
	}
	if len(m.PollOptions) > 0 {
		for _, s := range m.PollOptions {
			l = len(s)
			n += 1 + l + sovMessages(uint64(l))
		}
	}
	l = len(m.Payload)
	if l > 0 {
		n += 1 + l + sovMessages(uint64(l))
	}
	return n
}

func (m *MsgVotePoll) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ChannelID != 0 {
		n += 1 + sovMessages(uint64(m.ChannelID))
	}
	if m.MessageIndex != 0 {
		n += 1 + sovMessages(uint64(m.MessageIndex))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovMessages(uint64(l))
	}
	if m.Value != 0 {
		n += 1 + sovMessages(uint64(m.Value))
	}
	l = len(m.Payload)
	if l > 0 {
		n += 1 + l + sovMessages(uint64(l))
	}
	return n
}

func sovMessages(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMessages(x uint64) (n int) {
	return sovMessages(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgCreateChannel) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessages
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
			return fmt.Errorf("proto: MsgCreateChannel: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateChannel: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMessages
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMessages
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = append(m.Creator[:0], dAtA[iNdEx:postIndex]...)
			if m.Creator == nil {
				m.Creator = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
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
				return ErrInvalidLengthMessages
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessages
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
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
				return ErrInvalidLengthMessages
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessages
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payload", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMessages
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMessages
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Payload = append(m.Payload[:0], dAtA[iNdEx:postIndex]...)
			if m.Payload == nil {
				m.Payload = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessages(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMessages
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
func (m *MsgSendMessage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessages
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
			return fmt.Errorf("proto: MsgSendMessage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgSendMessage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChannelID", wireType)
			}
			m.ChannelID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ChannelID |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMessages
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMessages
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = append(m.Creator[:0], dAtA[iNdEx:postIndex]...)
			if m.Creator == nil {
				m.Creator = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Content", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
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
				return ErrInvalidLengthMessages
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessages
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Content = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tags", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
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
				return ErrInvalidLengthMessages
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessages
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tags = append(m.Tags, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PollOptions", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
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
				return ErrInvalidLengthMessages
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMessages
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PollOptions = append(m.PollOptions, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payload", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMessages
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMessages
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Payload = append(m.Payload[:0], dAtA[iNdEx:postIndex]...)
			if m.Payload == nil {
				m.Payload = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessages(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMessages
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
func (m *MsgVotePoll) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMessages
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
			return fmt.Errorf("proto: MsgVotePoll: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgVotePoll: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChannelID", wireType)
			}
			m.ChannelID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ChannelID |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MessageIndex", wireType)
			}
			m.MessageIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MessageIndex |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMessages
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMessages
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = append(m.Creator[:0], dAtA[iNdEx:postIndex]...)
			if m.Creator == nil {
				m.Creator = []byte{}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			m.Value = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Value |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payload", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMessages
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthMessages
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthMessages
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Payload = append(m.Payload[:0], dAtA[iNdEx:postIndex]...)
			if m.Payload == nil {
				m.Payload = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMessages(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMessages
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
func skipMessages(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMessages
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
					return 0, ErrIntOverflowMessages
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
					return 0, ErrIntOverflowMessages
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
				return 0, ErrInvalidLengthMessages
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMessages
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMessages
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMessages        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMessages          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMessages = fmt.Errorf("proto: unexpected end of group")
)
