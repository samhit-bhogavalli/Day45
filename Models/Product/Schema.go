package Product

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ProductName string `json:"product_name"`
	Price       uint   `json:"price"`
	Quantity    uint   `json:"quantity"`
}

func (p *Product) TableName() string {
	return "product"
}
