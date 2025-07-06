package db

import (
	"fmt"
	"log"

	"danek.com/telephone/config"
	"danek.com/telephone/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDBConn(cfg *config.Config) (*gorm.DB, error) {

	dbConf := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DataBase.DBHost,
		cfg.DataBase.DBPort,
		cfg.DataBase.DBUser,
		cfg.DataBase.DBPassword,
		cfg.DataBase.DBName,
		cfg.DataBase.DBSSLMode,
	)
	log.Printf("Trying to connect DB with %s", dbConf)
	db, err := gorm.Open(postgres.Open(dbConf), &gorm.Config{})
	if err != nil {
		return nil, domain.FormatErr("DB connection", err)
	}
	if err := AutoMigrate(db); err != nil {
		return nil, domain.FormatErr("Auto Migration", err)
	}
	return db, nil
}
