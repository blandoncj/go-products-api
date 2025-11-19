package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/blandoncj/go-products-api/services/read-service/internal/repository"
	"github.com/blandoncj/go-products-api/services/read-service/internal/service"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewHandler() http.Handler {
	// variables de entorno esperadas
	user := os.Getenv("MONGO_ROOT_USERNAME")
	pass := os.Getenv("MONGO_ROOT_PASSWORD")
	host := os.Getenv("MONGO_HOST")
	port := os.Getenv("MONGO_PORT")
	if port == "" {
		port = "27017"
	}
	dbname := os.Getenv("MONGO_DB")
	if dbname == "" {
		dbname = "parcialdb"
	}

	if user == "" || pass == "" || host == "" {
		panic("missing MongoDB environment variables for read service")
	}

	mongoURI := fmt.Sprintf("mongodb://%s:%s@%s:%s/?authSource=admin", user, pass, host, port)

	clientOpts := options.Client().ApplyURI(mongoURI)
	// conectar con timeout
	ctxConn, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctxConn, clientOpts)
	if err != nil {
		panic(fmt.Sprintf("mongo connect error: %v", err))
	}
	// ping
	ctxPing, cancelPing := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelPing()
	if err := client.Ping(ctxPing, nil); err != nil {
		panic(fmt.Sprintf("mongo ping error: %v", err))
	}

	db := client.Database(dbname)
	repo := repository.NewProductRepository(db)
	svc := service.NewProductService(repo)

	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("Read service OK"))
	})

	mux.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		products, err := svc.GetAll(r.Context())
		if err != nil {
			http.Error(w, "error reading products: "+err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(products)
	})

	return mux
}
