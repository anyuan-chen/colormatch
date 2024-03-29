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

// SpotifyAPIServiceClient is the client API for SpotifyAPIService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SpotifyAPIServiceClient interface {
	GetSearchResult(ctx context.Context, in *GetSearchResultRequest, opts ...grpc.CallOption) (*GetSearchResultResponse, error)
	GetTrackAudioFeatures(ctx context.Context, in *GetTrackAudioFeaturesRequest, opts ...grpc.CallOption) (*GetTrackAudioFeaturesResponse, error)
	GetTrackAudioAnalysis(ctx context.Context, in *GetTrackAudioAnalysisRequest, opts ...grpc.CallOption) (*GetTrackAudioAnalysisResponse, error)
	GetRecommendations(ctx context.Context, in *GetRecommendationsRequest, opts ...grpc.CallOption) (*GetRecommendationsResponse, error)
	GetColorMetadataForSpotifyAsset(ctx context.Context, in *GetColorMetadataForSpotifyAssetRequest, opts ...grpc.CallOption) (*GetColorMetadataForSpotifyAssetResponse, error)
	PingTokenValidity(ctx context.Context, in *PingTokenValidityRequest, opts ...grpc.CallOption) (*PingTokenValidityResponse, error)
}

type spotifyAPIServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSpotifyAPIServiceClient(cc grpc.ClientConnInterface) SpotifyAPIServiceClient {
	return &spotifyAPIServiceClient{cc}
}

