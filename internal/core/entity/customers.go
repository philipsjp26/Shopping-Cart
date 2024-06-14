package entity

import (
	"time"
)

type Customers struct {
	Id        uint      `json:"id"`
	Email     string    `json:"email"`
	Phone     *string   `json:"phone,omitempty"`
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CustomerCartProducts struct {
	ShoppingCartId int
	Name           string
	Description    string
	Quantity       int
	Price          float64
	CreatedAt      time.Time
}
