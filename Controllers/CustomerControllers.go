package Controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/samhit-bhogavalli/Day45/Models/Customer"
	"net/http"
)

func CreateCustomer(c *gin.Context) {
	var customer Customer.Customer
	c.BindJSON(&customer)
	err := Customer.CreateCustomer(&customer)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, customer)
	}
}
