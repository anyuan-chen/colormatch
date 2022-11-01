package image_service

import (
	"context"
	"fmt"
	"image/color"
	"math"

	colorsv1 "github.com/anyuan-chen/colormatch/gen/proto/go/colors/v1"
	imagev1 "github.com/anyuan-chen/colormatch/gen/proto/go/images/v1"
	sharedv1 "github.com/anyuan-chen/colormatch/gen/proto/go/shared/v1"
)

type PaletteMatchingServer struct {
	Color_Service colorsv1.PaletteMatchingServiceClient
	imagev1.UnimplementedPaletteMatchingServiceServer
}

func RGBToHex(c color.Color) string {
	r, g, b, _ := c.RGBA()
	return fmt.Sprintf("#%02x%02x%02x", r, g, b)
}
func (iss *PaletteMatchingServer) GetBackgroundColor(ctx context.Context, req *imagev1.GetBackgroundColorRequest) (*imagev1.GetBackgroundColorResponse, error) {
	error_response := &imagev1.GetBackgroundColorResponse{Color: &sharedv1.Color{}}
	imgData := req.Image
	palette := req.Palette
	backgroundPalette := req.BackgroundColors
	closest_palette, err := GetFrequencyArray(imgData, palette, iss.Color_Service)
	closest_background_palette := make(map[*sharedv1.Color]int32)
	for _, backgroundColor := range backgroundPalette.Color {
		closest_background_palette[backgroundColor] = closest_palette[backgroundColor]
	}
	var maxKey *sharedv1.Color = &sharedv1.Color{R: 0, G: 0, B: 0, A: 1}
	var maxVal int32 = math.MinInt32
	for key, val := range closest_background_palette {
		if val > maxVal {
			maxVal = val
			maxKey = key
		}
	}
	if err != nil {
		return error_response, err
	}
	return &imagev1.GetBackgroundColorResponse{Color: maxKey}, nil
}

func (iss *PaletteMatchingServer) GetHighlightColor(ctx context.Context, req *imagev1.GetHighlightColorRequest) (*imagev1.GetHighlightColorResponse, error) {
	return &imagev1.GetHighlightColorResponse{Color: nil}, nil
}
