package services

import (
	"kasir-api/models"
	"kasir-api/repositories"
)

type CategoryService struct {
	repo repositories.CategoryRepository
}

func NewCategoryService(repo *repositories.CategoryRepository) *CategoryService {
	return &CategoryService{repo: *repo}
}

func (s *CategoryService) GetCategories() ([]models.Category, error) {
	// panggil fungsi GetCategories pada Repository
	return s.repo.GetCategories()
}

func (s *CategoryService) CreateCategory(category *models.Category) error {
	// panggil fungsi CreateCategory pada Repository
	return s.repo.CreateCategory(category)
}

func (s *CategoryService) GetCategoryById(id int) (*models.Category, error) {
	// panggil fungsi GetCategoryById pada Repository
	return s.repo.GetCategoryById(id)
}

func (s *CategoryService) UpdateCategory(category *models.Category) error {
	// panggil fungsi UpdateCategory pada Repository
	return s.repo.UpdateCategory(category)
}

func (s *CategoryService) DeleteCategory(id int) error {
	// panggil fungsi DeleteCategory pada Repository
	return s.repo.DeleteCategory(id)
}
