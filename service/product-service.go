package service

import (
	"gin/model"
	"gin/repository"

	"github.com/gin-gonic/gin"
)

type RegisterProductRequestData struct {
	ProductId   string `json:"productId" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	Category    string `json:"category"`
	ImageUrl    string `json:"imageUrl" binding:"required"`
	Price       int    `json:"price" bindinig:"required"`
	Condition   string `json:"condition" binding:"required"`
	OwnerId     string `json:"ownerId" binding:"required"`
}

type productService struct {
	productRepository *repository.ProductRepository
	reviewRepository  *repository.ReviewRepository
}

func NewProductService(productRepository *repository.ProductRepository, reviewRepository *repository.ReviewRepository) *productService {
	return &productService{productRepository: productRepository, reviewRepository: reviewRepository}
}

func (productService *productService) RegisterProduct(c *gin.Context) {
	var requestData RegisterProductRequestData
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var product model.Product
	product.Create(requestData.ProductId, requestData.Title, requestData.Description, requestData.Category,
		requestData.ImageUrl, requestData.Price, requestData.Condition, requestData.OwnerId)
	productService.productRepository.Save(&product)

	c.JSON(201, gin.H{
		"message": "Succesfully Created!",
	})
}

func (productService *productService) UpdateProduct(c *gin.Context) {
	productId := c.Param("id")
	targetProduct, productErr := productService.productRepository.FindById(productId)
	if productErr != nil {
		c.JSON(500, gin.H{"error": productErr.Error()})
		return
	}
	var requestData RegisterProductRequestData
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	targetProduct.Title = requestData.Title
	targetProduct.Description = requestData.Description
	targetProduct.Category = requestData.Category
	targetProduct.ImageUrl = requestData.ImageUrl
	targetProduct.Price = requestData.Price
	targetProduct.Condition = requestData.Condition
	productService.productRepository.Update(productId, targetProduct)

	c.JSON(201, gin.H{
		"message": "Succesfully Created!",
	})
}

func (productService *productService) DeleteProduct(c *gin.Context) {
	productId := c.Param("id")
	productService.productRepository.Delete(productId, &model.Product{})

	c.JSON(201, gin.H{
		"message": "Succesfully Created!",
	})
}

func (productService *productService) GetProducts(c *gin.Context) {
	foundProducts, productErr := productService.productRepository.FindAll()
	if productErr != nil {
		c.JSON(500, gin.H{"error": productErr.Error()})
		return
	}

	var productResponseData []map[string]interface{}
	for _, product := range foundProducts {
		responseProduct := map[string]interface{}{
			"id":          product.ProductId,
			"title":       product.Title,
			"description": product.Description,
			"category":    product.Category,
			"image_url":   product.ImageUrl,
			"price":       product.Price,
			"condition":   product.Condition,
			"owner_id":    product.OwnerId,
		}
		productResponseData = append(productResponseData, responseProduct)
	}

	c.JSON(200, gin.H{
		"message": "OK",
		"product": productResponseData,
	})
}

func (productService *productService) GetProductsById(c *gin.Context) {
	productId := c.Param("id")

	foundProduct, productErr := productService.productRepository.FindById(productId)
	if productErr != nil {
		c.JSON(500, gin.H{"error": productErr.Error()})
		return
	}

	foundReviews, reviewErr := productService.reviewRepository.FindById(productId)
	if reviewErr != nil {
		c.JSON(500, gin.H{"error": reviewErr.Error()})
		return
	}

	var reviewResponseData []map[string]interface{}
	for _, review := range foundReviews {
		responseReview := map[string]interface{}{
			"productId": review.ProductId,
			"title":     review.Title,
			"userId":    review.UserId,
			"userName":  review.UserName,
			"satLevel":  review.SatLevel,
			"comment":   review.Comment,
		}
		reviewResponseData = append(reviewResponseData, responseReview)
	}

	c.JSON(200, gin.H{
		"message": "OK",
		"product": map[string]interface{}{
			"id":          foundProduct.ProductId,
			"title":       foundProduct.Title,
			"description": foundProduct.Description,
			"category":    foundProduct.Category,
			"image_url":   foundProduct.ImageUrl,
			"price":       foundProduct.Price,
			"condition":   foundProduct.Condition,
			"owner_id":    foundProduct.OwnerId,
			"review":      reviewResponseData,
		},
	})
}
