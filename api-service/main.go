package main

import (
	"context"
	"log"
	"net/http"
	"time"

	recipesPb "github.com/dins-app/web-services/proto/internal-recipes-service"
	"github.com/labstack/echo"
	"google.golang.org/grpc"
)

const (
	recipesAddr = "localhost:50051"
)

var (
	recipesClient recipesPb.RecipesClient
)

func main() {

	// connect to recipes service
	recipesService := connectRecipesServer()
	defer recipesService.Close()

	// create new instance of echo server
	e := echo.New()

	// register / GET route
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// register /recipe GET route
	e.GET("/recipe", getRecipe)

	// start http server
	e.Logger.Fatal(e.Start(":8080"))
}

func connectRecipesServer() *grpc.ClientConn {

	// Set up a connection to the server.
	conn, err := grpc.Dial(recipesAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	// register client
	recipesClient = recipesPb.NewRecipesClient(conn)

	return conn
}

func getRecipe(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := recipesClient.Get(ctx, &recipesPb.Request{})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"Status": "500",
			"Error":  "Unable to get recipes",
		})
	}

	return c.JSON(200, r.Recipes)
}
