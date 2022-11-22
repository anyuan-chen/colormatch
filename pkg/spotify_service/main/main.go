package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	lis, err := net.Listen("tcp", os.Getenv("SPOTIFY_SERVICE_PORT"))
	if err != nil {
		log.Fatalf("could not listen on tcp port")
	}
	fmt.Print(lis.Addr().String())

}
