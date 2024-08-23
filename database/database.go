package database

import (
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"shopinbiz.telegram.bot/config"
	"shopinbiz.telegram.bot/models"
)

func Init() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s", config.Config("POSTGRES_HOST"), config.Config("POSTGRES_USER"), config.Config("POSTGRES_PASSWORD"), config.Config("POSTGRES_PORT"), "postgres")
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	if err := DB.Exec("DROP DATABASE IF EXISTS shopinbiz_bot_db;").Error; err != nil {
		panic(err)
	}
	if err := DB.Exec("CREATE DATABASE shopinbiz_bot_db").Error; err != nil {
		panic(err)
	}
	dsn = fmt.Sprintf("host=%s user=%s password=%s port=%s dbname=%s", config.Config("POSTGRES_HOST"), config.Config("POSTGRES_USER"), config.Config("POSTGRES_PASSWORD"), config.Config("POSTGRES_PORT"), "postgres")
	if err != nil {
		log.Fatal(err)
	}
	DB.AutoMigrate(&models.Task{})
	return DB
}
