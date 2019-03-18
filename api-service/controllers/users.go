package controllers

import (
	"net/http"
	"os"

	"github.com/dins-app/web-services/api-service/models"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// RegisterUsersController registers user controller routes
func (c *Controller) RegisterUsersController() {

	// below routes do not require a jwt
	c.Server.Echo.POST("/v1/users/create", c.CreateUser)
	c.Server.Echo.POST("/v1/users/authorize", c.AuthorizeUser)

	routes := c.Server.Echo.Group("/v1/users")
	routes.Use(middleware.JWT([]byte(os.Getenv("JWT_SECRET"))))
	{
		routes.GET("", c.GetUser)
	}
}

// CreateUser creates a user using the identity service
func (c *Controller) CreateUser(e echo.Context) error {
	return e.JSON(http.StatusOK, models.Response{
		Message: "OK",
		Status:  http.StatusOK,
	})
}

// AuthorizeUser authorizes a user using the identity service
// if authorized by the identity service, jwt is generated and
// set to token
func (c *Controller) AuthorizeUser(e echo.Context) error {
	return e.JSON(http.StatusOK, models.Response{
		Message: "OK",
		Status:  http.StatusOK,
	})
}

// GetUser gets a user from identity service
// using authorized jwt token
func (c *Controller) GetUser(e echo.Context) error {
	return e.JSON(http.StatusOK, models.Response{
		Message: "OK",
		Status:  http.StatusOK,
	})
}
