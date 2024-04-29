package aircrafts

import (
	"net/http"

	"events/pkg/models"
	"events/pkg/transport/http/respond"

	"github.com/gin-gonic/gin"
)

func Get(c *gin.Context) {
	uriParams := &URIParams{}

	if err := c.ShouldBindUri(uriParams); err != nil {
		respond.WithError(c, http.StatusBadRequest, err)
		return
	}

	// lookup

	location := &models.LocationEvent{}

	// location, err := cache.Backend.GetAircraftLocation(uriParams.AircraftID)

	// if err != nil {
	//   respond.WithError(c, http.StatusNotFound, err)
	//   return
	// }

	// success

	respond.With(c, http.StatusOK, location)
}
