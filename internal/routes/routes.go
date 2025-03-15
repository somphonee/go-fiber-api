
package routes
import (
	"github.com/gofiber/fiber/v2"
	"github.com/somphonee/go-fiber-api/internal/handlers"
	"github.com/somphonee/go-fiber-api/internal/middleware"
)

func SetupRoutes(app *fiber.App, productHandler *handlers.ProductHandler, userHandler *handlers.UserHandler,  authHandler *handlers.AuthHandler) {
	// Product Routes
	api := app.Group("/api")
	// Authentication routes
	auth := api.Group("/auth")
	auth.Post("/register", authHandler.Register)
	auth.Post("/login", authHandler.Login)
	
	// Protected routes
	protected := api.Group("/", middleware.Protected())

	// Product routes
	products := protected.Group("/products")
	products.Get("/", productHandler.GetAllProductsPaginated) // ใช้ pagination
	products.Get("/all", productHandler.GetAllProducts) // ดึงข้อมูลทั้งหมดโดยไม่มี pagination
	products.Get("/search", productHandler.SearchProducts) // ค้นหาสินค้า
	products.Get("/:id", productHandler.GetProductByID)
	products.Post("/", productHandler.CreateProduct)
	products.Put("/:id", productHandler.UpdateProduct)
	products.Delete("/:id", productHandler.DeleteProduct)

	// User Routes
	users := protected.Group("/users")
	users.Get("/", userHandler.GetAllUsers)
	users.Get("/:id", userHandler.GetUserByID)
	users.Post("/", userHandler.CreateUser)
	users.Put("/:id", userHandler.UpdateUser)
	users.Delete("/:id", userHandler.DeleteUser)
}