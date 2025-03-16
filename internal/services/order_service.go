package services

import (
	"errors"
	"time"
	"github.com/somphonee/go-fiber-api/internal/models"
	"github.com/somphonee/go-fiber-api/internal/repository"
)

type OrderService struct {
	OrderRepo      *repository.OrderRepository
	OrderItemRepo  *repository.OrderItemRepository
}

func NewOrderService(orderRepo *repository.OrderRepository, orderItemRepo *repository.OrderItemRepository) *OrderService {
	return &OrderService{
		OrderRepo:      orderRepo,
		OrderItemRepo:  orderItemRepo,
	}
}

// CreateOrder creates a new order with order items
func (s *OrderService) CreateOrder(userID uint, items []models.OrderItem) (*models.Order, error) {
	// Calculate total price for the order
	totalPrice := 0.0
	for _, item := range items {
		totalPrice += item.Price * float64(item.Quantity)
	}

	// Create new order
	order := models.Order{
		UserID:   userID,
		Total:    totalPrice,
		Status:   "pending", // Default status
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Insert the order into the database
	err := s.OrderRepo.CreateOrder(&order)
	if err != nil {
		return nil, err
	}

	// Add order items
	for _, item := range items {
		item.OrderID = order.ID
		item.CreatedAt = time.Now()
		item.UpdatedAt = time.Now()
		err = s.OrderItemRepo.CreateOrderItem(&item)
		if err != nil {
			return nil, err
		}
	}

	// Return the created order
	return &order, nil
}

// GetOrder retrieves an order by its ID
func (s *OrderService) GetOrder(orderID uint) (*models.Order, error) {
	order, err := s.OrderRepo.GetOrder(orderID)
	if err != nil {
		return nil, err
	}
	return order, nil
}

// GetOrders retrieves all orders for a specific user
func (s *OrderService) GetOrders(userID uint) ([]models.Order, error) {
	orders, err := s.OrderRepo.GetOrders(userID)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

// UpdateOrderStatus updates the status of an existing order
func (s *OrderService) UpdateOrderStatus(orderID uint, status string) error {
	// Check if the order exists
	_, err := s.GetOrder(orderID)
	if err != nil {
		return errors.New("order not found")
	}

	// Update the order status
	err = s.OrderRepo.UpdateOrderStatus(orderID, status)
	if err != nil {
		return err
	}
	return nil
}

// DeleteOrder soft deletes an order
func (s *OrderService) DeleteOrder(orderID uint) error {
	// Check if the order exists
	_, err := s.GetOrder(orderID)
	if err != nil {
		return errors.New("order not found")
	}

	// Soft delete the order
	err = s.OrderRepo.DeleteOrder(orderID)
	if err != nil {
		return err
	}

	// Soft delete the related order items
	orderItems, err := s.OrderItemRepo.GetOrderItems(orderID)
	if err != nil {
		return err
	}

	for _, item := range orderItems {
		err := s.OrderItemRepo.DeleteOrderItem(item.ID)
		if err != nil {
			return err
		}
	}

	return nil
}

// GetOrderItems retrieves all items for a given order
func (s *OrderService) GetOrderItems(orderID uint) ([]models.OrderItem, error) {
	orderItems, err := s.OrderItemRepo.GetOrderItems(orderID)
	if err != nil {
		return nil, err
	}
	return orderItems, nil
}

// UpdateOrderItem updates an order item
func (s *OrderService) UpdateOrderItem(orderItem *models.OrderItem) error {
	// Update the order item
	err := s.OrderItemRepo.UpdateOrderItem(orderItem)
	if err != nil {
		return err
	}
	return nil
}