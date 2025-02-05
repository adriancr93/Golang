package services

import (
	"microservice-go/internal/models"
	"microservice-go/internal/repositories"
)

func GetAllProducts() ([]models.Product, error) {
	return repositories.GetAllProducts()
}

func GetProductByID(id uint) (models.Product, error) {
	return repositories.GetProductByID(id)
}

func CreateProduct(product *models.Product) error {
	return repositories.CreateProduct(product)
}

func UpdateProduct(product *models.Product) error {
	return repositories.UpdateProduct(product)
}

func DeleteProduct(id uint) error {
	return repositories.DeleteProduct(id)
}