func (c *spotifyAPIServiceClient) GetSearchResult(ctx context.Context, in *GetSearchResultRequest, opts ...grpc.CallOption) (*GetSearchResultResponse, error) {
	out := new(GetSearchResultResponse)
	err := c.cc.Invoke(ctx, "/spotify.v1.SpotifyAPIService/GetSearchResult", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spotifyAPIServiceClient) GetTrackAudioFeatures(ctx context.Context, in *GetTrackAudioFeaturesRequest, opts ...grpc.CallOption) (*GetTrackAudioFeaturesResponse, error) {
	out := new(GetTrackAudioFeaturesResponse)
	err := c.cc.Invoke(ctx, "/spotify.v1.SpotifyAPIService/GetTrackAudioFeatures", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spotifyAPIServiceClient) GetTrackAudioAnalysis(ctx context.Context, in *GetTrackAudioAnalysisRequest, opts ...grpc.CallOption) (*GetTrackAudioAnalysisResponse, error) {
	out := new(GetTrackAudioAnalysisResponse)
	err := c.cc.Invoke(ctx, "/spotify.v1.SpotifyAPIService/GetTrackAudioAnalysis", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spotifyAPIServiceClient) GetRecommendations(ctx context.Context, in *GetRecommendationsRequest, opts ...grpc.CallOption) (*GetRecommendationsResponse, error) {
	out := new(GetRecommendationsResponse)
	err := c.cc.Invoke(ctx, "/spotify.v1.SpotifyAPIService/GetRecommendations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spotifyAPIServiceClient) GetColorMetadataForSpotifyAsset(ctx context.Context, in *GetColorMetadataForSpotifyAssetRequest, opts ...grpc.CallOption) (*GetColorMetadataForSpotifyAssetResponse, error) {
	out := new(GetColorMetadataForSpotifyAssetResponse)
	err := c.cc.Invoke(ctx, "/spotify.v1.SpotifyAPIService/GetColorMetadataForSpotifyAsset", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spotifyAPIServiceClient) PingTokenValidity(ctx context.Context, in *PingTokenValidityRequest, opts ...grpc.CallOption) (*PingTokenValidityResponse, error) {
	out := new(PingTokenValidityResponse)
	err := c.cc.Invoke(ctx, "/spotify.v1.SpotifyAPIService/PingTokenValidity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SpotifyAPIServiceServer is the server API for SpotifyAPIService service.
// All implementations should embed UnimplementedSpotifyAPIServiceServer
// for forward compatibility
type SpotifyAPIServiceServer interface {
	GetSearchResult(context.Context, *GetSearchResultRequest) (*GetSearchResultResponse, error)
	GetTrackAudioFeatures(context.Context, *GetTrackAudioFeaturesRequest) (*GetTrackAudioFeaturesResponse, error)
	GetTrackAudioAnalysis(context.Context, *GetTrackAudioAnalysisRequest) (*GetTrackAudioAnalysisResponse, error)
	GetRecommendations(context.Context, *GetRecommendationsRequest) (*GetRecommendationsResponse, error)
	GetColorMetadataForSpotifyAsset(context.Context, *GetColorMetadataForSpotifyAssetRequest) (*GetColorMetadataForSpotifyAssetResponse, error)
	PingTokenValidity(context.Context, *PingTokenValidityRequest) (*PingTokenValidityResponse, error)
}

// UnimplementedSpotifyAPIServiceServer should be embedded to have forward compatible implementations.
type UnimplementedSpotifyAPIServiceServer struct {
}

func (UnimplementedSpotifyAPIServiceServer) GetSearchResult(context.Context, *GetSearchResultRequest) (*GetSearchResultResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSearchResult not implemented")
}
func (UnimplementedSpotifyAPIServiceServer) GetTrackAudioFeatures(context.Context, *GetTrackAudioFeaturesRequest) (*GetTrackAudioFeaturesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTrackAudioFeatures not implemented")
}
func (UnimplementedSpotifyAPIServiceServer) GetTrackAudioAnalysis(context.Context, *GetTrackAudioAnalysisRequest) (*GetTrackAudioAnalysisResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTrackAudioAnalysis not implemented")
}
func (UnimplementedSpotifyAPIServiceServer) GetRecommendations(context.Context, *GetRecommendationsRequest) (*GetRecommendationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRecommendations not implemented")
}
func (UnimplementedSpotifyAPIServiceServer) GetColorMetadataForSpotifyAsset(context.Context, *GetColorMetadataForSpotifyAssetRequest) (*GetColorMetadataForSpotifyAssetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetColorMetadataForSpotifyAsset not implemented")
}
func (UnimplementedSpotifyAPIServiceServer) PingTokenValidity(context.Context, *PingTokenValidityRequest) (*PingTokenValidityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PingTokenValidity not implemented")
}

// UnsafeSpotifyAPIServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SpotifyAPIServiceServer will
// result in compilation errors.
type UnsafeSpotifyAPIServiceServer interface {
	mustEmbedUnimplementedSpotifyAPIServiceServer()
}

func RegisterSpotifyAPIServiceServer(s grpc.ServiceRegistrar, srv SpotifyAPIServiceServer) {
	s.RegisterService(&SpotifyAPIService_ServiceDesc, srv)
}

func _SpotifyAPIService_GetSearchResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSearchResultRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpotifyAPIServiceServer).GetSearchResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spotify.v1.SpotifyAPIService/GetSearchResult",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpotifyAPIServiceServer).GetSearchResult(ctx, req.(*GetSearchResultRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpotifyAPIService_GetTrackAudioFeatures_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTrackAudioFeaturesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpotifyAPIServiceServer).GetTrackAudioFeatures(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spotify.v1.SpotifyAPIService/GetTrackAudioFeatures",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpotifyAPIServiceServer).GetTrackAudioFeatures(ctx, req.(*GetTrackAudioFeaturesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpotifyAPIService_GetTrackAudioAnalysis_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTrackAudioAnalysisRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpotifyAPIServiceServer).GetTrackAudioAnalysis(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spotify.v1.SpotifyAPIService/GetTrackAudioAnalysis",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpotifyAPIServiceServer).GetTrackAudioAnalysis(ctx, req.(*GetTrackAudioAnalysisRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpotifyAPIService_GetRecommendations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRecommendationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpotifyAPIServiceServer).GetRecommendations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spotify.v1.SpotifyAPIService/GetRecommendations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpotifyAPIServiceServer).GetRecommendations(ctx, req.(*GetRecommendationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpotifyAPIService_GetColorMetadataForSpotifyAsset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetColorMetadataForSpotifyAssetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpotifyAPIServiceServer).GetColorMetadataForSpotifyAsset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spotify.v1.SpotifyAPIService/GetColorMetadataForSpotifyAsset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpotifyAPIServiceServer).GetColorMetadataForSpotifyAsset(ctx, req.(*GetColorMetadataForSpotifyAssetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpotifyAPIService_PingTokenValidity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingTokenValidityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpotifyAPIServiceServer).PingTokenValidity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/spotify.v1.SpotifyAPIService/PingTokenValidity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpotifyAPIServiceServer).PingTokenValidity(ctx, req.(*PingTokenValidityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SpotifyAPIService_ServiceDesc is the grpc.ServiceDesc for SpotifyAPIService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SpotifyAPIService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "spotify.v1.SpotifyAPIService",
	HandlerType: (*SpotifyAPIServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetSearchResult",
			Handler:    _SpotifyAPIService_GetSearchResult_Handler,
		},
		{
			MethodName: "GetTrackAudioFeatures",
			Handler:    _SpotifyAPIService_GetTrackAudioFeatures_Handler,
		},
		{
			MethodName: "GetTrackAudioAnalysis",
			Handler:    _SpotifyAPIService_GetTrackAudioAnalysis_Handler,
		},
		{
			MethodName: "GetRecommendations",
			Handler:    _SpotifyAPIService_GetRecommendations_Handler,
		},
		{
			MethodName: "GetColorMetadataForSpotifyAsset",
			Handler:    _SpotifyAPIService_GetColorMetadataForSpotifyAsset_Handler,
		},
		{
			MethodName: "PingTokenValidity",
			Handler:    _SpotifyAPIService_PingTokenValidity_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "spotify/v1/spotify.proto",
}
