package main

import "fmt"

func displayMoves(g gameData) {
	for i, move := range g.remote.Tableau.EnumerateMoves() {
		fmt.Printf("%2d: %s\n", i+1, move)
	}
}
