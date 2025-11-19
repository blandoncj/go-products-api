package service

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MockDeleteRepository struct {
	mock.Mock
}

func (m *MockDeleteRepository) DeleteByID(ctx context.Context, id any) (*mongo.DeleteResult, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*mongo.DeleteResult), args.Error(1)
}

func TestProductService_DeleteProduct_Success(t *testing.T) {
	// Arrange
	mockRepo := new(MockDeleteRepository)
	service := NewProductService(mockRepo)
	ctx := context.Background()
	productID := primitive.NewObjectID()

	mockRepo.On("DeleteByID", ctx, productID).Return(&mongo.DeleteResult{DeletedCount: 1}, nil)

	// Act
	err := service.DeleteProduct(ctx, productID)

	// Assert - Regla de negocio: Producto existente debe eliminarse correctamente
	assert.NoError(t, err, "El producto debe eliminarse sin errores")
	mockRepo.AssertExpectations(t)
}

func TestProductService_DeleteProduct_NotFound(t *testing.T) {
	// Arrange
	mockRepo := new(MockDeleteRepository)
	service := NewProductService(mockRepo)
	ctx := context.Background()
	productID := primitive.NewObjectID()

	mockRepo.On("DeleteByID", ctx, productID).Return(&mongo.DeleteResult{DeletedCount: 0}, nil)

	// Act
	err := service.DeleteProduct(ctx, productID)

	// Assert - Regla de negocio: Eliminar producto inexistente no genera error (idempotencia)
	assert.NoError(t, err, "No debe generar error si el producto no existe")
	mockRepo.AssertExpectations(t)
}

func TestProductService_DeleteProduct_DatabaseError(t *testing.T) {
	// Arrange
	mockRepo := new(MockDeleteRepository)
	service := NewProductService(mockRepo)
	ctx := context.Background()
	productID := primitive.NewObjectID()

	mockRepo.On("DeleteByID", ctx, productID).Return(nil, errors.New("error de conexión"))

	// Act
	err := service.DeleteProduct(ctx, productID)

	// Assert - Regla de negocio: Errores de BD deben propagarse
	assert.Error(t, err, "Debe retornar error cuando hay problema de conexión")
	assert.Contains(t, err.Error(), "conexión")
	mockRepo.AssertExpectations(t)
}
