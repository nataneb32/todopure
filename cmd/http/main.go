package main

import (
  "github.com/labstack/echo/v4"
	"github.com/nataneb32/puretodo-api/cmd/http/api"
)
func main() {
  r := echo.New()

  api.ConfigRouter(r.Group("/"))

  r.Start(":8080")
}
