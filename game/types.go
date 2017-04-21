package game

import (
	"fmt"

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

// String shows the move in human readable form
// Note that displayed coordinates start at 1
func (m MoveType) String() string {
	return fmt.Sprintf("(%2d, %2d) -> %2d",
		m.FromCol+1,
		m.FromRow+1,
		m.ToCol+1,
	)
}

// EvaluatedMoveType is MoveType with evaluation
type EvaluatedMoveType struct {
	MoveType

	// FromCount is the size of the slice to be moved from the origin
	FromCount int

	// ToCount is the number of Cards at the destination which have the same
	// Suit as the incoming slice and which continue the Rank sequence.
	// If the move is made, there will be FromCount + ToCount Cards avaialble
	// to move from the destination
	ToCount int
}

// String shows the move in human readable form
// Note that displayed coordinates start at 1
func (m EvaluatedMoveType) String() string {
	return fmt.Sprintf("(%2d, %2d) -> %2d: from: %d + to: %d",
		m.FromCol+1,
		m.FromRow+1,
		m.ToCol+1,
		m.FromCount,
		m.ToCount,
	)
}
