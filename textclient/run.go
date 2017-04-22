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

	gamelib "github.com/dougfort/arachne/game"
)

var (
	moveRegex *regexp.Regexp
)

//
func init() {
	moveRegex = regexp.MustCompile(`^\D*(\d+)\s*,\s*\D*(\d+)\D*(\d+)`)
}

type gameData struct {
	remote *gamelib.Game // remote game
}

// run is the actual main body of the program
// it returns an exit code to main
func run() int {
	var err error
	var exitCode int
	var game gameData

	log.Printf("info: start")
	fmt.Println("arachne starts")
	fmt.Println("")
	if game, err = newGame(); err != nil {
		fmt.Printf("newGame failed: %s\n", err)
		return -1
	}
	displayGameData(game)
	fmt.Println("")
	fmt.Print(">")

	scanner := bufio.NewScanner(os.Stdin)
RUN_LOOP:
	for scanner.Scan() {

		rawLine := scanner.Text()
		splitLine := strings.SplitN(rawLine, " ", 2)
		if len(splitLine) == 0 {
			fmt.Println("unarsable command")
			continue RUN_LOOP
		}
		switch splitLine[0] {
		case "new":
			fmt.Println("starting new game")
			if game, err = newGame(); err != nil {
				fmt.Printf("newGame failed: %s\n", err)
				break RUN_LOOP
			}
			displayGameData(game)
		case "display":
			displayTableauStrings(game)
		case "scan":
			displayMoves(game)
		case "move":
			var move gamelib.MoveType
			if move, err = parseMoveComand(splitLine); err != nil {
				fmt.Printf("unparseable move command: %s\n", err)
				continue RUN_LOOP
			}
			if err = processMove(game, move); err != nil {
				fmt.Printf("unparseable move command: %s\n", err)
				continue RUN_LOOP
			}
			displayGameData(game)
		case "deal":
			if game.remote.Deck.RemainingCards() == 0 {
				fmt.Println("no cards available to deal")
				continue RUN_LOOP
			}
			if game.remote.Tableau.EmptyStack() {
				fmt.Println("cannot deal over empty stack")
				continue RUN_LOOP
			}
			if err = game.remote.Deal(); err != nil {
				fmt.Printf("deal failed: %s\n", err)
				continue RUN_LOOP
			}
			displayGameData(game)
		case "quit":
			fmt.Println("quitting")
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

func displayGameData(game gameData) {
	fmt.Printf("cards remaining: %d\n", game.remote.Deck.RemainingCards())
	fmt.Println("")
	displayTableauStrings(game)
	fmt.Println("")
	displayMoves(game)
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
