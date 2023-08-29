package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Category struct {
	db    *sql.DB
	ID    string
	Title string
}

func NewCategory(db *sql.DB) *Category {
	return &Category{db: db}
}

func (c *Category) Create(title string) (Category, error) {
	id := uuid.New().String()
	_, err := c.db.Exec("INSERT INTO categories (id, title) VALUES ($1, $2)", id, title)

	if err != nil {
		return Category{}, err
	}

	return Category{ID: id, Title: title}, nil
}

func (c *Category) FindAll() ([]Category, error) {
	rows, err := c.db.Query("SELECT id, title FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	categories := []Category{}

	for rows.Next() {
		var id, title string
		if err := rows.Scan(&id, &title); err != nil {
			return nil, err
		}
		categories = append(categories, Category{ID: id, Title: title})
	}

	return categories, nil
}

func (c *Category) FindByProductID(productID string) (Category, error) {
	var id, title string
	err := c.db.QueryRow("SELECT c.id, c.title FROM categories c JOIN products p ON c.id = p.category_id WHERE p.id = $1", productID).Scan(&id, &title)
	if err != nil {
		return Category{}, err
	}

	return Category{ID: id, Title: title}, nil
}
