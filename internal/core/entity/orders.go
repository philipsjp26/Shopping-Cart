package entity

import (
	"time"
)

type Order struct {
	Id            int
	TransactionId string
	CustomerId    int
	Status        string
	PaymentMethod string
	OrderDate     time.Time
	OrderItems    []OrderItem // Perubahan: menggunakan slice OrderItem
}

type OrderItem struct {
	OrderId   int
	ProductId int
	Quantity  int
	Price     float64
}
