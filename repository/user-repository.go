package repository

import (
	"gin/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (userRepository *UserRepository) Save(user *model.User) error {
	return userRepository.DB.Create(user).Error
}

func (userRepository *UserRepository) Update(id string, user *model.User) error {
	return userRepository.DB.Where("user_id=?", id).Save(user).Error
}

func (userRepository *UserRepository) Delete(user *model.User, id string) error {
	return userRepository.DB.Where("user_id=?", id).Delete(&user).Error
}

func (userRepository *UserRepository) FindById(id string) (user *model.User, err error) {
	if err = userRepository.DB.Where("user_id=?", id).First(&user).Error; err != nil {
		return
	}
	return
}

func (userRepository *UserRepository) FindAll() (user []*model.User, err error) {
	if err = userRepository.DB.Find(&user).Error; err != nil {
		return
	}
	return
}
