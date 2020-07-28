package api

import (
	"crud-go/api/controllers"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var server = controllers.Server{}

// Run ...
func Run() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}
	server.Initialize(os.Getenv("DB_DRIVER_PSQL"), os.Getenv("DB_USER_PSQL"), os.Getenv("DB_PASSWORD_PSQL"), os.Getenv("DB_PORT_PSQL"), os.Getenv("DB_HOST_PSQL"), os.Getenv("DB_NAME_PSQL"))

	// seed.Load(server.DB)
	server.Run(":8080")
}
