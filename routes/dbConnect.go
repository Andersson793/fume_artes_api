package routes

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db, err = gorm.Open(postgres.Open(os.Getenv("PG_STRING")), &gorm.Config{})

func DbConnect() {

	godotenv.Load(".env.local")

	if err != nil {
		fmt.Printf("A error ocurred: %v \n", err)
	}

}
