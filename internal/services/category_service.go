package services

import (
	"github.com/vCif3r/ecommerce-api/internal/models"
	"github.com/vCif3r/ecommerce-api/internal/repositories"
)

type CategoryService interface {
	GetAllCategories(page, limit int) ([]models.Category, error)
	GetCategoryByID(id uint) (*models.Category, error) 
	GetAllCategoriesList() ([]models.Category, error)
	GetAllCategoriesListTree() ([]models.Category, error)
}
type categoryService struct {
	repo repositories.CategoryRepository
}
func NewCategoryService(repo repositories.CategoryRepository) CategoryService {
	return &categoryService{repo: repo}
}

// metodos
func (s *categoryService) GetAllCategories(page, limit int) ([]models.Category, error) {
	return s.repo.FindAll(page, limit)
}

func (s *categoryService) GetCategoryByID(id uint) (*models.Category, error) {
	return s.repo.FindByID(id)
}

func (s *categoryService) GetAllCategoriesList() ([]models.Category, error) {
	return s.repo.GetAllCategoriesList()
}

func (s *categoryService) GetAllCategoriesListTree() ([]models.Category, error) {
	return s.repo.GetAllCategoriesListTree()
}