package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/dougfort/arachne/game"
)

type gameData struct {
	remote game.Game // remote game
}

// run is the actual main body of the program
// it returns an exit code to main
func run() int {
	var err error
	var exitCode int
	var game gameData

	log.Printf("info: start")
	fmt.Println("arachne starts")

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("")
	fmt.Printf(">")

RUN_LOOP:
	for scanner.Scan() {

		displayTableau(game)

		switch scanner.Text() {
		case "new":
			fmt.Println("starting new game")
			if game, err = newGame(); err != nil {
				fmt.Printf("newGame failed: %s", err)
				break RUN_LOOP
			}
			log.Printf("debug: game = %v", game)
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
