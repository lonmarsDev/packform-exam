package datastore

import "time"

//defines  models for result
type Result struct {
	OrderName         string
	CreatedAt         time.Time
	CustomerID        string
	PricePerUnit      float64
	Quantity          int
	Product           string
	DeliveredQuantity int
}
