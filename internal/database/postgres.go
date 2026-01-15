package database

import (
	"log"

	"backend-event-api/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	dsn := "host=localhost user=postgres password=postgres dbname=eventdb port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}

	err = db.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatal("migration failed")
	}

	log.Println("Database connected & migrated")
	return db
}
