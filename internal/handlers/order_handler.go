package handlers

import (
	"github.com/somphonee/go-fiber-api/internal/services"
	"github.com/gofiber/fiber/v2"
	"github.com/somphonee/go-fiber-api/internal/models"
	"strconv"
)

type OrderHandler struct {
	service *services.OrderService
}

func NewOrderHandler(service *services.OrderService) *OrderHandler {
	return &OrderHandler{service: service}
}

// CreateOrder creates a new order
func (h *OrderHandler) CreateOrder(c *fiber.Ctx) error {
		var req struct {
		UserID uint              `json:"user_id"`
		Items  []models.OrderItem `json:"items"`
	}


	if err := c.BodyParser(&req); err != nil {
	
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse order items",
		})
	}

	// Extract userID from the request context or JWT token (assumed)


	order, err := h.service.CreateOrder(req.UserID, req.Items)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create order",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(order)
}

// GetOrder retrieves a specific order by ID
func (h *OrderHandler) GetOrder(c *fiber.Ctx) error {
	orderID, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid order ID",
		})
	}

	order, err := h.service.GetOrder(uint(orderID))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Order not found",
		})
	}

	return c.JSON(order)
}

// GetOrders retrieves all orders for a specific user
func (h *OrderHandler) GetOrders(c *fiber.Ctx) error {
	// Extract userID from the request context or JWT token (assumed)
	userID := uint(1) // Example: Replace with actual userID from JWT

	orders, err := h.service.GetOrders(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve orders",
		})
	}

	return c.JSON(orders)
}

// UpdateOrderStatus updates the status of an order
func (h *OrderHandler) UpdateOrderStatus(c *fiber.Ctx) error {
	orderID, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid order ID",
		})
	}

	status := c.FormValue("status")
	if status == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Status is required",
		})
	}

	err = h.service.UpdateOrderStatus(uint(orderID), status)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update order status",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Order status updated successfully",
	})
}

// DeleteOrder soft deletes an order
func (h *OrderHandler) DeleteOrder(c *fiber.Ctx) error {
	orderID, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid order ID",
		})
	}

	err = h.service.DeleteOrder(uint(orderID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to delete order",
		})
	}

	return c.Status(fiber.StatusNoContent).Send(nil)
}

// GetOrderItems retrieves all items for a given order
func (h *OrderHandler) GetOrderItems(c *fiber.Ctx) error {
	orderID, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid order ID",
		})
	}

	orderItems, err := h.service.GetOrderItems(uint(orderID))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve order items",
		})
	}

	return c.JSON(orderItems)
}

// UpdateOrderItem updates an order item
func (h *OrderHandler) UpdateOrderItem(c *fiber.Ctx) error {
	orderItem := new(models.OrderItem)
	if err := c.BodyParser(orderItem); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse order item",
		})
	}

	err := h.service.UpdateOrderItem(orderItem)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to update order item",
		})
	}

	return c.Status(fiber.StatusOK).JSON(orderItem)
}
