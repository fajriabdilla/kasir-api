package services

import (
	"kasir-api/models"
	"kasir-api/repositories"
)

type ProductService struct {
	repo *repositories.ProductRepository
}

func NewProductService(repo *repositories.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) GetProducts(name string) ([]models.Product, error) {
	// ini adalah fungsi yang ada di Repository Product
	// repositories/product_repository.go
	return s.repo.GetProducts(name)
}

// buat fungsi CreateProduct
// kirim argument struct nya
func (s *ProductService) CreateProduct(data *models.Product) error {
	// panggil fungsi CreateProduct pada Repository // repositories/product_repository.go
	// kirim datanya
	return s.repo.CreateProduct(data)
}

func (s *ProductService) GetProductById(id int) (*models.Product, error) {
	return s.repo.GetProductById(id)
}

func (s *ProductService) UpdateProduct(product *models.Product) error {
	return s.repo.UpdateProduct(product)
}

func (s *ProductService) DeleteProduct(id int) error {
	// panggil fungsi DeleteProduct pada Repository / repositories/product_repository.go
	return s.repo.DeleteProduct(id)
}
