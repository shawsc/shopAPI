package main

import (
	"net/http"
	"practice/shop/models"
	"practice/shop/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//r.SetTrustedProxies([]string{"192.168.86.63"}) // localhost, change if local ip changes

	models.ConnectDatabase()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	r.GET("/product/:id", routes.GetProduct)
	r.GET("/product", routes.GetAllProducts)

	r.POST("/product", routes.CreateProduct)

	r.Run()
}
