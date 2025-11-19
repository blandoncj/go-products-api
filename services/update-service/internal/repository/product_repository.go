package repository

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UpdateRepository struct {
	collection *mongo.Collection
}

func NewUpdateRepository(db *mongo.Database) *UpdateRepository {
	return &UpdateRepository{collection: db.Collection("products")}
}

func (r *UpdateRepository) UpdateByID(ctx context.Context, id interface{}, update bson.M) (*mongo.UpdateResult, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	opts := options.Update().SetUpsert(false)
	res, err := r.collection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": update}, opts)
	return res, err
}
