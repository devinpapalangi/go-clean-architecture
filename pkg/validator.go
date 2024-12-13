package pkg

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strings"
)

func DecodeAndValidate(ctx *gin.Context, target interface{}) error {
	if err := json.NewDecoder(ctx.Request.Body).Decode(target); err != nil {
		return &CustomError{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
	}

	// Validate struct
	validator := validator.New()
	if err := validator.Struct(target); err != nil {

		validationErrorsSlice := strings.Split(err.Error(), "\n")

		return &CustomError{
			Code:    http.StatusBadRequest,
			Message: validationErrorsSlice[0],
		}
	}

	return nil
}

func HandleRequestValidation(ctx *gin.Context, target interface{}) error {
	err := DecodeAndValidate(ctx, target)
	return err
}
