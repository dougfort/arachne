package main

import (
	"testing"

	"github.com/dougfort/gocards"
	"github.com/dougfort/gocards/standard"
)

func TestDisplayCards(t *testing.T) {
	// test that every card has a key in the map, and that every value is unique
	valueSet := make(map[rune]struct{})

	for suit := standard.Clubs; suit <= standard.Spades; suit++ {
		for rank := standard.Ace; rank <= standard.King; rank++ {
			card := gocards.Card{Suit: suit, Rank: rank}
			v, ok := displayCards[card]
			if !ok {
				t.Fatalf("unmatched card: %s", card.Value())
			}
			_, ok = valueSet[v]
			if ok {
				t.Fatalf("duplicate value %s", card.Value())
			}
			valueSet[v] = struct{}{}
		}
	}
}
