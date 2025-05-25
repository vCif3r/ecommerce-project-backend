package handlers

import (
	"net/http"
	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
	"github.com/vCif3r/ecommerce-api/internal/repositories"
	"github.com/vCif3r/ecommerce-api/internal/services"
)

type ProductImageHandler struct {
	service services.ProductImageService
}

func NewProductImageHandler(service services.ProductImageService) *ProductImageHandler {
	return &ProductImageHandler{service: service}
}

func RegisterProductImageRoutes(router *gin.RouterGroup, db *gorm.DB) {
	productImageRepo := repositories.NewProductImageRepository(db)
	productImageService := services.NewProductImageService(productImageRepo)
	productImageHandler := NewProductImageHandler(productImageService)

	productsImages := router.Group("/products/images")
	{
		productsImages.GET("", productImageHandler.GetProductsImages)
	}
}

func (h *ProductImageHandler) GetProductsImages(c *gin.Context) {
	productsImages, err := h.service.GetAllProductImages()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, productsImages)
}