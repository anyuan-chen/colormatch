package color_service

import (
	"context"

	"github.com/anyuan-chen/colormatch/protos/colors/v1"
)

type ColorServiceServer struct {
	colors.UnimplementedPaletteMatchingServiceServer
}
type XYZ struct {
	x float64
	y float64
	z float64
}
type CIE76 struct {
	l float64
	a float64
	b float64
}

func (s *ColorServiceServer) MatchColor(ctx context.Context, req *colors.MatchColorRequest) (*colors.MatchColorResponse, error) {
	user_defined_color := req.GetColor()
	palette := req.GetPalette().Color

	for _, color := range palette {
		
	}

	//hex -> rgb
	//rgb -> l*a*b*
	//l*a*b* -> euclidian distance between l, a, and b
	//shortest euclidian distance "wins"

	return &colors.MatchColorResponse{Color: &colors.Color{HexCode: "hi"}}, nil
}
