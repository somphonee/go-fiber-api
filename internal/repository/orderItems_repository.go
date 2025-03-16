package repository

import (
	"time"
	"gorm.io/gorm"
	"github.com/somphonee/go-fiber-api/internal/models"
)

type OrderItemRepository struct {
	DB *gorm.DB
}

func NewOrderItemRepository(db *gorm.DB) *OrderItemRepository {
	return &OrderItemRepository{
		DB: db,
	}
}

// CreateOrderItem creates a new order item in the database
func (r *OrderItemRepository) CreateOrderItem(orderItem *models.OrderItem) error {
	query := `
		INSERT INTO order_items (order_id, product_id, quantity, price, created_at, updated_at)
		VALUES (:order_id, :product_id, :quantity, :price, :created_at, :updated_at);
	`
	params := map[string]interface{}{
		"order_id":   orderItem.OrderID,
		"product_id": orderItem.ProductID,
		"quantity":   orderItem.Quantity,
		"price":      orderItem.Price,
		"created_at": orderItem.CreatedAt,
		"updated_at": orderItem.UpdatedAt,
	}
	result := r.DB.Exec(query, params)
	return result.Error
}

// GetOrderItems retrieves all items for a given order
func (r *OrderItemRepository) GetOrderItems(orderID uint) ([]models.OrderItem, error) {
	query := `
		SELECT id, order_id, product_id, quantity, price, created_at, updated_at, deleted_at
		FROM order_items WHERE order_id = :order_id;
	`
	params := map[string]interface{}{
		"order_id": orderID,
	}
	var orderItems []models.OrderItem
	result := r.DB.Raw(query, params).Scan(&orderItems)
	if result.Error != nil {
		return nil, result.Error
	}
	return orderItems, nil
}

// UpdateOrderItem updates an order item based on its ID
func (r *OrderItemRepository) UpdateOrderItem(orderItem *models.OrderItem) error {
	query := `
		UPDATE order_items SET quantity = :quantity, price = :price, updated_at = :updated_at
		WHERE id = :id;
	`
	params := map[string]interface{}{
		"quantity":   orderItem.Quantity,
		"price":      orderItem.Price,
		"updated_at": time.Now(),
		"id":         orderItem.ID,
	}
	result := r.DB.Exec(query, params)
	return result.Error
}

// DeleteOrderItem soft deletes an order item (sets deleted_at)
func (r *OrderItemRepository) DeleteOrderItem(orderItemID uint) error {
	query := `
		UPDATE order_items SET deleted_at = :deleted_at WHERE id = :id;
	`
	params := map[string]interface{}{
		"deleted_at": time.Now(),
		"id":         orderItemID,
	}
	result := r.DB.Exec(query, params)
	return result.Error
}

// GetOrderItem retrieves a single order item by its ID
func (r *OrderItemRepository) GetOrderItem(orderItemID uint) (*models.OrderItem, error) {
	query := `
		SELECT id, order_id, product_id, quantity, price, created_at, updated_at, deleted_at
		FROM order_items WHERE id = :id;
	`
	params := map[string]interface{}{
		"id": orderItemID,
	}
	var orderItem models.OrderItem
	result := r.DB.Raw(query, params).Scan(&orderItem)
	if result.Error != nil {
		return nil, result.Error
	}
	return &orderItem, nil
}