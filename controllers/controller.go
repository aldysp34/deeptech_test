package controllers

import (
	"github.com/aldysp34/deeptech-test/middlewares"
	"github.com/aldysp34/deeptech-test/repositories"
	service "github.com/aldysp34/deeptech-test/services"
	"github.com/gorilla/mux"
)

func NewAdminRouter(router *mux.Router) {
	adminRepo := repositories.NewAdminRepository()
	service := service.NewAdminService(adminRepo)
	controllers := NewAdminController(service)

	admin := router.PathPrefix("/admin").Subrouter()

	admin.HandleFunc("/register", controllers.Create).Methods("POST")
	admin.HandleFunc("/login", controllers.Login).Methods("POST")

	adminAuth := admin.NewRoute().Subrouter()
	adminAuth.Use(middlewares.AuthMiddleware)
	adminAuth.HandleFunc("/profile", controllers.UpdateProfile).Methods("PUT")
	adminAuth.HandleFunc("/list", controllers.List).Methods("GET")
	adminAuth.HandleFunc("/{id}", controllers.GetByID).Methods("GET")
	adminAuth.HandleFunc("/{id}", controllers.Delete).Methods("DELETE")
}

func NewCategoryRouter(router *mux.Router) {
	repo := repositories.NewCategoryRepository()
	service := service.NewCategoryService(repo)
	controller := NewCategoryController(service)

	category := router.PathPrefix("/categories").Subrouter()
	category.HandleFunc("", controller.Create).Methods("POST")
	category.HandleFunc("", controller.List).Methods("GET")
	category.HandleFunc("/{id}", controller.GetByID).Methods("GET")
	category.HandleFunc("/{id}", controller.Update).Methods("PUT")
	category.HandleFunc("/{id}", controller.Delete).Methods("DELETE")
}

func NewProductRouter(router *mux.Router) {
	productRepo := repositories.NewProductRepository()
	productService := service.NewProductService(productRepo)
	productController := NewProductController(productService)

	product := router.PathPrefix("/products").Subrouter()
	product.HandleFunc("", productController.Create).Methods("POST")
	product.HandleFunc("", productController.List).Methods("GET")
	product.HandleFunc("/{id}", productController.GetByID).Methods("GET")
	product.HandleFunc("/{id}", productController.Update).Methods("PUT")
	product.HandleFunc("/{id}", productController.Delete).Methods("DELETE")
}

func NewTransactionRouter(router *mux.Router) {
	transactionService := service.NewTransactionService()
	transactionController := NewTransactionController(transactionService)

	transaction := router.PathPrefix("/transactions").Subrouter()
	transaction.Use(middlewares.AuthMiddleware)
	transaction.HandleFunc("", transactionController.Create).Methods("POST")
}
