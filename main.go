package main

import (
	_ "kasir-api/docs" // Import generated docs
	"kasir-api/handler"
	"kasir-api/repository"
	"log"
	"net/http"

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
	// Initialize Repository
	categoryRepo := repository.NewInMemoryCategoryRepository()

	// Initialize Handler
	categoryHandler := handler.NewCategoryHandler(categoryRepo)

	// Setup Router
	mux := http.NewServeMux()

	// Register Routes
	mux.HandleFunc("GET /categories", categoryHandler.GetCategories)
	mux.HandleFunc("POST /categories", categoryHandler.CreateCategory)
	mux.HandleFunc("GET /categories/{id}", categoryHandler.GetCategory)
	mux.HandleFunc("PUT /categories/{id}", categoryHandler.UpdateCategory)
	mux.HandleFunc("DELETE /categories/{id}", categoryHandler.DeleteCategory)

	// Health Check
	mux.HandleFunc("GET /health", handler.HealthCheck)

	// Swagger
	mux.HandleFunc("GET /swagger/", httpSwagger.WrapHandler)

	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
