package repository

import (
	"time"
	"gorm.io/gorm"
	"github.com/somphonee/go-fiber-api/internal/models"
)

type OrderRepository struct {
	DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{DB: db}
}

// CreateOrder creates a new order in the database and returns its ID
func (r *OrderRepository) CreateOrder(order *models.Order) error {
	query := `
		INSERT INTO orders (user_id, total, status, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?) RETURNING id;
	`
	// รับค่า ID ของ order ที่ถูกสร้าง
	result := r.DB.Raw(query, order.UserID, order.Total, order.Status, order.CreatedAt, order.UpdatedAt).Scan(&order.ID)
	return result.Error
}

// GetOrder retrieves an order by its ID
func (r *OrderRepository) GetOrder(orderID uint) (*models.Order, error) {
	query := `
		SELECT id, user_id, total, status, created_at, updated_at, deleted_at
		FROM orders WHERE id = ?;
	`
	var order models.Order
	result := r.DB.Raw(query, orderID).Scan(&order)
	if result.Error != nil {
		return nil, result.Error
	}
	return &order, nil
}

// GetOrders retrieves all orders for a specific user
func (r *OrderRepository) GetOrders(userID uint) ([]models.Order, error) {
	query := `
		SELECT id, user_id, total, status, created_at, updated_at, deleted_at
		FROM orders WHERE user_id = ?;
	`
	var orders []models.Order
	result := r.DB.Raw(query, userID).Scan(&orders)
	if result.Error != nil {
		return nil, result.Error
	}
	return orders, nil
}

// UpdateOrderStatus updates the status of an order
func (r *OrderRepository) UpdateOrderStatus(orderID uint, status string) error {
	query := `
		UPDATE orders SET status = ?, updated_at = ? WHERE id = ?;
	`
	result := r.DB.Exec(query, status, time.Now(), orderID)
	return result.Error
}

// DeleteOrder soft deletes an order by marking it as deleted (using deleted_at)
func (r *OrderRepository) DeleteOrder(orderID uint) error {
	query := `
		UPDATE orders SET deleted_at = ? WHERE id = ?;
	`
	result := r.DB.Exec(query, time.Now(), orderID)
	return result.Error
}
