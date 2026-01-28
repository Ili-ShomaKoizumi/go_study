package repository

import (
	"gin/model"

	"gorm.io/gorm"
)

type ProductRepository struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{DB: db}
}

func (productRepository *ProductRepository) Save(product *model.Product) error {
	return productRepository.DB.Create(product).Error
}

func (productRepository *ProductRepository) Update(id string, product *model.Product) error {
	return productRepository.DB.Where("product_id=?", id).Save(product).Error
}

func (productRepository *ProductRepository) Delete(id string, product *model.Product) error {
	return productRepository.DB.Where("product_id=?", id).Delete(product).Error
}

func (productRepository *ProductRepository) FindAll() (product []*model.Product, err error) {
	if err = productRepository.DB.Find(&product).Error; err != nil {
		return
	}
	return
}

func (productRepository *ProductRepository) FindById(id string) (product *model.Product, err error) {
	if err = productRepository.DB.First(&product, id).Error; err != nil {
		return
	}
	return
}
