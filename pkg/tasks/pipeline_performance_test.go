package tasks_test

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"testing"
	"time"

	"events/pkg/models"
	"events/pkg/tasks/consumer"
	"events/pkg/tasks/producer"
	"events/pkg/utils"

	"github.com/stretchr/testify/assert"

	_ "events/testing"
)

// TODO: refactor me
// TODO: add fuzzing to force an error count with bad data

func TestPipelinePerformance(t *testing.T) {
	maxAircrafts := int(1e4)
	maxEvents := int(1e4)
	maxStations := int(1e6)

	threads := 10

	wg := sync.WaitGroup{}
	errs := make(chan error)
	sent := make(chan int)

	start := time.Now()
	stop := time.Time{}

	for i := 0; i < threads; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			count := maxEvents / threads

			for i := 0; i < count; i++ {
				aircraftID := fmt.Sprintf("aircraft-%d", rand.Intn(maxAircrafts)+1)
				stationID := fmt.Sprintf("station-%d", rand.Intn(maxStations)+1)

				event := &models.LocationEvent{
					AircraftID: aircraftID,
					Latitude:   1.0,
					Longitude:  1.0,
					StationID:  stationID,
					Timestamp:  utils.Ref(time.Now()),
				}

				err := consumer.Process(event)

				if err != nil {
					errs <- err
					continue
				}

				err = producer.Process(event)

				if err != nil {
					errs <- err
				}
			}

			sent <- count
		}()
	}

	go func() {
		wg.Wait()

		stop = time.Now()

		close(errs)
		close(sent)
	}()

	// await

	await := sync.WaitGroup{}

	// errors

	await.Add(1)
	errors := 0

	go func() {
		defer await.Done()

		for err := range errs {
			errors++
			// log.Fatalln(err)
			log.Println(err)
		}
	}()

	// count

	await.Add(1)
	total := 0

	go func() {
		defer await.Done()
		for count := range sent {
			total += count
		}
	}()

	await.Wait()

	// duration

	duration := stop.Sub(start)
	durationInSeconds := float64(duration) / float64(time.Second)

	// rate

	rate := float64(total) / durationInSeconds

	fmt.Printf(
		"\n-----\nduration: %v\nsent: %d\n => %f events/second\n-----\n",
		duration,
		total,
		rate,
	)

	assert.Equal(t, maxEvents, total)
	assert.Equal(t, 0, errors, "unexpected errors found")
}
