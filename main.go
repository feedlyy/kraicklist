package main

import (
	"challenge.haraj.com.sa/kraicklist/delivery"
	"challenge.haraj.com.sa/kraicklist/entity"
	"challenge.haraj.com.sa/kraicklist/repository"
	services2 "challenge.haraj.com.sa/kraicklist/services"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	search := entity.Searcher{}
	recordRepo := repository.NewRecordRepository(search)
	service := services2.NewRecordServices(recordRepo)
	delivery.NewRecordHandler(service)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3001"
	}
	// start server
	fmt.Printf("Server is listening on %s...", port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		log.Fatalf("unable to start server due: %v", err)
	}
}
