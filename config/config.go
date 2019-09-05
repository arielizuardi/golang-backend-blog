package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

var appConfig config

type config struct {
	appPort                   int
	dbHost                    string
	dbPort                    int
	dbUsername                string
	dbPassword                string
	dbName                    string
	dbMaxPoolSize             int
	dbMaxIdleConn             int
	dbMaxOpenConn             int
	dbMaxConnLifetimeDuration time.Duration
	logLevel                  string
}

// Load load config
func Load() error {
	viper.AutomaticEnv()
	appConfig = config{
		appPort:                   viper.GetInt("PORT"),
		dbHost:                    viper.GetString("DB_HOST"),
		dbPort:                    viper.GetInt("DB_PORT"),
		dbUsername:                viper.GetString("DB_USER"),
		dbPassword:                viper.GetString("DB_PASS"),
		dbName:                    viper.GetString("DB_NAME"),
		dbMaxIdleConn:             viper.GetInt("DB_MAX_IDLE_CONN"),
		dbMaxOpenConn:             viper.GetInt("DB_MAX_OPEN_CONN"),
		dbMaxConnLifetimeDuration: viper.GetDuration("DB_CONN_MAX_LIFETIME"),
		logLevel:                  viper.GetString("LOG_LEVEL"),
	}

	return nil
}

// DBConnectionURL returns db connecction url
func DBConnectionURL() string {
	return fmt.Sprintf("postgres://%s:%s@%s:5432/%s?sslmode=disable", appConfig.dbUsername, appConfig.dbPassword, appConfig.dbHost, appConfig.dbName)
}
