package models

type Product struct {
	ID        int      `json:"id"`
	UserID    int      `json:"user_id"`
	Name      string   `json:"name"`
	Category  string   `json:"category"`
	Price     int      `json:"price"`
	Condition string   `json:"condition"`
	Images    []string `json:"images"`
}
