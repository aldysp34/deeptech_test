package controllers

import (
	"encoding/json"
	"net/http"

	model "github.com/aldysp34/deeptech-test/models"
	service "github.com/aldysp34/deeptech-test/services"
	"github.com/aldysp34/deeptech-test/utils"
	"github.com/go-playground/validator/v10"
)

type TransactionController struct {
	service *service.TransactionService
}

func NewTransactionController(service *service.TransactionService) *TransactionController {
	return &TransactionController{service: service}
}

func (c *TransactionController) Create(w http.ResponseWriter, r *http.Request) {
	var req model.CreateTransactionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.Error(w, http.StatusBadRequest, "invalid request format")
		return
	}

	if err := validator.New().Struct(req); err != nil {
		utils.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	err := c.service.CreateTransaction(req)
	if err != nil {
		utils.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	utils.Success(w, "transaction created successfully")
}
