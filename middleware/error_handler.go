package middleware

import (
	"errors"
	"net/http"

	"github.com/anyTV/gomodules/response"
	"github.com/anyTV/gomodules/ferrors"
	logger "github.com/anyTV/gomodules/logging"
	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		for _, err := range c.Errors {
			var e ferrors.HttpError
			switch {

			case errors.As(err.Err, &e):

				logger.Warnf(e.Error())

				c.AbortWithStatusJSON(e.Status, e)
			default:

				logger.Warnf("Internal error: %s", e.Error())

				c.AbortWithStatusJSON(
					http.StatusInternalServerError,
					gin.H{"code": response.InternalServerError},
				)
			}
		}
	}
}
