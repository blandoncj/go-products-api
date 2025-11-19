package main

import (
	"log"
	"net/http"
	"os"

	"github.com/blandoncj/go-products-api/services/read-service/internal/controller"
)

func main() {
	port := os.Getenv("READ_SERVICE_PORT")
	if port == "" {
		port = "8082"
	}

	handler := controller.NewHandler()
	log.Printf("Read service listening on :%s", port)
	if err := http.ListenAndServe(":"+port, handler); err != nil {
		log.Fatalf("read service failed: %v", err)
	}
}
