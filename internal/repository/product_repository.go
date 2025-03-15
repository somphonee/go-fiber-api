package repository

import (
	"gorm.io/gorm"
	"github.com/somphonee/go-fiber-api/internal/models"
	"github.com/somphonee/go-fiber-api/internal/utils"
)


type ProductRepository struct {
	DB *gorm.DB
}


func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{DB: db}
}

// FindAllPaginated ดึงข้อมูลสินค้าทั้งหมดแบบ pagination
func (r *ProductRepository) FindAllPaginated(pagination *utils.Pagination) error {
	var products []models.Product
	return utils.Paginate(r.DB, &products, pagination)
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

// FindByName ค้นหาสินค้าจากชื่อ
func (r *ProductRepository) FindByName(name string) ([]models.Product, error) {
	var products []models.Product
	err := r.DB.Where("name LIKE ?", "%"+name+"%").Find(&products).Error
	return products, err
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