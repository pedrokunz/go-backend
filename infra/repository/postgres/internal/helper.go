package internal

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	client *gorm.DB
)

func Connect() (*gorm.DB, error) {
	if client != nil {
		return client, nil
	}

	newClient, err := gorm.Open(postgres.Open(getDSN()), &gorm.Config{})

	client = newClient

	err = configDB()

	return client, err
}

func getDSN() string {
	dsn := fmt.Sprintf("host=%s user=%s password=%s sslmode=%s dbname=restaurant port=5432 TimeZone=America/Sao_Paulo",
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_SSLMODE"))

	return dsn
}

func configDB() error {
	sqlDB, err := client.DB()
	if err != nil {
		return err
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return nil
}
