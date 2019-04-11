package connectors

import (
	"log"

	"github.com/dins-app/web-services/api-service/models"
	identityPb "github.com/dins-app/web-services/proto/identity-service"
	"google.golang.org/grpc"
)

// ConnectIdentity connects to identity gRPC server
func ConnectIdentity(s *models.Server) *grpc.ClientConn {
	// Set up a connection to the server.
	conn, err := grpc.Dial(s.IdentityServerAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	// register client
	s.IdentityClient = identityPb.NewIdentityClient(conn)
	log.Println("Connected", s.IdentityServerAddr)
	return conn
}
