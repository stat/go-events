package stream_test

import (
	"events/pkg/env"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStream(t *testing.T) {
	// s := &stream.Stream[models.LocationEvent, models.LocationEvent]{
	// Options: &Options[
	// }

	// vars

	vars, err := env.Load()
	require.Nil(t, err)

}
