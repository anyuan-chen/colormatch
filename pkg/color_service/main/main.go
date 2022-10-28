package main

import (
	"log"
	"net"
	"os"

	colorsv1 "github.com/anyuan-chen/colormatch/gen/proto/go/colors/v1"
	"github.com/anyuan-chen/colormatch/pkg/color_service"
	"google.golang.org/grpc"
)

func main() {
	list, err := net.Listen("tcp", os.Getenv("COLOR_SERVICE_PORT"))
	if err != nil {
		log.Fatalf("service failed to listen on port 6060")
	}
	server := &color_service.ColorServiceServer{}
	grpcServer := grpc.NewServer()
	colorsv1.RegisterPaletteMatchingServiceServer(grpcServer, server)
	if err := grpcServer.Serve(list); err != nil {
		log.Fatalf("service crashed - fatal")
	}
}
