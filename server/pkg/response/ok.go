package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const DEFAULT_OK_STATUS = http.StatusOK

type OkResponse struct {
	Ok   bool        `json:"ok" example:"true"`
	Data interface{} `json:"data"`
}

// Ok is a helper function to send a 204 OK response
func Ok(c *gin.Context, data interface{}, code ...int) {
	statusCode := DEFAULT_OK_STATUS

	if len(code) > 0 {
		statusCode = code[0]
	}

	body := gin.H{"ok": true, "data": data}

	c.JSON(statusCode, body)
}

// NoContent is a helper function to send a 204 No Content response
func NoContent(c *gin.Context) {
	c.Status(http.StatusNoContent)
}
