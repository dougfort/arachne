package game

import (
	"fmt"
	"testing"

	"github.com/dougfort/gocards"
	"github.com/dougfort/gocards/standard"
)

var (
	aceOfClubs   = gocards.Card{Suit: standard.Clubs, Rank: standard.Ace}
	testTableaus = []Tableau{
		Tableau{},
		Tableau{
			StackType{Cards: gocards.Cards{aceOfClubs}},
			StackType{},
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
		{"move to same stack", testTableaus[1], MoveType{1, 0, 1}, true},
		{"move to empty stack", testTableaus[1], MoveType{0, 0, 1}, true},
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
