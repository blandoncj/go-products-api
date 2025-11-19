package repository

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type DeleteRepository struct {
	collection *mongo.Collection
}

func NewDeleteRepository(db *mongo.Database) *DeleteRepository {
	return &DeleteRepository{collection: db.Collection("products")}
}

func (r *DeleteRepository) DeleteByID(ctx context.Context, id interface{}) (*mongo.DeleteResult, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	res, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	return res, err
}
