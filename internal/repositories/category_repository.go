package repositories

import (
	"github.com/vCif3r/ecommerce-api/internal/models"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	FindAll(page, limit int) ([]models.Category, error)
	GetAllCategoriesList() ([]models.Category, error) 
	GetAllCategoriesListTree() ([]models.Category, error) 
	FindByID(id uint) (*models.Category, error) 
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db: db}
}


// métodos
func (r *categoryRepository) FindAll(page, limit int) ([]models.Category, error) {
	var categories []models.Category
	offset := (page - 1) * limit

	err := r.db.
		Offset(offset).Limit(limit).
		Find(&categories).Error

	return categories, err
}

func (r *categoryRepository) FindByID(id uint) (*models.Category, error) {
	var category models.Category
	err := r.db.Preload("Children").Preload("Parent").First(&category, id).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (r *categoryRepository) GetAllCategoriesList() ([]models.Category, error) {
	var categories []models.Category

	err := r.db.
		Preload("Children").
		Where("parent_id IS NULL").
		Order("name ASC").
		Find(&categories).Error

	return categories, err
}

func (r *categoryRepository) GetAllCategoriesListTree() ([]models.Category, error) {
	var categories []models.Category

	err := r.db.
		Preload("Children", func(db *gorm.DB) *gorm.DB {
			return db.Preload("Children")
		}).
		Where("parent_id IS NULL").
		Order("name ASC").
		Find(&categories).Error

	return categories, err
}

