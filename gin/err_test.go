package ginerr

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestCreateGinController(t *testing.T) {
	r := gin.Default()
	r.GET("/", CreateGinController(func(ctx *gin.Context) ApiError {
		return NewDefaultError(400, "validate_error", "Validate error.")
	}))

	w := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/", nil)
	r.ServeHTTP(w, request)

	assert.Equal(t, 400, w.Code)
	assert.Regexp(t, "validate_error", w.Body.String())
}
