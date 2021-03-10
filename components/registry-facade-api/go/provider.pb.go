// Code generated by protoc-gen-go. DO NOT EDIT.
// source: provider.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type GetImageSpecRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetImageSpecRequest) Reset()         { *m = GetImageSpecRequest{} }
func (m *GetImageSpecRequest) String() string { return proto.CompactTextString(m) }
func (*GetImageSpecRequest) ProtoMessage()    {}
func (*GetImageSpecRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6a9f3c02af3d1c8, []int{0}
}

func (m *GetImageSpecRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetImageSpecRequest.Unmarshal(m, b)
}
func (m *GetImageSpecRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetImageSpecRequest.Marshal(b, m, deterministic)
}
func (m *GetImageSpecRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetImageSpecRequest.Merge(m, src)
}
func (m *GetImageSpecRequest) XXX_Size() int {
	return xxx_messageInfo_GetImageSpecRequest.Size(m)
}
func (m *GetImageSpecRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetImageSpecRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetImageSpecRequest proto.InternalMessageInfo

func (m *GetImageSpecRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetImageSpecResponse struct {
	Spec                 *ImageSpec `protobuf:"bytes,1,opt,name=spec,proto3" json:"spec,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *GetImageSpecResponse) Reset()         { *m = GetImageSpecResponse{} }
func (m *GetImageSpecResponse) String() string { return proto.CompactTextString(m) }
func (*GetImageSpecResponse) ProtoMessage()    {}
func (*GetImageSpecResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6a9f3c02af3d1c8, []int{1}
}

func (m *GetImageSpecResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetImageSpecResponse.Unmarshal(m, b)
}
func (m *GetImageSpecResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetImageSpecResponse.Marshal(b, m, deterministic)
}
func (m *GetImageSpecResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetImageSpecResponse.Merge(m, src)
}
func (m *GetImageSpecResponse) XXX_Size() int {
	return xxx_messageInfo_GetImageSpecResponse.Size(m)
}
func (m *GetImageSpecResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetImageSpecResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetImageSpecResponse proto.InternalMessageInfo

func (m *GetImageSpecResponse) GetSpec() *ImageSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func init() {
	proto.RegisterType((*GetImageSpecRequest)(nil), "registryfacade.GetImageSpecRequest")
	proto.RegisterType((*GetImageSpecResponse)(nil), "registryfacade.GetImageSpecResponse")
}

func init() {
	proto.RegisterFile("provider.proto", fileDescriptor_c6a9f3c02af3d1c8)
}

var fileDescriptor_c6a9f3c02af3d1c8 = []byte{
	// 180 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0x28, 0xca, 0x2f,
	0xcb, 0x4c, 0x49, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x2b, 0x4a, 0x4d, 0xcf,
	0x2c, 0x2e, 0x29, 0xaa, 0x4c, 0x4b, 0x4c, 0x4e, 0x4c, 0x49, 0x95, 0xe2, 0xcf, 0xcc, 0x4d, 0x4c,
	0x4f, 0x2d, 0x2e, 0x48, 0x4d, 0x86, 0x28, 0x50, 0x52, 0xe5, 0x12, 0x76, 0x4f, 0x2d, 0xf1, 0x04,
	0x89, 0x06, 0x17, 0xa4, 0x26, 0x07, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0xf1, 0x71, 0x31,
	0x65, 0xa6, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x31, 0x65, 0xa6, 0x28, 0xb9, 0x72, 0x89,
	0xa0, 0x2a, 0x2b, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x15, 0xd2, 0xe5, 0x62, 0x01, 0x19, 0x06, 0x56,
	0xc9, 0x6d, 0x24, 0xa9, 0x87, 0x6a, 0x9d, 0x1e, 0x42, 0x03, 0x58, 0x99, 0x51, 0x36, 0x17, 0x0f,
	0x88, 0x17, 0x00, 0x75, 0xa4, 0x50, 0x34, 0x17, 0x0f, 0xb2, 0xb1, 0x42, 0xca, 0xe8, 0x06, 0x60,
	0x71, 0x9b, 0x94, 0x0a, 0x7e, 0x45, 0x10, 0x97, 0x29, 0x31, 0x38, 0xb1, 0x46, 0x31, 0x27, 0x16,
	0x64, 0x26, 0xb1, 0x81, 0x3d, 0x6a, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0xd4, 0xed, 0x92, 0xfd,
	0x1b, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SpecProviderClient is the client API for SpecProvider service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SpecProviderClient interface {
	// GetImageSpec provides the image spec for a particular ID. What the ID referes to is specific to
	// the spec provider. For example, in case of ws-manager providing the spec, the ID is a
	// workspace instance ID.
	GetImageSpec(ctx context.Context, in *GetImageSpecRequest, opts ...grpc.CallOption) (*GetImageSpecResponse, error)
}

type specProviderClient struct {
	cc grpc.ClientConnInterface
}

func NewSpecProviderClient(cc grpc.ClientConnInterface) SpecProviderClient {
	return &specProviderClient{cc}
}

func (c *specProviderClient) GetImageSpec(ctx context.Context, in *GetImageSpecRequest, opts ...grpc.CallOption) (*GetImageSpecResponse, error) {
	out := new(GetImageSpecResponse)
	err := c.cc.Invoke(ctx, "/registryfacade.SpecProvider/GetImageSpec", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SpecProviderServer is the server API for SpecProvider service.
type SpecProviderServer interface {
	// GetImageSpec provides the image spec for a particular ID. What the ID referes to is specific to
	// the spec provider. For example, in case of ws-manager providing the spec, the ID is a
	// workspace instance ID.
	GetImageSpec(context.Context, *GetImageSpecRequest) (*GetImageSpecResponse, error)
}

// UnimplementedSpecProviderServer can be embedded to have forward compatible implementations.
type UnimplementedSpecProviderServer struct {
}

func (*UnimplementedSpecProviderServer) GetImageSpec(ctx context.Context, req *GetImageSpecRequest) (*GetImageSpecResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetImageSpec not implemented")
}

func RegisterSpecProviderServer(s *grpc.Server, srv SpecProviderServer) {
	s.RegisterService(&_SpecProvider_serviceDesc, srv)
}

func _SpecProvider_GetImageSpec_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetImageSpecRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpecProviderServer).GetImageSpec(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/registryfacade.SpecProvider/GetImageSpec",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpecProviderServer).GetImageSpec(ctx, req.(*GetImageSpecRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SpecProvider_serviceDesc = grpc.ServiceDesc{
	ServiceName: "registryfacade.SpecProvider",
	HandlerType: (*SpecProviderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetImageSpec",
			Handler:    _SpecProvider_GetImageSpec_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "provider.proto",
}
