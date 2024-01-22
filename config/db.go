package config

import (
	"context"
	"fmt"
	"os"

	"github.com/samsul96maarif/auth-service/lib"
	local_log "github.com/samsul96maarif/auth-service/lib/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDB() (db *lib.Database, err error) {
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", host, username, password, dbname, port)
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err != nil {
		return nil, err
	}
	local_log.Info(context.Background(), "Succesfully connected", make(map[string]interface{}))
	return &lib.Database{DB: conn}, nil
}
