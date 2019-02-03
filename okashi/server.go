package main

import (
	"log"
	"net"

	"github.com/revenue-hack/grpc-sample/okashi/okashi"
	"google.golang.org/grpc"
)

func main() {
	listenPort, err := net.Listen("tcp", ":10101")
	if err != nil {
		log.Fatalln(err)
	}
	server := grpc.NewServer()
	okashiService := &okashi.OkashiService{}

	okashi.RegisterOkashiServiceServer(server, okashiService)
	server.Serve(listenPort)
}
