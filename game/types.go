package game

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

// MoveType describes the transfer of a slice of Cards from from the end of one
// Stack to the end of another Stack within a Tableau
type MoveType struct {

	// FromCol is the index of the Stack where the move originates
	FromCol int

	// FromRow is the index of the first Card in a slice to be moved
	FromRow int

	// TOCol is the index of the stack to which the move slice is appended
	ToCol int
}
