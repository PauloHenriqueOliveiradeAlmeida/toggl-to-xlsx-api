package Middlewares

import (
	"net/http"
	"reflect"
	"toggl-xlsx-back/src/Application/Errors"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Next()
		lastError := context.Errors.Last()
		if lastError == nil {
			return
		}

		errorTypes := map[reflect.Type]int{
			reflect.TypeOf(Errors.BadRequest{}):         http.StatusBadRequest,
			reflect.TypeOf(Errors.ServiceUnavailable{}): http.StatusServiceUnavailable,
			reflect.TypeOf(Errors.InternalError{}):      http.StatusInternalServerError,
		}
		errorCode := errorTypes[reflect.TypeOf(lastError.Err).Elem()]
		context.AbortWithStatusJSON(errorCode, gin.H{"message": lastError.Error()})
	}
}
