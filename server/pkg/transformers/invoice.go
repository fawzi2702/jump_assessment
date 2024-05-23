package transformers

import (
	"github.com/this_is_iz/jump_server/internal/utils"
	"github.com/this_is_iz/jump_server/pkg/models"
)

// TransformInvoice transforms an invoice model into an invoice response
func TransformInvoice(invoice models.Invoice) models.InvoiceResponse {
	invoiceAmount, err := utils.AddDecimalPoint(invoice.Amount)
	if err != nil {
		invoiceAmount = 0
	}

	transformedUser := TransformUser(*invoice.User)

	return models.InvoiceResponse{
		Invoice: invoice,
		Amount:  invoiceAmount,
		User:    &transformedUser,
	}
}

// TransformInvoices transforms a slice of invoice models into a slice of invoice responses
func TransformInvoices(invoices []models.Invoice) []models.InvoiceResponse {
	return utils.Map(invoices, TransformInvoice)
}
