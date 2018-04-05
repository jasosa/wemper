package main

import (
	"context"
	"github.com/jasosa/wemper/pkg/service"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"time"
)

func init() {

}

func main() {

	//Initialize logger
	logger := log.New()
	logger.Formatter = &log.JSONFormatter{}
	logger.Out = os.Stdout
	logger.Level = log.DebugLevel

	conf := service.Config{
		Logger: logger,
		DBName: getEnv("DBNAME", "wempathy"),
		DBUser: getEnv("DBUSER", "wempathy"),
		DBPwd:  getEnv("DBPWD", "wempathy2018"),
		DBHost: getEnv("DBHOST", "tcp(localhost:3306)"),
	}

	//configuring server
	svc := service.New(conf)
	addr := getAddress()
	server := svc.Server(addr)

	//subscribing to os.interrupt signals
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	//starting server in it's own goroutine
	go func() {
		if err := server.ListenAndServe(); err != nil {
			logger.WithError(err).Error()
			stop <- os.Interrupt
		}
	}()

	<-stop //blocking until interrupt arrives

	logger.WithField("address", addr).Info("Shutting down the server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	server.Shutdown(ctx)

	logger.WithField("address", addr).Info("Server gratefully stopped...")
}

func getAddress() string {
	addr := ":" + os.Getenv("PORT")
	if addr == ":" {
		addr = ":8080"
	}
	return addr
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
