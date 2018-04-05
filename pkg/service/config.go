package service

import (
	log "github.com/sirupsen/logrus"
)

//Config stores all the configuration for the service
type Config struct {
	Logger *log.Logger
	DBUser string
	DBPwd  string
	DBName string
	DBHost string
}
