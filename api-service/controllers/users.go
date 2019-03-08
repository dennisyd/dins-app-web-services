package controllers

import (
	"context"

	"github.com/labstack/echo"
	"github.com/team-morpheus/lasagna-msa/api-service/models"
	iPb "github.com/team-morpheus/lasagna-msa/identity-service/proto"
)

// UsersHandler is our users controller struct
type UsersHandler struct {
	api *models.API
}

// RegisterUsersRoutes registers controller routes with api
func RegisterUsersRoutes(api *models.API) {
	h := UsersHandler{api}
	routes := api.Echo.Group("/api/v1/users")
	{
		routes.POST("", h.CreateUser)
	}
}

// CreateUser gets recipes from internal recipes svc
func (h *UsersHandler) CreateUser(c echo.Context) error {
	user := iPb.User{}

	// bind req body to user
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(400, map[string]string{"error": err.Error()})
	}

	// create user using client
	resp, err := h.api.ISvc.Create(context.Background(), &user)
	if err != nil {
		return c.JSON(400, map[string]string{"error": err.Error()})
	}
	return c.JSON(200, resp)
}
