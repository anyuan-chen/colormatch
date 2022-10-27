package color_service

import (
	"context"
	"errors"
	"math"

	colorsv1 "github.com/anyuan-chen/colormatch/gen/proto/go/colors/v1"
	sharedv1 "github.com/anyuan-chen/colormatch/gen/proto/go/shared/v1"
)

type ColorServiceServer struct {
	colorsv1.UnimplementedPaletteMatchingServiceServer
}

func (s *ColorServiceServer) MatchColor(ctx context.Context, req *colorsv1.MatchColorRequest) (*colorsv1.MatchColorResponse, error) {
	user_defined_color := req.GetColor()
	palette := req.GetPalette().Color

	dist := math.MaxFloat64
	var closestColor *sharedv1.Color
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
	return &colorsv1.MatchColorResponse{Color: closestColor}, nil
}
