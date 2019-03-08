package controllers

import (
	"context"

	"github.com/labstack/echo"
	"github.com/team-morpheus/lasagna-msa/api-service/models"
	irPb "github.com/team-morpheus/lasagna-msa/internal-recipes-service/proto"
)

// RecipesHandler is our recipes controller struct
type RecipesHandler struct {
	api *models.API
}

// RegisterRecipesRoutes registers controller routes with api
func RegisterRecipesRoutes(api *models.API) {
	h := RecipesHandler{api}
	routes := api.Echo.Group("/api/v1/recipes")
	{
		routes.GET("", h.GetRecipes)
	}
}

// GetRecipes gets recipes from internal recipes svc
func (h *RecipesHandler) GetRecipes(c echo.Context) error {
	// get recipes using client
	recipes, err := h.api.IrSvc.GetRecipes(context.Background(), &irPb.Request{})
	if err != nil {
		return c.JSON(400, map[string]string{"error": err.Error()})
	}
	return c.JSON(200, recipes)
}
