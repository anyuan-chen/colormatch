package main

import (
	"log"
	"net"
	"os"

	colorsv1 "github.com/anyuan-chen/colormatch/gen/proto/go/colors/v1"
	imagesv1 "github.com/anyuan-chen/colormatch/gen/proto/go/images/v1"
	"github.com/anyuan-chen/colormatch/pkg/image_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	list, err := net.Listen("tcp", os.Getenv("IMAGE_SERVICE_PORT"))
	if err != nil {
		log.Fatalf("the image service won't listen")
	}
	color_client_grpc, err := grpc.Dial(os.Getenv("BASE_URL")+os.Getenv("COLOR_SERVICE_PORT"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("the color client won't start up")
	}
	defer color_client_grpc.Close()
	color_client := colorsv1.NewPaletteMatchingServiceClient(color_client_grpc)
	server := &image_service.PaletteMatchingServer{Color_Service: color_client}
	grpcServer := grpc.NewServer()
	imagesv1.RegisterPaletteMatchingServiceServer(grpcServer, server)
	if err := grpcServer.Serve(list); err != nil {
		log.Fatalf("failed to listen ff15")
	}
}
