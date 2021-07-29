package Customer

import (
	"github.com/samhit-bhogavalli/Day45/Models/Order"
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	Name   string        `json:"name"`
	Orders []Order.Order `json:"orders"`
}

func (c *Customer) TableName() string {
	return "customer"
}
