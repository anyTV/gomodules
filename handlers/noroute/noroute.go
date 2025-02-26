package noroute

import (
	"net/http"

	"github.com/anyTV/gomodules/config"
	"github.com/anyTV/gomodules/ferrors"
	"github.com/anyTV/gomodules/response"
	"github.com/gin-gonic/gin"
)

func N(c *gin.Context) {
	c.Error(
		ferrors.NewHttpError(
			http.StatusNotFound,
			response.NothingTodoHere,
			config.EmptyString,
		),
	)

	return
}
