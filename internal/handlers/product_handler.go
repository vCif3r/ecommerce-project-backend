package handlers

import (
	"net/http"
	"strconv"
	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
	"github.com/vCif3r/ecommerce-api/internal/repositories"
	"github.com/vCif3r/ecommerce-api/internal/services"
)

type ProductHandler struct {
	service services.ProductService
}

func NewProductHandler(service services.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func RegisterProductRoutes(router *gin.RouterGroup, db *gorm.DB) {
	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepo)
	productHandler := NewProductHandler(productService)

	products := router.Group("/products")
	{
		products.GET("", productHandler.GetProducts)
		products.GET("/:id", productHandler.GetProduct)
		products.GET("/recent", productHandler.GetNewProducts)
		products.GET("/category/:idCategory", productHandler.GetProductsByCategoryID)
		products.GET("/search", productHandler.SearchProduct)
		products.GET("/:id/recommendations", productHandler.GetProductsRecomended)
		//products.POST("", productHandler.CreateProduct)
		//products.PUT("/:id", productHandler.UpdateProduct)
		//products.DELETE("/:id", productHandler.DeleteProduct)
	}
}

func (h *ProductHandler) GetProducts(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	products, err := h.service.GetAllProducts(page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": products,
		"meta": gin.H{
			"page":  page,
			"limit": limit,
		},
	})
}

// GetProduct handles GET /products/:id
func (h *ProductHandler) GetProduct(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	product, err := h.service.GetProductByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, product)
}


func (h *ProductHandler) GetNewProducts(c *gin.Context) {
	products, err := h.service.GetNewProducts()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}

func (h *ProductHandler) GetProductsByCategoryID(c *gin.Context) {
	idParam := c.Param("idCategory")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}

	products, err := h.service.GetProductsByCategory(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, products)
}

func (h *ProductHandler) SearchProduct(c *gin.Context) {
	queryParam := c.DefaultQuery("query", "")
	products, err := h.service.SearchProduct(queryParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, products)
}

func (h *ProductHandler) GetProductsRecomended(c *gin.Context) {
	idProductParam := c.Param("id")

	idProduct, err := strconv.Atoi(idProductParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	products, err := h.service.GetProductsRecomended(uint(idProduct))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, products)
}