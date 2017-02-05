package game

import (
	"testing"

	"github.com/dougfort/arachne/types"
)

func TestGame(t *testing.T) {
	var expectedHidden = [types.TableauWidth]int{5, 4, 4, 5, 4, 4, 5, 4, 4, 5}

	g := New()

	var totalHidden int
	for col := 0; col < types.TableauWidth; col++ {
		if len(g.HiddenCards[col]) != expectedHidden[col] {
			t.Fatalf("col: %d; hidden cards mismatch: %d != %d",
				col, len(g.HiddenCards[col]), expectedHidden[col])
		}
		if g.Tableau[col].HiddenCount != expectedHidden[col] {
			t.Fatalf("col: %d; hidden count mismatch: %d != %d",
				col, g.Tableau[col].HiddenCount, expectedHidden[col])
		}
		totalHidden += len(g.HiddenCards[col])
	}

	// TODO: fix magic numbers 4 * 13
	expectedRemaining := deckCount*(4*13) - (totalHidden + types.TableauWidth)

	if g.Deck.RemainingCards() != expectedRemaining {
		t.Fatalf("invalid remaining: expected %d, found %d",
			expectedRemaining, g.Deck.RemainingCards())
	}
}
