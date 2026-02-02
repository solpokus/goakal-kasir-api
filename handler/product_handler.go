package handler

import (
	"encoding/json"
	"kasir-api/model"
	"kasir-api/repository"
	"net/http"
	"strconv"
)

type ProductHandler struct {
	repo repository.ProductRepository
}

func NewProductHandler(repo repository.ProductRepository) *ProductHandler {
	return &ProductHandler{repo: repo}
}

// GetProducts godoc
// @Summary      Get all products
// @Description  Get a list of all products
// @Tags         products
// @Produce      json
// @Success      200  {array}   model.Product
// @Router       /api/produk [get]
func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.repo.FindAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

// GetProduct godoc
// @Summary      Get a product by ID
// @Description  Get details of a specific product
// @Tags         products
// @Produce      json
// @Param        id   path      int  true  "Product ID"
// @Success      200  {object}  model.Product
// @Failure      404  {string}  string  "Product not found"
// @Router       /api/produk/{id} [get]
func (h *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	product, err := h.repo.FindByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

// CreateProduct godoc
// @Summary      Create a new product
// @Description  Create a new product with the input payload
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        product  body      model.Product  true  "Product Content"
// @Success      201       {object}  model.Product
// @Failure      400       {string}  string  "Invalid input"
// @Router       /api/produk [post]
func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product model.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdProduct, err := h.repo.Create(product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdProduct)
}

// UpdateProduct godoc
// @Summary      Update a product
// @Description  Update a product's details by ID
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        id        path      int             true  "Product ID"
// @Param        product  body      model.Product  true  "Product Content"
// @Success      200       {object}  model.Product
// @Failure      400       {string}  string  "Invalid input or ID"
// @Router       /api/produk/{id} [put]
func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	var product model.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedProduct, err := h.repo.Update(id, product)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedProduct)
}

// DeleteProduct godoc
// @Summary      Delete a product
// @Description  Delete a product by ID
// @Tags         products
// @Produce      json
// @Param        id   path      int  true  "Product ID"
// @Success      204  {object}  nil
// @Failure      400  {string}  string  "Invalid ID"
// @Router       /api/produk/{id} [delete]
func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	if err := h.repo.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
