package main

import (
	db "ecommerce/config"
	"ecommerce/controller"
	"ecommerce/middleware"
	"ecommerce/repository"
	"ecommerce/service"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize database connection
	db.InitializeDatabase()

	// Create repository, service, and controller
	productRepo := repository.NewProductRepository(db.GetDB())
	productService := service.NewProductService(productRepo)
	productController := controller.NewProductController(productService)

	// Initialize Gin router
	r := gin.Default()

	// Custom Middlewares
	r.Use(middleware.LoggingMiddlewareGin())
	authMiddleware := middleware.AuthMiddlewareGin(db.GetDB())

	// Routes with Authentication Middleware
	authorized := r.Group("/")
	authorized.Use(authMiddleware)
	{
		authorized.POST("/products", productController.CreateProduct)
		authorized.GET("/products/:id", productController.GetProduct)
		authorized.GET("/products", productController.GetAllProducts)
		authorized.PUT("/products/:id", productController.UpdateProduct)
		authorized.DELETE("/products/:id", productController.DeleteProduct)
	}

	// Start server
	r.Run(":8080")
}
