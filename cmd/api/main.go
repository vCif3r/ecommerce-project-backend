package main

import (
	"log"
	"github.com/vCif3r/ecommerce-api/internal/handlers"
	"github.com/vCif3r/ecommerce-api/pkg/config"
	"github.com/vCif3r/ecommerce-api/pkg/database"
	"github.com/vCif3r/ecommerce-api/pkg/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	// Cargar configuración
	cfg := config.LoadConfig()

	// Inicializar base de datos
	db, err := database.NewDB(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Inicializar Redis (opcional para cache)
	// redisClient := cache.NewRedisClient(cfg)

	// Configurar router
	router := gin.Default()

	// Middlewares
	router.Use(middleware.CORS())

	// Rutas
	api := router.Group("/api/v1")
	{
		handlers.RegisterProductRoutes(api, db)
		handlers.RegisterCategoryRoutes(api, db)
		handlers.RegisterProductImageRoutes(api, db)
		handlers.RegisterAuthRoutes(api, db)
		//handlers.RegisterUserRoutes(api, db)
		//handlers.RegisterOrderRoutes(api, db)
		//handlers.RegisterAuthRoutes(api, db, cfg.JWTSecret)
	}

	// Iniciar servidor
	log.Printf("Starting server on port %s", cfg.ServerPort)
	if err := router.Run(":" + cfg.ServerPort); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}