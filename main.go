package main

import (
	"cashier_indrajaya/infrastructure/database"
	"cashier_indrajaya/infrastructure/repositories"
	"cashier_indrajaya/interfaces/http/handlers"
	"cashier_indrajaya/interfaces/http/middlewares"
	"cashier_indrajaya/usecases"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	// Setup database
	db, err := database.SetupDatabase()
	if err != nil {
		log.Fatal("Failed to setup database:", err)
	}
	defer db.Close()

	// Initialize repositories
	productRepo := repositories.NewSQLiteProductRepository(db)
	saleRepo := repositories.NewSQLiteSaleRepository(db)

	// Initialize services
	productService := usecases.NewProductService(productRepo)
	saleService := usecases.NewSaleService(saleRepo, productRepo)

	// Initialize handlers
	productHandler := handlers.NewProductHandler(productService)
	saleHandler := handlers.NewSaleHandler(saleService)

	// Setup routes
	r := mux.NewRouter()
	
	// API routes
	api := r.PathPrefix("/api").Subrouter()
	
	// Product routes
	api.HandleFunc("/products", productHandler.GetProducts).Methods("GET")
	api.HandleFunc("/products/{id}", productHandler.GetProduct).Methods("GET")
	api.HandleFunc("/products", productHandler.CreateProduct).Methods("POST")
	api.HandleFunc("/products/{id}", productHandler.UpdateProduct).Methods("PUT")
	api.HandleFunc("/products/{id}", productHandler.DeleteProduct).Methods("DELETE")
	
	// Sale routes
	api.HandleFunc("/sales", saleHandler.CreateSale).Methods("POST")
	api.HandleFunc("/sales", saleHandler.GetSales).Methods("GET")

	 
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./web/views/")))

	// Apply CORS middleware
	handler := middlewares.CorsMiddleware(r)

	fmt.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}