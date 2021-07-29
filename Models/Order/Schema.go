package Order

import (
	"github.com/samhit-bhogavalli/Day45/Models/Product"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	CustomerID uint `json:"customer_id"`
	ProductID  uint `json:"product_id"`
	Product    Product.Product
	Quantity   uint   `json:"quantity"`
	Status     string `json:"status"`
	TotalPrice uint `json:"total_price"`
}

func (o *Order) TableName() string {
	return "order"
}
