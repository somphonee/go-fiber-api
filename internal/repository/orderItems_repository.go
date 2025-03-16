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
	return &OrderItemRepository{DB: db}
}

// CreateOrderItem creates a new order item in the database
func (r *OrderItemRepository) CreateOrderItem(orderItem *models.OrderItem) error {
	query := `
		INSERT INTO order_items (order_id, product_id, quantity, price, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id;
	`
	err := r.DB.Raw(query,
		orderItem.OrderID,
		orderItem.ProductID,
		orderItem.Quantity,
		orderItem.Price,
		time.Now(),
		time.Now(),
	).Scan(&orderItem.ID).Error

	return err
}

// GetOrderItems retrieves all items for a given order
func (r *OrderItemRepository) GetOrderItems(orderID uint) ([]models.OrderItem, error) {
	query := `
		SELECT id, order_id, product_id, quantity, price, created_at, updated_at, deleted_at
		FROM order_items WHERE order_id = $1;
	`
	var orderItems []models.OrderItem
	err := r.DB.Raw(query, orderID).Scan(&orderItems).Error
	if err != nil {
		return nil, err
	}
	return orderItems, nil
}

// UpdateOrderItem updates an order item based on its ID
func (r *OrderItemRepository) UpdateOrderItem(orderItem *models.OrderItem) error {
	query := `
		UPDATE order_items 
		SET quantity = $1, price = $2, updated_at = $3
		WHERE id = $4;
	`
	err := r.DB.Exec(query,
		orderItem.Quantity,
		orderItem.Price,
		time.Now(),
		orderItem.ID,
	).Error

	return err
}

// DeleteOrderItem soft deletes an order item (sets deleted_at)
func (r *OrderItemRepository) DeleteOrderItem(orderItemID uint) error {
	query := `
		UPDATE order_items SET deleted_at = $1 WHERE id = $2;
	`
	err := r.DB.Exec(query, time.Now(), orderItemID).Error
	return err
}

// GetOrderItem retrieves a single order item by its ID
func (r *OrderItemRepository) GetOrderItem(orderItemID uint) (*models.OrderItem, error) {
	query := `
		SELECT id, order_id, product_id, quantity, price, created_at, updated_at, deleted_at
		FROM order_items WHERE id = $1;
	`
	var orderItem models.OrderItem
	err := r.DB.Raw(query, orderItemID).Scan(&orderItem).Error
	if err != nil {
		return nil, err
	}
	return &orderItem, nil
}
