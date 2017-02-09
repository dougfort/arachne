package game

import (
	"testing"

	"github.com/dougfort/gocards"

	"github.com/dougfort/arachne/types"
)

func TestGame(t *testing.T) {
	var expectedHidden = [types.TableauWidth]int{5, 4, 4, 5, 4, 4, 5, 4, 4, 5}

	cardCount := make(map[gocards.Card]int)
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
		for _, card := range g.HiddenCards[col] {
			cardCount[card]++
		}
	}

	// we expect each hidden card to occur 1-2 times
	for card, count := range cardCount {
		if count > 2 {
			t.Fatalf("hidden card error: %s: count=%d", card.Value(), count)
		}
	}

	// TODO: fix magic numbers 4 * 13
	expectedRemaining := deckCount*(4*13) - (totalHidden + types.TableauWidth)

	if g.Deck.RemainingCards() != expectedRemaining {
		t.Fatalf("invalid remaining: expected %d, found %d",
			expectedRemaining, g.Deck.RemainingCards())
	}

	cardCount = make(map[gocards.Card]int)
	for _, tab := range g.Tableau {
		if len(tab.Cards) != 1 {
			t.Fatalf("invalid len: %d", len(tab.Cards))
		}
		cardCount[tab.Cards[0]]++
	}

	for card, count := range cardCount {
		if count > 2 {
			t.Fatalf("too many %s cards: %d", card.Value(), count)
		}
	}
}
