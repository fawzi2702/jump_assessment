package models

type Transaction struct {
	ID        ID      `json:"id"`
	InvoiceID ID      `json:"invoice_id"`
	Amount    float64 `json:"amount"`
	Reference string  `json:"reference"`
}

type TransactionModel struct {
	baseModel
}

func NewTransactionModel() *TransactionModel {
	return &TransactionModel{
		baseModel{
			model: DB.Model(&Transaction{}),
		},
	}
}

type CreateTransactionRequest struct {
	InvoiceID ID      `json:"invoice_id" binding:"required"`
	Amount    float64 `json:"amount" binding:"required"`
	Reference string  `json:"reference" binding:"required"`
}
