package repository

import (
	"gin/model"

	"gorm.io/gorm"
)

type PurHistoryRepository struct {
	DB *gorm.DB
}

func NewPurHistoryRepository(db *gorm.DB) *PurHistoryRepository {
	return &PurHistoryRepository{DB: db}
}

func (purHistoryRepository *PurHistoryRepository) Save(purHistory *model.PurHistory) error {
	return purHistoryRepository.DB.Create(purHistory).Error
}

func (purHistoryRepository *PurHistoryRepository) Update(userId string, productId string, purHistory *model.PurHistory) error {
	return purHistoryRepository.DB.Where("user_id=?", userId).Where("product_id=?", productId).Save(purHistory).Error
}

func (purHistoryRepository *PurHistoryRepository) FindById(userId string, productId ...string) (purHistory *model.PurHistory, err error) {
	if productId != nil {
		if err = purHistoryRepository.DB.Where("product_id=?", productId).Where("user_id=?", userId).First(&purHistory).Error; err != nil {
			return
		}
	} else {
		if err = purHistoryRepository.DB.Where("user_id=?", userId).First(&purHistory).Error; err != nil {
			return
		}
	}
	return
}
