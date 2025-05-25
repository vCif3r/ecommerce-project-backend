package repositories

import (
	"github.com/vCif3r/ecommerce-api/internal/models"
	"gorm.io/gorm"
)

type ProductRepository interface {
	FindAll(page, limit int) ([]models.Product, error)
	FindByID(id uint) (*models.Product, error)
	Create(product *models.Product) error
	Update(product *models.Product) error
	Delete(id uint) error
	Search(query string, page, limit int) ([]models.Product, error)
	GetNewProducts() ([]models.Product, error)
	GetProductsByCategory(categoryID uint) ([]models.Product, error)
	SearchProduct(query string) ([]models.Product, error)
	GetProductsRecomended(idProduct uint) ([]models.Product, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) FindAll(page, limit int) ([]models.Product, error) {
	var products []models.Product
	offset := (page - 1) * limit

	err := r.db.Preload("Category").Preload("Images").
		Offset(offset).Limit(limit).
		Find(&products).Error

	return products, err
}

func (r *productRepository) FindByID(id uint) (*models.Product, error) {
	var product models.Product
	err := r.db.Preload("Category").Preload("Images").First(&product, id).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *productRepository) Create(product *models.Product) error {
	return r.db.Create(product).Error
}

func (r *productRepository) Update(product *models.Product) error {
	return r.db.Save(product).Error
}

func (r *productRepository) Delete(id uint) error {
	return r.db.Delete(&models.Product{}, id).Error
}

func (r *productRepository) Search(query string, page, limit int) ([]models.Product, error) {
	var products []models.Product
	offset := (page - 1) * limit
	err := r.db.Preload("Category").Preload("Images").
		Where("name LIKE ?", "%"+query+"%").
		Offset(offset).Limit(limit).
		Find(&products).Error
	return products, err
}


func (r *productRepository) GetNewProducts() ([]models.Product, error) {
	var products []models.Product
	err := r.db.Preload("Category").Preload("Images").Order("created_at DESC").Limit(10).Find(&products).Error
	return products, err
}

// func (r *productRepository) GetProductsByCategoryID(idCategory uint, page int, limit int) ([]models.Product, error) {
// 	var products []models.Product
// 	offset := (page -1) * limit
// 	err := r.db.Preload("Category").
// 	Where("category_id = ?", idCategory).
// 	Offset(offset).Limit(limit).
// 	Find(&products).Error
// 	return products, err
// }

func (r *productRepository) GetProductsByCategory(categoryID uint) ([]models.Product, error) {
    var products []models.Product

    // Primero obtenemos la categoría y sus descendientes
    var categoryIDs []uint
    if err := r.db.Model(&models.Category{}).
        Where("id = ? OR parent_id = ?", categoryID, categoryID).
        Pluck("id", &categoryIDs).Error; err != nil {
        return nil, err
    }

    // Luego obtenemos los productos de esas categorías
    err := r.db.Where("category_id IN ?", categoryIDs).
        Preload("Category").
        Preload("Images").
        Find(&products).Error

    return products, err
}

func (r *productRepository) SearchProduct(query string) ([]models.Product, error) {
	var products []models.Product
	err := r.db.
		Joins("JOIN categories ON categories.id = products.category_id").
		Where("categories.name ILIKE ? OR products.name ILIKE ?", query+"%", query+"%").
		Preload("Category").Preload("Images").
		Limit(10).
		Find(&products).Error
	return products, err
}

func (r *productRepository) GetProductsRecomended(idProduct uint) ([]models.Product, error) {
	var product models.Product

	// Buscar el producto original
	if err := r.db.First(&product, idProduct).Error; err != nil {
		return nil, err
	}

	// Buscar otros productos de la misma categoría (excluyendo el original)
	var products []models.Product
	err := r.db.Preload("Category").Preload("Images").
		Where("category_id = ? AND id != ?", product.CategoryID, idProduct).
		Limit(10).
		Find(&products).Error

	return products, err
}
