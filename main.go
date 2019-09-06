package main

import (
	"log"

	"github.com/arielizuardi/golang-backend-blog/config"
	"github.com/arielizuardi/golang-backend-blog/server"
	"github.com/jinzhu/gorm"
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
