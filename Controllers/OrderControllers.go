package Controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/samhit-bhogavalli/Day45/Models/Order"
	"github.com/samhit-bhogavalli/Day45/Models/Product"
	"net/http"
	"strconv"
)

func PlaceOrder(c *gin.Context) {
	var order Order.Order
	c.BindJSON(&order)
	var product Product.Product
	if er := Product.GetProductById(strconv.Itoa(int(order.ProductID)), &product); er != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}
	order.TotalPrice = order.Quantity*product.Price
	order.Status = "order placed"
	if err := Order.PlaceOrder(&order); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, order)
	}
}

func GetOrderById(c *gin.Context) {
	var order Order.Order
	id := c.Params.ByName("id")
	err := Order.GetOrderById(id, &order)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, order)
	}
}

func GetOrderByCustomerId(c *gin.Context) {
	var orders []Order.Order
	id := c.Params.ByName("id")
	err := Order.GetOrderByCustomerId(id, &orders)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, orders)
	}
}
