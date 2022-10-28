package image_service

import (
	"image/color"

	sharedv1 "github.com/anyuan-chen/colormatch/gen/proto/go/shared/v1"
)

func SharedColorToRGBA(s *sharedv1.Color) color.RGBA {
	return color.RGBA{R: uint8(s.R), G: uint8(s.G), B: uint8(s.B), A: uint8(s.A)}
}
func RGBAToSharedColor(c color.RGBA) *sharedv1.Color {
	return &sharedv1.Color{R: int32(c.R), G: int32(c.G), B: int32(c.B), A: int32(c.A)}
}
