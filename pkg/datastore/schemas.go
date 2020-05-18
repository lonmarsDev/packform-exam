package datastore

import "time"

//Defines schema for deliveries
type Delivery struct {
	ID                int `gorm:"primary_key"`
	OrderItemID       int
	DeliveredQuantity int
}

//defines schema for order items
type OrderItem struct {
	ID           int `gorm:"primary_key"`
	OrderID      int
	PricePerUnit float64
	Quantity     int
	Product      string
}

//defines schema for orders
type Order struct {
	ID         int `gorm:"primary_key"`
	CreatedAt  time.Time
	OrderName  string
	CustomerID string
}
