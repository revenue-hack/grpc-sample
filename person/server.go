package main

import (
	"net"
	"log"
	"google.golang.org/grpc"
	"github.com/revenue-hack/grpc-sample/person/service"
	"github.com/revenue-hack/grpc-sample/person/protocol"
	"os"
	"os/signal"
)

func main() {
	listenPort, err := net.Listen("tcp", ":9999")
	if err != nil {
		log.Fatalln(err)
	}
	server := grpc.NewServer()
	personService := &service.PersonService{}
	protocol.RegisterPersonServiceServer(server, personService)


	go func() {
		server.Serve(listenPort)
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<- quit
	log.Println("grpc server kill")
	server.Stop()
}
