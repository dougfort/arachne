package main

import (
	"fmt"

	"github.com/pkg/errors"

	"github.com/dougfort/arachne/internal/game"
)

func displayMoves(tableau game.Tableau, orderer game.Orderer) error {
	moves := tableau.EnumerateMoves()
	if err := orderer.Order(moves); err != nil {
		return errors.Wrap(err, "Order(moves)")
	}

	for i, move := range moves {
		fmt.Printf("%2d: %s\n", i+1, move)
	}

	return nil
}
