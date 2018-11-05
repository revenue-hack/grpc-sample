package main

import (
	"net"
	"log"
	"github.com/go-kit/kit/transport/grpc"
)

func main() {

	listenPort, err := net.Listen("tcp", ":19003")
	if err != nil {
		log.Fatalln(err)
	}
	server := grpc.NewServer()
}
