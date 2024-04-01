package server

import (
	"github.com/gin-gonic/gin"
)

func Engine() (*gin.Engine, error) {
	if instance == nil {
		return nil, ServerNotInitializedError
	}

	return instance, nil
}
