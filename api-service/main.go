package main

import (
	"context"

	"github.com/labstack/echo"
	microclient "github.com/micro/go-micro/client"
	irsPb "github.com/team-morpheus/lasagna-msa/internal-recipes-service/proto"
)

const PORT = "9203"

func main() {

	// create new instance of echo server
	e := echo.New()

	// create new pb client
	client := irsPb.NewInternalRecipesService("lasagna.internal.recipes.service", microclient.DefaultClient)

	// register routes
	e.GET("/", func(c echo.Context) error {
		// get recipes using client
		recipes, err := client.GetRecipes(context.WithTimeout, &irsPb.Request{})
		if err != nil {
			return c.JSON(400, map[string]string{"error": err})
		}
		return c.JSON(200, recipes)
	})

	// start http server wrapped with fatal helper func
	e.Logger.Fatal(e.Start(":" + PORT))
}
