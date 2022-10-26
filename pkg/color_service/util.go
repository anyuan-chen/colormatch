package color_service

import (
	"image/color"
	"math"
)

func HexToRgb(hex string) (color.RGBA, error) {
	return color.RGBA{}, nil
}

func RgbToCIE76(rgba color.RGBA) (CIE76, error) {
	return CIE76{}, nil
}

func CIE76ColorDifference(color1 CIE76, color2 CIE76) float64 {
	aDelta := math.Pow((color1.a - color2.a), 2)
	bDelta := math.Pow((color1.b - color2.b), 2)
	lDelta := math.Pow((color1.l - color2.l), 2)
	return math.Sqrt(aDelta + bDelta + lDelta)
}
