package main

import (
	"log"
	"net"
	"os"

	imagesv1 "github.com/anyuan-chen/colormatch/gen/proto/go/images/v1"
	spotifyv1 "github.com/anyuan-chen/colormatch/gen/proto/go/spotify/v1"
	api_auth "github.com/anyuan-chen/colormatch/pkg/api/auth_api"
	"github.com/anyuan-chen/colormatch/pkg/spotify_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	lis, err := net.Listen("tcp", os.Getenv("SPOTIFY_SERVICE_PORT"))
	if err != nil {
		log.Fatalf("could not listen on tcp port")
	}
	images_client_grpc, err := grpc.Dial(os.Getenv("BASE_URL")+os.Getenv("IMAGE_SERVICE_PORT"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("the color client won't start up")
	}
	defer images_client_grpc.Close()
	images_client := imagesv1.NewPaletteMatchingServiceClient(images_client_grpc)
	server := &spotify_service.SpotifyRetrievalServer{
		Images_Service: images_client,
		Auth:           api_auth.Authenticator,
	}
	//
	grpcServer := grpc.NewServer()
	//tie the protobuf server to GRPC
	spotifyv1.RegisterSpotifyImageColorMatchingServiceServer(grpcServer, server)
	//have the GRPC server listen on the port set earlier
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to listen ff15")
	}

}
