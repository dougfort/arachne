package game

import (
	"fmt"
	"github.com/dougfort/gocards"
)

// ValidateMove returns nil if the move is valid in the Tableau
// A Move is valid if
// 1. The slice of Cards at 'From' is all of the same Suit
// 2. The slice of Cards at 'From' is in order by Rank from top to bottom
// 3. The bottom Card at 'To' is successor by Rank of the top Card moved
func (t Tableau) ValidateMove(m MoveType) error {
	var err error
	var s gocards.Cards

	if s, err = t.getSliceToMove(m); err != nil {
		return err
	}

	if !stackIndexValid(m.ToCol) || m.ToCol == m.FromCol {
		return fmt.Errorf("invalid ToCol: %d", m.FromCol)
	}

	if len(t[m.ToCol].Cards) > 0 {
		bottomCard := t.getBottomCard(m.ToCol)
		if s[0].Rank != bottomCard.Rank-1 {
			return fmt.Errorf("Rank of move slice top (%d) does not fit ToCol bottom (%d)",
				s[0].Rank, bottomCard.Rank)
		}
	}

	return nil
}

func stackIndexValid(index int) bool {
	return index >= 0 && index < TableauWidth
}

func (t Tableau) getSliceToMove(m MoveType) (gocards.Cards, error) {
	if !stackIndexValid(m.FromCol) {
		return nil, fmt.Errorf("m.FromCol invalid: %d", m.FromCol)
	}

	if !(m.FromRow >= 0 && m.FromRow < len(t[m.FromCol].Cards)) {
		return nil, fmt.Errorf("m.FromRow invalid: %d", m.FromRow)
	}

	s := t[m.FromCol].Cards[m.FromRow:]

	var prev gocards.Card
	// interate from the top (highest Rank) card downto the bottom (lowest Rank)
	for i, card := range s {
		if i > 0 {
			if card.Suit != prev.Suit {
				return nil, fmt.Errorf("move slice not all the same Suit at %d", i)
			}
			if card.Rank != prev.Rank-1 {
				return nil, fmt.Errorf("move slice out of order %d %d at %d",
					prev.Rank, card.Rank, i)
			}
		}
		prev = card
	}

	return s, nil
}

func (t Tableau) getBottomCard(col int) gocards.Card {
	return t[col].Cards[len(t[col].Cards)-1]
}
