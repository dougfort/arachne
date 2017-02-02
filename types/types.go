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
