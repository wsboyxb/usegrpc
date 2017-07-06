package main

import (
	"fmt"
	"github.com/wsboyxb/usegrpc/echo/echo"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"time"
)

type server struct {
}

func (*server) SayHi(ctx context.Context, in *echo.Request) (*echo.Reply, error) {
	return &echo.Reply{
		Message:   fmt.Sprintf("%s: %s @%d", "Hi", in.UserName, in.Timestamp),
		Timestamp: time.Now().UnixNano(),
	}, nil
}

func main() {
	s := grpc.NewServer()
	echo.RegisterEchoServer(s, &server{})
	reflection.Register(s)

	lis, err := net.Listen("tcp", ":12345")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
