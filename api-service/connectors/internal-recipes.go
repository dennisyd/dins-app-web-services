package connectors

import (
	"log"

	"github.com/dins-app/web-services/api-service/models"
	internalRecipesPb "github.com/dins-app/web-services/proto/internal-recipes-service"
	"google.golang.org/grpc"
)

// ConnectInternalRecipes connects to internal-recipes gRPC server
func ConnectInternalRecipes(s *models.Server) *grpc.ClientConn {
	// Set up a connection to the server.
	conn, err := grpc.Dial(s.InternalRecipesServerAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	// register client
	s.InternalRecipesClient = internalRecipesPb.NewInternalRecipesClient(conn)
	log.Println("Connected", s.InternalRecipesServerAddr)
	return conn
}
