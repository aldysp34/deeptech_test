package model

type TransactionItemRequest struct {
	ProductID int64 `json:"product_id" validate:"required"`
	Quantity  int   `json:"quantity" validate:"required,gt=0"`
}

type CreateTransactionRequest struct {
	Type  string                   `json:"type" validate:"required,oneof=in out"`
	Items []TransactionItemRequest `json:"items" validate:"required,dive"`
}
