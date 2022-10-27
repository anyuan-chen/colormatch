package color_service

import (
	"errors"
	"image/color"

	"github.com/jkl1337/go-chromath"
)

func HexToRgb(hex string) (color.RGBA, error) {
	rgba := color.RGBA{}
	if hex[0] != '#' {
		return rgba, errors.New("invalid hex code - does not start with a #")
	}
	hexToByte := func(b byte) byte {
		switch {
		case b >= '0' && b <= '9':
			return b - '0'
		case b >= 'a' && b <= 'f':
			return b - 'a' + 10
		case b >= 'A' && b <= 'F':
			return b - 'A' + 10
		}
		return 0
	}
	switch len(hex) {
	case 4:
		rgba.R = hexToByte(hex[1]) * 17
		rgba.G = hexToByte(hex[2]) * 17
		rgba.B = hexToByte(hex[3]) * 17
	case 7:
		rgba.R = hexToByte(hex[1])<<4 + hexToByte(hex[2])
		rgba.G = hexToByte(hex[3])<<4 + hexToByte(hex[4])
		rgba.B = hexToByte(hex[5])<<4 + hexToByte(hex[6])
	default:
		return rgba, errors.New("incorrect length of hex code")
	}
	return rgba, nil
}

func RgbToXYZ(rgba color.RGBA) (chromath.XYZ, error) {
	rgbToXYZ := chromath.NewRGBTransformer(&chromath.SpaceAdobeRGB,
		&chromath.AdaptationBradford, chromath.SpaceAdobeRGB.IlluminantRef,
		&chromath.Scaler8bClamping, 1.0, nil)
	chromathPoint := [3]float64{float64(rgba.R), float64(rgba.G), float64(rgba.B)}
	chromathRGBA := chromath.RGB(chromathPoint)
	c1xyz := rgbToXYZ.Convert(chromathRGBA)
	return c1xyz, nil
}

func XYZToCIE76(xyz chromath.XYZ) (chromath.Lab, error) {
	lab2xyz := chromath.NewLabTransformer(chromath.SpaceAdobeRGB.IlluminantRef)
	c1lab := lab2xyz.Invert(xyz)
	return c1lab, nil
}
