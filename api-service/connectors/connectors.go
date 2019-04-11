package connectors

import (
	"github.com/dins-app/web-services/api-service/models"
	"google.golang.org/grpc"
)

// ConnectServices connects to gRPC services
func ConnectServices(s *models.Server) []*grpc.ClientConn {
	connections := []*grpc.ClientConn{}
	connections = append(connections, ConnectIdentity(s))
	connections = append(connections, ConnectInternalRecipes(s))
	return connections
}
