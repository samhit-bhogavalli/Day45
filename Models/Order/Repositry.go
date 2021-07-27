package Order

import (
	"fmt"
	"github.com/samhit-bhogavalli/Day45/Config"
	"gorm.io/gorm/clause"
)

func PlaceOrder(order *Order) error {
	if err := Config.DB.Create(order).Error; err != nil {
		return err
	} else {
		WaitQueue.Mutex.Lock()
		WaitQueue.waitArr.PushBack(order)
		WaitQueue.Mutex.Unlock()
		fmt.Printf("order with id %d is placed in the queue\n", order.ID)
		return nil
	}
}

func GetOrderByCustomerId(customerId string, order *[]Order) error {
	if err := Config.DB.Where("customer_id = ?", customerId).Preload(clause.Associations).Find(order).Error; err != nil {
		return err
	}
	return nil
}

func GetOrderById(id string, order *Order) error {
	if err := Config.DB.Where("id = ?", id).Preload(clause.Associations).First(order).Error; err != nil {
		return err
	}
	return nil
}

func UpdateOrder(order *Order) error {
	if err := Config.DB.Save(order).Error; err != nil {
		return err
	}
	return nil
}
