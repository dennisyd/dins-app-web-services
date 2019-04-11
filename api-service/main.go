package main

import (
	"os"

	"github.com/dins-app/web-services/api-service/connectors"
	"github.com/dins-app/web-services/api-service/controllers"

	"github.com/dins-app/web-services/api-service/models"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var (
	server = models.Server{
		InternalRecipesServerAddr: "localhost:50051",
		IdentityServerAddr:        "localhost:50052",
	}
)

func init() {
	if v, ok := os.LookupEnv("RECIPES_SERVICE_ADDRESS"); ok {
		server.InternalRecipesServerAddr = v
	}
	if v, ok := os.LookupEnv("IDENTITY_SERVICE_ADDRESS"); ok {
		server.IdentityServerAddr = v
	}
}

func main() {

	// connect to grpc services and defer close
	conn := connectors.ConnectServices(&server)
	for conn := range serviceConnections {
		defer conn.Close()
	}

	// create new instance of echo server
	server.Echo = echo.New()

	// register logging middleware
	server.Echo.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${method}  ${uri}  ${latency_human}  ${status}\n",
	}))

	// register controller routes with server
	controllers.Register(&server)

	// start http server
	server.Echo.Logger.Fatal(server.Echo.Start(":8080"))
}
