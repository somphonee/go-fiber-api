package services
import (
	"github.com/somphonee/go-fiber-api/internal/repository"
	"github.com/somphonee/go-fiber-api/internal/models"
	"github.com/somphonee/go-fiber-api/internal/utils"
)

type ProductService struct {
	repo *repository.ProductRepository
}
func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}
// GetAllProductsPaginated ดึงข้อมูลสินค้าแบบแบ่งหน้า
func (s *ProductService) GetAllProductsPaginated(pagination *utils.Pagination) error {
	return s.repo.FindAllPaginated(pagination)
}

// SearchProducts ค้นหาสินค้าจากชื่อ
func (s *ProductService) SearchProducts(name string) ([]models.Product, error) {
	return s.repo.FindByName(name)
}

func (s *ProductService) GetAllProducts() ([]models.Product, error) {
	return s.repo.FindAll()
}

func (s *ProductService) GetProductByID(id uint) (models.Product, error) {
	return s.repo.FindByID(id)
}

func (s *ProductService) CreateProduct(product *models.Product) error {
	return s.repo.Create(product)
}

func (s *ProductService) UpdateProduct(product *models.Product) error {
	return s.repo.Update(product)
}

func (s *ProductService) DeleteProduct(id uint) error {
	return s.repo.Delete(id)
}