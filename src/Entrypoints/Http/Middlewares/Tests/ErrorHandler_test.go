package Tests

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"toggl-xlsx-back/src/Application/Errors"
	"toggl-xlsx-back/src/Entrypoints/Http/Middlewares"
)

func TestWithErrorThrown(test *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.Use(Middlewares.ErrorHandler())
	error := &Errors.BadRequest{Message: "Bad request"}
	router.GET("/with-error", func(context *gin.Context) {
		context.Error(error)
		return
	})

	request, _ := http.NewRequest(http.MethodGet, "/with-error", nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)
	responseBody, _ := json.Marshal(error)
	assert.Equal(test, string(responseBody), recorder.Body.String())
}

func TestWithoutErrorThrown(test *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.Use(Middlewares.ErrorHandler())

	response := gin.H{"message": "Success"}
	router.GET("/without-error", func(context *gin.Context) {
		context.JSON(http.StatusOK, response)
	})

	request, _ := http.NewRequest(http.MethodGet, "/without-error", nil)
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	assert.Equal(test, http.StatusOK, recorder.Code)

	responseBody, _ := json.Marshal(response)
	assert.Equal(test, string(responseBody), recorder.Body.String())
}
