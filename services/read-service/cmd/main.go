package main

import (
	"log"
	"net/http"
	"os"
	"parcial3/services/delete/internal/controller"
)

func main() {
	port := "8084"
	if p := GetEnv("DELETE_PORT", "8084"); p != "" {
		port = p
	}
	handler := controller.NewHandler()
	log.Printf("Delete service listening on :%s", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}

func GetEnv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
