package models

import (
	"github.com/labstack/echo"
	irsPb "github.com/team-morpheus/lasagna-msa/internal-recipes-service/proto"
)

// API will hold our connections to the different services
type API struct {
	IrSvc irsPb.InternalRecipesService
	Echo  *echo.Echo
}
