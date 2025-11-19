package service

import (
	"context"

	"github.com/blandoncj/go-products-api/services/update-service/internal/repository"
	"go.mongodb.org/mongo-driver/bson"
)

type ProductService struct {
	repo *repository.UpdateRepository
}

func NewProductService(repo *repository.UpdateRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) UpdateProduct(ctx context.Context, id interface{}, name string, description string) error {
	update := bson.M{}

	if name != "" {
		update["name"] = name
	}
	update["description"] = description
	_, err := s.repo.UpdateByID(ctx, id, update)
	return err
}
