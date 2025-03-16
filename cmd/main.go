package main

import (
	"log"
	"os"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/somphonee/go-fiber-api/config"
	"github.com/somphonee/go-fiber-api/internal/handlers"
	"github.com/somphonee/go-fiber-api/internal/repository"
	"github.com/somphonee/go-fiber-api/internal/routes"
	"github.com/somphonee/go-fiber-api/internal/services"
	"github.com/somphonee/go-fiber-api/migrations"

)
func main() {
	// Load environment variables
	config.LoadEnv()

	// Connect to database
	db := config.ConnectDB()

	// Run migrations
	if err := migrations.Migrate(db); err != nil {
		log.Fatal("Migration failed:", err)
	}

	// Initialize repositories
	productRepo := repository.NewProductRepository(db)
	userRepo := repository.NewUserRepository(db)
  orderRepo := repository.NewOrderRepository(db)
	orderItemRepo := repository.NewOrderItemRepository(db)

	
	
	// Initialize services
	productService := services.NewProductService(productRepo)
	userService := services.NewUserService(userRepo)
	authService := services.NewAuthService(userRepo)
	orderService := services.NewOrderService(orderRepo,orderItemRepo)

	// Initialize handlers
	productHandler := handlers.NewProductHandler(productService)
	userHandler := handlers.NewUserHandler(userService)
	authHandler := handlers.NewAuthHandler(authService)
	orderHandler := handlers.NewOrderHandler(orderService)

	// Create fiber app
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		},
	})

	// Middlewares
	app.Use(logger.New())
	app.Use(cors.New())
	app.Use(recover.New())

	// Setup routes
	routes.SetupRoutes(app, productHandler, userHandler,authHandler,orderHandler)

	// Add health check endpoint
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	// Get port from environment variable
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	// Start server
	log.Printf("Server starting on port %s...\n", port)
	log.Fatal(app.Listen(":" + port))
}