package main

import (
	"fmt"

	"github.com/pkg/errors"

	"github.com/dougfort/arachne/internal/game"
)

func displayMoves(tableau game.Tableau, orderer game.Orderer) error {
	eSlice := tableau.EnumerateMoves()

	rSlice, err := orderer.Order(eSlice)
	if err != nil {
		return errors.Wrap(err, "Order(eSlice)")
	}

	for i, move := range rSlice {
		fmt.Printf("%2d: %s\n", i+1, move)
	}

	return nil
}
