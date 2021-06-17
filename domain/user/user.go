package userDomain

import (
	"github.com/nataneb32/puretodo-api/infrastructure/commands"
	"gorm.io/gorm"
)

type UserService struct {
  db *gorm.DB
  createUser CreateUser
  findUser FindUser
}

func New(db *gorm.DB) *UserService {
  userCommands := &commands.UserCommands{}

  return &UserService{
    db,
    userCommands,
    userCommands,
  }
}


