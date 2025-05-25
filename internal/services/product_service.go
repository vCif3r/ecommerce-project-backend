package services

import (
	"github.com/vCif3r/ecommerce-api/internal/models"
	"github.com/vCif3r/ecommerce-api/internal/repositories"
)

type ProductService interface {
	GetAllProducts(page, limit int) ([]models.Product, error)
	GetProductByID(id uint) (*models.Product, error)
	//CreateProduct(product *models.Product) error
	//UpdateProduct(id uint, product *models.Product) error
	//DeleteProduct(id uint) error
	//SearchProducts(query string, page, limit int) ([]models.Product, error)
	GetNewProducts() ([]models.Product, error)
	GetProductsByCategory(categoryID uint) ([]models.Product, error) 
	SearchProduct(query string) ([]models.Product, error) 
	GetProductsRecomended(idProduct uint) ([]models.Product, error)
}

type productService struct {
	repo repositories.ProductRepository
}

func NewProductService(repo repositories.ProductRepository) ProductService {
	return &productService{repo: repo}
}

func (s *productService) GetAllProducts(page, limit int) ([]models.Product, error) {
	return s.repo.FindAll(page, limit)
}

func (s *productService) GetProductByID(id uint) (*models.Product, error) {
	return s.repo.FindByID(id)
}

func (s *productService) GetNewProducts() ([]models.Product, error) {
	return s.repo.GetNewProducts()
}

func (s *productService) GetProductsByCategory(categoryID uint) ([]models.Product, error) {
	return s.repo.GetProductsByCategory(categoryID)
}

func (s *productService) SearchProduct(query string) ([]models.Product, error) {
	return s.repo.SearchProduct(query)
}

func (s *productService) GetProductsRecomended(idProduct uint) ([]models.Product, error) {
	return s.repo.GetProductsRecomended(idProduct)
}