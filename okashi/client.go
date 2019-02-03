package main

import (
	"context"
	"fmt"
	"log"

	"github.com/revenue-hack/grpc-sample/okashi/okashi"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:10101", grpc.WithInsecure())
	if err != nil {
		log.Fatal("client connection error:", err)
	}
	defer conn.Close()

	client := okashi.NewOkashiServiceClient(conn)
	message := &okashi.GetOkashiMessage{Name: "hgoehoge", LikeFreq: 11}
	res, err := client.GetFavoriteOkashi(context.TODO(), message)
	fmt.Printf("result: %#v\n", res)
	fmt.Printf("error: %#v\n", err)
}
