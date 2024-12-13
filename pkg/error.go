package pkg

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-clean-architecture/common"
	"net/http"
)

type CustomError struct {
	Code    int
	Message string
}

func (e *CustomError) Error() string {
	return e.Message
}

func NewCustomError(code int, message string) *CustomError {
	return &CustomError{
		Code:    code,
		Message: message,
	}
}

func HandleError(c *gin.Context, err error) {
	var httpErr *CustomError
	if errors.As(err, &httpErr) {
		c.JSON(httpErr.Code, common.ErrorResponse{
			Error: httpErr.Message,
		})
	} else {
		c.JSON(http.StatusInternalServerError, common.ErrorResponse{
			Error: err.Error(),
		})
	}
}
