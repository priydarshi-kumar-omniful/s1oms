package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// OrderStatus represents the possible statuses of an order
type OrderStatus string

const (
	OnHold   OrderStatus = "on_hold"
	NewOrder OrderStatus = "new_order"
)

type Order struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty" csv:"-"` // CSV does not handle ObjectID
	TenantID    string             `bson:"tenant_id" json:"tenant_id" csv:"tenant_id"` // Multi-tenancy
	CustomerID  string             `bson:"customer_id" json:"customer_id" csv:"customer_id"`
	Items       []OrderItem        `bson:"items" json:"items" csv:"-"` // Items as a slice, not a direct CSV column
	TotalAmount float64            `bson:"total_amount" json:"total_amount" csv:"-"` // Derived field, not in CSV
	Status      OrderStatus        `bson:"status" json:"status" csv:"status"`
}

// OrderItem represents an item within an order
type OrderItem struct {
	ProductID   string  `bson:"product_id" json:"product_id" csv:"product_id"`
	SKUID       string  `bson:"sku_id" json:"sku_id" csv:"sku_id"`
	Name        string  `bson:"name" json:"name" csv:"name"`
	Quantity    int     `bson:"quantity" json:"quantity" csv:"quantity"`
	WarehouseID string  `bson:"warehouse_id" json:"warehouse_id" csv:"warehouse_id"` // Reference to warehouse in WMS (PostgreSQL)
	Price       float64 `bson:"price" json:"price" csv:"price"`
}



type KafkaResponseOrderMessage struct {
	OrderItemsID    string `json:"order_items_id"`
	OrderID         string `json:"OrderID"`
	SKUID           string `json:"sku_id"`
	QuantityOrdered int    `json:"quantity_ordered"`
	HubID           string `json:"hub_id"`
	SellerID        string `json:"seller_id"`
}