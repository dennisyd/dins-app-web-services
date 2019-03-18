package main

import (
	"net/http"
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
	}
)

func init() {
	if v, ok := os.LookupEnv("RECIPES_SERVICE_ADDRESS"); ok {
		server.InternalRecipesServerAddr = v
	}
}

func main() {

	// connect to recipes service, if connection is lost
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

	// create new instance of echo server
	server.Echo = echo.New()

	// register logging middleware
	server.Echo.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${method}  ${uri}  ${latency_human}  ${status}\n",
	}))

	// register controller routes with server
	controllers.Register(&server)

	// register / GET route
	server.Echo.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// start http server
	server.Echo.Logger.Fatal(server.Echo.Start(":8080"))
}
