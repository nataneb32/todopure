package userDomain

import (
	"github.com/nataneb32/puretodo-api/entities"
	"gorm.io/gorm"
)

type CreateUser interface {
  Create(*gorm.DB, *entities.User) error
}

type FindUser interface {
  FindOneByUsername(db *gorm.DB, username string, u *entities.User) error
}
