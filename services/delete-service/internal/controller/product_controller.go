package controller

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/blandoncj/go-products-api/services/delete-service/internal/repository"
	"github.com/blandoncj/go-products-api/services/delete-service/internal/service"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewHandler() http.Handler {
	user := os.Getenv("MONGO_ROOT_USERNAME")
	pass := os.Getenv("MONGO_ROOT_PASSWORD")
	host := os.Getenv("MONGO_HOST")
	port := os.Getenv("MONGO_PORT")
	if port == "" {
		port = "27017"
	}
	dbname := os.Getenv("MONGO_DB")

	mongoURI := fmt.Sprintf("mongodb://%s:%s@%s:%s/?authSource=admin", user, pass, host, port)
	clientOpts := options.Client().ApplyURI(mongoURI)
	ctxConn, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctxConn, clientOpts)
	if err != nil {
		panic(err)
	}
	ctxPing, cancelPing := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelPing()
	if err := client.Ping(ctxPing, nil); err != nil {
		panic(err)
	}

	db := client.Database(dbname)
	repo := repository.NewDeleteRepository(db)
	svc := service.NewProductService(repo)

	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Delete service OK"))
	})

	mux.HandleFunc("/products/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		idHex := r.URL.Path[len("/products/"):]
		objID, err := primitive.ObjectIDFromHex(idHex)
		if err != nil {
			http.Error(w, "invalid id format", http.StatusBadRequest)
			return
		}
		if err := svc.DeleteProduct(r.Context(), objID); err != nil {
			http.Error(w, "delete error: "+err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status":"deleted"}`))
	})

	return mux
}
