package main

import (
	"github.com/samhit-bhogavalli/Day45/Config"
	"github.com/samhit-bhogavalli/Day45/Models/Customer"
	"github.com/samhit-bhogavalli/Day45/Models/Order"
	"github.com/samhit-bhogavalli/Day45/Models/Product"
	"github.com/samhit-bhogavalli/Day45/Models/Retailer"
	"github.com/samhit-bhogavalli/Day45/Routes"
)

var err error

func main() {
	Config.DBInit()

	Config.DB.AutoMigrate(&Retailer.Retailer{})
	Config.DB.AutoMigrate(&Customer.Customer{})
	Config.DB.AutoMigrate(&Product.Product{})
	Config.DB.AutoMigrate(&Order.Order{})

	r := Routes.SetRouter()
	go Order.SelectOrder()
	//running
	//TODO concurrency handling with mutex
	r.Run("localhost:8080")

}
