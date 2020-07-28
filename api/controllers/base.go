package controllers

import (
	"crud-go/api/models"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Server ...
type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

//Initialize ...
func (server *Server) Initialize(DbDriver, DbUser, DbPassword, DbPort, DbHost, DbName string) {
	var err error

	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DbHost, DbPort, DbUser, DbName, DbPassword)
	server.DB, err = gorm.Open(DbDriver, DBURL)
	if err != nil {
		fmt.Printf("Cannot connect to %s database", DbDriver)
		log.Fatal("This is the error: ", err)
	} else {
		fmt.Printf("We are connected to the %s database", DbDriver)
	}
	server.DB.Debug().AutoMigrate(&models.User{}) // database migration
	server.Router = mux.NewRouter()

	server.initializeRoutes()
}

// Run ...
func (server *Server) Run(addr string) {
	fmt.Printf("Listening on port %s", addr)
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
