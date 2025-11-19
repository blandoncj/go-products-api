package service

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockProductRepository struct {
	mock.Mock
}

func (m *MockProductRepository) Create(ctx context.Context, product any) error {
	args := m.Called(ctx, product)
	return args.Error(0)
}

func TestProductService_Create_Success(t *testing.T) {
	mockRepo := new(MockProductRepository)
	service := &ProductService{Repo: mockRepo}
	ctx := context.Background()

	product := map[string]any{
		"name":        "Laptop",
		"description": "High performance laptop",
		"price":       1500.00,
		"stock":       10,
	}

	mockRepo.On("Create", ctx, product).Return(nil)

	err := service.Create(ctx, product)

	assert.NoError(t, err, "El producto debe crearse sin errores")
	mockRepo.AssertExpectations(t)
}

func TestProductService_Create_InvalidProduct(t *testing.T) {
	mockRepo := new(MockProductRepository)
	service := &ProductService{Repo: mockRepo}
	ctx := context.Background()

	invalidProduct := map[string]any{
		"name":  "Invalid Product",
		"price": -100.00,
		"stock": 5,
	}

	mockRepo.On("Create", ctx, invalidProduct).Return(errors.New("precio no puede ser negativo"))

	err := service.Create(ctx, invalidProduct)

	assert.Error(t, err, "Debe fallar cuando el precio es negativo")
	assert.Contains(t, err.Error(), "precio no puede ser negativo")
	mockRepo.AssertExpectations(t)
}

func TestProductService_Create_StockZero(t *testing.T) {
	mockRepo := new(MockProductRepository)
	service := &ProductService{Repo: mockRepo}
	ctx := context.Background()

	product := map[string]any{
		"name":  "Out of Stock Product",
		"price": 50.00,
		"stock": 0,
	}

	mockRepo.On("Create", ctx, product).Return(nil)

	err := service.Create(ctx, product)

	assert.NoError(t, err, "Producto con stock 0 puede crearse (para pre-ordenes)")
	mockRepo.AssertExpectations(t)
}
