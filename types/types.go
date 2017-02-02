package types

import (
	"github.com/dougfort/gocards"
)

// StackType represents one stack of cards in the Tableau
type StackType struct {
	HiddenCount int
	Cards       gocards.Cards
}

// TableauWidth is the number of stacks in the Tableau
const TableauWidth = 10

// Tableau is the outer (visible) game layout
type Tableau [TableauWidth]StackType

// HiddenCards represents the Cards that are not visible in the Tableau
type HiddenCards [TableauWidth]gocards.Card

// Game represents a complete, playable, game
type Game struct {
	Deck        gocards.PlayableDeck
	Tableau     Tableau
	HiddenCards HiddenCards
}
