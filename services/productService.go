package services

import (
	"errors"

	"testing/models"
	"testing/repositories"
)

type ProductServiceInterface interface {
	Insert(productID, productName string, price float64, stock int) error
}

type ProductService struct {
	repo repositories.ProductRepositoryInterface
}

func NewProductService(repo repositories.ProductRepositoryInterface) ProductService {
	return ProductService{
		repo: repo,
	}
}

func (s ProductService) Insert(productID, productName string, price float64, stock int) error {
	if len(productID) == 0 {
		return errors.New("productID can not be null")
	}

	err := s.repo.Add(models.Product{
		ID:    productID,
		Name:  productName,
		Price: float64(price),
		Stock: stock,
	})

	if err != nil {
		return err
	}

	return nil
}
