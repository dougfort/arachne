package game

import (
	"fmt"
	"testing"

	"github.com/dougfort/gocards"
)

func TestNewGame(t *testing.T) {
	var expectedHidden = [TableauWidth]int{5, 4, 4, 5, 4, 4, 5, 4, 4, 5}

	cardCount := make(map[gocards.Card]int)
	g := New()

	var totalHidden int
	for col := 0; col < TableauWidth; col++ {
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
	expectedRemaining := deckCount*(4*13) - (totalHidden + TableauWidth)

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

func TestGameMove(t *testing.T) {
	testCases := []struct {
		desc        string
		before      Game
		moves       []MoveType
		after       Game
		expectError bool
	}{
		{
			desc: "single from, empty to",
			before: Game{
				Deck: nil,
				Tableau: Tableau{
					StackType{Cards: gocards.Cards{aceOfClubs}},
				},
			},
			moves: []MoveType{
				MoveType{FromCol: 0, FromRow: 0, ToCol: 1},
			},
			after: Game{
				Deck: nil,
				Tableau: Tableau{
					StackType{},
					StackType{Cards: gocards.Cards{aceOfClubs}},
				},
			},
		},
		{
			desc: "simple capture",
			before: Game{
				Deck: nil,
				Tableau: Tableau{
					StackType{Cards: gocards.Cards{
						queenOfClubs,
						jackOfClubs,
						tenOfClubs,
						nineOfClubs,
						eightOfClubs,
						sevenOfClubs,
						sixOfClubs,
						fiveOfClubs,
						fourOfClubs,
						threeOfClubs,
						twoOfClubs,
						aceOfClubs,
					}},
					StackType{Cards: gocards.Cards{kingOfClubs}},
				},
			},
			moves: []MoveType{
				MoveType{FromCol: 0, FromRow: 0, ToCol: 1},
			},
			after: Game{
				Deck:    nil,
				Tableau: Tableau{},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s expect error: %t", tc.desc, tc.expectError), func(t *testing.T) {
			var err error
		MOVE_LOOP:
			for _, move := range tc.moves {
				if _, err = tc.before.Move(move); err != nil {
					if tc.expectError {
						break MOVE_LOOP
					}
					t.Fatalf("tc.before.Move(%s) failed: %s", move, err)
				}
			}
			if err != nil {
				if !tableausEqual(tc.before.Tableau, tc.after.Tableau) {
					t.Fatalf("tableaux mismatch after move(s)")
				}
			}
		})
	}
}

func tableausEqual(t1, t2 Tableau) bool {
	for i := 0; i < TableauWidth; i++ {
		if len(t1[i].Cards) != len(t2[i].Cards) {
			return false
		}
		if t1[i].HiddenCount != t2[i].HiddenCount {
			return false
		}
		for j := 0; j < len(t1[i].Cards); j++ {
			if !t1[i].Cards[j].Equal(t2[i].Cards[j]) {
				return false
			}
		}
	}

	return true
}
