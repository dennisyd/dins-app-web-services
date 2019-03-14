package main

import (
	"log"
	"net"
	"os"
	"time"

	pb "github.com/dins-app/web-services/proto/internal-recipes-service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

var (
	address = "127.0.0.1:50051"
)

func init() {
	if v, ok := os.LookupEnv("SERVICE_ADDRESS"); ok {
		address = v
	}
}

func main() {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer(
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: 5 * time.Minute, // <--- This fixes it!
		}),
	)
	pb.RegisterRecipesServer(s, &Handler{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
