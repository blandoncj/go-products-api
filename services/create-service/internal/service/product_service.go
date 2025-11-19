package service

import (
	"context"

	"github.com/blandoncj/go-products-api/services/create-service/internal/repository"
)

type ProductService struct {
	Repo repository.ProductRepositoryInterface
}

func (s *ProductService) Create(ctx context.Context, product any) error {
	return s.Repo.Create(ctx, product)
}
