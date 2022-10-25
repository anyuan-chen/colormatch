package color_service

import (
	"context"

	"github.com/anyuan-chen/colormatch/protos/colors/v1"
)

type ColorServiceServer struct {
	colors.UnimplementedPaletteMatchingServiceServer
}

func (s *ColorServiceServer) MatchColor(ctx context.Context, req *colors.MatchColorRequest) (*colors.MatchColorResponse, error) {
	return &colors.MatchColorResponse{Color: &colors.Color{HexCode: "hi"}}, nil
}
