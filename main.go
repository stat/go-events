package main

import (
	"events/app"
	"events/pkg/env"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// vars

	vars, err := env.Load()

	if err != nil {
		panic(err)
	}

	// initialize

	err = app.Initialize(vars)

	if err != nil {
		panic(err)
	}

	// await termination

	cExit := make(chan os.Signal, 1)
	signal.Notify(cExit, os.Interrupt, syscall.SIGTERM)

	<-cExit

	// terminate

	if err := app.Terminate(vars); err != nil {
		panic(err)
	}

	// log shutdown

	log.Println("gracefully shutdown!")
}
