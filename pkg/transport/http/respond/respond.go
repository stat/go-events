package respond

import (
	"events/pkg/utils"

	"github.com/gin-gonic/gin"
)

type Response[T any] struct {
	Data    *T      `json:"data"`
	Error   *string `json:"error"`
	Message *string `json:"message"`
}

func With[T any](c *gin.Context, code int, data *T) {
	c.JSON(
		code,
		&Response[T]{
			Data: data,
		},
	)
}

func WithError(c *gin.Context, code int, err error) {
	c.JSON(
		code,
		&Response[interface{}]{
			Error: utils.Ref(err.Error()),
		},
	)
}
