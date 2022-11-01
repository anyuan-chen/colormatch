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
	//sets up a listener on the local network
	list, err := net.Listen("tcp", os.Getenv("IMAGE_SERVICE_PORT"))
	if err != nil {
		log.Fatalf("the image service won't listen")
	}
	//create a connection with the color client dependency
	color_client_grpc, err := grpc.Dial(os.Getenv("BASE_URL")+os.Getenv("COLOR_SERVICE_PORT"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("the color client won't start up")
	}
	defer color_client_grpc.Close()
	//create a color client from the connection created earlier
	color_client := colorsv1.NewPaletteMatchingServiceClient(color_client_grpc)
	//create a protobuf server with the dependencies required
	server := &image_service.PaletteMatchingServer{Color_Service: color_client}
	grpcServer := grpc.NewServer()
	//tie the protobuf server to GRPC
	imagesv1.RegisterPaletteMatchingServiceServer(grpcServer, server)
	//have the GRPC server listen on the port set earlier
	if err := grpcServer.Serve(list); err != nil {
		log.Fatalf("failed to listen ff15")
	}
}
