package service

import (
	"context"

	"github.com/blandoncj/go-products-api/services/create-service/internal/repository"
)

type ProductService struct {
	Repo *repository.ProductRepository
}

func (s *ProductService) Create(ctx context.Context, product interface{}) error {
	return s.Repo.Create(ctx, product)
}
