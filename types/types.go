package types

import (
	"github.com/dougfort/gocards"
)

// StackType reresesnts on stack of cards in the Tableau
type StackType struct {
	HiddenCount int
	Cards       gocards.Cards
}

// TableauWidth is the number of stacks in the Tableau
const TableauWidth = 10

// Tableau is the game layout
type Tableau []StackType
