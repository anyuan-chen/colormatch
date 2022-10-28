package main

import (
	"net"
	"os"
)

func main() {
	list, err := net.Listen("tcp", os.Getenv("IMAGE_SERVICE_PORT"))
}
