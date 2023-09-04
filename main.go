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

	// Product routes
	r.GET("/products/", routes.GetProduct)          // Form input with request param
	r.GET("/products/:id", routes.GetProduct)       // JSON input
	r.GET("/products", routes.BatchGetProducts)     // JSON array of ids
	r.POST("/products", routes.CreateProduct)       // JSON input
	r.PATCH("/products/:id", routes.UpdateProduct)  // JSON input
	r.DELETE("/products/:id", routes.DeleteProduct) // Form input

	// Review routes
	r.GET("/reviews/:id", routes.GetReview)
	r.GET("/reviewsbyproduct/:product_id", routes.GetReviewsByProductID)
	r.POST("/reviews", routes.CreateReview)
	r.PATCH("/reviews/:id", routes.UpdateReview)
	r.DELETE("/reviews/:id", routes.DeleteReview)

	r.Run()
}
