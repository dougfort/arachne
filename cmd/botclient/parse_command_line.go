package main

import (
	"flag"
	"fmt"

	"github.com/ardanlabs/kit/cfg"
)

// parseCommandLine stores commandline arguments in cfg
func parseCommandLine() {
	var seed int64
	var maxTurns int

	flag.Int64Var(&seed, "seed", 0, "seed to be used randomizing the game deck")
	flag.IntVar(&maxTurns, "max-turns", -1, "maximum number of turns to play")
	flag.Parse()

	cfg.SetString("SEED", fmt.Sprintf("%d", seed))
	cfg.SetInt("MAX_TURNS", maxTurns)
}
