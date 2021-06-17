package user

import (
	"github.com/labstack/echo/v4"
	"github.com/nataneb32/puretodo-api/domain/user"
	"github.com/nataneb32/puretodo-api/infrastructure"
)

var userService *userDomain.UserService = userDomain.New(infrastructure.DB)

func ConfigRoutes(g *echo.Group) {
  g.POST("/create", CreateHandler)
}
