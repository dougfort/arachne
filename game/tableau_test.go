package game

import (
	"fmt"
	"testing"

	"github.com/dougfort/gocards"
	"github.com/dougfort/gocards/standard"
)

var (
	aceOfClubs    = gocards.Card{Suit: standard.Clubs, Rank: standard.Ace}
	twoOfClubs    = gocards.Card{Suit: standard.Clubs, Rank: standard.Two}
	threeOfClubs  = gocards.Card{Suit: standard.Clubs, Rank: standard.Three}
	twoOfHearts   = gocards.Card{Suit: standard.Hearts, Rank: standard.Two}
	threeOfSpades = gocards.Card{Suit: standard.Spades, Rank: standard.Three}

	testTableaus = []Tableau{
		Tableau{},
		Tableau{
			StackType{Cards: gocards.Cards{aceOfClubs}},
			StackType{},
		},
		Tableau{
			StackType{Cards: gocards.Cards{aceOfClubs}},
			StackType{Cards: gocards.Cards{twoOfHearts}},
		},
		Tableau{
			StackType{Cards: gocards.Cards{twoOfHearts, aceOfClubs}},
			StackType{},
		},
		Tableau{
			StackType{Cards: gocards.Cards{aceOfClubs, twoOfClubs}},
			StackType{},
		},
		Tableau{
			StackType{Cards: gocards.Cards{twoOfClubs, aceOfClubs}},
			StackType{Cards: gocards.Cards{twoOfClubs}},
		},
		Tableau{
			StackType{Cards: gocards.Cards{twoOfClubs, aceOfClubs}},
			StackType{Cards: gocards.Cards{threeOfClubs}},
		},
	}
)

func TestTableauMove(t *testing.T) {
	testCases := []struct {
		desc        string
		tab         Tableau
		move        MoveType
		expectValid bool
	}{
		{"empty tableau", testTableaus[0], MoveType{0, 0, 1}, false},
		{"move to same stack", testTableaus[1], MoveType{1, 0, 1}, false},
		{"invalid from col", testTableaus[1], MoveType{-1, 0, 1}, false},
		{"invalid to col", testTableaus[1], MoveType{0, 0, -1}, false},
		{"empty slice", testTableaus[1], MoveType{1, 0, 0}, false},
		{"move to empty stack", testTableaus[1], MoveType{0, 0, 1}, true},
		{"move to wrong Rank", testTableaus[2], MoveType{1, 0, 0}, false},
		{"move to right Rank", testTableaus[2], MoveType{0, 0, 1}, true},
		{"suit mismatch", testTableaus[3], MoveType{0, 0, 1}, false},
		{"rank mismatch: from", testTableaus[4], MoveType{0, 0, 1}, false},
		{"rank mismatch: to", testTableaus[5], MoveType{0, 0, 1}, false},
		{"move to valid card", testTableaus[6], MoveType{0, 0, 1}, true},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s %t", tc.desc, tc.expectValid), func(t *testing.T) {
			err := tc.tab.ValidateMove(tc.move)
			if tc.expectValid && err != nil {
				t.Fatalf("expected valid move: %s", err)
			}
			if !tc.expectValid && err == nil {
				t.Fatalf("expected invalid move")
			}
		})
	}
}
