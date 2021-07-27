package Controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/samhit-bhogavalli/Day45/Models/Product"

	"net/http"
)

func CreateProduct(c *gin.Context) {
	var product Product.Product
	c.BindJSON(&product)
	err := Product.CreateProduct(&product)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, product)
	}
}

func UpdateProduct(c *gin.Context) {
	var product Product.Product
	id := c.Params.ByName("id")
	err := Product.GetProductById(id, &product)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}
	c.BindJSON(&product)
	err = Product.UpdateProduct(id, &product)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, product)
	}
}

func GetProductById(c *gin.Context) {
	var product Product.Product
	id := c.Params.ByName("id")
	err := Product.GetProductById(id, &product)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, product)
	}
}

func GetAllProducts(c *gin.Context) {
	var products []Product.Product
	err := Product.GetAllProducts(&products)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, products)
	}
}
