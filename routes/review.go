package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"practice/shop/models"
)

type CreateReviewInput struct {
	ProductID   uint   `json:"product_id" binding:"required"`
	Rating      uint   `json:"rating" binding:"required"`
	Author      string `json:"author" binding:"required"`
	Description string `json:"description" binding:"required"`
	Date        string `json:"date" binding:"required"`
}

type GetReviewInput struct {
	ID uint `form:"id" json:"id"`
}

type GetReviewsByProductInput struct {
	ProductID uint `form:"product_id" json:"product_id"`
}

type UpdateReviewInput struct {
	ID          uint   `json:"id" binding:"required"`
	ProductID   uint   `json:"product_id"`
	Rating      uint   `json:"rating"`
	Author      string `json:"author"`
	Description string `json:"description"`
	Date        string `json:"date"`
}

func CreateReview(c *gin.Context) {
	var input CreateReviewInput

	// Validate input
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "malformed request " + err.Error()})
		return
	}

	// Copy input into review model
	review := models.Review{
		ProductID:   input.ProductID,
		Rating:      input.Rating,
		Author:      input.Author,
		Description: input.Description,
		Date:        input.Date,
	}

	// Validate product exists
	var product models.Product
	err = models.DB.First(&product, input.ProductID).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "product does not exist, " + err.Error()})
		return
	}

	// Insert product into database
	err = models.DB.Create(&review).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": review})
}

func GetReview(c *gin.Context) {
	var review models.Review
	var input GetReviewInput

	// Validate input
	err := c.Bind(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "malformed request " + err.Error()})
		return
	}

	// Fetch ID in database
	err = models.DB.First(&review, input).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"review": review})
}

func GetReviewsByProductID(c *gin.Context) {
	var reviews []models.Review
	var product models.Product
	var input GetReviewsByProductInput

	// Validate input
	err := c.ShouldBindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "malformed request " + err.Error()})
		return
	}

	// Validate product exists
	err = models.DB.First(&product, input.ProductID).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "product does not exist, " + err.Error()})
		return
	}

	// Fetch reviews from database
	err = models.DB.Where("product_id = ?", input.ProductID).Find(&reviews).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"reviews": reviews})
}

func UpdateReview(c *gin.Context) {
	var review models.Review
	var input UpdateReviewInput

	// Validate input
	err := c.Bind(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "malformed request " + err.Error()})
		return
	}

	// Get review, if it exists
	err = models.DB.First(&review, input.ID).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Review not found."})
		return
	}

	// Update record
	models.DB.Model(&review).Updates(input)

	c.JSON(http.StatusOK, gin.H{"update": "success"})
}

func DeleteReview(c *gin.Context) {
	var review models.Review
	var input GetReviewInput

	// Validate input
	err := c.Bind(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "malformed request " + err.Error()})
		return
	}

	// Get review, if it exists
	err = models.DB.First(&review, input.ID).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Review not found."})
		return
	}

	// Delete record
	models.DB.Delete(&review)

	c.JSON(http.StatusOK, gin.H{"delete": "success"})
}
