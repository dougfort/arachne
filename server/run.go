package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

// run is the actual main body of the program
// it returns an exit code to main
func run() int {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	log.Printf("info: wait signal")
	s := <-sigChan
	log.Printf("info: signal %v", s)

	return 0
}
