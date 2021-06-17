package infrastructure

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func db() *gorm.DB {
  db, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")))
  if err != nil {
    panic(err)
  }
 
  return db
}

var DB *gorm.DB = db()
