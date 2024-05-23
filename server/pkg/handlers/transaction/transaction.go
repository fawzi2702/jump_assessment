package transaction

import (
	"github.com/gin-gonic/gin"
	"github.com/this_is_iz/jump_server/internal/utils"
	"github.com/this_is_iz/jump_server/pkg/models"
	"github.com/this_is_iz/jump_server/pkg/response"
)

// NewTransaction is a transaction controller that handles the creation of a new transaction
func NewTransaction(c *gin.Context) {
	invoiceModel := models.NewInvoiceModel()
	invoiceModel.StartSession()

	var req models.CreateTransactionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, "invalid request body")
		return
	}

	invoice, err := invoiceModel.GetInvoiceByID(req.InvoiceID)
	if err != nil {
		response.InternalServerError(c)
		return
	} else if invoice == nil {
		response.NotFound(c)
		return
	}

	transactionAmount, err := utils.RemoveDecimalPoint(req.Amount)
	if err != nil {
		response.BadRequest(c, "invalid amount")
		return
	}

	if invoice.Status != models.INVOICE_STATUS_PENDING {
		response.Unprocessable(c, "invoice already paid")
		return
	} else if invoice.Amount != transactionAmount {
		response.BadRequest(c, "invalid amount")
		return
	}

	ok, err := invoiceModel.UpdateInvoiceStatus(invoice.InvoiceId, models.INVOICE_STATUS_PAID)
	if err != nil || !ok {
		response.InternalServerError(c)
		return
	}

	response.NoContent(c)
}
