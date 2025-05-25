package services

import (
	"github.com/vCif3r/ecommerce-api/internal/models"
	"github.com/vCif3r/ecommerce-api/internal/repositories"
)

type ProductImageService interface {
	GetAllProductImages() ([]models.ProductImage, error)
}

type productImageService struct {
	repo repositories.ProductImageRepository
}

func NewProductImageService(repo repositories.ProductImageRepository) ProductImageService {
	return &productImageService{repo: repo};
}

func (s *productImageService) GetAllProductImages() ([]models.ProductImage, error) {
	return s.repo.FindAll();
}