package service

import (
	"gin/model"
	"gin/repository"
	"time"

	"github.com/gin-gonic/gin"
)

type RegisterPurHistoryRequestData struct {
	UserId         string    `json:"userId" binding:"required"`
	ProductId      string    `json:"productId" binding:"required"`
	Title          string    `json:"title" binding:"required"`
	Price          int       `json:"price" binding:"required"`
	Quantity       int       `json:"quantity" binding:"required"`
	Category       string    `json:"category" binding:"required"`
	Condition      string    `json:"condition" binding:"required"`
	DeliveryStatus string    `json:"deliverystatus" binding:"required"`
	OwnerId        string    `json:"ownerid" binding:"required"`
	DtPurchase     time.Time `json:"dtpurchase" binding:"required"`
}

type purHistoryService struct {
	purHistoryRepository *repository.PurHistoryRepository
}

func NewPurHistoryService(purHistoryRepository *repository.PurHistoryRepository) *purHistoryService {
	return &purHistoryService{purHistoryRepository: purHistoryRepository}
}

func (purHistoryService *purHistoryService) RegisterPurHistory(c *gin.Context) {
	var requestData RegisterPurHistoryRequestData
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var purHistory model.PurHistory
	purHistory.Create(requestData.UserId, requestData.ProductId, requestData.Title, requestData.Price,
		requestData.Quantity, requestData.Category, requestData.Condition, requestData.DeliveryStatus,
		requestData.OwnerId)
	purHistoryService.purHistoryRepository.Save(&purHistory)

	c.JSON(201, gin.H{
		"message": "Succesfully Created!",
	})
}

func (purHistoryService *purHistoryService) UpdatePurHistory(c *gin.Context) {
	UserId := c.Param("userId")
	ProductId := c.Param("productId")
	targetPurHistory, purHistoryErr := purHistoryService.purHistoryRepository.FindById(UserId, ProductId)
	if purHistoryErr != nil {
		c.JSON(500, gin.H{"error": purHistoryErr.Error()})
		return
	}
	var requestData RegisterPurHistoryRequestData
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	targetPurHistory.DeliveryStatus = requestData.DeliveryStatus
	purHistoryService.purHistoryRepository.Update(UserId, ProductId, targetPurHistory)

	c.JSON(201, gin.H{
		"message": "Succesfully Created!",
	})
}

func (purHistoryService *purHistoryService) GetPurHistorysById(c *gin.Context) {
	UserId := c.Param("id")
	foundPurHistory, purHistoryErr := purHistoryService.purHistoryRepository.FindById(UserId)
	if purHistoryErr != nil {
		c.JSON(500, gin.H{"error": purHistoryErr.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "OK",
		"purHistory": map[string]interface{}{
			"userId":         foundPurHistory.UserId,
			"productId":      foundPurHistory.ProductId,
			"title":          foundPurHistory.Title,
			"price":          foundPurHistory.Price,
			"quantity":       foundPurHistory.Quantity,
			"category":       foundPurHistory.Category,
			"condition":      foundPurHistory.Condition,
			"deliveryStatus": foundPurHistory.DeliveryStatus,
			"ownerId":        foundPurHistory.OwnerId,
			"dtPurchase":     foundPurHistory.DtPurchase,
		},
	})
}
