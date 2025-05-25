package repositories

import (
	"github.com/vCif3r/ecommerce-api/internal/models"
	"gorm.io/gorm"
)

type ProductImageRepository interface {
	FindAll() ([]models.ProductImage, error)
}

type productImageRepository struct {
	db *gorm.DB
}

func NewProductImageRepository(db *gorm.DB) ProductImageRepository {
	return &productImageRepository{db: db}
}

func (r *productImageRepository) FindAll() ([]models.ProductImage, error) {
	var result []models.ProductImage
	err := r.db.Find(&result).Error
	return result, err
}
