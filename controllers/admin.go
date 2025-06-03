package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	model "github.com/aldysp34/deeptech-test/models"
	service "github.com/aldysp34/deeptech-test/services"
	"github.com/aldysp34/deeptech-test/utils"
	"github.com/gorilla/mux"
)

type AdminController struct {
	Service *service.AdminService
}

func NewAdminController(s *service.AdminService) *AdminController {
	return &AdminController{
		Service: s,
	}
}

func (ac *AdminController) Create(w http.ResponseWriter, r *http.Request) {
	var req model.CreateAdminRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := ac.Service.CreateAdmin(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Admin created"})
}

func (ac *AdminController) Login(w http.ResponseWriter, r *http.Request) {
	var req model.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token, err := ac.Service.Login(req.Email, req.Password)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func (ac *AdminController) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	adminID, _ := utils.ExtractAdminIDFromRequest(r)

	var req model.UpdateProfileRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := ac.Service.UpdateProfile(adminID, req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Profile updated"})
}

func (ac *AdminController) List(w http.ResponseWriter, r *http.Request) {
	admins, err := ac.Service.ListAdmins()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(admins)
}

func (ac *AdminController) GetByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	admin, err := ac.Service.GetAdminByID(id)
	if err != nil {
		http.Error(w, "Admin not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(admin)
}

func (ac *AdminController) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(mux.Vars(r)["id"], 10, 64)
	err := ac.Service.DeleteAdmin(id)
	if err != nil {
		http.Error(w, "Failed to delete admin", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"message": "Admin deleted"})
}
