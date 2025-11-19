package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Product struct {
	ID          interface{} `bson:"_id,omitempty" json:"_id"`
	Name        string      `bson:"name" json:"name"`
	Description int         `bson:"description" json:"description"`
	Price       float64     `bson:"price" json:"price"`
	Stock       int         `bson:"stock" json:"stock"`
}

type ProductRepository struct {
	collection *mongo.Collection
}

func NewProductRepository(db *mongo.Database) *ProductRepository {
	return &ProductRepository{
		collection: db.Collection("products"),
	}
}

func (r *ProductRepository) FindAll(ctx context.Context) ([]Product, error) {
	cursor, err := r.collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var products []Product
	if err := cursor.All(ctx, &products); err != nil {
		return nil, err
	}

	if products == nil {
		products = []Product{}
	}

	return products, nil
}
