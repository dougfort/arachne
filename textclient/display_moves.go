package main

import "fmt"

func displayMoves(g gameData) {
	for i, move := range g.remote.Tableau.EnumerateMoves() {
		fmt.Printf("%2d: (%2d, %2d) -> %2d\n",
			i+1,
			move.FromCol+1,
			move.FromRow+1,
			move.ToCol+1,
		)
	}
}
