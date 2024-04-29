package aircrafts_events

import (
	"events/pkg/models"
	"net/http"

	"events/pkg/db/queue"
	"events/pkg/tasks/consumer"
	"events/pkg/transport/http/respond"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	// bind JSON body

	event := &models.LocationEvent{}

	if err := c.ShouldBindJSON(event); err != nil {
		respond.WithError(c, http.StatusBadRequest, err)
		return
	}

	// enqueue

	task, err := consumer.NewTask(event)

	if err != nil {
		respond.WithError(c, http.StatusInternalServerError, err)
		return
	}

	queue, err := queue.Instance()

	if err != nil {
		respond.WithError(c, http.StatusServiceUnavailable, err)
		return
	}

	info, err := queue.Enqueue(task)

	if err != nil {
		respond.WithError(c, http.StatusServiceUnavailable, err)
		return
	}

	// success

	respond.With(c, http.StatusCreated, info)
}
