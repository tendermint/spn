// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: campaign/query.proto

package types

import (
	context "context"
	fmt "fmt"
	query "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// this line is used by starport scaffolding # 3
type QueryGetCampaignChainsRequest struct {
	CampaignID uint64 `protobuf:"varint,1,opt,name=campaignID,proto3" json:"campaignID,omitempty"`
}

func (m *QueryGetCampaignChainsRequest) Reset()         { *m = QueryGetCampaignChainsRequest{} }
func (m *QueryGetCampaignChainsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGetCampaignChainsRequest) ProtoMessage()    {}
func (*QueryGetCampaignChainsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a55190e2afa5f29, []int{0}
}
func (m *QueryGetCampaignChainsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetCampaignChainsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetCampaignChainsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetCampaignChainsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetCampaignChainsRequest.Merge(m, src)
}
func (m *QueryGetCampaignChainsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetCampaignChainsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetCampaignChainsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetCampaignChainsRequest proto.InternalMessageInfo

func (m *QueryGetCampaignChainsRequest) GetCampaignID() uint64 {
	if m != nil {
		return m.CampaignID
	}
	return 0
}

type QueryGetCampaignChainsResponse struct {
	CampaignChains CampaignChains `protobuf:"bytes,1,opt,name=campaignChains,proto3" json:"campaignChains"`
}

func (m *QueryGetCampaignChainsResponse) Reset()         { *m = QueryGetCampaignChainsResponse{} }
func (m *QueryGetCampaignChainsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGetCampaignChainsResponse) ProtoMessage()    {}
func (*QueryGetCampaignChainsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a55190e2afa5f29, []int{1}
}
func (m *QueryGetCampaignChainsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetCampaignChainsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetCampaignChainsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetCampaignChainsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetCampaignChainsResponse.Merge(m, src)
}
func (m *QueryGetCampaignChainsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetCampaignChainsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetCampaignChainsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetCampaignChainsResponse proto.InternalMessageInfo

func (m *QueryGetCampaignChainsResponse) GetCampaignChains() CampaignChains {
	if m != nil {
		return m.CampaignChains
	}
	return CampaignChains{}
}

type QueryAllCampaignChainsRequest struct {
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllCampaignChainsRequest) Reset()         { *m = QueryAllCampaignChainsRequest{} }
func (m *QueryAllCampaignChainsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAllCampaignChainsRequest) ProtoMessage()    {}
func (*QueryAllCampaignChainsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a55190e2afa5f29, []int{2}
}
func (m *QueryAllCampaignChainsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllCampaignChainsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllCampaignChainsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllCampaignChainsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllCampaignChainsRequest.Merge(m, src)
}
func (m *QueryAllCampaignChainsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllCampaignChainsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllCampaignChainsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllCampaignChainsRequest proto.InternalMessageInfo

func (m *QueryAllCampaignChainsRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type QueryAllCampaignChainsResponse struct {
	CampaignChains []CampaignChains    `protobuf:"bytes,1,rep,name=campaignChains,proto3" json:"campaignChains"`
	Pagination     *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllCampaignChainsResponse) Reset()         { *m = QueryAllCampaignChainsResponse{} }
func (m *QueryAllCampaignChainsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAllCampaignChainsResponse) ProtoMessage()    {}
func (*QueryAllCampaignChainsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7a55190e2afa5f29, []int{3}
}
func (m *QueryAllCampaignChainsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllCampaignChainsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllCampaignChainsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllCampaignChainsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllCampaignChainsResponse.Merge(m, src)
}
func (m *QueryAllCampaignChainsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllCampaignChainsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllCampaignChainsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllCampaignChainsResponse proto.InternalMessageInfo

func (m *QueryAllCampaignChainsResponse) GetCampaignChains() []CampaignChains {
	if m != nil {
		return m.CampaignChains
	}
	return nil
}

func (m *QueryAllCampaignChainsResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryGetCampaignChainsRequest)(nil), "tendermint.spn.campaign.QueryGetCampaignChainsRequest")
	proto.RegisterType((*QueryGetCampaignChainsResponse)(nil), "tendermint.spn.campaign.QueryGetCampaignChainsResponse")
	proto.RegisterType((*QueryAllCampaignChainsRequest)(nil), "tendermint.spn.campaign.QueryAllCampaignChainsRequest")
	proto.RegisterType((*QueryAllCampaignChainsResponse)(nil), "tendermint.spn.campaign.QueryAllCampaignChainsResponse")
}

func init() { proto.RegisterFile("campaign/query.proto", fileDescriptor_7a55190e2afa5f29) }

var fileDescriptor_7a55190e2afa5f29 = []byte{
	// 438 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0xcf, 0xaa, 0x13, 0x31,
	0x14, 0xc6, 0x27, 0xf7, 0x5e, 0x5d, 0x44, 0xb8, 0x60, 0xb8, 0xa0, 0x14, 0x8d, 0x32, 0x0b, 0xaf,
	0x76, 0x91, 0xd0, 0x2a, 0x75, 0x23, 0x48, 0xff, 0x60, 0x71, 0xa7, 0x03, 0x6e, 0xdc, 0x48, 0x66,
	0x0c, 0xe9, 0xc0, 0x4c, 0x92, 0x36, 0xa9, 0x5a, 0xc4, 0x85, 0x3e, 0x81, 0xe0, 0xb3, 0xb8, 0xd6,
	0x65, 0x97, 0x05, 0x37, 0xae, 0x44, 0x5a, 0x1f, 0x44, 0x3a, 0x93, 0x71, 0x3a, 0xa5, 0xd3, 0x2a,
	0xee, 0xc2, 0xe4, 0x9c, 0xef, 0x7c, 0xbf, 0x2f, 0x67, 0xe0, 0x59, 0xc4, 0x52, 0xcd, 0x62, 0x21,
	0xe9, 0x78, 0xca, 0x27, 0x33, 0xa2, 0x27, 0xca, 0x2a, 0x74, 0xc5, 0x72, 0xf9, 0x92, 0x4f, 0xd2,
	0x58, 0x5a, 0x62, 0xb4, 0x24, 0x45, 0x51, 0xe3, 0x9a, 0x50, 0x4a, 0x24, 0x9c, 0x32, 0x1d, 0x53,
	0x26, 0xa5, 0xb2, 0xcc, 0xc6, 0x4a, 0x9a, 0xbc, 0xad, 0xd1, 0x8c, 0x94, 0x49, 0x95, 0xa1, 0x21,
	0x33, 0x3c, 0xd7, 0xa3, 0xaf, 0x5a, 0x21, 0xb7, 0xac, 0x45, 0x35, 0x13, 0xb1, 0xcc, 0x8a, 0x5d,
	0xed, 0x99, 0x50, 0x42, 0x65, 0x47, 0xba, 0x3e, 0xb9, 0xaf, 0xf8, 0x8f, 0x9d, 0xe2, 0xf0, 0x22,
	0x1a, 0xb1, 0xb8, 0x98, 0xe0, 0x3f, 0x84, 0xd7, 0x9f, 0xae, 0x75, 0x87, 0xdc, 0xf6, 0x5d, 0x41,
	0x3f, 0xbb, 0x0f, 0xf8, 0x78, 0xca, 0x8d, 0x45, 0x18, 0xc2, 0xa2, 0xf3, 0xf1, 0xe0, 0x2a, 0xb8,
	0x09, 0x6e, 0x9f, 0x04, 0x1b, 0x5f, 0xfc, 0xd7, 0x10, 0xd7, 0x09, 0x18, 0xad, 0xa4, 0xe1, 0xe8,
	0x19, 0x3c, 0x8d, 0x2a, 0x37, 0x99, 0xca, 0xa5, 0xf6, 0x39, 0xa9, 0x09, 0x85, 0x54, 0x85, 0x7a,
	0x27, 0xf3, 0x1f, 0x37, 0xbc, 0x60, 0x4b, 0xc4, 0x17, 0xce, 0x79, 0x37, 0x49, 0x76, 0x3b, 0x7f,
	0x04, 0x61, 0x19, 0x92, 0x9b, 0x79, 0x8b, 0xe4, 0x89, 0x92, 0x75, 0xa2, 0x24, 0x7f, 0x21, 0x97,
	0x28, 0x79, 0xc2, 0x04, 0x77, 0xbd, 0xc1, 0x46, 0xa7, 0xff, 0x15, 0x38, 0xc4, 0x1d, 0x93, 0xf6,
	0x20, 0x1e, 0xff, 0x37, 0x22, 0x1a, 0x56, 0x08, 0x8e, 0x5c, 0x6a, 0x87, 0x08, 0x72, 0x4f, 0x9b,
	0x08, 0xed, 0xf7, 0xc7, 0xf0, 0x42, 0x86, 0x80, 0xbe, 0x00, 0x78, 0x5a, 0x9d, 0x8d, 0x3a, 0xb5,
	0x26, 0xf7, 0x6e, 0x46, 0xe3, 0xfe, 0x3f, 0xf7, 0xe5, 0xce, 0xfc, 0x07, 0x1f, 0xbe, 0xfd, 0xfa,
	0x74, 0xd4, 0x41, 0xf7, 0x68, 0x29, 0x40, 0x8d, 0x2e, 0x57, 0x94, 0x56, 0x73, 0xa0, 0x6f, 0xcb,
	0x7d, 0x7b, 0x87, 0x3e, 0x03, 0x78, 0xb9, 0x2a, 0xdc, 0x4d, 0x92, 0x43, 0x10, 0x75, 0x4b, 0x72,
	0x08, 0xa2, 0xf6, 0xc9, 0x7d, 0x9a, 0x41, 0xdc, 0x41, 0xe7, 0x7f, 0x09, 0xd1, 0x1b, 0xcc, 0x97,
	0x18, 0x2c, 0x96, 0x18, 0xfc, 0x5c, 0x62, 0xf0, 0x71, 0x85, 0xbd, 0xc5, 0x0a, 0x7b, 0xdf, 0x57,
	0xd8, 0x7b, 0xde, 0x14, 0xb1, 0x1d, 0x4d, 0x43, 0x12, 0xa9, 0x74, 0x5b, 0xec, 0x4d, 0x29, 0x67,
	0x67, 0x9a, 0x9b, 0xf0, 0x62, 0xf6, 0xdb, 0xde, 0xfd, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x3f, 0xd5,
	0x84, 0x36, 0x67, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Queries a campaignChains by index.
	CampaignChains(ctx context.Context, in *QueryGetCampaignChainsRequest, opts ...grpc.CallOption) (*QueryGetCampaignChainsResponse, error)
	// Queries a list of campaignChains items.
	CampaignChainsAll(ctx context.Context, in *QueryAllCampaignChainsRequest, opts ...grpc.CallOption) (*QueryAllCampaignChainsResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) CampaignChains(ctx context.Context, in *QueryGetCampaignChainsRequest, opts ...grpc.CallOption) (*QueryGetCampaignChainsResponse, error) {
	out := new(QueryGetCampaignChainsResponse)
	err := c.cc.Invoke(ctx, "/tendermint.spn.campaign.Query/CampaignChains", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) CampaignChainsAll(ctx context.Context, in *QueryAllCampaignChainsRequest, opts ...grpc.CallOption) (*QueryAllCampaignChainsResponse, error) {
	out := new(QueryAllCampaignChainsResponse)
	err := c.cc.Invoke(ctx, "/tendermint.spn.campaign.Query/CampaignChainsAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Queries a campaignChains by index.
	CampaignChains(context.Context, *QueryGetCampaignChainsRequest) (*QueryGetCampaignChainsResponse, error)
	// Queries a list of campaignChains items.
	CampaignChainsAll(context.Context, *QueryAllCampaignChainsRequest) (*QueryAllCampaignChainsResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) CampaignChains(ctx context.Context, req *QueryGetCampaignChainsRequest) (*QueryGetCampaignChainsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CampaignChains not implemented")
}
func (*UnimplementedQueryServer) CampaignChainsAll(ctx context.Context, req *QueryAllCampaignChainsRequest) (*QueryAllCampaignChainsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CampaignChainsAll not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_CampaignChains_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetCampaignChainsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).CampaignChains(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tendermint.spn.campaign.Query/CampaignChains",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).CampaignChains(ctx, req.(*QueryGetCampaignChainsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_CampaignChainsAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllCampaignChainsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).CampaignChainsAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tendermint.spn.campaign.Query/CampaignChainsAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).CampaignChainsAll(ctx, req.(*QueryAllCampaignChainsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tendermint.spn.campaign.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CampaignChains",
			Handler:    _Query_CampaignChains_Handler,
		},
		{
			MethodName: "CampaignChainsAll",
			Handler:    _Query_CampaignChainsAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "campaign/query.proto",
}

func (m *QueryGetCampaignChainsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetCampaignChainsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetCampaignChainsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.CampaignID != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.CampaignID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *QueryGetCampaignChainsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetCampaignChainsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetCampaignChainsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.CampaignChains.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryAllCampaignChainsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllCampaignChainsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllCampaignChainsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryAllCampaignChainsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllCampaignChainsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllCampaignChainsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.CampaignChains) > 0 {
		for iNdEx := len(m.CampaignChains) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CampaignChains[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryGetCampaignChainsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CampaignID != 0 {
		n += 1 + sovQuery(uint64(m.CampaignID))
	}
	return n
}

func (m *QueryGetCampaignChainsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.CampaignChains.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryAllCampaignChainsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAllCampaignChainsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.CampaignChains) > 0 {
		for _, e := range m.CampaignChains {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryGetCampaignChainsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryGetCampaignChainsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetCampaignChainsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CampaignID", wireType)
			}
			m.CampaignID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CampaignID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryGetCampaignChainsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryGetCampaignChainsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetCampaignChainsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CampaignChains", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.CampaignChains.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryAllCampaignChainsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryAllCampaignChainsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllCampaignChainsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryAllCampaignChainsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryAllCampaignChainsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllCampaignChainsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CampaignChains", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CampaignChains = append(m.CampaignChains, CampaignChains{})
			if err := m.CampaignChains[len(m.CampaignChains)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
