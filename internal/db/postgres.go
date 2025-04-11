package db

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/OmastarWasabi/people-api/internal/models"
)

func InitDB() *gorm.DB {
	dsn := "host=localhost user=postgres password=secret dbname=people_db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("не удалось подключиться к базе", err)
	}

	err = db.AutoMigrate(models.User{})
	if err != nil {
		log.Fatal("ошибка миграции", err)
	}

	fmt.Println("Подключение к базе успешно!")
	return db
}