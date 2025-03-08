package routes

import (
    "github.com/ahmadammarm/go-rest-api/handlers"
    "github.com/ahmadammarm/go-rest-api/internals/repositories"
    "github.com/ahmadammarm/go-rest-api/internals/services"
    "github.com/gorilla/mux"
    "gorm.io/gorm"
)

// Router untuk produk
func ProductRouter(r *mux.Router, db *gorm.DB) {
    productRepo := &repositories.ProductRepository{DB: db}
    productService := &services.ProductService{Repo: productRepo}
    productHandler := &handlers.ProductHandler{Service: productService}

    r.HandleFunc("/", productHandler.GetAllProducts).Methods("GET")
    r.HandleFunc("/", productHandler.CreateProduct).Methods("POST")
    r.HandleFunc("/{id}", productHandler.GetProductById).Methods("GET")
    r.HandleFunc("/{id}", productHandler.UpdateProduct).Methods("PUT")
    r.HandleFunc("/{id}", productHandler.DeleteProduct).Methods("DELETE")
}

// Router untuk kategori
func CategoryRouter(r *mux.Router, db *gorm.DB) {
    categoryRepo := &repositories.CategoryRepository{DB: db}
    categoryService := &services.CategoryService{Repo: categoryRepo}
    categoryHandler := &handlers.CategoryHandler{Service: categoryService}

    r.HandleFunc("/", categoryHandler.GetAllCategories).Methods("GET")
    r.HandleFunc("/", categoryHandler.CreateCategory).Methods("POST")
    r.HandleFunc("/{id}", categoryHandler.GetCategoryById).Methods("GET")
    r.HandleFunc("/{id}", categoryHandler.UpdateCategory).Methods("PUT")
    r.HandleFunc("/{id}", categoryHandler.DeleteCategory).Methods("DELETE")
}

func SetupRoutes(db *gorm.DB) *mux.Router {
    routes := mux.NewRouter()

    // Router untuk produk
    productRouter := routes.PathPrefix("/products").Subrouter()
    ProductRouter(productRouter, db)

    // Router untuk kategori
    categoryRouter := routes.PathPrefix("/categories").Subrouter()
    CategoryRouter(categoryRouter, db)

    return routes
}