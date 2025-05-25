package handlers

import (
	"net/http"
	"strconv"
	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
	"github.com/vCif3r/ecommerce-api/internal/repositories"
	"github.com/vCif3r/ecommerce-api/internal/services"
)

type CategoryHandler struct {
	service services.CategoryService
}

func NewCategoryHandler(service services.CategoryService) *CategoryHandler {
	return &CategoryHandler{service: service}
}

func RegisterCategoryRoutes(router *gin.RouterGroup, db *gorm.DB) {
	categoryRepo := repositories.NewCategoryRepository(db)
	categoryService := services.NewCategoryService(categoryRepo)
	categoryHandler := NewCategoryHandler(categoryService)

	categories := router.Group("/categories")
	{
		categories.GET("", categoryHandler.GetCategories)
		categories.GET("/:id", categoryHandler.GetCategoryByID)
		categories.GET("/list", categoryHandler.GetCategoriesList)
		categories.GET("/list/tree", categoryHandler.GetCategoriesListTree)
	}
}

func (h *CategoryHandler) GetCategories(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	categories, err := h.service.GetAllCategories(page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": categories,
		"meta": gin.H{
			"page":  page,
			"limit": limit,
		},
	})
}

func (h *CategoryHandler) GetCategoryByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
		return
	}

	category, err := h.service.GetCategoryByID(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if category == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	c.JSON(http.StatusOK, category)	
}


func (h *CategoryHandler) GetCategoriesList(c *gin.Context) {
	categories, err := h.service.GetAllCategoriesList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, categories)
}

func (h *CategoryHandler) GetCategoriesListTree(c *gin.Context) {
	categories, err := h.service.GetAllCategoriesListTree()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, categories)
}