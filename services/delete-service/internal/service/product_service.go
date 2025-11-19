package service

import (
	"context"

	"github.com/blandoncj/go-products-api/services/delete-service/internal/repository"
)

type ProductService struct {
	repo repository.ProductRepositoryInterface
}

func NewProductService(repo repository.ProductRepositoryInterface) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) DeleteProduct(ctx context.Context, id interface{}) error {
	_, err := s.repo.DeleteByID(ctx, id)
	return err
}
