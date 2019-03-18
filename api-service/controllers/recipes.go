package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dins-app/web-services/api-service/models"
	recipesPb "github.com/dins-app/web-services/proto/internal-recipes-service"
	"github.com/labstack/echo"
)

// RegisterRecipesController registers recipes controller routes
func (c *Controller) RegisterRecipesController() {
	routes := c.Server.Echo.Group("/v1/recipes")
	{
		routes.GET("", c.GetRecipes)
		routes.POST("", c.CreateRecipe)
	}
}

// GetRecipes gets all recipes from the internal recipes service
func (c *Controller) GetRecipes(e echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := (*c.Server.InternalRecipesClient).Get(ctx, &recipesPb.Request{})
	if err != nil {
		log.Println(err)
		return e.JSON(http.StatusInternalServerError, models.Response{
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("Unable to get recipes: %s", err),
		})
	}

	return e.JSON(http.StatusOK, r)
}

// CreateRecipe creates a recipe with the internal recipes service
func (c *Controller) CreateRecipe(e echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	recipe := recipesPb.Recipe{}

	err := e.Bind(&recipe)
	if err != nil {
		log.Println(err)
		return e.JSON(http.StatusInternalServerError, models.Response{
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("Unable to create recipe: %s", err),
		})
	}

	r, err := (*c.Server.InternalRecipesClient).Create(ctx, &recipe)
	if err != nil {
		log.Println(err)
		return e.JSON(http.StatusInternalServerError, models.Response{
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("Unable to create recipe: %s", err),
		})
	}

	return e.JSON(200, r)
}
