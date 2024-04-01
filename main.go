package main

import (
	"grid/app"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	vars, err := app.Initialize()

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
