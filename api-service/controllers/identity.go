package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/dins-app/web-services/api-service/models"
	identityPb "github.com/dins-app/web-services/proto/identity-service"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// RegisterIdentitiesController registers identity controller routes
func (c *Controller) RegisterIdentitiesController() {

	// below routes do not require a jwt
	c.Server.Echo.POST("/v1/identities/create", c.CreateUser)
	c.Server.Echo.POST("/v1/identities/authorize", c.AuthorizeUser)

	routes := c.Server.Echo.Group("/v1/identities")
	routes.Use(middleware.JWT([]byte(os.Getenv("JWT_SECRET"))))
	{
		routes.GET("", c.GetUser)
	}
}

// CreateUser creates a user using the identity service
func (c *Controller) CreateUser(e echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	user := identityPb.User{}

	err := e.Bind(&user)
	if err != nil {
		log.Println(err)
		return e.JSON(http.StatusInternalServerError, models.Response{
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("Unable to create user: %s", err),
		})
	}

	r, err := c.Server.IdentityClient.Create(ctx, &user)
	if err != nil {
		log.Println(err)
		return e.JSON(http.StatusInternalServerError, models.Response{
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("Unable to create user: %s", err),
		})
	}

	// generate jwt token for response
	err = generateJwt(r)
	if err != nil {
		log.Println(err)
		return e.JSON(http.StatusInternalServerError, models.Response{
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("Unable to create user: %s", err),
		})
	}

	return e.JSON(200, r)
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

// generate jwt token
func generateJwt(r *identityPb.Response) error {
	var err error

	// create token
	token := jwt.New(jwt.SigningMethodHS256)

	// set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = r.User.Id
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// generate encoded token and set Token field
	r.Token, err = token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return err
	}

	// if all went well, return nil
	return nil
}
