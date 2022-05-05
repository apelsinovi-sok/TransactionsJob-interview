package DB

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New() *gorm.DB {
	db, err := gorm.Open(postgres.Open("host=localhost user=user password=123 dbname=postgres port=5432 sslmode=disable"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
