package Controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/samhit-bhogavalli/Day45/Models/Order"
	"github.com/samhit-bhogavalli/Day45/Models/Retailer"
	"net/http"
)

func GetAllTransactions(c *gin.Context) {
	var orders []Order.Order
	if err := Retailer.GetTransactions(&orders); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, orders)
	}
}
