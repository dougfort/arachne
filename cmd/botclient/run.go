package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/go-kit/kit/log"

	"github.com/ardanlabs/kit/cfg"

	"github.com/dougfort/arachne/internal/client"
)

// run is the actual main body of the program
// it returns an exit code to main
func run() int {
	const cfgNamespace = "arachne"
	const defaultAddress = ":10000"
	var address string
	var seedStr string
	var seed int64
	var maxTurns int
	var logger log.Logger
	var c client.Client
	var lg client.LocalGame
	var err error

	err = cfg.Init(cfg.EnvProvider{Namespace: cfgNamespace})
	if err != nil {
		panic(err)
	}

	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))

	parseCommandLine()

	if address, err = cfg.String("ADDRESS"); err != nil {
		address = defaultAddress
	}

	c, err = client.New(address)
	if err != nil {
		logger.Log("error", fmt.Sprintf("client.New(): %v", err))
		return -1
	}

	// TODO: rig kit to handle int64, etc
	seedStr, err = cfg.String("SEED")
	if err == nil {
		if seed, err = strconv.ParseInt(seedStr, 10, 64); err != nil {
			logger.Log("error", fmt.Sprintf("strconv.ParseInt: %v", err))
			return -1
		}
	}

	if seed == 0 {
		if lg, err = c.NewGame(); err != nil {
			logger.Log("error", fmt.Sprintf("c.NewGame(): %v", err))
			return -1
		}
		logger.Log("game-type", "new", "seed", lg.Seed)
	} else {
		if lg, err = c.ReplayGame(seed); err != nil {
			logger.Log("error", fmt.Sprintf("c.NewGame(): %v", err))
			return -1
		}
		logger.Log("game-type", "replay", "seed", lg.Seed)
	}

	maxTurns = cfg.MustInt("MAX_TURNS")

	var turn int
	turnValuer := func() log.Valuer {
		return func() interface{} {
			return turn
		}
	}()
	logger = log.With(logger, "turn", turnValuer)

TURN_LOOP:
	for turn = 1; maxTurns == -1 || turn <= maxTurns; turn++ {
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
