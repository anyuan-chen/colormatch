package color_service_test

import (
	"image/color"
	"testing"

	"github.com/anyuan-chen/colormatch/pkg/color_service"
	"github.com/stretchr/testify/assert"
)

func TestHexColorDifference(t *testing.T) {
	easy, err := color_service.HexColorDifference("#000000", "#000000")
	assert.Nil(t, err)
	assert.Equal(t, 0.0, easy)
}
func TestRGBColorDifference(t *testing.T) {
	easy, err := color_service.RGBColorDifference(color.RGBA{R: 0, G: 0, B: 0}, color.RGBA{R: 0, G: 0, B: 0})
	assert.Nil(t, err)
	assert.Equal(t, 0.0, easy)
	med, err := color_service.RGBColorDifference(color.RGBA{R: 123, G: 4, B: 5}, color.RGBA{R: 253, G: 121, B: 123})
	assert.Nil(t, err)
	assert.Equal(t, 42.93981408224605, med)
	hard, err := color_service.RGBColorDifference(color.RGBA{R: 123, G: 255, B: 54}, color.RGBA{R: 253, G: 121, B: 123})
	assert.Nil(t, err)
	assert.Equal(t, 85.49479206935433, hard)
}

