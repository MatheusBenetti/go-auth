package models

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

var DB *gorm.DB

func ConnectDB(db Database) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", db.Host, db.User, db.Password, db.DBName, db.Port)
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %w", err)
	}

	if err := database.AutoMigrate(&User{}); err != nil {
		return nil, fmt.Errorf("failed to auto-migrate database: %w", err)
	}

	fmt.Println("Migrated database")
	DB = database
	return database, nil
}
