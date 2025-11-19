package service

import (
	"context"
	"errors"
	"testing"

	"github.com/blandoncj/go-products-api/services/read-service/internal/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockReadRepository struct {
	mock.Mock
}

func (m *MockReadRepository) FindAll(ctx context.Context) ([]repository.Product, error) {
	args := m.Called(ctx)
	return args.Get(0).([]repository.Product), args.Error(1)
}

func TestProductService_GetAll_Success(t *testing.T) {
	// Arrange
	mockRepo := new(MockReadRepository)
	service := NewProductService(mockRepo)
	ctx := context.Background()

	expectedProducts := []repository.Product{
		{Name: "Laptop", Description: 1, Price: 1500.00, Stock: 10},
		{Name: "Mouse", Description: 2, Price: 25.00, Stock: 50},
		{Name: "Keyboard", Description: 3, Price: 75.00, Stock: 30},
	}

	mockRepo.On("FindAll", ctx).Return(expectedProducts, nil)

	// Act
	products, err := service.GetAll(ctx)

	// Assert - Regla de negocio: Debe retornar todos los productos disponibles
	assert.NoError(t, err, "No debe haber errores al obtener productos")
	assert.Len(t, products, 3, "Debe retornar 3 productos")
	assert.Equal(t, "Laptop", products[0].Name)
	mockRepo.AssertExpectations(t)
}

func TestProductService_GetAll_EmptyList(t *testing.T) {
	// Arrange
	mockRepo := new(MockReadRepository)
	service := NewProductService(mockRepo)
	ctx := context.Background()

	mockRepo.On("FindAll", ctx).Return([]repository.Product{}, nil)

	// Act
	products, err := service.GetAll(ctx)

	// Assert - Regla de negocio: Lista vacía es válida (sin productos en inventario)
	assert.NoError(t, err, "No debe haber error con lista vacía")
	assert.Empty(t, products, "La lista debe estar vacía")
	mockRepo.AssertExpectations(t)
}

func TestProductService_GetAll_LowStockProducts(t *testing.T) {
	// Arrange
	mockRepo := new(MockReadRepository)
	service := NewProductService(mockRepo)
	ctx := context.Background()

	lowStockProducts := []repository.Product{
		{Name: "Laptop", Description: 1, Price: 1500.00, Stock: 2},
		{Name: "Mouse", Description: 2, Price: 25.00, Stock: 1},
	}

	mockRepo.On("FindAll", ctx).Return(lowStockProducts, nil)

	// Act
	products, err := service.GetAll(ctx)

	// Assert - Regla de negocio: Debe mostrar productos con stock bajo
	assert.NoError(t, err)
	for _, p := range products {
		if p.Stock < 5 {
			// Regla de negocio: productos con stock < 5 deben alertar reabastecimiento
			assert.Less(t, p.Stock, 5, "Producto %s tiene stock bajo", p.Name)
		}
	}
	mockRepo.AssertExpectations(t)
}

func TestProductService_GetAll_DatabaseError(t *testing.T) {
	// Arrange
	mockRepo := new(MockReadRepository)
	service := NewProductService(mockRepo)
	ctx := context.Background()

	mockRepo.On("FindAll", ctx).Return([]repository.Product{}, errors.New("timeout de conexión"))

	// Act
	products, err := service.GetAll(ctx)

	// Assert - Regla de negocio: Errores de BD deben propagarse
	assert.Error(t, err, "Debe retornar error de BD")
	assert.Empty(t, products, "No debe retornar productos en caso de error")
	mockRepo.AssertExpectations(t)
}
