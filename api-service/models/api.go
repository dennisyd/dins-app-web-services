package models

import (
	"github.com/labstack/echo"
	iPb "github.com/team-morpheus/lasagna-msa/internal-recipes-service/proto"
	irPb "github.com/team-morpheus/lasagna-msa/internal-recipes-service/proto"
)

// API will hold our connections to the different services
type API struct {
	IrSvc irPb.InternalRecipesService
	ISvc  iPb.IdentityService
	Echo  *echo.Echo
}
