package main

import (
	"fmt"
	"os"

	"github.com/go-kit/kit/log"

	"github.com/dougfort/arachne/internal/client"
)

// run is the actual main body of the program
// it returns an exit code to main
func run() int {
	var logger log.Logger
	var cfg config
	var c client.Client
	var lg client.LocalGame
	var err error

	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))

	cfg = parseCommandLine()

	c, err = client.New()
	if err != nil {
		logger.Log("error", fmt.Sprintf("client.New(): %v", err))
		return -1
	}

	if cfg.seed == 0 {
		if lg, err = c.NewGame(); err != nil {
			logger.Log("error", fmt.Sprintf("c.NewGame(): %v", err))
			return -1
		}
	} else {
		if lg, err = c.ReplayGame(cfg.seed); err != nil {
			logger.Log("error", fmt.Sprintf("c.NewGame(): %v", err))
			return -1
		}
	}

	logger.Log("seed", lg.Seed)

	return 0
}
