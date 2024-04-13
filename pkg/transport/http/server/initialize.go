package server

import (
	"errors"
	"fmt"

	"grid/pkg/env"
	"grid/pkg/lifecycle"
	"grid/pkg/transport/http/middleware/cors"
	"grid/pkg/transport/http/routes"
	"grid/pkg/transport/http/validator"

	"github.com/gin-gonic/gin"
)

var (
	Initializers = lifecycle.Fns[gin.Engine]{
		cors.Initialize,
		routes.Initialize,
		validator.Initialize,
	}

	ServerAlreadyInitializedError = errors.New("HTTP server has already been initialized")
	ServerNotInitializedError     = errors.New("HTTP server has not been initialized")

	instance *gin.Engine
)

func Initialize(env *env.Vars) error {
	// check

	if instance != nil {
		return ServerAlreadyInitializedError
	}

	engine, err := New()

	if err != nil {
		return err
	}

	// run

	go func() {
		engine.Run(fmt.Sprintf(":%s", env.HTTPServerPort))
	}()

	instance = engine

	// success

	return nil
}

func New() (*gin.Engine, error) {
	// create

	engine := gin.Default()

	// initialize

	err := Initializers.Execute(engine)

	if err != nil {
		return nil, err
	}

	// success

	return engine, nil
}
