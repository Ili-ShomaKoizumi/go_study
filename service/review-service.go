package service

import (
	"gin/model"
	"gin/repository"

	"github.com/gin-gonic/gin"
)

type RegisterReviewRequestData struct {
	ProductId string `json:"productId" binding:"required"`
	Title     string `json:"title" binding:"required"`
	UserId    string `json:"userId" binding:"required"`
	UserName  string `json:"userName" binding:"required"`
	SatLevel  int    `json:"SatlLevel" binding:"required"`
	Comment   string `json:"Comment" binding:"required"`
}

type reviewService struct {
	reviewRepository *repository.ReviewRepository
}

func NewReviewService(reviewRepository *repository.ReviewRepository) *reviewService {
	return &reviewService{reviewRepository: reviewRepository}
}

func (reviewService *reviewService) RegisterReview(c *gin.Context) {
	var requestData RegisterReviewRequestData
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var review model.Review
	review.Create(requestData.ProductId, requestData.Title, requestData.UserId, requestData.UserName,
		requestData.SatLevel, requestData.Comment)
	reviewService.reviewRepository.Save(&review)

	c.JSON(201, gin.H{
		"message": "Succesfully Created!",
	})
}
