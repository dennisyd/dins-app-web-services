package main

import (
	"log"
	"net"
	"os"
	"time"

	pb "github.com/dins-app/web-services/proto/identity-service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

var (
	address = "127.0.0.1:50052"
)

func init() {
	if v, ok := os.LookupEnv("SERVICE_ADDRESS"); ok {
		address = v
	}
}

func main() {

	// connect to mongodb and defer session close
	dbSession := ConnectDB()
	defer dbSession.Close()

	// start listening on addr
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// start new grpc service with 5 min connection timeout
	s := grpc.NewServer(
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: 5 * time.Minute,
		}),
	)

	// register handlers with grpc
	pb.RegisterIdentityServer(s, &Handler{})

	// serve grpc over listener
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
