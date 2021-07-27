package Order

import (
	"container/list"
	"fmt"
	"github.com/samhit-bhogavalli/Day45/Models/Product"
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
	Mutex sync.Mutex
}

var WaitQueue Queue
var CustMap CustomerMap

var ProductMap sync.Map

//implement atomicity
func ProcessOrder(channel chan Order, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	for order := range channel {
		time.Sleep(time.Second * 15)
		var product Product.Product
		flag := false
		err := Product.GetProductById(strconv.Itoa(int(order.ProductID)), &product)
		mutex , ok := ProductMap.Load(int(order.ProductID))
		if !ok {
			ProductMap.Store(int(order.ProductID), &sync.Mutex{})
		}
		mutex , _ =  ProductMap.Load(int(order.ProductID))
		mutex.(*sync.Mutex).Lock()
		if err != nil {
			order.Status = "failed"
		} else {
			if product.Quantity < order.Quantity {
				order.Status = "failed"
			} else {
				flag = true
				//Product.GetProductById(strconv.Itoa(int(product.ID)),&product)
				product.Quantity = product.Quantity - order.Quantity
				order.Status = "success"
				if err = Product.UpdateProduct(strconv.Itoa(int(order.ProductID)), &product); err != nil {
					flag = false
					order.Status = "failed"
				}
			}
		}
		if err := UpdateOrder(&order); err != nil {
			if flag {
				Product.GetProductById(strconv.Itoa(int(product.ID)),&product)
				product.Quantity += order.Quantity
			}
			Product.UpdateProduct(strconv.Itoa(int(order.ProductID)), &product)
		}
		mutex.(*sync.Mutex).Unlock()
		ProductMap.Delete(int(order.ProductID))
		fmt.Printf("order with id %d is executed\n",order.ID)
		CustMap.Mutex.Lock()
		CustMap.Cmap[int(order.CustomerID)] = time.Now()
		CustMap.Mutex.Unlock()
	}
}

func SelectOrder() {
	activeOrders := make(chan Order, 10)
	//time.Sleep(time.Minute * 5)
	fmt.Println("Order Processing Started")
	CustMap.Cmap = make(map[int]time.Time)
	var waitGroup sync.WaitGroup
	for i := 0; i < 10; i++ {
		waitGroup.Add(1)
		go ProcessOrder(activeOrders, &waitGroup)
	}
	go func() {
		for {
			WaitQueue.Mutex.Lock()
			for node := WaitQueue.waitArr.Front(); node != nil; node = node.Next() {
				//fmt.Println(int(node.Value.(*Order).ID))
				placed := time.Now()
				CustMap.Mutex.Lock()
				prevTime := CustMap.Cmap[int(node.Value.(*Order).CustomerID)]
				CustMap.Mutex.Unlock()
				//fmt.Println(placed.Sub(prevTime))
				if placed.Sub(prevTime) >= time.Duration(time.Minute*2) {
					fmt.Printf("order with id %d is inserted in activeOrders\n",int(node.Value.(*Order).ID))
					activeOrders <- *node.Value.(*Order)
					WaitQueue.waitArr.Remove(node)
				}
			}
			WaitQueue.Mutex.Unlock()
			//time.Sleep(time.Second * 10)
		}
	}()
	waitGroup.Wait()
}
