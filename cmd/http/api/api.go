package api

import (
	"github.com/labstack/echo/v4"
	"github.com/nataneb32/puretodo-api/cmd/http/api/user"
)

func ConfigRouter(g *echo.Group) {
  user.ConfigRoutes(g.Group("/app/user"))
}
