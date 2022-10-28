package color_service

import (
	"errors"
	"image/color"

	"github.com/jkl1337/go-chromath"
	"github.com/jkl1337/go-chromath/deltae"
)

func HexColorDifference(color1 string, color2 string) (float64, error) {
	rgb1, err := HexToRgb(color1)
	if err != nil {
		return 0, errors.New("bad hex code color 1")
	}
	rgb2, err := HexToRgb(color2)
	if err != nil {
		return 0, errors.New("bad hex code colo 2")
	}
	xyz1, err := RgbToXYZ(rgb1)
	if err != nil {
		return 0, errors.New("problem from rgb to xyz 1")
	}
	xyz2, err := RgbToXYZ(rgb2)
	if err != nil {
		return 0, errors.New("problem from rgb to xyz 2")
	}
	lab1, err := XYZToCIE76(xyz1)
	if err != nil {
		return 0, errors.New("problem from xyz to lab1 1")
	}
	lab2, err := XYZToCIE76(xyz2)
	if err != nil {
		return 0, errors.New("problem from xyz to lab1 2")
	}
	return deltae.CIE2000(lab1, lab2, &deltae.KLChDefault), nil
}
func RGBColorDifference(color1 color.RGBA, color2 color.RGBA) (float64, error) {
	xyz1, err := RgbToXYZ(color1)
	if err != nil {
		return 0, errors.New("problem from rgb to xyz 1")
	}
	xyz2, err := RgbToXYZ(color2)
	if err != nil {
		return 0, errors.New("problem from rgb to xyz 2")
	}
	lab1, err := XYZToCIE76(xyz1)
	if err != nil {
		return 0, errors.New("problem from xyz to lab1 1")
	}
	lab2, err := XYZToCIE76(xyz2)
	if err != nil {
		return 0, errors.New("problem from xyz to lab1 2")
	}
	return deltae.CIE2000(lab1, lab2, &deltae.KLChDefault), nil
}

func CIE76ColorDifference(color1 chromath.Lab, color2 chromath.Lab) float64 {
	return deltae.CIE2000(color1, color2, &deltae.KLChDefault)
}
