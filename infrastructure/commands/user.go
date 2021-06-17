package commands

import (
	"github.com/nataneb32/puretodo-api/entities"
	"gorm.io/gorm"
)

type UserCommands struct {}

func (*UserCommands) Create(db *gorm.DB, u *entities.User) error {
  return db.Create(u).Error
}

func (*UserCommands) FindOneByUsername(db *gorm.DB, username string, u *entities.User) error {
  return db.Where(&entities.User{Username: username}).First(u).Error
}
