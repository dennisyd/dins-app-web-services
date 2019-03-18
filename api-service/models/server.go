package models

import (
	identityPb "github.com/dins-app/web-services/proto/identity-service"
	internalRecipesPb "github.com/dins-app/web-services/proto/internal-recipes-service"
	"github.com/labstack/echo"
)

// Server is our server model
type Server struct {
	Echo                      *echo.Echo
	InternalRecipesServerAddr string
	InternalRecipesClient     *internalRecipesPb.InternalRecipesClient
	IdentityServerAddr        string
	IdentityClient            *identityPb.IdentityClient
}
