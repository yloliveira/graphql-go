package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Product struct {
	db         *sql.DB
	ID         string
	Title      string
	Price      int
	CategoryID string
}

func NewProduct(db *sql.DB) *Product {
	return &Product{db: db}
}

func (p *Product) Create(title string, price int, categoryID string) (Product, error) {
	id := uuid.New().String()
	_, err := p.db.Exec("INSERT INTO products (id, title, price, category_id) VALUES ($1, $2, $3, $4)", id, title, price, categoryID)

	if err != nil {
		return Product{}, err
	}

	return Product{ID: id, Title: title, Price: price, CategoryID: categoryID}, nil
}
