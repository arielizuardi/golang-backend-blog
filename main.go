package main

import (
	"log"

	"github.com/arielizuardi/sph-backend-coding-challenge/config"
	"github.com/arielizuardi/sph-backend-coding-challenge/server"
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
