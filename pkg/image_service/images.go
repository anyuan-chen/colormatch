package image_service

import (
	"context"

	imagev1 "github.com/anyuan-chen/colormatch/gen/proto/go/images/v1"
	sharedv1 "github.com/anyuan-chen/colormatch/gen/proto/go/shared/v1"
)

type ImageServiceServer struct {
	imagev1.UnimplementedPaletteMatchingServiceServer
}

func (i *ImageServiceServer) GetBackgroundColor(ctx context.Context, req *imagev1.GetBackgroundColorRequest) (imagev1.GetBackgroundColorResponse, error) {
	//image := req.Image
	//palette = req.Palette
	var backgroundColor *sharedv1.Color
	return imagev1.GetBackgroundColorResponse{Color: backgroundColor}, nil
}
