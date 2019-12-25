// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/proto/api/api.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	types "github.com/iotexproject/high-table/proto/golang/types"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type GetDelegateRequest struct {
	DelegateID           uint64   `protobuf:"varint,1,opt,name=delegateID,proto3" json:"delegateID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDelegateRequest) Reset()         { *m = GetDelegateRequest{} }
func (m *GetDelegateRequest) String() string { return proto.CompactTextString(m) }
func (*GetDelegateRequest) ProtoMessage()    {}
func (*GetDelegateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_15eae2866f8410ca, []int{0}
}

func (m *GetDelegateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDelegateRequest.Unmarshal(m, b)
}
func (m *GetDelegateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDelegateRequest.Marshal(b, m, deterministic)
}
func (m *GetDelegateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDelegateRequest.Merge(m, src)
}
func (m *GetDelegateRequest) XXX_Size() int {
	return xxx_messageInfo_GetDelegateRequest.Size(m)
}
func (m *GetDelegateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDelegateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetDelegateRequest proto.InternalMessageInfo

func (m *GetDelegateRequest) GetDelegateID() uint64 {
	if m != nil {
		return m.DelegateID
	}
	return 0
}

type GetDelegateResponse struct {
	Delegate             *types.Delegate `protobuf:"bytes,1,opt,name=delegate,proto3" json:"delegate,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *GetDelegateResponse) Reset()         { *m = GetDelegateResponse{} }
func (m *GetDelegateResponse) String() string { return proto.CompactTextString(m) }
func (*GetDelegateResponse) ProtoMessage()    {}
func (*GetDelegateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_15eae2866f8410ca, []int{1}
}

