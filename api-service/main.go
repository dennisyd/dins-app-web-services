package main

import (
	"net/url"
	"os"
	"time"

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

	// connect to internal recipes service, if connection is lost
	// retry connection every 3 seconds
	go func() {
		for {
			if server.InternalRecipesClient == nil {
				conn := connectors.ConnectInternalRecipes(&server)
				defer conn.Close()
			}
			// sleep 3 seconds
			time.Sleep(3 * time.Second)
		}
	}()

	// connect to identity service, if connection is lost
	// retry connection every 3 seconds
	go func() {
		for {
			if server.IdentityClient == nil {
				conn := connectors.ConnectIdentity(&server)
				defer conn.Close()
			}
			// sleep 3 seconds
			time.Sleep(3 * time.Second)
		}
	}()

	// create new instance of echo server
	server.Echo = echo.New()

	// parse url for drone server
	droneURL, err := url.Parse(os.Getenv("DRONE_SERVER_URL"))
	if err != nil {
		server.Echo.Logger.Fatal(err)
	}

	// configure proxy middleware
	targets := []*middleware.ProxyTarget{
		{
			URL: droneURL,
		},
	}

	// register group for drone
	server.Echo.Group("/", middleware.Proxy(middleware.NewRoundRobinBalancer(targets)))

	// register logging middleware
	server.Echo.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${method}  ${uri}  ${latency_human}  ${status}\n",
	}))

	// register controller routes with server
	controllers.Register(&server)

	// start http server
	server.Echo.Logger.Fatal(server.Echo.Start(":8080"))
}
