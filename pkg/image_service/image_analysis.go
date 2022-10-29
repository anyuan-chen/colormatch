package image_service

import (
	"context"
	"image/color"

	colorsv1 "github.com/anyuan-chen/colormatch/gen/proto/go/colors/v1"
	sharedv1 "github.com/anyuan-chen/colormatch/gen/proto/go/shared/v1"
)

func GetDistanceMatrix(imgData *sharedv1.Image, palette *sharedv1.Palette, color_service colorsv1.PaletteMatchingServiceClient) ([][]color.RGBA, error) {
	img, err := BytesToImage(&imgData.ImageData)
	if err != nil {
		return nil, err
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
			closest_palette_color, err := color_service.MatchColor(context.Background(), request)
			if err != nil {
				return nil, err
			}
			closest_palette[i][j] = SharedColorToRGBA(closest_palette_color.Color)
		}
	}
	return closest_palette, nil
}
