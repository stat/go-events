package aircrafts

import (
	"net/http"

	"grid/pkg/repos/cache"
	"grid/pkg/transport/http/respond"

	"github.com/gin-gonic/gin"
)

type URIParams struct {
	AircraftID string `uri:"aircraft_id"`
}

func Get(c *gin.Context) {
	uriParams := &URIParams{}

	if err := c.ShouldBindUri(uriParams); err != nil {
		respond.WithError(c, http.StatusBadRequest, err)
		return
	}

	// lookup

	location, err := cache.Backend.GetAircraftLocation(uriParams.AircraftID)

	if err != nil {
		respond.WithError(c, http.StatusNotFound, err)
		return
	}

	respond.With(c, http.StatusOK, location)
}
