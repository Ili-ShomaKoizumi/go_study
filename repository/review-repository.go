package repository

import (
	"gin/model"

	"gorm.io/gorm"
)

type ReviewRepository struct {
	DB *gorm.DB
}

func NewReviewRepository(db *gorm.DB) *ReviewRepository {
	return &ReviewRepository{DB: db}
}

func (reviewRepository *ReviewRepository) Save(review *model.Review) error {
	return reviewRepository.DB.Create(review).Error
}

func (reviewRepository *ReviewRepository) FindById(id string) (review []*model.Review, err error) {
	if err = reviewRepository.DB.Where("product_id=?", id).Find(&review).Error; err != nil {
		return
	}
	return
}
