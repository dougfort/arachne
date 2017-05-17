package main

import (
	"fmt"

	"github.com/dougfort/arachne/internal/game"

	"github.com/dougfort/gocards"
	"github.com/dougfort/gocards/standard"
)

func displayTableauStrings(t game.Tableau) {
	const stringBack = "(.....)"
	cardStrings := map[gocards.Card]string{
		gocards.Card{Suit: standard.Clubs, Rank: standard.Ace}:      "(C,  A)",
		gocards.Card{Suit: standard.Clubs, Rank: standard.Two}:      "(C,  2)",
		gocards.Card{Suit: standard.Clubs, Rank: standard.Three}:    "(C,  3)",
		gocards.Card{Suit: standard.Clubs, Rank: standard.Four}:     "(C,  4)",
		gocards.Card{Suit: standard.Clubs, Rank: standard.Five}:     "(C,  5)",
		gocards.Card{Suit: standard.Clubs, Rank: standard.Six}:      "(C,  6)",
		gocards.Card{Suit: standard.Clubs, Rank: standard.Seven}:    "(C,  7)",
		gocards.Card{Suit: standard.Clubs, Rank: standard.Eight}:    "(C,  8)",
		gocards.Card{Suit: standard.Clubs, Rank: standard.Nine}:     "(C,  9)",
		gocards.Card{Suit: standard.Clubs, Rank: standard.Ten}:      "(C, 10)",
		gocards.Card{Suit: standard.Clubs, Rank: standard.Jack}:     "(C,  J)",
		gocards.Card{Suit: standard.Clubs, Rank: standard.Queen}:    "(C,  Q)",
		gocards.Card{Suit: standard.Clubs, Rank: standard.King}:     "(C,  K)",
		gocards.Card{Suit: standard.Diamonds, Rank: standard.Ace}:   "(D,  A)",
		gocards.Card{Suit: standard.Diamonds, Rank: standard.Two}:   "(D,  2)",
		gocards.Card{Suit: standard.Diamonds, Rank: standard.Three}: "(D,  3)",
		gocards.Card{Suit: standard.Diamonds, Rank: standard.Four}:  "(D,  4)",
		gocards.Card{Suit: standard.Diamonds, Rank: standard.Five}:  "(D,  5)",
		gocards.Card{Suit: standard.Diamonds, Rank: standard.Six}:   "(D,  6)",
		gocards.Card{Suit: standard.Diamonds, Rank: standard.Seven}: "(D,  7)",
		gocards.Card{Suit: standard.Diamonds, Rank: standard.Eight}: "(D,  8)",
		gocards.Card{Suit: standard.Diamonds, Rank: standard.Nine}:  "(D,  9)",
		gocards.Card{Suit: standard.Diamonds, Rank: standard.Ten}:   "(D, 10)",
		gocards.Card{Suit: standard.Diamonds, Rank: standard.Jack}:  "(D,  J)",
		gocards.Card{Suit: standard.Diamonds, Rank: standard.Queen}: "(D,  Q)",
		gocards.Card{Suit: standard.Diamonds, Rank: standard.King}:  "(D,  K)",
		gocards.Card{Suit: standard.Hearts, Rank: standard.Ace}:     "(H,  A)",
		gocards.Card{Suit: standard.Hearts, Rank: standard.Two}:     "(H,  2)",
		gocards.Card{Suit: standard.Hearts, Rank: standard.Three}:   "(H,  3)",
		gocards.Card{Suit: standard.Hearts, Rank: standard.Four}:    "(H,  4)",
		gocards.Card{Suit: standard.Hearts, Rank: standard.Five}:    "(H,  5)",
		gocards.Card{Suit: standard.Hearts, Rank: standard.Six}:     "(H,  6)",
		gocards.Card{Suit: standard.Hearts, Rank: standard.Seven}:   "(H,  7)",
		gocards.Card{Suit: standard.Hearts, Rank: standard.Eight}:   "(H,  8)",
		gocards.Card{Suit: standard.Hearts, Rank: standard.Nine}:    "(H,  9)",
		gocards.Card{Suit: standard.Hearts, Rank: standard.Ten}:     "(H, 10)",
		gocards.Card{Suit: standard.Hearts, Rank: standard.Jack}:    "(H,  J)",
		gocards.Card{Suit: standard.Hearts, Rank: standard.Queen}:   "(H,  Q)",
		gocards.Card{Suit: standard.Hearts, Rank: standard.King}:    "(H,  K)",
		gocards.Card{Suit: standard.Spades, Rank: standard.Ace}:     "(S,  A)",
		gocards.Card{Suit: standard.Spades, Rank: standard.Two}:     "(S,  2)",
		gocards.Card{Suit: standard.Spades, Rank: standard.Three}:   "(S,  3)",
		gocards.Card{Suit: standard.Spades, Rank: standard.Four}:    "(S,  4)",
		gocards.Card{Suit: standard.Spades, Rank: standard.Five}:    "(S,  5)",
		gocards.Card{Suit: standard.Spades, Rank: standard.Six}:     "(S,  6)",
		gocards.Card{Suit: standard.Spades, Rank: standard.Seven}:   "(S,  7)",
		gocards.Card{Suit: standard.Spades, Rank: standard.Eight}:   "(S,  8)",
		gocards.Card{Suit: standard.Spades, Rank: standard.Nine}:    "(S,  9)",
		gocards.Card{Suit: standard.Spades, Rank: standard.Ten}:     "(S, 10)",
		gocards.Card{Suit: standard.Spades, Rank: standard.Jack}:    "(S,  J)",
		gocards.Card{Suit: standard.Spades, Rank: standard.Queen}:   "(S,  Q)",
		gocards.Card{Suit: standard.Spades, Rank: standard.King}:    "(S,  K)",
	}

	image := make([]string, game.TableauWidth)

	fmt.Printf("   %-7s %-7s %-7s %-7s %-7s %-7s %-7s %-7s %-7s %-7s\n",
		"   1", "   2", "   3", "   4", "   5",
		"   6", "   7", "   8", "   9", "  10",
	)

ROW_LOOP:
	for row := 0; ; row++ {
		var found bool
		for col := 0; col < game.TableauWidth; col++ {
			if row < t[col].HiddenCount {
				image[col] = stringBack
				found = true
			} else {
				visibleRow := row - t[col].HiddenCount
				if visibleRow < len(t[col].Cards) {
					image[col] = cardStrings[t[col].Cards[visibleRow]]
					found = true
				} else {
					image[col] = ""
				}
			}
		}
		if !found {
			break ROW_LOOP
		}
		fmt.Printf("%2d %-7s %-7s %-7s %-7s %-7s %-7s %-7s %-7s %-7s %-7s\n",
			row+1,
			image[0], image[1], image[2], image[3], image[4],
			image[5], image[6], image[7], image[8], image[9],
		)

	}
}
