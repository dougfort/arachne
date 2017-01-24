package api

import "github.com/dougfort/arachne/types"

// GameInfo reports the state of a game
type GameInfo struct {

	// Token identifies the game to the server
	Token []byte

	// RemainingCards is the number of cards left to deal
	RemainingCards int

	// Tableau is the actual cards laid out in stacks
	Tableau types.Tableau
}

// ArachneAPI describes the entry points used to communicate with arachneserv
type ArachneAPI interface {

	// NewGame requests the start of a new game
	NewGame() (GameInfo, error)
}
