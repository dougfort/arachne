package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/ardanlabs/kit/cfg"
	"github.com/ardanlabs/kit/log"

	"github.com/dougfort/arachne/internal/client"
	"github.com/dougfort/arachne/internal/game"
)

// run is the actual main body of the program
// it returns an exit code to main
func run() int {
	const fname = "run"
	const cfgNamespace = "arachne"
	var address string
	var seedStr string
	var seed int64
	var logCtx string
	var maxTurns int
	var c client.Client
	var lg client.LocalGame
	var orderer game.Orderer
	var err error

	err = cfg.Init(cfg.EnvProvider{Namespace: cfgNamespace})
	if err != nil {
		panic(err)
	}

	parseCommandLine()

	logLevel := func() int {
		ll, err := cfg.Int("LOGGING_LEVEL")
		if err != nil {
			return log.DEV
		}
		return ll
	}
	log.Init(os.Stderr, logLevel, log.Ldefault)

	address = cfg.MustString("ADDRESS")

	c, err = client.New(address)
	if err != nil {
		log.Error(logCtx, fname, err, "client.New(%s)", address)
		return -1
	}

	// TODO: rig kit to handle int64, etc
	seedStr, err = cfg.String("SEED")
	if err == nil {
		if seed, err = strconv.ParseInt(seedStr, 10, 64); err != nil {
			log.Error(logCtx, fname, err, "strconv.ParseInt(%s)", seedStr)
			return -1
		}
	}

	if seed == 0 {
		if lg, err = c.NewGame(); err != nil {
			log.Error(logCtx, fname, err, "c.NewGame()")
			return -1
		}
		logCtx = fmt.Sprintf("%d", lg.Seed)
		log.User(logCtx, fname, "new game")
	} else {
		if lg, err = c.ReplayGame(seed); err != nil {
			log.Error(logCtx, fname, err, "c.NewGame()")
			return -1
		}
		logCtx = fmt.Sprintf("%d", lg.Seed)
		log.User(logCtx, fname, "replay")
	}

	orderer = game.NewHighestMove()
	maxTurns = cfg.MustInt("MAX_TURNS")
	completedMoves := make(map[game.MoveType]struct{})

TURN_LOOP:
	for turn := 1; maxTurns == -1 || turn <= maxTurns; turn++ {

		var availableMoves []game.EvaluatedMoveType
		enumeratedMoves := lg.Tableau.EnumerateMoves()
		for _, move := range enumeratedMoves {
			_, ok := completedMoves[move.MoveType]
			if !ok {
				availableMoves = append(availableMoves, move)
			}
		}

		if len(availableMoves) == 0 {
			if lg.CardsRemaining == 0 {
				log.User(logCtx, fname, "ends in deadlock")
				break TURN_LOOP
			}
			log.User(logCtx, fname, "deal")
			if lg, err = c.Deal(); err != nil {
				log.Error(logCtx, fname, err, "Deal()")
				return -1
			}
			completedMoves = make(map[game.MoveType]struct{})
			continue TURN_LOOP
		}

		orderedMoves, err := orderer.Order(availableMoves)
		if err != nil {
			log.Error(logCtx, fname, err, "orderer.Order")
			return -1
		}

		// Pick the most highly rated move
		move := orderedMoves[len(orderedMoves)-1]

		if lg, err = c.Move(move.MoveType); err != nil {
			log.Error(logCtx, fname, err, "Move(%s)", move.MoveType)
			return -1
		}

		log.User(logCtx, fname, move.String())
		completedMoves[move.EvaluatedMoveType.MoveType] = struct{}{}
	}

	return 0
}
