package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	model "github.com/aldysp34/deeptech-test/models"
	service "github.com/aldysp34/deeptech-test/services"
	"github.com/gorilla/mux"
)

type CategoryController struct {
	Service *service.CategoryService
}

func NewCategoryController(service *service.CategoryService) *CategoryController {
	return &CategoryController{Service: service}
}

func (c *CategoryController) Create(w http.ResponseWriter, r *http.Request) {
	var category model.Category
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := c.Service.Create(&category); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(category)
}

func (c *CategoryController) List(w http.ResponseWriter, r *http.Request) {
	categories, err := c.Service.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(categories)
}

func (c *CategoryController) GetByID(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, _ := strconv.ParseInt(idStr, 10, 64)

	category, err := c.Service.GetByID(id)
	if err != nil {
		http.Error(w, "Category not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(category)
}

func (c *CategoryController) Update(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, _ := strconv.ParseInt(idStr, 10, 64)

	var category model.Category
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	category.ID = id
	if err := c.Service.Update(&category); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(category)
}

func (c *CategoryController) Delete(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, _ := strconv.ParseInt(idStr, 10, 64)

	if err := c.Service.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
