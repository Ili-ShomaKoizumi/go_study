package service

import (
	"gin/model"
	"gin/repository"

	"github.com/gin-gonic/gin"
)

type RegisterUserRequestData struct {
	UserId          string `json:"UserId" binding:"required"`
	UserName        string `json:"userName" binding:"required"`
	DisplayName     string `json:"displayName" binding:"required"`
	Email           string `json:"email" binding:"required"`
	ProfileImageUrl string `json:"profileImageUrl"`
	Description     string `json:"description"`
}

type userService struct {
	userRepository *repository.UserRepository
}

func NewUserService(userRepository *repository.UserRepository) *userService {
	return &userService{userRepository: userRepository}
}

func (userService *userService) RegisterUser(c *gin.Context) {
	var requestData RegisterUserRequestData
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var user model.User
	user.Create(requestData.UserId, requestData.UserName, requestData.DisplayName,
		requestData.Email, requestData.ProfileImageUrl, requestData.Description)
	userService.userRepository.Save(&user)

	c.JSON(201, gin.H{
		"message": "Succesfully Created!",
	})
}

func (userService *userService) UpdateUser(c *gin.Context) {
	userId := c.Param("id")
	targetUser, userErr := userService.userRepository.FindById(userId)
	if userErr != nil {
		c.JSON(500, gin.H{"error": userErr.Error()})
		return
	}
	var requestData RegisterUserRequestData
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	targetUser.UserName = requestData.UserName
	targetUser.DisplayName = requestData.DisplayName
	targetUser.Email = requestData.Email
	targetUser.ProfileImageUrl = requestData.ProfileImageUrl
	targetUser.Description = requestData.Description
	userService.userRepository.Update(userId, targetUser)

	c.JSON(201, gin.H{
		"message": "Succesfully Created!",
	})
}

func (userService *userService) DeleteUser(c *gin.Context) {
	userId := c.Param("id")
	userService.userRepository.Delete(&model.User{}, userId)

	c.JSON(201, gin.H{
		"message": "Succesfully Created!",
	})
}

func (userService *userService) GetUsersById(c *gin.Context) {
	userId := c.Param("id")
	foundUser, userErr := userService.userRepository.FindById(userId)
	if userErr != nil {
		c.JSON(500, gin.H{"error": userErr.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "OK",
		"user": map[string]interface{}{
			"id":              foundUser.UserId,
			"username":        foundUser.UserName,
			"displayname":     foundUser.DisplayName,
			"email":           foundUser.Email,
			"profileimageurl": foundUser.ProfileImageUrl,
			"description":     foundUser.Description,
		},
	})
}

func (userService *userService) GetUsers(c *gin.Context) {
	foundUsers, userErr := userService.userRepository.FindAll()
	if userErr != nil {
		c.JSON(500, gin.H{"error": userErr.Error()})
		return
	}

	var userResponseData []map[string]interface{}
	for _, user := range foundUsers {
		responseUser := map[string]interface{}{
			"id":              user.UserId,
			"displayname":     user.DisplayName,
			"profileimageurl": user.ProfileImageUrl,
			"description":     user.Description,
		}
		userResponseData = append(userResponseData, responseUser)
	}
	c.JSON(200, gin.H{
		"message": "OK",
		"user":    userResponseData,
	})
}
