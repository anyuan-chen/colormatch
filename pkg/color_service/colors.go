package color_service

import (
	"context"
	"errors"
	"math"

	"github.com/anyuan-chen/colormatch/protos/colors/v1"
)

type ColorServiceServer struct {
	colors.UnimplementedPaletteMatchingServiceServer
}

func (s *ColorServiceServer) MatchColor(ctx context.Context, req *colors.MatchColorRequest) (*colors.MatchColorResponse, error) {
	user_defined_color := req.GetColor()
	palette := req.GetPalette().Color

	dist := math.MaxFloat64
	var closestColor *colors.Color
	for _, color := range palette {
		cur_dist, err := HexColorDifference(color.HexCode, user_defined_color.HexCode)
		if err != nil {
			return nil, errors.New("")
		}
		if cur_dist > dist {
			dist = cur_dist
			closestColor = color
		}
	}
	return &colors.MatchColorResponse{Color: closestColor}, nil
}
