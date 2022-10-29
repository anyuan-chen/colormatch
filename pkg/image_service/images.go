package image_service

import (
	"context"
	"fmt"
	"image/color"

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
	default_context := context.Background()
	error_response := &imagev1.GetBackgroundColorResponse{Color: &sharedv1.Color{}}
	imgData := req.Image
	palette := req.Palette
	img, err := BytesToImage(&imgData.ImageData)
	if err != nil {
		return error_response, err
	}
	closest_palette := make([][]color.RGBA, imgData.Height)
	for i := range closest_palette {
		closest_palette[i] = make([]color.RGBA, imgData.Width)
	}
	for i := 0; i < int(imgData.Height); i++ {
		for j := 0; j < int(imgData.Width); j++ {
			current_pixel_color := img.At(i, j)
			R, G, B, A := current_pixel_color.RGBA()
			color := &sharedv1.Color{R: int32(R), G: int32(G), B: int32(B), A: int32(A)}
			request := &colorsv1.MatchColorRequest{Color: color, Palette: palette}
			closest_palette_color, err := iss.Color_Service.MatchColor(default_context, request)
			if err != nil {
				return nil, err
			}
			closest_palette[i][j] = SharedColorToRGBA(closest_palette_color.Color)
		}
	}

	if err != nil {
		return error_response, err
	}

	var backgroundColor *sharedv1.Color
	return &imagev1.GetBackgroundColorResponse{Color: backgroundColor}, nil
}

func (iss *PaletteMatchingServer) GetHighlightColor(ctx context.Context, req *imagev1.GetHighlightColorRequest) (*imagev1.GetHighlightColorResponse, error) {
	return &imagev1.GetHighlightColorResponse{Color: nil}, nil
}
