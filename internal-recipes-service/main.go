package main

import (
	"fmt"
	"log"
	"net"
	"time"

	pb "github.com/dins-app/web-services/proto/internal-recipes-service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

const (
	port = ":50051"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost%s", port))
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
