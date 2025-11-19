package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type ProductRepositoryInterface interface {
	Create(ctx context.Context, product any) error
}

type ProductRepository struct {
	Collection *mongo.Collection
}

func NewProductRepository(db *mongo.Database) *ProductRepository {
	return &ProductRepository{Collection: db.Collection("products")}
}

func (r *ProductRepository) Create(ctx context.Context, product any) error {
	_, err := r.Collection.InsertOne(ctx, product)
	return err
}
