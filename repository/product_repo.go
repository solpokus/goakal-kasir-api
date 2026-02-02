package repository

import (
	"database/sql"
	"kasir-api/model"
)

type ProductRepository interface {
	FindAll() ([]model.Product, error)
	FindByID(id int) (model.Product, error)
	Create(product model.Product) (model.Product, error)
	Update(id int, product model.Product) (model.Product, error)
	Delete(id int) error
}

type PostgresProductRepository struct {
	db *sql.DB
}

func NewPostgresProductRepository(db *sql.DB) *PostgresProductRepository {
	return &PostgresProductRepository{db: db}
}

func (r *PostgresProductRepository) FindAll() ([]model.Product, error) {
	rows, err := r.db.Query("SELECT id, name, price, stock, category_id FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []model.Product
	for rows.Next() {
		var p model.Product
		if err := rows.Scan(&p.ID, &p.Name, &p.Price, &p.Stock, &p.CategoryID); err != nil {
			return nil, err
		}
		products = append(products, p)
	}
	return products, nil
}

func (r *PostgresProductRepository) FindByID(id int) (model.Product, error) {
	var p model.Product
	err := r.db.QueryRow("SELECT id, name, price, stock, category_id FROM products WHERE id = $1", id).
		Scan(&p.ID, &p.Name, &p.Price, &p.Stock, &p.CategoryID)
	if err != nil {
		return model.Product{}, err
	}
	return p, nil
}

func (r *PostgresProductRepository) Create(product model.Product) (model.Product, error) {
	err := r.db.QueryRow(
		"INSERT INTO products (name, price, stock, category_id) VALUES ($1, $2, $3, $4) RETURNING id",
		product.Name, product.Price, product.Stock, product.CategoryID,
	).Scan(&product.ID)

	if err != nil {
		return model.Product{}, err
	}
	return product, nil
}

func (r *PostgresProductRepository) Update(id int, product model.Product) (model.Product, error) {
	_, err := r.db.Exec(
		"UPDATE products SET name = $1, price = $2, stock = $3, category_id = $4 WHERE id = $5",
		product.Name, product.Price, product.Stock, product.CategoryID, id,
	)
	if err != nil {
		return model.Product{}, err
	}
	product.ID = id
	return product, nil
}

func (r *PostgresProductRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM products WHERE id = $1", id)
	return err
}
