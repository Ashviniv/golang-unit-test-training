package services

import (
	"errors"
	"testing"
	mocks2 "testing/mocks/repositories"
	"testing/models"
)

func TestInsert(t *testing.T) {
	repo := mocks2.ProductRepositoryInterface{}
	service := NewProductService(&repo)

	t.Run("success", func(t *testing.T) {
		repo.On("Add",models.Product{
			ID:    "test",
			Name:  "mac",
			Price: 10,
			Stock: 10,
		}).Return(nil).Once()

		err := service.Insert("test", "mac", 10,10)
		if err != nil {
			t.Errorf("Failed to insert %v", err.Error())
		}
	})

	t.Run("failure", func(t *testing.T) {
		repo.On("Add",models.Product{
			ID:    "test",
			Name:  "mac",
			Price: 10,
			Stock: 10,
		}).Return(errors.New("failed to add"))

		err := service.Insert("test", "mac", 10,10)
		if err == nil {
			t.Errorf("Expecting error, got no error")
		}
	})

}
