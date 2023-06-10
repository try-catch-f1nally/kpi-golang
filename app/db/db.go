package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"kpi-golang/app/models"
	"kpi-golang/app/utils"
	"log"
)

func Init() *gorm.DB {
	host := utils.GetEnvVar("DB_HOST", "localhost")
	port := utils.GetEnvVar("DB_PORT", "5432")
	user := utils.GetEnvVar("DB_USER", "user")
	password := utils.GetEnvVar("DB_PASSWORD", "password")
	dbname := utils.GetEnvVar("DB_DBNAME", "demo")
	sslMode := utils.GetEnvVar("DB_SSLMODE", "disable")
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", host, port, user, password, dbname, sslMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&models.Order{}, &models.Product{}, &models.Review{}, &models.User{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}
