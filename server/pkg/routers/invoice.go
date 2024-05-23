package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/this_is_iz/jump_server/pkg/handlers/invoice"
)

func SetupInvoiceRouter(r *gin.Engine) {
	router := r.Group("/invoice")
	{
		router.GET("/", invoice.GetInvoices)
		router.POST("/", invoice.CreateInvoice)
	}
}
