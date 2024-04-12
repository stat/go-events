package env_test

import (
	"testing"

	"grid/pkg/env"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoad(t *testing.T) {
	vars, err := env.Load()

	require.Nil(t, err)
	assert.NotEmpty(t, vars.ASYNQConcurrency)
	assert.NotEmpty(t, vars.HTTPServerPort)
	assert.NotEmpty(t, vars.PostgresUser)
	assert.NotEmpty(t, vars.PostgresPassword)
	assert.NotEmpty(t, vars.PostgresHost)
	assert.NotEmpty(t, vars.PostgresPort)
	assert.NotEmpty(t, vars.PostgresDB)
	assert.NotEmpty(t, vars.PostgresSchema)
	assert.NotEmpty(t, vars.PostgresTimezone)
	assert.NotEmpty(t, vars.RedisDB)
	assert.NotEmpty(t, vars.RedisDBAsynq)
	assert.NotEmpty(t, vars.RedisDBCache)
	assert.NotEmpty(t, vars.RedisDBEvents)
	assert.NotEmpty(t, vars.RedisHost)
	assert.NotEmpty(t, vars.RedisPort)
	assert.NotEmpty(t, vars.WebSocketPort)

	assert.True(t, vars.InitAll)
	assert.False(t, vars.InitWorkers)

}
