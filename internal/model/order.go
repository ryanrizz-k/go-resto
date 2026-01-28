package model

type OrderStatus string

type Order struct {
	ID            string         `gorm:"primaryKey" json:"id"`
	UserID        string         `json:"user_id"`
	Status        OrderStatus    `json:"status"`
	ProductOrders []ProductOrder `json:"product_orders"`
	ReferenceID   string         `gorm:"unique" json:"reference_id"`
}

type ProductOrderStatus string

type ProductOrder struct {
	ID         string             `gorm:"primaryKey" json:"id"`
	OrderID    string             `json:"order_id"`
	OrderCode  string             `json:"order_code"`
	Quantity   int                `json:"quantity"`
	TotalPrice int64              `json:"total_price"`
	Status     ProductOrderStatus `json:"status"`
}

type OrderMenuProductRequest struct {
	OrderCode string `json:"order_code"`
	Quantity  int    `json:"quantity"`
}

type OrderMenuRequest struct {
	UserID        string                    `json:"-"`
	OrderProducts []OrderMenuProductRequest `json:"order_products"`
	ReferenceID   string                    `json:"reference_id"`
}

type GetOrderInfoRequest struct {
	UserID  string `json:"-"`
	OrderID string `json:"order_id"`
}
