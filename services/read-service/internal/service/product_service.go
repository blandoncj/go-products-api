package service

import (
	"context"

	"github.com/blandoncj/go-products-api/services/read-service/internal/repository"
)

type ProductService struct {
	repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) GetAll(ctx context.Context) ([]repository.Product, error) {
	return s.repo.FindAll(ctx)
}
