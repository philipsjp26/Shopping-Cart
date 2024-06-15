package entity

import "time"

type Cart struct {
	Id             int       `json:"id"`
	ProductId      int       `json:"product_id"`
	ShoppingCartId int       `json:"shopping_cart_id"`
	Quantity       int       `json:"quantity"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type ProductsOfCart struct {
	ProductId int
	Name      string
	Stock     int
	Price     float64
	TotalQty  int
}
