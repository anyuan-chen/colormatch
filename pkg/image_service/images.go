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

// RGBToHex returns a six-digit hex code given an RGB color
func RGBToHex(c color.Color) string {
	r, g, b, _ := c.RGBA()
	return fmt.Sprintf("#%02x%02x%02x", r, g, b)
}

// GetBackgroundColor returns a sharedv1.Color object representing the most common
// background color within a given palette. If the backgroundcolor palette given does
// not have any colors that are the closest in proximity to colors in the images,
// the functions returns white.
func (iss *PaletteMatchingServer) GetBackgroundColor(ctx context.Context, req *imagev1.GetBackgroundColorRequest) (*imagev1.GetBackgroundColorResponse, error) {
	error_response := &imagev1.GetBackgroundColorResponse{Color: &sharedv1.Color{}}
	imgData := req.Image
	palette := req.Palette
	backgroundPalette := req.BackgroundColors
	closest_palette, err := GetFrequencyArray(imgData, palette, iss.Color_Service)
	closest_background_palette := make(map[*sharedv1.Color]int32)
	//filters out all of the non-background colors
	for _, backgroundColor := range backgroundPalette.Color {
		closest_background_palette[backgroundColor] = closest_palette[backgroundColor]
	}
	//default values set for if there are no background colors
	var maxKey *sharedv1.Color = &sharedv1.Color{R: 255, G: 255, B: 255, A: 1}
	var maxVal int32 = math.MinInt32
	//extracts most frequent background color
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
