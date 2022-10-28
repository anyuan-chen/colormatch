package image_service

import (
	"bytes"
	"image"
	"image/jpeg"
)

func BytesToImage(rawData *[]byte) (image.Image, error) {
	img, _, err := image.Decode(bytes.NewReader(*rawData))
	if err != nil {
		return nil, err
	}
	return img, nil
}
func ImageToJPEG(image image.Image) ([]byte, error) {
	buf := new(bytes.Buffer)
	err := jpeg.Encode(buf, image, nil)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
