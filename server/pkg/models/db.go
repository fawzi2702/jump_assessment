package models

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/this_is_iz/jump_server/internal/environement"
)

var DB *gorm.DB

// InitializeDB initializes the database connection
func InitializeDB() error {
	dsn, err := environement.Get("DB_DSN")
	if err != nil {
		return err
	}

	database, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("error opening database: %v", err)
	} else {
		fmt.Println("Database connection established")
	}

	if err := database.AutoMigrate(models...); err != nil {
		return fmt.Errorf("error migrating database: %v", err)
	}

	DB = database

	return nil
}

// CloseDB closes the database connection
func CloseDB() {
	sqlDB, err := DB.DB()
	if err != nil {
		fmt.Println("Error getting database connection")
		return
	}

	sqlDB.Close()
}
