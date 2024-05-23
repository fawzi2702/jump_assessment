package models

import (
	"github.com/this_is_iz/jump_server/internal/utils"
	"gorm.io/gorm"
)

const INVOICE_STATUS_PENDING = "pending"
const INVOICE_STATUS_PAID = "paid"

type Invoice struct {
	InvoiceId ID     `json:"invoice_id" gorm:"primarykey;column:id"`
	UserId    ID     `json:"user_id" gorm:"column:user_id"`
	Status    string `json:"status" gorm:"column:status;type:enum('pending', 'paid');default:'pending'"`
	Label     string `json:"label" gorm:"column:label" binding:"required"`
	Amount    int    `json:"amount" gorm:"column:amount" binding:"required"`
	User      *User  `json:"user,omitempty" gorm:"foreignKey:user_id;references:UserId;tableName:users"`
}

type InvoiceResponse struct {
	Invoice
	Amount float64       `json:"amount"`
	User   *UserResponse `json:"user"`
}

type InvoiceModel struct {
	baseModel
}

// NewInvoiceModel creates a new InvoiceModel
func NewInvoiceModel() *InvoiceModel {
	return &InvoiceModel{
		baseModel{
			model: DB.Model(&Invoice{}).Preload("User", func(db *gorm.DB) *gorm.DB {
				return db.Select("id, first_name, last_name, balance")
			}),
		},
	}
}

// GetInvoiceByID retrieves an invoice by its ID
func (m *InvoiceModel) GetInvoiceByID(invoiceId ID) (*Invoice, error) {
	var invoice Invoice

	err := m.model.Where("id = ?", invoiceId).First(&invoice).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return &invoice, nil
}

// GetInvoicesByUserID retrieves all invoices for a given user
func (m *InvoiceModel) GetInvoicesByUserID(userID ID) (*[]Invoice, error) {
	var invoices []Invoice

	err := m.model.Where("user_id = ?", userID).Find(&invoices).Error
	if err != nil {
		return nil, err
	}

	return &invoices, nil
}

type CreateInvoiceRequest struct {
	UserId ID      `json:"user_id"`
	Label  string  `json:"label"`
	Amount float64 `json:"amount"`
}

// InsertInvoice inserts a new invoice into the invoices table
func (m *InvoiceModel) InsertInvoice(input *CreateInvoiceRequest) (bool, error) {
	amount, err := utils.RemoveDecimalPoint(input.Amount)
	if err != nil {
		return false, err
	}

	new := &Invoice{
		UserId: input.UserId,
		Label:  input.Label,
		Amount: amount,
		Status: INVOICE_STATUS_PENDING,
	}

	err = m.model.Create(new).Error
	if err != nil {
		return false, err
	}

	return true, nil
}

type UpdateInvoiceRequest struct {
	InvoiceId ID      `json:"id" binding:"required"`
	Amount    float64 `json:"amount" binding:"required"`
}

// UpdateInvoiceStatus updates the status of an invoice
func (m *InvoiceModel) UpdateInvoiceStatus(invoiceId ID, status string) (bool, error) {
	err := m.model.Where("id = ?", invoiceId).Update("status", status).Error
	if err != nil {
		return false, err
	}

	return true, nil
}
