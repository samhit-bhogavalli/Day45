package Order

import (
	"container/list"
	"fmt"
	"github.com/samhit-bhogavalli/Day45/Config"
	"github.com/samhit-bhogavalli/Day45/Models/Product"
	"gorm.io/gorm"
	"strconv"
	"sync"
	"time"
)

type Queue struct {
	Mutex   sync.Mutex
	waitArr list.List
}

type CustomerMap struct {
	Cmap  map[int]time.Time
	ACmap map[int]int
	Mutex sync.Mutex
}

var WaitQueue Queue
var CustMap CustomerMap

func ProcessOrder(activeOrders chan Order, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	for order := range activeOrders {
		//time for processing order
		time.Sleep(time.Second * 15)
		var product Product.Product
		err := Config.DB.Transaction(func(tx *gorm.DB) error {
			if err := tx.Where("id = ?", strconv.Itoa(int(order.ProductID))).First(&product).Error; err != nil {
				return err
			}
			if product.Quantity < order.Quantity {
				order.Status = "failed"
			} else {
				order.Status = "success"
				product.Quantity -= order.Quantity
			}

			if err := tx.Save(&product).Error; err != nil {
				return err
			}
			if err := tx.Save(&order).Error; err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			order.Status = "failed"
			UpdateOrder(&order)
		}
		fmt.Printf("order with id %d is executed\n",order.ID)
		CustMap.Mutex.Lock()
		CustMap.Cmap[int(order.CustomerID)] = time.Now()
		CustMap.ACmap[int(order.CustomerID)] = 0
		CustMap.Mutex.Unlock()
	}
}

func SelectOrder() {
	activeOrders := make(chan Order, 10)
	fmt.Println("Order Processing Started")
	CustMap.Cmap = make(map[int]time.Time)
	CustMap.ACmap = make(map[int]int)
	var waitGroup sync.WaitGroup
	for i := 0; i < 10; i++ {
		waitGroup.Add(1)
		go ProcessOrder(activeOrders, &waitGroup)
	}
	go func() {
		for {
			for node := WaitQueue.waitArr.Front(); node != nil; node = node.Next() {
				current := time.Now()
				CustMap.Mutex.Lock()
				prevTime := CustMap.Cmap[int(node.Value.(*Order).CustomerID)]
				CustMap.Mutex.Unlock()
				if current.Sub(prevTime) >= time.Duration(time.Minute*2) && CustMap.ACmap[int(node.Value.(*Order).CustomerID)] == 0 {
					CustMap.ACmap[int(node.Value.(*Order).CustomerID)] = 1
					fmt.Printf("order with id %d is inserted in activeOrders\n",int(node.Value.(*Order).ID))
					activeOrders <- *node.Value.(*Order)
					WaitQueue.waitArr.Remove(node)
				}
			}
			time.Sleep(time.Second * 10)
		}
	}()
	waitGroup.Wait()
}
