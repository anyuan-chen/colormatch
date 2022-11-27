// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: spotify/v1/spotify.proto

package spotifyv1

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

// SpotifyImageColorMatchingServiceClient is the client API for SpotifyImageColorMatchingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SpotifyImageColorMatchingServiceClient interface {
	GetColorMetadataForSpotifyAsset(ctx context.Context, in *GetColorMetadataForSpotifyAssetRequest, opts ...grpc.CallOption) (*GetColorMetadataForSpotifyAssetResponse, error)
	PingTokenValidity(ctx context.Context, in *PingTokenValidityRequest, opts ...grpc.CallOption) (*PingTokenValidityResponse, error)
}

type spotifyImageColorMatchingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSpotifyImageColorMatchingServiceClient(cc grpc.ClientConnInterface) SpotifyImageColorMatchingServiceClient {
	return &spotifyImageColorMatchingServiceClient{cc}
}

func (c *spotifyImageColorMatchingServiceClient) GetColorMetadataForSpotifyAsset(ctx context.Context, in *GetColorMetadataForSpotifyAssetRequest, opts ...grpc.CallOption) (*GetColorMetadataForSpotifyAssetResponse, error) {
	out := new(GetColorMetadataForSpotifyAssetResponse)
	err := c.cc.Invoke(ctx, "/spotify.v1.SpotifyImageColorMatchingService/GetColorMetadataForSpotifyAsset", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spotifyImageColorMatchingServiceClient) PingTokenValidity(ctx context.Context, in *PingTokenValidityRequest, opts ...grpc.CallOption) (*PingTokenValidityResponse, error) {
	out := new(PingTokenValidityResponse)
	err := c.cc.Invoke(ctx, "/spotify.v1.SpotifyImageColorMatchingService/PingTokenValidity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SpotifyImageColorMatchingServiceServer is the server API for SpotifyImageColorMatchingService service.
// All implementations should embed UnimplementedSpotifyImageColorMatchingServiceServer
// for forward compatibility
type SpotifyImageColorMatchingServiceServer interface {
	GetColorMetadataForSpotifyAsset(context.Context, *GetColorMetadataForSpotifyAssetRequest) (*GetColorMetadataForSpotifyAssetResponse, error)
	PingTokenValidity(context.Context, *PingTokenValidityRequest) (*PingTokenValidityResponse, error)
}

// UnimplementedSpotifyImageColorMatchingServiceServer should be embedded to have forward compatible implementations.
type UnimplementedSpotifyImageColorMatchingServiceServer struct {
}

func (UnimplementedSpotifyImageColorMatchingServiceServer) GetColorMetadataForSpotifyAsset(context.Context, *GetColorMetadataForSpotifyAssetRequest) (*GetColorMetadataForSpotifyAssetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetColorMetadataForSpotifyAsset not implemented")
}
func (UnimplementedSpotifyImageColorMatchingServiceServer) PingTokenValidity(context.Context, *PingTokenValidityRequest) (*PingTokenValidityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PingTokenValidity not implemented")
}

// UnsafeSpotifyImageColorMatchingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SpotifyImageColorMatchingServiceServer will
// result in compilation errors.
type UnsafeSpotifyImageColorMatchingServiceServer interface {
	mustEmbedUnimplementedSpotifyImageColorMatchingServiceServer()
}

func RegisterSpotifyImageColorMatchingServiceServer(s grpc.ServiceRegistrar, srv SpotifyImageColorMatchingServiceServer) {
	s.RegisterService(&SpotifyImageColorMatchingService_ServiceDesc, srv)
}

func _SpotifyImageColorMatchingService_GetColorMetadataForSpotifyAsset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetColorMetadataForSpotifyAssetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpotifyImageColorMatchingServiceServer).GetColorMetadataForSpotifyAsset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spotify.v1.SpotifyImageColorMatchingService/GetColorMetadataForSpotifyAsset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpotifyImageColorMatchingServiceServer).GetColorMetadataForSpotifyAsset(ctx, req.(*GetColorMetadataForSpotifyAssetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpotifyImageColorMatchingService_PingTokenValidity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingTokenValidityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpotifyImageColorMatchingServiceServer).PingTokenValidity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spotify.v1.SpotifyImageColorMatchingService/PingTokenValidity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpotifyImageColorMatchingServiceServer).PingTokenValidity(ctx, req.(*PingTokenValidityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SpotifyImageColorMatchingService_ServiceDesc is the grpc.ServiceDesc for SpotifyImageColorMatchingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SpotifyImageColorMatchingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "spotify.v1.SpotifyImageColorMatchingService",
	HandlerType: (*SpotifyImageColorMatchingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetColorMetadataForSpotifyAsset",
			Handler:    _SpotifyImageColorMatchingService_GetColorMetadataForSpotifyAsset_Handler,
		},
		{
			MethodName: "PingTokenValidity",
			Handler:    _SpotifyImageColorMatchingService_PingTokenValidity_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spotify/v1/spotify.proto",
}
