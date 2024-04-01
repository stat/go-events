package aircrafts_events

import (
	"grid/pkg/models"
	"net/http"

	"grid/pkg/db/queue"
	"grid/pkg/tasks/consumer"
	"grid/pkg/transport/http/respond"

	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	// bind JSON body

	event := &models.ADSB{}

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
