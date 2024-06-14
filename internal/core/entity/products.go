package entity

import "time"

type (
	Products struct {
		ID          uint      `json:"id,omitempty"`
		CategoryID  int       `json:"category_id,omitempty"`
		Name        string    `json:"name,omitempty"`
		Description *string   `json:"description,omitempty"`
		Quantity    int       `json:"quantity,omitempty"`
		Price       float64   `json:"price,omitempty"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	}
)
