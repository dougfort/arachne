package main

import (
	//	"context"

	"github.com/dougfort/arachne/game"

	pb "github.com/dougfort/arachne/arachne"
)

func arachne2pb(tableau game.Tableau) []*pb.Stack {
	stack := make([]*pb.Stack, game.TableauWidth)
	for col := 0; col < game.TableauWidth; col++ {
		stack[col] = new(pb.Stack)
		stack[col].HiddenCount = int32(tableau[col].HiddenCount)
		cardsLen := len(tableau[col].Cards)
		stack[col].Cards = make([]*pb.Card, cardsLen)
		for row := 0; row < cardsLen; row++ {
			localCard := tableau[col].Cards[row]
			stack[col].Cards[row] =
				&pb.Card{
					Suit: int32(localCard.Suit),
					Rank: int32(localCard.Rank),
				}
		}
	}

	return stack
}
