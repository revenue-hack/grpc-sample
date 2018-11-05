package main

import (
	"net"
	"log"
	"google.golang.org/grpc"
	"github.com/revenue-hack/grpc-sample/cat/service"
	"github.com/revenue-hack/grpc-sample/cat/pb"
)

func main() {

	listenPort, err := net.Listen("tcp", ":19003")
	if err != nil {
		log.Fatalln(err)
	}
	server := grpc.NewServer()
	catService := &service.MyCatService{}

	pb.RegisterCatServer(server, catService)
	server.Serve(listenPort)
}
