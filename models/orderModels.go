package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// OrderStatus represents the possible statuses of an order
type OrderStatus string

const (
	Pending   OrderStatus = "pending"
	Shipped   OrderStatus = "shipped"
	Delivered OrderStatus = "delivered"
	Canceled  OrderStatus = "canceled"
)

// Order represents an order in MongoDB
type Order struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	TenantID    string             `bson:"tenant_id" ` // Multi-tenancy
	CustomerID  string             `bson:"customer_id"`
	Items       []OrderItem        `bson:"items"`
	TotalAmount float64            `bson:"total_amount"`
	Status      OrderStatus        `bson:"status"`
}

// OrderItem represents an item within an order
type OrderItem struct {
	ProductID   string  ` bson:"product_id"`
	SKUID       string  ` bson:"sku_id"`
	Name        string  ` bson:"name"`
	Quantity    int     ` bson:"quantity"`
	WarehouseID string  ` bson:"warehouse_id"` // Reference to warehouse in WMS (PostgreSQL)
	Price       float64 ` bson:"price"`
}


