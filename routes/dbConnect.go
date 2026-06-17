package routes

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DbConnect() *gorm.DB {

	godotenv.Load(".env")

	var db, err = gorm.Open(postgres.Open(os.Getenv("PG_STRING")), &gorm.Config{})

	if err != nil {
		fmt.Printf("A error ocurred: %v \n", err)
	}

	return db

}
