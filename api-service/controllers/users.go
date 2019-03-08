package controllers

import (
	"github.com/labstack/echo"
	"github.com/team-morpheus/lasagna-msa/api-service/models"
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

	return c.JSON(200, map[string]string{"message": "OK"})
}
