package ginerr

import "github.com/gin-gonic/gin"

// ApiError is interface of ApiError
type ApiError interface {
	GetStatusCode() int
	GetCode() string
	GetMessage() string
}

// ApiController is function type we wrapped
type ApiController = func(ctx *gin.Context) ApiError

// CreateGinController trans our controller into gin controller
func CreateGinController(apiController ApiController) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		apiError := apiController(ctx)
		if apiError != nil {
			ctx.JSON(apiError.GetStatusCode(), gin.H{"message": apiError.GetMessage(), "code": apiError.GetCode()})
		}
	}
}

// DefaultError is default error struct impl ApiError
type DefaultError struct {
	StatusCode int
	Code       string
	Message    string
}

// GetStatusCode impl GetStatusCode
func (de *DefaultError) GetStatusCode() int {
	return de.StatusCode
}

// GetCode impl GetCode
func (de *DefaultError) GetCode() string {
	return de.Code
}

// GetMessage impl GetMessage
func (de *DefaultError) GetMessage() string {
	return de.Message
}

// NewDefaultError return a new default error
func NewDefaultError(statusCode int, code string, message string) *DefaultError {
	return &DefaultError{StatusCode: statusCode, Code: code, Message: message}
}
