package game

import (
	"fmt"
	"testing"

	"github.com/dougfort/gocards"
	"github.com/dougfort/gocards/standard"
)

type moveSet map[MoveType]struct{}

var (
	aceOfClubs    = gocards.Card{Suit: standard.Clubs, Rank: standard.Ace}
	twoOfClubs    = gocards.Card{Suit: standard.Clubs, Rank: standard.Two}
	threeOfClubs  = gocards.Card{Suit: standard.Clubs, Rank: standard.Three}
	fourOfClubs   = gocards.Card{Suit: standard.Clubs, Rank: standard.Four}
	fiveOfClubs   = gocards.Card{Suit: standard.Clubs, Rank: standard.Five}
	sixOfClubs    = gocards.Card{Suit: standard.Clubs, Rank: standard.Six}
	sevenOfClubs  = gocards.Card{Suit: standard.Clubs, Rank: standard.Seven}
	eightOfClubs  = gocards.Card{Suit: standard.Clubs, Rank: standard.Eight}
	nineOfClubs   = gocards.Card{Suit: standard.Clubs, Rank: standard.Nine}
	tenOfClubs    = gocards.Card{Suit: standard.Clubs, Rank: standard.Ten}
	jackOfClubs   = gocards.Card{Suit: standard.Clubs, Rank: standard.Jack}
	queenOfClubs  = gocards.Card{Suit: standard.Clubs, Rank: standard.Queen}
	kingOfClubs   = gocards.Card{Suit: standard.Clubs, Rank: standard.King}
	twoOfHearts   = gocards.Card{Suit: standard.Hearts, Rank: standard.Two}
	threeOfSpades = gocards.Card{Suit: standard.Spades, Rank: standard.Three}

	testTableaus = []Tableau{
		// 0
		Tableau{},

		// 1
		Tableau{
			StackType{Cards: gocards.Cards{aceOfClubs}},
		},

		// 2
		Tableau{
			StackType{Cards: gocards.Cards{aceOfClubs}},
			StackType{Cards: gocards.Cards{twoOfHearts}},
		},

		// 3
		Tableau{
			StackType{Cards: gocards.Cards{twoOfHearts, aceOfClubs}},
			StackType{},
		},

		// 4
		Tableau{
			StackType{Cards: gocards.Cards{aceOfClubs, twoOfClubs}},
			StackType{},
		},

		// 5
		Tableau{
			StackType{Cards: gocards.Cards{twoOfClubs, aceOfClubs}},
			StackType{Cards: gocards.Cards{twoOfClubs}},
		},

		// 6
		Tableau{
			StackType{Cards: gocards.Cards{twoOfClubs, aceOfClubs}},
			StackType{Cards: gocards.Cards{threeOfClubs}},
		},

		// 7
	}
)

func TestValidateMove(t *testing.T) {
	testCases := []struct {
		desc        string
		tab         Tableau
		move        MoveType
		expectValid bool
	}{
		{"empty tableau", testTableaus[0], MoveType{0, 0, 1}, false},
		{"invalid to col", testTableaus[1], MoveType{0, 0, -1}, false},
		{"invalid from col", testTableaus[1], MoveType{-1, 0, 1}, false},
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

func TestEnumerateMoves(t *testing.T) {
	testCases := []struct {
		desc          string
		tab           Tableau
		expectedMoves []MoveType
	}{
		{"empty tableau", testTableaus[0], nil},
		{"move to empty stack", testTableaus[1], []MoveType{
			MoveType{0, 0, 1},
			MoveType{0, 0, 2},
			MoveType{0, 0, 3},
			MoveType{0, 0, 4},
			MoveType{0, 0, 5},
			MoveType{0, 0, 6},
			MoveType{0, 0, 7},
			MoveType{0, 0, 8},
			MoveType{0, 0, 9},
		}},
		{"move to right Rank", testTableaus[2], []MoveType{
			MoveType{0, 0, 1},
			MoveType{0, 0, 2},
			MoveType{0, 0, 3},
			MoveType{0, 0, 4},
			MoveType{0, 0, 5},
			MoveType{0, 0, 6},
			MoveType{0, 0, 7},
			MoveType{0, 0, 8},
			MoveType{0, 0, 9},
			MoveType{1, 0, 2},
			MoveType{1, 0, 3},
			MoveType{1, 0, 4},
			MoveType{1, 0, 5},
			MoveType{1, 0, 6},
			MoveType{1, 0, 7},
			MoveType{1, 0, 8},
			MoveType{1, 0, 9},
		}},
		{"move to valid card", testTableaus[6], []MoveType{
			MoveType{0, 0, 1},
			MoveType{0, 0, 2},
			MoveType{0, 0, 3},
			MoveType{0, 0, 4},
			MoveType{0, 0, 5},
			MoveType{0, 0, 6},
			MoveType{0, 0, 7},
			MoveType{0, 0, 8},
			MoveType{0, 0, 9},
			MoveType{0, 1, 2},
			MoveType{0, 1, 3},
			MoveType{0, 1, 4},
			MoveType{0, 1, 5},
			MoveType{0, 1, 6},
			MoveType{0, 1, 7},
			MoveType{0, 1, 8},
			MoveType{0, 1, 9},
			MoveType{1, 0, 2},
			MoveType{1, 0, 3},
			MoveType{1, 0, 4},
			MoveType{1, 0, 5},
			MoveType{1, 0, 6},
			MoveType{1, 0, 7},
			MoveType{1, 0, 8},
			MoveType{1, 0, 9},
		}},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%s", tc.desc), func(t *testing.T) {
			moves := tc.tab.EnumerateMoves()
			s1, s2 := compareMoves(moves, tc.expectedMoves)
			if len(s1) != 0 || len(s2) != 0 {
				t.Fatalf("expected moves: %s, found: %s",
					moveSetStrings(s1), moveSetStrings(s2))
			}
		})
	}
}

// compareMoves returns (moves in m1 not in m2, moves in m2 not in m1)
func compareMoves(m1, m2 []MoveType) (moveSet, moveSet) {
	s1 := make(map[MoveType]struct{})
	s2 := make(map[MoveType]struct{})

	for _, m := range m1 {
		s1[m] = struct{}{}
	}

	for _, m := range m2 {
		s2[m] = struct{}{}
	}

	for _, m := range m1 {
		delete(s2, m)
	}

	for _, m := range m2 {
		delete(s1, m)
	}

	return s1, s2
}

func moveString(m MoveType) string {
	return fmt.Sprintf("[(%d,%d)->%d]", m.FromCol, m.FromRow, m.ToCol)
}

func moveSetStrings(s moveSet) []string {
	result := make([]string, len(s))
	var i int
	for m := range s {
		result[i] = moveString(m)
		i++
	}
	return result
}
