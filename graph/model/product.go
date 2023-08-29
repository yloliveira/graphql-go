package model

type Product struct {
	ID       string    `json:"id"`
	Title    string    `json:"title"`
	Price    int       `json:"price"`
	Category *Category `json:"category"`
}
