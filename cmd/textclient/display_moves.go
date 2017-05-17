package main

import (
	"fmt"

	"github.com/dougfort/arachne/internal/game"
)

func displayMoves(tableau game.Tableau) {
	for i, move := range tableau.EnumerateMoves() {
		fmt.Printf("%2d: %s\n", i+1, move)
	}
}
