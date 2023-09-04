package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"practice/shop/models"
)

type ProductInput struct {
	Name        string  `json:"name" binding:"required"`
	Brand       string  `json:"brand" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Category    string  `json:"category"`
	Price       float32 `json:"price" binding:"required"`
}

func CreateProduct(c *gin.Context) {
	var input ProductInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "malformed request " + err.Error()})
		return
	}

	product := models.Product{
		Name:        input.Name,
		Brand:       input.Brand,
		Description: input.Description,
		Category:    input.Category,
		Price:       input.Price,
	}

	result := models.DB.Create(&product)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": product})
}

func GetAllProducts(c *gin.Context) {
	var product []models.Product

	result := models.DB.Find(&product)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": product})
}

func GetProduct(c *gin.Context) {

}

func BatchGetProducts(c *gin.Context) {

}
