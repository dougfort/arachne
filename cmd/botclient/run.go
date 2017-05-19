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
		logger.Log("game-type", "new", "seed", lg.Seed)
	} else {
		if lg, err = c.ReplayGame(cfg.seed); err != nil {
			logger.Log("error", fmt.Sprintf("c.NewGame(): %v", err))
			return -1
		}
		logger.Log("game-type", "replay", "seed", lg.Seed)
	}

	var turn int
	turnValuer := func() log.Valuer {
		return func() interface{} {
			return turn
		}
	}()
	logger = log.With(logger, "turn", turnValuer)

TURN_LOOP:
	for turn = 1; cfg.maxTurns == -1 || turn <= cfg.maxTurns; turn++ {
		availableMoves := lg.Tableau.EnumerateMoves()
		if len(availableMoves) == 0 {
			if lg.CardsRemaining == 0 {
				logger.Log("end", "deadlock")
				break TURN_LOOP
			}
			logger.Log("deal", "")
			if lg, err = c.Deal(); err != nil {
				logger.Log("error", fmt.Sprintf("Deal() %v", err))
				return -1
			}
			continue TURN_LOOP
		}

		// TODO: analyze the moves,instead of picking the first one
		move := availableMoves[0]

		if lg, err = c.Move(move.MoveType); err != nil {
			logger.Log("error", fmt.Sprintf("Move(%s): %v", move.MoveType, err))
			return -1
		}

		logger.Log("move", move)
	}

	return 0
}
