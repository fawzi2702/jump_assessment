package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/this_is_iz/jump_server/pkg/handlers/transaction"
)

func SetupTransactionRouter(r *gin.Engine) {
	router := r.Group("/transaction")
	{
		router.POST("/", transaction.NewTransaction)
	}
}