func (m *GetDelegateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDelegateResponse.Unmarshal(m, b)
}
func (m *GetDelegateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDelegateResponse.Marshal(b, m, deterministic)
}
func (m *GetDelegateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDelegateResponse.Merge(m, src)
}
func (m *GetDelegateResponse) XXX_Size() int {
	return xxx_messageInfo_GetDelegateResponse.Size(m)
}
func (m *GetDelegateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDelegateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetDelegateResponse proto.InternalMessageInfo

func (m *GetDelegateResponse) GetDelegate() *types.Delegate {
	if m != nil {
		return m.Delegate
	}
	return nil
}

type UpdateDelegateRequest struct {
	Delegate             *types.Delegate `protobuf:"bytes,1,opt,name=delegate,proto3" json:"delegate,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *UpdateDelegateRequest) Reset()         { *m = UpdateDelegateRequest{} }
func (m *UpdateDelegateRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateDelegateRequest) ProtoMessage()    {}
func (*UpdateDelegateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_15eae2866f8410ca, []int{2}
}

func (m *UpdateDelegateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateDelegateRequest.Unmarshal(m, b)
}
func (m *UpdateDelegateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateDelegateRequest.Marshal(b, m, deterministic)
}
func (m *UpdateDelegateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateDelegateRequest.Merge(m, src)
}
func (m *UpdateDelegateRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateDelegateRequest.Size(m)
}
func (m *UpdateDelegateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateDelegateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateDelegateRequest proto.InternalMessageInfo

func (m *UpdateDelegateRequest) GetDelegate() *types.Delegate {
	if m != nil {
		return m.Delegate
	}
	return nil
}

func init() {
	proto.RegisterType((*GetDelegateRequest)(nil), "api.GetDelegateRequest")
	proto.RegisterType((*GetDelegateResponse)(nil), "api.GetDelegateResponse")
	proto.RegisterType((*UpdateDelegateRequest)(nil), "api.UpdateDelegateRequest")
}

func init() { proto.RegisterFile("proto/proto/api/api.proto", fileDescriptor_15eae2866f8410ca) }

var fileDescriptor_15eae2866f8410ca = []byte{
	// 280 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0xd1, 0x4a, 0xc3, 0x30,
	0x14, 0x86, 0x57, 0x14, 0x91, 0x33, 0x50, 0x88, 0xa8, 0xb3, 0x82, 0x8c, 0x5e, 0x09, 0x62, 0x02,
	0x9b, 0x3e, 0x80, 0xa5, 0x2a, 0xbb, 0x93, 0x89, 0x37, 0xde, 0xa5, 0xdd, 0x31, 0x8d, 0x74, 0x4b,
	0x6c, 0x4f, 0xc5, 0xbd, 0x89, 0x8f, 0x2b, 0x4d, 0xec, 0xe8, 0x9c, 0x37, 0xbb, 0x48, 0x20, 0x27,
	0xdf, 0x9f, 0xf3, 0x9f, 0xfc, 0x70, 0x66, 0x4b, 0x43, 0x46, 0xf8, 0x5d, 0x5a, 0xdd, 0x2c, 0xee,
	0x4e, 0x6c, 0x47, 0x5a, 0x1d, 0x9e, 0x2b, 0x63, 0x54, 0x81, 0x1e, 0x48, 0xeb, 0x37, 0x81, 0x73,
	0x4b, 0x4b, 0x4f, 0x84, 0xc3, 0xae, 0x98, 0x96, 0x16, 0x2b, 0x31, 0xc3, 0x02, 0x95, 0x24, 0xf4,
	0x44, 0x74, 0x03, 0xec, 0x11, 0x29, 0xf9, 0x2d, 0x4e, 0xf1, 0xa3, 0xc6, 0x8a, 0xd8, 0x05, 0x40,
	0xcb, 0x4d, 0x92, 0x41, 0x30, 0x0c, 0x2e, 0x77, 0xa7, 0x9d, 0x4a, 0x14, 0xc3, 0xd1, 0x9a, 0xaa,
	0xb2, 0x66, 0x51, 0x21, 0xbb, 0x82, 0xfd, 0x16, 0x72, 0xa2, 0xfe, 0xe8, 0x90, 0xbb, 0xae, 0x7c,
	0x85, 0xae, 0x80, 0x28, 0x81, 0xe3, 0x17, 0x3b, 0x93, 0x84, 0x7f, 0x9b, 0x6f, 0xf3, 0xca, 0xe8,
	0x3b, 0x00, 0xb8, 0x7b, 0x9a, 0x3c, 0x63, 0xf9, 0xa9, 0x33, 0x64, 0x31, 0xf4, 0x3b, 0xc6, 0xd8,
	0x29, 0x6f, 0x7e, 0x6b, 0x73, 0xc0, 0x70, 0xb0, 0x79, 0xe1, 0x67, 0x88, 0x7a, 0xec, 0x01, 0x0e,
	0xd6, 0x8d, 0xb1, 0xd0, 0xd1, 0xff, 0xba, 0x0d, 0x4f, 0xb8, 0x0f, 0x80, 0xb7, 0x01, 0xf0, 0xfb,
	0x26, 0x80, 0xa8, 0x17, 0xdf, 0xbe, 0x8e, 0x95, 0xa6, 0xbc, 0x4e, 0x79, 0x66, 0xe6, 0x42, 0x1b,
	0xc2, 0x2f, 0x5b, 0x9a, 0x77, 0xcc, 0x48, 0xe4, 0x5a, 0xe5, 0xd7, 0x24, 0xd3, 0x36, 0x37, 0xa1,
	0x4c, 0x21, 0x17, 0xaa, 0xc9, 0x36, 0xdd, 0x73, 0x95, 0xf1, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x92, 0x82, 0x36, 0x52, 0xf9, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// APIServiceClient is the client API for APIService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type APIServiceClient interface {
	GetDelegate(ctx context.Context, in *GetDelegateRequest, opts ...grpc.CallOption) (*GetDelegateResponse, error)
	UpdateDelegate(ctx context.Context, in *UpdateDelegateRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type aPIServiceClient struct {
	cc *grpc.ClientConn
}

func NewAPIServiceClient(cc *grpc.ClientConn) APIServiceClient {
	return &aPIServiceClient{cc}
}

func (c *aPIServiceClient) GetDelegate(ctx context.Context, in *GetDelegateRequest, opts ...grpc.CallOption) (*GetDelegateResponse, error) {
	out := new(GetDelegateResponse)
	err := c.cc.Invoke(ctx, "/api.APIService/GetDelegate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIServiceClient) UpdateDelegate(ctx context.Context, in *UpdateDelegateRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/api.APIService/UpdateDelegate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// APIServiceServer is the server API for APIService service.
type APIServiceServer interface {
	GetDelegate(context.Context, *GetDelegateRequest) (*GetDelegateResponse, error)
	UpdateDelegate(context.Context, *UpdateDelegateRequest) (*empty.Empty, error)
}

// UnimplementedAPIServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAPIServiceServer struct {
}

func (*UnimplementedAPIServiceServer) GetDelegate(ctx context.Context, req *GetDelegateRequest) (*GetDelegateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDelegate not implemented")
}
func (*UnimplementedAPIServiceServer) UpdateDelegate(ctx context.Context, req *UpdateDelegateRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDelegate not implemented")
}

func RegisterAPIServiceServer(s *grpc.Server, srv APIServiceServer) {
	s.RegisterService(&_APIService_serviceDesc, srv)
}

func _APIService_GetDelegate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDelegateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServiceServer).GetDelegate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.APIService/GetDelegate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServiceServer).GetDelegate(ctx, req.(*GetDelegateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _APIService_UpdateDelegate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDelegateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServiceServer).UpdateDelegate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.APIService/UpdateDelegate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServiceServer).UpdateDelegate(ctx, req.(*UpdateDelegateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _APIService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.APIService",
	HandlerType: (*APIServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDelegate",
			Handler:    _APIService_GetDelegate_Handler,
		},
		{
			MethodName: "UpdateDelegate",
			Handler:    _APIService_UpdateDelegate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/proto/api/api.proto",
}
