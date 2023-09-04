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

	r.GET("/product", routes.GetProduct)        // Form input with request param
	r.GET("/product/:id", routes.GetProduct)    // JSON input
	r.GET("/products", routes.BatchGetProducts) // JSON array of ids

	r.POST("/product", routes.CreateProduct) // JSON input

	r.PATCH("/product/:id", routes.UpdateProduct) // JSON input

	r.DELETE("/product/:id", routes.DeleteProduct) // Form input

	r.Run()
}
