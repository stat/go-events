package aircrafts

import (
	"net/http"

	"grid/pkg/transport/http/respond"

	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	respond.With(c, http.StatusOK, "")
}
