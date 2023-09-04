package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"practice/shop/models"
)

type CreateProductInput struct {
	Name        string  `json:"name" binding:"required"`
	Brand       string  `json:"brand" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Category    string  `json:"category"`
	Price       float32 `json:"price" binding:"required"`
}

type GetProductInput struct {
	ID uint `form:"id" json:"id"`
}

type UpdateProductInput struct {
	ID          uint    `json:"id" binding:"required"`
	Name        string  `json:"name"`
	Brand       string  `json:"brand"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Price       float32 `json:"price"`
}

func CreateProduct(c *gin.Context) {
	var input CreateProductInput

	// Validate input
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "malformed request " + err.Error()})
		return
	}

	// Copy input into product model
	product := models.Product{
		Name:        input.Name,
		Brand:       input.Brand,
		Description: input.Description,
		Category:    input.Category,
		Price:       input.Price,
	}

	// Insert product into database
	result := models.DB.Create(&product)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": product})
}

func GetProduct(c *gin.Context) {
	var product models.Product
	var input GetProductInput

	// Validate input
	err := c.Bind(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "malformed request " + err.Error()})
		return
	}

	// Fetch ID in database
	result := models.DB.First(&product, input)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"product": product})
}

func BatchGetProducts(c *gin.Context) {
	var products []models.Product
	var input []GetProductInput

	// Validate input
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "malformed request " + err.Error()})
		return
	}

	// Copy IDs from array of JSON objects to a slice of uints
	ids := make([]uint, len(input))
	for i, v := range input {
		ids[i] = v.ID
	}

	// Fetch products from database
	result := models.DB.Find(&products, ids)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"product": products})
}

func UpdateProduct(c *gin.Context) {
	var product models.Product
	var input UpdateProductInput

	// Validate input
	err := c.Bind(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "malformed request " + err.Error()})
		return
	}

	// Get product, if it exists
	result := models.DB.First(&product, input.ID)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product not found."})
		return
	}

	// Update record
	models.DB.Model(&product).Updates(input)

	c.JSON(http.StatusOK, gin.H{"update": "success"})
}

func DeleteProduct(c *gin.Context) {
	var product models.Product
	var input GetProductInput

	// Validate input
	err := c.Bind(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "malformed request " + err.Error()})
		return
	}

	// Get product, if it exists
	result := models.DB.First(&product, input.ID)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Product not found."})
		return
	}

	// Delete record
	models.DB.Delete(&product)

	c.JSON(http.StatusOK, gin.H{"delete": "success"})
}
