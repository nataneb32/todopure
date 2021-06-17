package user

import (
	"github.com/labstack/echo/v4"
	"github.com/nataneb32/puretodo-api/entities"
)

func CreateHandler(c echo.Context) error {
  var r entities.User
  err := c.Bind(&r)
  if err != nil {
    return err
  }

  err = userService.Create(&r)
  if err != nil {
    return err
  }
  
  c.JSON(200, r)
  return nil
}
