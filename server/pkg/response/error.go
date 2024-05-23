package response

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Ok    bool   `json:"ok" example:"false"`
	Error string `json:"error"`
}

// Error is an helper function to send an error response
func Error(c *gin.Context, status int, err error) {
	body := gin.H{"ok": false, "error": err.Error()}

	c.JSON(status, body)
}

// InternalServerError is a helper function to send a 500 Internal Server Error response
func InternalServerError(c *gin.Context, err ...error) {
	e := errors.New("internal Server Error")

	if len(err) > 0 {
		e = err[0]
	}

	Error(c, http.StatusInternalServerError, e)
}

// BadRequest is a helper function to send a 400 Bad Request response
func BadRequest(c *gin.Context, reason string) {
	err := errors.New(reason)

	Error(c, http.StatusBadRequest, err)
}

// ValidationError is a helper function to send a 400 Bad Request response
func ValidationError(c *gin.Context, reason string) {}

// NotFound is a helper function to send a 404 Not Found response
func NotFound(c *gin.Context, reason ...string) {
	err := errors.New("not found")

	if len(reason) > 0 {
		err = errors.New(reason[0])
	}

	Error(c, http.StatusNotFound, err)
}

// Unprocessable is a helper function to send a 422 Unprocessable Entity response
func Unprocessable(c *gin.Context, reason string) {
	err := errors.New(reason)

	Error(c, http.StatusUnprocessableEntity, err)
}
