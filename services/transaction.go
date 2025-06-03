package service

import (
	"database/sql"
	"fmt"

	"github.com/aldysp34/deeptech-test/config"
	model "github.com/aldysp34/deeptech-test/models"
)

type TransactionService struct {
	db *sql.DB
}

func NewTransactionService() *TransactionService {
	db := config.Connect()

	return &TransactionService{
		db: db,
	}
}

func (s *TransactionService) CreateTransaction(req model.CreateTransactionRequest) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	result, err := tx.Exec("INSERT INTO transactions (type) VALUES (?)", req.Type)
	if err != nil {
		tx.Rollback()
		return err
	}

	transactionID, _ := result.LastInsertId()

	for _, item := range req.Items {
		var stock int
		err := tx.QueryRow("SELECT stock FROM products WHERE id = ?", item.ProductID).Scan(&stock)
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("product %d not found", item.ProductID)
		}

		if req.Type == "out" && item.Quantity > stock {
			tx.Rollback()
			return fmt.Errorf("stock for product %d is insufficient", item.ProductID)
		}

		if req.Type == "out" {
			_, err = tx.Exec("UPDATE products SET stock = stock - ? WHERE id = ?", item.Quantity, item.ProductID)
		} else {
			_, err = tx.Exec("UPDATE products SET stock = stock + ? WHERE id = ?", item.Quantity, item.ProductID)
		}

		if err != nil {
			tx.Rollback()
			return err
		}

		_, err = tx.Exec(`INSERT INTO transaction_items (transaction_id, product_id, quantity)
                          VALUES (?, ?, ?)`, transactionID, item.ProductID, item.Quantity)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}
