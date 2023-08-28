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
