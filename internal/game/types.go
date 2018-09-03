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
	return fmt.Sprintf("(%2d, %2d) -> %2d",
		m.FromCol+1,
		m.FromRow+1,
		m.ToCol+1,
	)
}

// RankedMoveType is EvaluatedMoveType assigned a rank
type RankedMoveType struct {
	EvaluatedMoveType

	// Rank is an arbitrary number assigned by the Orderer
	// The higher the Rank, the more valuable the move is presumed to be
	Rank int
}

// String shows the ranked move in human readable form
// Note that displayed coordinates start at 1
func (m RankedMoveType) String() string {
	return fmt.Sprintf("(%2d, %2d) -> %2d: rank = %d",
		m.FromCol+1,
		m.FromRow+1,
		m.ToCol+1,
		m.Rank,
	)
}

// Orderer an interface for objects which put possible moves in some
// (preferential) order
type Orderer interface {

	// Order puts possible moves in some (preferential) order
	Order([]EvaluatedMoveType) ([]RankedMoveType, error)
}
