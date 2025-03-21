package config

import (
	"gorm.io/driver/postgres"
	//_ "gorm.io/driver/postgres"
	"os"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect(){
	godotenv.Load(".env")
	dbURL := os.Getenv("DB_URL")
	d, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil{
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB{
	return db
}