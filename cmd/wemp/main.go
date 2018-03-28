package main

import (
	"context"
	"github.com/jasosa/wemper/pkg/service"
	"log"
	"os"
	"os/signal"
	"time"
)

func main() {

	//configuring server
	svc := service.New()
	addr := getAddress()
	server := svc.Server(addr)

	//subscribing to os.interrupt signals
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	//starting server in it's own goroutine
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Println(err)
			stop <- os.Interrupt
		}
	}()

	<-stop //blocking until interrupt arrives

	log.Println("\nShutting down the server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	server.Shutdown(ctx)

	log.Println("\nServer gratefully stopped...")
}

func getAddress() string {
	addr := ":" + os.Getenv("PORT")
	if addr == ":" {
		addr = ":8080"
	}
	return addr
}
