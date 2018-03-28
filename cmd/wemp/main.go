package main

import (
	"github.com/jasosa/wemper/pkg/service"
	"log"
)

func main() {

	svc := service.New()
	server := svc.Server(":8080")
	log.Fatal(server.ListenAndServe())
}
