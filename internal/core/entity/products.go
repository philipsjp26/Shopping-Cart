package entity

import "time"

type (
	Products struct {
		ID          uint      `json:"id"`
		CategoryID  int       `json:"category_id"`
		Name        string    `json:"name"`
		Description *string   `json:"description"`
		Quantity    int       `json:"quantity"`
		Price       float64   `json:"price"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	}
)
