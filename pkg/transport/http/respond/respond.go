package respond

import (
	"grid/pkg/utils"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Data    interface{} `json:"data"`
	Error   *string     `json:"error"`
	Message *string     `json:"message"`
}

func With(c *gin.Context, code int, data interface{}) {
	c.JSON(
		code,
		&Response{
			Data: data,
		},
	)
}

func WithError(c *gin.Context, code int, err error) {
	c.JSON(
		code,
		&Response{
			Error: utils.Ref(err.Error()),
		},
	)
}
