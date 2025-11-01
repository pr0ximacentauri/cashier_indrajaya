package routes

import (
	"github.com/gofiber/fiber/v2"
	"cashier_indrajaya/controllers"
	"cashier_indrajaya/repositories"
	"cashier_indrajaya/services"
)

func SetupRoutes(app *fiber.App) {
	// Inisialisasi repository, service, controller
	productRepo := repositories.NewProductRepository()
	productService := services.NewProductService(productRepo)
	productController := controllers.NewProductController(productService)

	transactionRepo := repositories.NewTransactionRepository()
	transactionService := services.NewTransactionService(transactionRepo)
	transactionController := controllers.NewTransactionController(transactionService)

	// Grouping API
	api := app.Group("/api")

	// Product Routes
	api.Get("/products", productController.GetAll)
	api.Post("/products", productController.Create)
	api.Put("/products/:id", productController.Update)
	api.Delete("/products/:id", productController.Delete)

	// Transaction Routes
	api.Get("/transactions", transactionController.GetAll)
	api.Post("/transactions", transactionController.Create)
}
