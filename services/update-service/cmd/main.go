package main

import (
	"log"
	"net/http"
	"os"

	"github.com/blandoncj/go-products-api/services/update-service/internal/controller"
)

func main() {
	port := os.Getenv("UPDATE_SERVICE_PORT")
	if port == "" {
		port = "8083"
	}
	handler := controller.NewHandler()
	log.Printf("Update service listening on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
