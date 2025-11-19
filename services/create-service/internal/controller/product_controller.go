package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/blandoncj/go-products-api/pkg/model"
	"github.com/blandoncj/go-products-api/services/create-service/internal/repository"
	"github.com/blandoncj/go-products-api/services/create-service/internal/service"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewHandler() http.Handler {
	user := os.Getenv("MONGO_ROOT_USERNAME")
	pass := os.Getenv("MONGO_ROOT_PASSWORD")
	host := os.Getenv("MONGO_HOST")
	dbName := os.Getenv("MONGO_DB")

	mongoURI := fmt.Sprintf("mongodb://%s:%s@%s:27017/?authSource=admin", user, pass, host)

	clientOpts := options.Client().ApplyURI(mongoURI)

	client, err := mongo.Connect(context.Background(), clientOpts)
	if err != nil {
		panic(fmt.Sprintf("error connecting to MongoDB: %v", err))
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := client.Ping(ctx, nil); err != nil {
		panic(fmt.Sprintf("cannot ping MongoDB: %v", err))
	}

	db := client.Database(dbName)
	repo := repository.NewProductRepository(db)
	svc := &service.ProductService{Repo: repo}

	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Create service OK"))
	})

	mux.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		var product model.Product
		if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := svc.Create(r.Context(), product); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(product)
	})

	return mux
}
