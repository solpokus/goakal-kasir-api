package repository

import (
	"errors"
	"kasir-api/model" // Adjust import path if module name is different, relying on go.mod 'kasir-api'
	"sync"
)

var (
	ErrCategoryNotFound = errors.New("category not found")
)

type CategoryRepository interface {
	FindAll() ([]model.Category, error)
	FindByID(id int) (model.Category, error)
	Create(category model.Category) (model.Category, error)
	Update(id int, category model.Category) (model.Category, error)
	Delete(id int) error
}

type InMemoryCategoryRepository struct {
	mu         sync.RWMutex
	categories map[int]model.Category
	lastID     int
}

func NewInMemoryCategoryRepository() *InMemoryCategoryRepository {
	return &InMemoryCategoryRepository{
		categories: make(map[int]model.Category),
		lastID:     0,
	}
}

func (r *InMemoryCategoryRepository) FindAll() ([]model.Category, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var categories []model.Category
	for _, c := range r.categories {
		categories = append(categories, c)
	}
	return categories, nil
}

func (r *InMemoryCategoryRepository) FindByID(id int) (model.Category, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	category, ok := r.categories[id]
	if !ok {
		return model.Category{}, ErrCategoryNotFound
	}

	return category, nil
}

func (r *InMemoryCategoryRepository) Create(category model.Category) (model.Category, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.lastID++
	category.ID = r.lastID
	r.categories[category.ID] = category

	return category, nil
}

func (r *InMemoryCategoryRepository) Update(id int, category model.Category) (model.Category, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.categories[id]; !ok {
		return model.Category{}, ErrCategoryNotFound
	}

	category.ID = id
	r.categories[id] = category

	return category, nil
}

func (r *InMemoryCategoryRepository) Delete(id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.categories[id]; !ok {
		return ErrCategoryNotFound
	}

	delete(r.categories, id)
	return nil
}
