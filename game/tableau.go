package game

import (
	"github.com/pkg/errors"

	"github.com/dougfort/gocards"
)

// EnumerateMoves lists all possible legal moves in the current
// Tableau.
// The returned slice is ordered by the perceived desirability of
// the move.
func (t Tableau) EnumerateMoves() []EvaluatedMoveType {
	var err error
	var moves []EvaluatedMoveType

FROM_LOOP:
	for from := 0; from < TableauWidth; from++ {
		if len(t[from].Cards) == 0 {
			continue FROM_LOOP
		}

	TO_LOOP:
		for to := 0; to < TableauWidth; to++ {

			if to == from {
				continue TO_LOOP
			}

			// we could weed out a lot of these before validating,
			// but why bother, it would just make the code more confusing
		ROW_LOOP:
			for i := 0; i < len(t[from].Cards); i++ {
				var evaluatedMove EvaluatedMoveType

				row := t[from].HiddenCount + i
				move := MoveType{FromCol: from, FromRow: row, ToCol: to}
				if evaluatedMove, err = t.EvaluateMove(move); err != nil {
					continue ROW_LOOP
				}
				moves = append(moves, evaluatedMove)
			}
		}
	}

	return moves
}

// EvaluateMove returns the evaluated move if the move is valid
// in the Tableau
//
// A move is valid if
//   1. The slice of Cards at 'From'  is all of the same Suit
//   2. The slice of Cards at 'From' is in order by Rank from top to bottom
//   3. The bottom Card at 'To' is successor by Rank of the top Card moved
//
func (t Tableau) EvaluateMove(m MoveType) (EvaluatedMoveType, error) {
	var err error
	var fromCount int
	var toCount int
	var ok bool
	var s gocards.Cards

	if s, err = t.getSliceToMove(m); err != nil {
		return EvaluatedMoveType{}, err
	}
	fromCount = len(s)

	if !toColValid(m) {
		return EvaluatedMoveType{},
			errors.Errorf("invalid ToCol: %d", m.FromCol)
	}

	if toCount, ok = t.evaluateMoveDest(m, s); !ok {
		return EvaluatedMoveType{},
			errors.Errorf("move slice does not fit dest")
	}

	return EvaluatedMoveType{
		MoveType:  m,
		FromCount: fromCount,
		ToCount:   toCount,
	}, nil
}

// EmptyStack returns true if the Tableau contains an empty Stack
func (t Tableau) EmptyStack() bool {
	for col := 0; col < TableauWidth; col++ {
		if len(t[col].Cards) == 0 {
			return true
		}
	}
	return false
}

func toColValid(m MoveType) bool {
	return stackIndexValid(m.ToCol) && m.ToCol != m.FromCol
}

func stackIndexValid(index int) bool {
	return index >= 0 && index < TableauWidth
}

func (t Tableau) getSliceToMove(m MoveType) (gocards.Cards, error) {
	if !stackIndexValid(m.FromCol) {
		return nil, errors.Errorf("m.FromCol invalid: %d", m.FromCol)
	}

	row := t.computeCardsRow(m)
	if !(row >= 0 && row < len(t[m.FromCol].Cards)) {
		return nil, errors.Errorf("computeCardsRow %d invalid: %s",
			row, m)
	}

	s := t[m.FromCol].Cards[row:]

	var prev gocards.Card
	// interate from the top (highest Rank) card downto the bottom (lowest Rank)
	for i, card := range s {
		if i > 0 {
			if card.Suit != prev.Suit {
				return nil, errors.Errorf("move slice not all the same Suit at %d", i)
			}
			if card.Rank != prev.Rank-1 {
				return nil, errors.Errorf("move slice out of order %d %d at %d",
					prev.Rank, card.Rank, i)
			}
		}
		prev = card
	}

	return s, nil
}

func (t Tableau) computeCardsRow(m MoveType) int {
	return m.FromRow - t[m.FromCol].HiddenCount
}

func (t Tableau) evaluateMoveDest(
	m MoveType,
	s gocards.Cards,
) (int, bool) {

	// if there are no cards in the dest stack, we can move
	// anything there.
	if len(t[m.ToCol].Cards) == 0 {
		return 0, true
	}

	var rankMatch bool
	var suitMatchCount int
	prevRank := s[0].Rank

	// index starts at the bottom card of the dest stack
DEST_LOOP:
	for i := len(t[m.ToCol].Cards) - 1; i >= 0; i-- {
		// we check Rank first, because if we have at least
		// one match on Rank, the move is valid, even if
		// Suit doesn't match.
		if t[m.ToCol].Cards[i].Rank != prevRank+1 {
			break DEST_LOOP
		}
		rankMatch = true
		prevRank = t[m.ToCol].Cards[i].Rank

		if t[m.ToCol].Cards[i].Suit != s[0].Suit {
			break DEST_LOOP
		}
		suitMatchCount++
	}

	return suitMatchCount, rankMatch
}

func (t Tableau) getBottomCard(col int) gocards.Card {
	return t[col].Cards[len(t[col].Cards)-1]
}
