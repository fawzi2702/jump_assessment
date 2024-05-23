package invoice

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/this_is_iz/jump_server/pkg/models"
	"github.com/this_is_iz/jump_server/pkg/response"
	"github.com/this_is_iz/jump_server/pkg/transformers"
)

func GetInvoices(c *gin.Context) {
	invoiceModel := models.NewInvoiceModel()

	userId, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		response.BadRequest(c, "user_id query param must be a valid integer")
		return
	}

	invoices, err := invoiceModel.GetInvoicesByUserID(userId)
	if err != nil {
		response.InternalServerError(c)
		return
	}

	response.Ok(c, transformers.TransformInvoices(*invoices))
}

func CreateInvoice(c *gin.Context) {
	invoiceModel := models.NewInvoiceModel()
	userModel := models.NewUserModel()

	var req models.CreateInvoiceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Printf("error: %v\n", err)
		response.BadRequest(c, "invalid request body")
		return
	}

	user, err := userModel.GetUserByID(req.UserId)
	if err != nil {
		response.InternalServerError(c)
		return
	} else if user == nil {
		response.BadRequest(c, "invalid user_id")
		return
	}

	ok, err := invoiceModel.InsertInvoice(&req)
	if err != nil || !ok {
		response.InternalServerError(c)
		return
	}

	response.NoContent(c)
}
