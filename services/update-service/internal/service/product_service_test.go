package service

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MockUpdateRepository struct {
	mock.Mock
}

func (m *MockUpdateRepository) UpdateByID(ctx context.Context, id interface{}, update bson.M) (*mongo.UpdateResult, error) {
	args := m.Called(ctx, id, update)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*mongo.UpdateResult), args.Error(1)
}

func TestProductService_UpdateProduct_Success(t *testing.T) {
	// Arrange
	mockRepo := new(MockUpdateRepository)
	service := NewProductService(mockRepo)
	ctx := context.Background()
	productID := primitive.NewObjectID()

	newName := "Laptop Pro"
	newDescription := "Updated description"

	expectedUpdate := bson.M{
		"name":        newName,
		"description": newDescription,
	}

	mockRepo.On("UpdateByID", ctx, productID, expectedUpdate).Return(&mongo.UpdateResult{ModifiedCount: 1}, nil)

	// Act
	err := service.UpdateProduct(ctx, productID, newName, newDescription)

	// Assert - Regla de negocio: Actualización exitosa debe modificar el producto
	assert.NoError(t, err, "La actualización debe ser exitosa")
	mockRepo.AssertExpectations(t)
}

func TestProductService_UpdateProduct_PartialUpdate(t *testing.T) {
	// Arrange
	mockRepo := new(MockUpdateRepository)
	service := NewProductService(mockRepo)
	ctx := context.Background()
	productID := primitive.NewObjectID()

	// Solo actualizar descripción (nombre vacío)
	newDescription := "Only description updated"

	expectedUpdate := bson.M{
		"description": newDescription,
	}

	mockRepo.On("UpdateByID", ctx, productID, expectedUpdate).Return(&mongo.UpdateResult{ModifiedCount: 1}, nil)

	// Act
	err := service.UpdateProduct(ctx, productID, "", newDescription)

	// Assert - Regla de negocio: Actualización parcial es válida (solo descripción)
	assert.NoError(t, err, "Actualización parcial debe ser exitosa")
	mockRepo.AssertExpectations(t)
}

func TestProductService_UpdateProduct_EmptyFields(t *testing.T) {
	// Arrange
	mockRepo := new(MockUpdateRepository)
	service := NewProductService(mockRepo)
	ctx := context.Background()
	productID := primitive.NewObjectID()

	expectedUpdate := bson.M{
		"description": "",
	}

	mockRepo.On("UpdateByID", ctx, productID, expectedUpdate).Return(&mongo.UpdateResult{ModifiedCount: 1}, nil)

	// Act
	err := service.UpdateProduct(ctx, productID, "", "")

	// Assert - Regla de negocio: Se permite actualizar con descripción vacía
	assert.NoError(t, err, "Actualización con campos vacíos debe funcionar")
	mockRepo.AssertExpectations(t)
}

func TestProductService_UpdateProduct_NotFound(t *testing.T) {
	// Arrange
	mockRepo := new(MockUpdateRepository)
	service := NewProductService(mockRepo)
	ctx := context.Background()
	productID := primitive.NewObjectID()

	update := bson.M{
		"name":        "New Name",
		"description": "New Description",
	}

	mockRepo.On("UpdateByID", ctx, productID, update).Return(&mongo.UpdateResult{ModifiedCount: 0}, nil)

	// Act
	err := service.UpdateProduct(ctx, productID, "New Name", "New Description")

	// Assert - Regla de negocio: Actualizar producto inexistente no genera error
	assert.NoError(t, err, "No debe fallar si el producto no existe")
	mockRepo.AssertExpectations(t)
}

func TestProductService_UpdateProduct_DatabaseError(t *testing.T) {
	// Arrange
	mockRepo := new(MockUpdateRepository)
	service := NewProductService(mockRepo)
	ctx := context.Background()
	productID := primitive.NewObjectID()

	update := bson.M{
		"name":        "New Name",
		"description": "New Description",
	}

	mockRepo.On("UpdateByID", ctx, productID, update).Return(nil, errors.New("fallo de escritura"))

	// Act
	err := service.UpdateProduct(ctx, productID, "New Name", "New Description")

	// Assert - Regla de negocio: Errores de BD deben propagarse
	assert.Error(t, err, "Debe retornar error de base de datos")
	assert.Contains(t, err.Error(), "escritura")
	mockRepo.AssertExpectations(t)
}
