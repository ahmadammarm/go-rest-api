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

type CategoryHandler struct {
    Service *services.CategoryService
}

func (h *CategoryHandler) GetAllCategories(w http.ResponseWriter, r *http.Request) {
    categories, err := h.Service.GetAllCategories()
    if err != nil {
        http.Error(w, "Gagal mendapatkan data semua kategori", http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(categories)
}

func (h *CategoryHandler) CreateCategory(w http.ResponseWriter, r *http.Request) {
    var category models.Category
    err := json.NewDecoder(r.Body).Decode(&category)
    if err != nil {
        http.Error(w, "Gagal membuat kategori", http.StatusInternalServerError)
        return
    }

    newCategory, err := h.Service.CreateCategory(category)
    if err != nil {
        http.Error(w, "Gagal membuat kategori", http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(newCategory)
}

func (h *CategoryHandler) GetCategoryById(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])
    category, err := h.Service.GetCategoryById(id)
    if err != nil {
        http.Error(w, fmt.Sprintf("Gagal mendapatkan data dari kategori %d", id), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(category)
}

func (h *CategoryHandler) UpdateCategory(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])
    var category models.Category
    err := json.NewDecoder(r.Body).Decode(&category)
    if err != nil {
        http.Error(w, "Gagal mengupdate kategori", http.StatusInternalServerError)
        return
    }

    categoryUpdate, err := h.Service.UpdateCategory(id, category)
    if err != nil {
        http.Error(w, "Gagal mengupdate kategori", http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(categoryUpdate)
}

func (h *CategoryHandler) DeleteCategory(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id, _ := strconv.Atoi(params["id"])
    err := h.Service.DeleteCategory(id)
    if err != nil {
        http.Error(w, fmt.Sprintf("Gagal menghapus kategori %d", id), http.StatusInternalServerError)
        return
    }
    json.NewEncoder(w).Encode(map[string]string{"message": "Berhasil menghapus kategori"})
}


