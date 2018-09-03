package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/pkg/errors"

	"github.com/ardanlabs/kit/cfg"

	"github.com/dougfort/arachne/internal/client"

	gamelib "github.com/dougfort/arachne/internal/game"
)

var (
	moveRegex *regexp.Regexp
)

//
func init() {
	moveRegex = regexp.MustCompile(`^\D*(\d+)\s*,\s*\D*(\d+)\D*(\d+)`)
}

// main is the main entry point
// all it does is exit with an error code
func main() {
	os.Exit(run())
}

// run is the actual main body of the program
// it returns an exit code to main
func run() int {
	const cfgNamespace = "arachne"
	const defaultAddress = ":10000"
	var address string
	var err error
	var exitCode int
	var c client.Client
	var lg client.LocalGame

	err = cfg.Init(cfg.EnvProvider{Namespace: cfgNamespace})
	if err != nil {
		panic(err)
	}

	if address, err = cfg.String("ADDRESS"); err != nil {
		address = defaultAddress
	}

	c, err = client.New(address)
	if err != nil {
		fmt.Printf("unable to create client: %v\n", err)
		return -1
	}

	fmt.Println("arachne starts")
	fmt.Println("")
	if lg, err = c.NewGame(); err != nil {
		fmt.Printf("NewGame failed: %v\n", err)
		return -1
	}

	orderer := gamelib.NewHighestMove()

	if err = displayGameData(lg, orderer); err != nil {
		fmt.Printf("displayGameData failed: %v\n", err)
		return -1
	}

	fmt.Println("")
	fmt.Print(">")

	scanner := bufio.NewScanner(os.Stdin)
RUN_LOOP:
	for scanner.Scan() {

		rawLine := scanner.Text()
		splitLine := strings.SplitN(rawLine, " ", 2)
		if len(splitLine) == 0 {
			fmt.Println("unparsable command")
			continue RUN_LOOP
		}
		switch splitLine[0] {
		case "new":
			fmt.Println("starting new random game")
			if lg, err = c.NewGame(); err != nil {
				fmt.Printf("NewGame failed: %v\n", err)
				break RUN_LOOP
			}
			if err = displayGameData(lg, orderer); err != nil {
				fmt.Printf("displayGameData failed: %v\n", err)
				break RUN_LOOP
			}
		case "replay":
			if len(splitLine) < 2 {
				fmt.Println("no seed value given for replay")
				continue RUN_LOOP
			}
			seed, err := strconv.ParseInt(splitLine[1], 10, 64)
			if err != nil {
				fmt.Printf("Unable to parse seed '%s'\n", splitLine[1])
				continue RUN_LOOP
			}
			fmt.Println("replaying game from seed %d", seed)
			if lg, err = c.ReplayGame(seed); err != nil {
				fmt.Printf("ReplayGame(%d) failed: %v\n", seed, err)
				break RUN_LOOP
			}
		case "display":
			displayTableauStrings(lg.Tableau)
		case "scan":
			if err = displayMoves(lg.Tableau, orderer); err != nil {
				fmt.Printf("Unable to display tableau: %s\n", err)
				continue RUN_LOOP
			}
		case "move":
			var move gamelib.MoveType
			if move, err = parseMoveComand(splitLine); err != nil {
				fmt.Printf("unparseable move command: %s\n", err)
				continue RUN_LOOP
			}
			if lg, err = c.Move(move); err != nil {
				fmt.Printf("move failed: %s\n", err)
				continue RUN_LOOP
			}
			if err = displayGameData(lg, orderer); err != nil {
				fmt.Printf("displayGameData failed: %v\n", err)
				break RUN_LOOP
			}
		case "deal":
			if lg.CardsRemaining == 0 {
				fmt.Println("no cards available to deal")
				continue RUN_LOOP
			}
			if lg.Tableau.EmptyStack() {
				fmt.Println("cannot deal over empty stack")
				continue RUN_LOOP
			}
			if lg, err = c.Deal(); err != nil {
				fmt.Printf("deal failed: %v\n", err)
				continue RUN_LOOP
			}
			if err = displayGameData(lg, orderer); err != nil {
				fmt.Printf("displayGameData failed: %v\n", err)
				break RUN_LOOP
			}
		case "quit":
			fmt.Println("quitting")
			if err := c.Close(); err != nil {
				fmt.Printf("Close failed: %v\n", err)
			}
			break RUN_LOOP

		default:
			fmt.Println("unknown command")
		}
		fmt.Println("")
		fmt.Printf(">")
	}
	log.Printf("info: end")

	return exitCode
}

func displayGameData(lg client.LocalGame, orderer gamelib.Orderer) error {
	fmt.Printf("seed: %d\n", lg.Seed)
	fmt.Println("")
	fmt.Printf("captures: %d\n", lg.CaptureCount)
	fmt.Println("")
	fmt.Printf("cards remaining: %d\n", lg.CardsRemaining)
	fmt.Println("")
	displayTableauStrings(lg.Tableau)
	fmt.Println("")
	if err := displayMoves(lg.Tableau, orderer); err != nil {
		return errors.Wrap(err, "displayMoves")
	}

	return nil
}

func parseMoveComand(splitLine []string) (gamelib.MoveType, error) {
	var move gamelib.MoveType
	var intVal int
	var err error

	if len(splitLine) < 2 {
		return gamelib.MoveType{}, errors.Errorf("invalid line for move: '%q'",
			splitLine)
	}

	parsedLine := moveRegex.FindStringSubmatch(splitLine[1])
	if len(parsedLine) != 4 {
		return gamelib.MoveType{}, errors.Errorf("unparseable line for move: '%s'",
			splitLine[1])
	}

	if intVal, err = strconv.Atoi(parsedLine[1]); err != nil {
		return gamelib.MoveType{}, errors.Wrapf(err, "invalid FromCol '%q'",
			parsedLine)
	}
	move.FromCol = intVal - 1

	if intVal, err = strconv.Atoi(parsedLine[2]); err != nil {
		return gamelib.MoveType{}, errors.Wrapf(err, "invalid FromRow '%q'",
			parsedLine)
	}
	move.FromRow = intVal - 1

	if intVal, err = strconv.Atoi(parsedLine[3]); err != nil {
		return gamelib.MoveType{}, errors.Wrapf(err, "invalid ToCol '%q'",
			parsedLine)
	}
	move.ToCol = intVal - 1

	return move, nil
}
