package handlers

import (
	"gorm.io/gorm"
	"github.com/vCif3r/ecommerce-api/internal/models"
	"github.com/vCif3r/ecommerce-api/internal/services"
	"github.com/vCif3r/ecommerce-api/internal/repositories"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type AuthHandler struct {
	service *services.AuthService
}

func NewAuthHandler(service *services.AuthService) *AuthHandler {
	return &AuthHandler{service: service}
}

func RegisterAuthRoutes(router *gin.RouterGroup, db *gorm.DB) {
	authRepo := repositories.NewAuthRepository(db)
	// Replace "your-secret-key" and 24*time.Hour with your actual secret and duration as needed
	authService := services.NewAuthService(authRepo, "your-secret-key", 24*time.Hour)
	authHandler := NewAuthHandler(authService)

	users := router.Group("/auth")
	{
		users.POST("/register", authHandler.Register)
		users.POST("/login", authHandler.Login)
	}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req models.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.service.Register(&req)
	if err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "email already registered" {
			status = http.StatusConflict
		}
		c.JSON(status, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := h.service.Login(&req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}