package routes

import (
	aircrafts "events/pkg/transport/http/handlers/aircrafts"
	aircrafts_events "events/pkg/transport/http/handlers/aircrafts/events"

	"github.com/gin-gonic/gin"
)

func Initialize(engine *gin.Engine) error {
	// register

	routes := engine.Group("v1.0")
	{
		routes.GET("/aircrafts", aircrafts.List)
		routes.GET("/aircrafts/:aircraft_id", aircrafts.Get)
		routes.GET("/aircrafts/:aircraft_id/events")

		routes.POST("/aircrafts/:aircraft_id/events", aircrafts_events.Create)
	}

	// success

	return nil
}
