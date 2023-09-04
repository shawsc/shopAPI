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

type ProductIdInput struct {
	ID uint `form:"id" json:"id"`
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

func GetProduct(c *gin.Context) {
	var product models.Product
	var input ProductIdInput

	err := c.Bind(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "malformed request " + err.Error()})
		return
	}

	result := models.DB.First(&product, input)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"product": product})
}

func BatchGetProducts(c *gin.Context) {
	var products []models.Product
	var input []ProductIdInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "malformed request " + err.Error()})
		return
	}

	ids := make([]uint, len(input))
	for i, v := range input {
		ids[i] = v.ID
	}

	result := models.DB.Find(&products, ids)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"product": products})
}
