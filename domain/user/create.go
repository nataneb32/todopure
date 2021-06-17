package userDomain

import (
	"fmt"

	"github.com/nataneb32/puretodo-api/entities"
)

func (it *UserService) Create(u *entities.User) error {
  //TODO: Validation of user

  err := it.createUser.Create(it.db, u)
  if err != nil {
    return fmt.Errorf("user create: %w", err)
  }
  return nil
}
