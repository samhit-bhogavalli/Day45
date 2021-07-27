package Retailer

import (
	"github.com/samhit-bhogavalli/Day45/Config"
	"github.com/samhit-bhogavalli/Day45/Models/Order"
	"gorm.io/gorm/clause"
)

func GetTransactions(orders *[]Order.Order) error {
	if err := Config.DB.Preload(clause.Associations).Find(orders).Error; err != nil {
		return err
	}
	return nil
}
