// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type NewCategory struct {
	Title string `json:"title"`
}

type NewProduct struct {
	Title      string `json:"title"`
	Price      int    `json:"price"`
	CategoryID string `json:"categoryId"`
}
