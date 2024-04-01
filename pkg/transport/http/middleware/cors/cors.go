package cors

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Initialize(engine *gin.Engine) error {
	engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"DELETE", "HEAD", "GET", "OPTIONS", "POST", "PUT", "PATCH"},
		AllowHeaders:     []string{"Authorization", "Content-Length", "Content-Type", "Origin"},
		ExposeHeaders:    []string{"Authorization", "Content-Length", "Content-Type", "Origin"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 12 * time.Hour,
	}))

	return nil
}
