package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"strconv"

	"github.com/ahmadammarm/go-rest-api/internals/models"
	"github.com/ahmadammarm/go-rest-api/internals/services"
	"github.com/gorilla/mux"
)

type ProductHandler struct {
	Service *services.ProductService
}

func (h *ProductHandler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.Service.GetAllProducts()
	if err != nil {
		http.Error(w, "Gagal mendapatkan data semua produk", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(products)
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, "Gagal membuat produk", http.StatusInternalServerError)
		return
	}

	newProduct, err := h.Service.CreateProduct(product)
	if err != nil {
		http.Error(w, "Gagal membuat produk", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(newProduct)
}

func (h *ProductHandler) GetProductById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	product, err := h.Service.GetProductById(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Gagal mendapatkan data dari produk %d", id), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(product)
}

func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	var product models.Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, fmt.Sprintf("Gagal mengubah data dari produk %d", id), http.StatusInternalServerError)
		return
	}

	updatedProduct, err := h.Service.UpdateProduct(id, product)
	if err != nil {
		http.Error(w, fmt.Sprintf("Gagal mendapatkan data dari produk %d", id), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(updatedProduct)
}

func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	err := h.Service.DeleteProduct(id)
	if err != nil {
		http.Error(w, fmt.Sprintf("Gagal menghapus produk %d", id), http.StatusInternalServerError)
		return
	}

	w.Write([]byte("Produk berhasil dihapus"))
}
