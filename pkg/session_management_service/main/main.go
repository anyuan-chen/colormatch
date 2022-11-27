package main

import (
	"log"
	"net"
	"os"

	session_managementv1 "github.com/anyuan-chen/colormatch/gen/proto/go/session_management/v1"
	"github.com/anyuan-chen/colormatch/pkg/session_management_service"
	"golang.org/x/oauth2"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", os.Getenv("SESSION_MANAGEMENT_PORT"))
	if err != nil {
		log.Fatal(err)
	}
	string_token := make(map[string]oauth2.Token)
	session_manager := &session_management_service.Session_Management_Service{
		Sessions: string_token,
	}
	server := grpc.NewServer()
	session_managementv1.RegisterSessionManagementServiceServer(server, session_manager)
	if err := server.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
