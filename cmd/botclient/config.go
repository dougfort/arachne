package main

import "flag"

type config struct {
	seed     int64
	maxTurns int
}

func parseCommandLine() config {
	var cfg config
	flag.Int64Var(&cfg.seed, "seed", 0, "seed to be used randomizing the game deck")
	flag.IntVar(&cfg.maxTurns, "max-turns", -1, "maximum number of turns to play")
	flag.Parse()

	return cfg
}
