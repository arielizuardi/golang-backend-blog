package main

import (
	"log"

	_ "github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/arielizuardi/golang-backend-blog/config"
	"github.com/arielizuardi/golang-backend-blog/server"
)

func main() {
	config.Load()

	dbStr := config.DBConnectionURL()
	dbConn, err := gorm.Open("postgres", dbStr)
	if err != nil {
		log.Fatal("Can't connect to postgres db. ", err)
	}

	serverReady := make(chan bool)
	s := server.Server{
		DBConn:      dbConn,
		Port:        config.Port(),
		ServerReady: serverReady,
	}

	s.Start()
}
