package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/blandoncj/go-products-api/services/create-service/internal/controller"
)

func main() {
	port := os.Getenv("CREATE_SERVICE_PORT")
	if port == "" {
		port = "8081"
	}

	handler := controller.NewHandler()
	addr := fmt.Sprintf(":%s", port)
	log.Printf("Create service listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, handler))
}
