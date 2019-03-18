package controllers

import "github.com/dins-app/web-services/api-service/models"

// Controller is our controllers struct
type Controller struct {
	Server *models.Server
}

// Register registers all controllers with the echo server
func Register(s *models.Server) {
	c := Controller{s}
	c.RegisterRecipesController()
}
