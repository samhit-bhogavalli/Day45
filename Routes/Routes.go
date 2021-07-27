package Routes

import (
	"github.com/gin-gonic/gin"
	"github.com/samhit-bhogavalli/Day45/Controllers"
)

func SetRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/ecommerce-api")
	{
		grp1.POST("/customer/create", Controllers.CreateCustomer)
		grp1.POST("/product", Controllers.CreateProduct)
		grp1.PATCH("/product/:id", Controllers.UpdateProduct)
		grp1.GET("/product/:id", Controllers.GetProductById)
		grp1.GET("/products", Controllers.GetAllProducts)
		grp1.POST("/order", Controllers.PlaceOrder)
		grp1.GET("/order/:id", Controllers.GetOrderById)
		grp1.GET("/customer/history/:id", Controllers.GetOrderByCustomerId)
		grp1.GET("/retailer/history", Controllers.GetAllTransactions)
	}
	return r
}
