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

func (p *Product) FindAll() ([]Product, error) {
	rows, err := p.db.Query("SELECT id, title, price, category_id FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := []Product{}

	for rows.Next() {
		var id, title, categoryID string
		var price int
		if err := rows.Scan(&id, &title, &price, &categoryID); err != nil {
			return nil, err
		}
		products = append(products, Product{ID: id, Title: title, Price: price, CategoryID: categoryID})
	}

	return products, nil
}

func (p *Product) FindByCategoryID(categoryID string) ([]Product, error) {
	rows, err := p.db.Query("SELECT id, title, price, category_id FROM products WHERE category_id = $1", categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := []Product{}

	for rows.Next() {
		var id, title, categoryID string
		var price int
		if err := rows.Scan(&id, &title, &price, &categoryID); err != nil {
			return nil, err
		}
		products = append(products, Product{ID: id, Title: title, Price: price, CategoryID: categoryID})
	}

	return products, nil
}
