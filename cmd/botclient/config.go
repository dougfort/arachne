package main

import "flag"

type config struct {
	seed int64
}

func parseCommandLine() config {
	var cfg config
	flag.Int64Var(&cfg.seed, "seed", 0, "seed to be used randomizing the game deck")
	flag.Parse()

	return cfg
}
