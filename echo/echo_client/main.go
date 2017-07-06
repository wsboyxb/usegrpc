package main

import (
	"github.com/wsboyxb/usegrpc/echo/echo"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:12345", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := echo.NewEchoClient(conn)

	reply, err := client.SayHi(context.Background(),
		&echo.Request{
			UserName:  "yang xb",
			Timestamp: time.Now().UnixNano(),
		})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("receive from server: %#v", reply)
}
