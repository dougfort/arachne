package main

import (
	"fmt"

	pb "github.com/dougfort/arachne/arachne"
)

func displayTableauStrings(g *pb.Game) {
	const stringBack = "(.....)"
	cardStrings := map[pb.Card]string{
		pb.Card{Suit: 1, Rank: 1}:  "(C,  A)",
		pb.Card{Suit: 1, Rank: 2}:  "(C,  2)",
		pb.Card{Suit: 1, Rank: 3}:  "(C,  3)",
		pb.Card{Suit: 1, Rank: 4}:  "(C,  4)",
		pb.Card{Suit: 1, Rank: 5}:  "(C,  5)",
		pb.Card{Suit: 1, Rank: 6}:  "(C,  6)",
		pb.Card{Suit: 1, Rank: 7}:  "(C,  7)",
		pb.Card{Suit: 1, Rank: 8}:  "(C,  8)",
		pb.Card{Suit: 1, Rank: 9}:  "(C,  9)",
		pb.Card{Suit: 1, Rank: 10}: "(C, 10)",
		pb.Card{Suit: 1, Rank: 11}: "(C,  J)",
		pb.Card{Suit: 1, Rank: 12}: "(C,  Q)",
		pb.Card{Suit: 1, Rank: 13}: "(C,  K)",
		pb.Card{Suit: 2, Rank: 1}:  "(D,  A)",
		pb.Card{Suit: 2, Rank: 2}:  "(D,  2)",
		pb.Card{Suit: 2, Rank: 3}:  "(D,  3)",
		pb.Card{Suit: 2, Rank: 4}:  "(D,  4)",
		pb.Card{Suit: 2, Rank: 5}:  "(D,  5)",
		pb.Card{Suit: 2, Rank: 6}:  "(D,  6)",
		pb.Card{Suit: 2, Rank: 7}:  "(D,  7)",
		pb.Card{Suit: 2, Rank: 8}:  "(D,  8)",
		pb.Card{Suit: 2, Rank: 9}:  "(D,  9)",
		pb.Card{Suit: 2, Rank: 10}: "(D, 10)",
		pb.Card{Suit: 2, Rank: 11}: "(D,  J)",
		pb.Card{Suit: 2, Rank: 12}: "(D,  Q)",
		pb.Card{Suit: 2, Rank: 13}: "(D,  K)",
		pb.Card{Suit: 3, Rank: 1}:  "(H,  A)",
		pb.Card{Suit: 3, Rank: 2}:  "(H,  2)",
		pb.Card{Suit: 3, Rank: 3}:  "(H,  3)",
		pb.Card{Suit: 3, Rank: 4}:  "(H,  4)",
		pb.Card{Suit: 3, Rank: 5}:  "(H,  5)",
		pb.Card{Suit: 3, Rank: 6}:  "(H,  6)",
		pb.Card{Suit: 3, Rank: 7}:  "(H,  7)",
		pb.Card{Suit: 3, Rank: 8}:  "(H,  8)",
		pb.Card{Suit: 3, Rank: 9}:  "(H,  9)",
		pb.Card{Suit: 3, Rank: 10}: "(H, 10)",
		pb.Card{Suit: 3, Rank: 11}: "(H,  J)",
		pb.Card{Suit: 3, Rank: 12}: "(H,  Q)",
		pb.Card{Suit: 3, Rank: 13}: "(H,  K)",
		pb.Card{Suit: 4, Rank: 1}:  "(S,  A)",
		pb.Card{Suit: 4, Rank: 2}:  "(S,  2)",
		pb.Card{Suit: 4, Rank: 3}:  "(S,  3)",
		pb.Card{Suit: 4, Rank: 4}:  "(S,  4)",
		pb.Card{Suit: 4, Rank: 5}:  "(S,  5)",
		pb.Card{Suit: 4, Rank: 6}:  "(S,  6)",
		pb.Card{Suit: 4, Rank: 7}:  "(S,  7)",
		pb.Card{Suit: 4, Rank: 8}:  "(S,  8)",
		pb.Card{Suit: 4, Rank: 9}:  "(S,  9)",
		pb.Card{Suit: 4, Rank: 10}: "(S, 10)",
		pb.Card{Suit: 4, Rank: 11}: "(S,  J)",
		pb.Card{Suit: 4, Rank: 12}: "(S,  Q)",
		pb.Card{Suit: 4, Rank: 13}: "(S,  K)",
	}

	var row int32
	var col int32

	tableauWidth := len(g.GetStack())
	image := make([]string, tableauWidth)

	fmt.Printf("   %-7s %-7s %-7s %-7s %-7s %-7s %-7s %-7s %-7s %-7s\n",
		"   1", "   2", "   3", "   4", "   5",
		"   6", "   7", "   8", "   9", "  10",
	)
ROW_LOOP:
	for {
		var found bool
		for col = 0; col < int32(tableauWidth); col++ {
			if row < g.Stack[col].HiddenCount {
				image[col] = stringBack
				found = true
			} else {
				visibleRow := row - g.Stack[col].HiddenCount
				if visibleRow < int32(len(g.Stack[col].Cards)) {
					image[col] = cardStrings[*g.Stack[col].Cards[visibleRow]]
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

		row++
	}
}
