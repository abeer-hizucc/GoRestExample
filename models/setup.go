package models

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func NewDatabaseConfig() DBConfig {
	return DBConfig{
		User:     "postgres",
		Password: "1234",
		DBName:   "GoTest",
		Port:     5433,
		SSLMode:  "disable",
		TimeZone: "Asia/Shanghai",
	}
}

func ConnectDatabase() error {
	config := NewDatabaseConfig()
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
		config.User, config.Password, config.DBName, config.Port,
		config.SSLMode, config.TimeZone)

	database, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to Database: %w", err)
	}
	err = database.AutoMigrate(&Book{})
	if err != nil {
		return fmt.Errorf("failed to migrate database schema: %w", err)
	}
	DB = database
	return nil
}
