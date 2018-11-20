package ginerr

import "github.com/gin-gonic/gin"

// ApiError is interface of ApiError
type ApiError interface {
	GetCode() int
	GetMessage() string
}

// ApiController is function type we wrapped
type ApiController = func(ctx *gin.Context) ApiError

// CreateGinController trans our controller into gin controller
func CreateGinController(apiController ApiController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		apiError := apiController(ctx)
		if apiError != nil {
			ctx.JSON(apiError.GetCode(), gin.H{"message": apiError.GetMessage()})
		}
	}
}

// DefaultError is default error struct impl ApiError
type DefaultError struct {
	Code    int
	Message string
}

// GetCode impl GetCode
func (de *DefaultError) GetCode() int {
	return de.Code
}

// GetMessage impl GetMessage
func (de *DefaultError) GetMessage() string {
	return de.Message
}

// NewDefaultError return a new default error
func NewDefaultError(code int, message string) *DefaultError {
	return &DefaultError{Code: code, Message: message}
}
