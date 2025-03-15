package repository

import (
	"gorm.io/gorm"
	"github.com/somphonee/go-fiber-api/internal/models"
)


type ProductRepository struct {
	DB *gorm.DB
}


func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{DB: db}
}

func (r *ProductRepository) FindAll() ([]models.Product, error) {
	var products []models.Product
	err := r.DB.Find(&products).Error
	return products, err
}

func (r *ProductRepository) FindByID(id uint) (models.Product, error) {
	var product models.Product
	err := r.DB.First(&product, id).Error
	return product, err
}

func (r *ProductRepository) Create(product *models.Product) error {
	return r.DB.Create(&product).Error
}

func (r *ProductRepository) Update(product *models.Product) error {
	return r.DB.Save(&product).Error
}

func (r *ProductRepository) Delete(id uint) error {
	return r.DB.Delete(&models.Product{}, id).Error
}