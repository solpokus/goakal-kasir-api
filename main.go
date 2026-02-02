package main

import (
	"log"
	"net/http"

	"kasir-api/config"
	"kasir-api/database"
	_ "kasir-api/docs" // Import generated docs
	"kasir-api/handler"
	"kasir-api/repository"

	httpSwagger "github.com/swaggo/http-swagger"
)

// @title           Kasir API
// @version         1.0
// @description     A simple API for managing categories in a Point of Sale system.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /

func main() {
	// Load Configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Connect to Database
	db, err := database.Connect(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Initialize Repositories
	categoryRepo := repository.NewInMemoryCategoryRepository()
	productRepo := repository.NewPostgresProductRepository(db)

	// Initialize Handlers
	categoryHandler := handler.NewCategoryHandler(categoryRepo)
	productHandler := handler.NewProductHandler(productRepo)

	// Setup Router
	mux := http.NewServeMux()

	// Register Routes - Category
	mux.HandleFunc("GET /categories", categoryHandler.GetCategories)
	mux.HandleFunc("POST /categories", categoryHandler.CreateCategory)
	mux.HandleFunc("GET /categories/{id}", categoryHandler.GetCategory)
	mux.HandleFunc("PUT /categories/{id}", categoryHandler.UpdateCategory)
	mux.HandleFunc("DELETE /categories/{id}", categoryHandler.DeleteCategory)

	// Register Routes - Product
	mux.HandleFunc("GET /api/produk", productHandler.GetProducts)
	mux.HandleFunc("POST /api/produk", productHandler.CreateProduct)
	mux.HandleFunc("GET /api/produk/{id}", productHandler.GetProduct)
	mux.HandleFunc("PUT /api/produk/{id}", productHandler.UpdateProduct)
	mux.HandleFunc("DELETE /api/produk/{id}", productHandler.DeleteProduct)

	// Health Check
	mux.HandleFunc("GET /health", handler.HealthCheck)

	// Swagger
	mux.HandleFunc("GET /swagger/", httpSwagger.WrapHandler)

	log.Printf("Server starting on path :%s", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, mux); err != nil {
		log.Fatal(err)
	}
}
