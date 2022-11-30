// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: other_service.proto

package other_service

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// OtherServiceClient is the client API for OtherService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OtherServiceClient interface {
	CreateOther(ctx context.Context, in *CreateOtherRequest, opts ...grpc.CallOption) (*OtherID, error)
	UpdateOther(ctx context.Context, in *UpdateOtherRequest, opts ...grpc.CallOption) (*OtherID, error)
}

type otherServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOtherServiceClient(cc grpc.ClientConnInterface) OtherServiceClient {
	return &otherServiceClient{cc}
}

func (c *otherServiceClient) CreateOther(ctx context.Context, in *CreateOtherRequest, opts ...grpc.CallOption) (*OtherID, error) {
	out := new(OtherID)
	err := c.cc.Invoke(ctx, "/other_service.OtherService/CreateOther", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *otherServiceClient) UpdateOther(ctx context.Context, in *UpdateOtherRequest, opts ...grpc.CallOption) (*OtherID, error) {
	out := new(OtherID)
	err := c.cc.Invoke(ctx, "/other_service.OtherService/UpdateOther", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OtherServiceServer is the server API for OtherService service.
// All implementations must embed UnimplementedOtherServiceServer
// for forward compatibility
type OtherServiceServer interface {
	CreateOther(context.Context, *CreateOtherRequest) (*OtherID, error)
	UpdateOther(context.Context, *UpdateOtherRequest) (*OtherID, error)
	mustEmbedUnimplementedOtherServiceServer()
}

// UnimplementedOtherServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOtherServiceServer struct {
}

func (UnimplementedOtherServiceServer) CreateOther(context.Context, *CreateOtherRequest) (*OtherID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOther not implemented")
}
func (UnimplementedOtherServiceServer) UpdateOther(context.Context, *UpdateOtherRequest) (*OtherID, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOther not implemented")
}
func (UnimplementedOtherServiceServer) mustEmbedUnimplementedOtherServiceServer() {}

// UnsafeOtherServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OtherServiceServer will
// result in compilation errors.
type UnsafeOtherServiceServer interface {
	mustEmbedUnimplementedOtherServiceServer()
}

func RegisterOtherServiceServer(s grpc.ServiceRegistrar, srv OtherServiceServer) {
	s.RegisterService(&OtherService_ServiceDesc, srv)
}

func _OtherService_CreateOther_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOtherRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OtherServiceServer).CreateOther(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/other_service.OtherService/CreateOther",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OtherServiceServer).CreateOther(ctx, req.(*CreateOtherRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OtherService_UpdateOther_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateOtherRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OtherServiceServer).UpdateOther(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/other_service.OtherService/UpdateOther",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OtherServiceServer).UpdateOther(ctx, req.(*UpdateOtherRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OtherService_ServiceDesc is the grpc.ServiceDesc for OtherService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OtherService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "other_service.OtherService",
	HandlerType: (*OtherServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOther",
			Handler:    _OtherService_CreateOther_Handler,
		},
		{
			MethodName: "UpdateOther",
			Handler:    _OtherService_UpdateOther_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "other_service.proto",
}
