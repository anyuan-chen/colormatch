// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: images/v1/images.proto

package imagesv1

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

// PaletteMatchingServiceClient is the client API for PaletteMatchingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PaletteMatchingServiceClient interface {
	GetBackgroundColor(ctx context.Context, in *GetBackgroundColorRequest, opts ...grpc.CallOption) (*GetBackgroundColorResponse, error)
	GetHighlightColor(ctx context.Context, in *GetHighlightColorRequest, opts ...grpc.CallOption) (*GetHighlightColorResponse, error)
}

type paletteMatchingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPaletteMatchingServiceClient(cc grpc.ClientConnInterface) PaletteMatchingServiceClient {
	return &paletteMatchingServiceClient{cc}
}

func (c *paletteMatchingServiceClient) GetBackgroundColor(ctx context.Context, in *GetBackgroundColorRequest, opts ...grpc.CallOption) (*GetBackgroundColorResponse, error) {
	out := new(GetBackgroundColorResponse)
	err := c.cc.Invoke(ctx, "/images.v1.PaletteMatchingService/GetBackgroundColor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *paletteMatchingServiceClient) GetHighlightColor(ctx context.Context, in *GetHighlightColorRequest, opts ...grpc.CallOption) (*GetHighlightColorResponse, error) {
	out := new(GetHighlightColorResponse)
	err := c.cc.Invoke(ctx, "/images.v1.PaletteMatchingService/GetHighlightColor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PaletteMatchingServiceServer is the server API for PaletteMatchingService service.
// All implementations should embed UnimplementedPaletteMatchingServiceServer
// for forward compatibility
type PaletteMatchingServiceServer interface {
	GetBackgroundColor(context.Context, *GetBackgroundColorRequest) (*GetBackgroundColorResponse, error)
	GetHighlightColor(context.Context, *GetHighlightColorRequest) (*GetHighlightColorResponse, error)
}

// UnimplementedPaletteMatchingServiceServer should be embedded to have forward compatible implementations.
type UnimplementedPaletteMatchingServiceServer struct {
}

func (UnimplementedPaletteMatchingServiceServer) GetBackgroundColor(context.Context, *GetBackgroundColorRequest) (*GetBackgroundColorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBackgroundColor not implemented")
}
func (UnimplementedPaletteMatchingServiceServer) GetHighlightColor(context.Context, *GetHighlightColorRequest) (*GetHighlightColorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHighlightColor not implemented")
}

// UnsafePaletteMatchingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PaletteMatchingServiceServer will
// result in compilation errors.
type UnsafePaletteMatchingServiceServer interface {
	mustEmbedUnimplementedPaletteMatchingServiceServer()
}

func RegisterPaletteMatchingServiceServer(s grpc.ServiceRegistrar, srv PaletteMatchingServiceServer) {
	s.RegisterService(&PaletteMatchingService_ServiceDesc, srv)
}

func _PaletteMatchingService_GetBackgroundColor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBackgroundColorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaletteMatchingServiceServer).GetBackgroundColor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/images.v1.PaletteMatchingService/GetBackgroundColor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaletteMatchingServiceServer).GetBackgroundColor(ctx, req.(*GetBackgroundColorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PaletteMatchingService_GetHighlightColor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHighlightColorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PaletteMatchingServiceServer).GetHighlightColor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/images.v1.PaletteMatchingService/GetHighlightColor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PaletteMatchingServiceServer).GetHighlightColor(ctx, req.(*GetHighlightColorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PaletteMatchingService_ServiceDesc is the grpc.ServiceDesc for PaletteMatchingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PaletteMatchingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "images.v1.PaletteMatchingService",
	HandlerType: (*PaletteMatchingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBackgroundColor",
			Handler:    _PaletteMatchingService_GetBackgroundColor_Handler,
		},
		{
			MethodName: "GetHighlightColor",
			Handler:    _PaletteMatchingService_GetHighlightColor_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "images/v1/images.proto",
}
