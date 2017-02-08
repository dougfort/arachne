package game

import (
	"github.com/dougfort/gocards"
	"github.com/dougfort/gocards/standard"

	"github.com/dougfort/arachne/types"
)

// HiddenCards represents the Cards that are not visible in the Tableau
type HiddenCards [types.TableauWidth]gocards.Cards

// Game represents a complete, playable, game
type Game struct {
	Deck        gocards.PlayableDeck
	Tableau     types.Tableau
	HiddenCards HiddenCards
}

const deckCount = 2

// New starts a new game
func New() Game {
	var game Game

	d := standard.NewDecks(deckCount)
	game.Deck = d.Shuffle()
	game.Tableau, game.HiddenCards = initialDeal(game.Deck)

	return game
}

// initialDeal deals the cards the way a human would
// returning the Tableau and HiddenCards
// . . . . . . . . . .
// . . . . . . . . . .
// . . . . . . . . . .
// . . . . . . . . . .
// . X X . X X . X X .
// X     X     X     X
//
// where '.' is a hidden card and 'X' is visible
func initialDeal(deck gocards.PlayableDeck) (types.Tableau, HiddenCards) {
	var tab = types.Tableau{
		types.StackType{HiddenCount: 5, Cards: make(gocards.Cards, 1)},
		types.StackType{HiddenCount: 4, Cards: make(gocards.Cards, 1)},
		types.StackType{HiddenCount: 4, Cards: make(gocards.Cards, 1)},
		types.StackType{HiddenCount: 5, Cards: make(gocards.Cards, 1)},
		types.StackType{HiddenCount: 4, Cards: make(gocards.Cards, 1)},
		types.StackType{HiddenCount: 4, Cards: make(gocards.Cards, 1)},
		types.StackType{HiddenCount: 5, Cards: make(gocards.Cards, 1)},
		types.StackType{HiddenCount: 4, Cards: make(gocards.Cards, 1)},
		types.StackType{HiddenCount: 4, Cards: make(gocards.Cards, 1)},
		types.StackType{HiddenCount: 5, Cards: make(gocards.Cards, 1)},
	}
	var hid = HiddenCards{
		make(gocards.Cards, tab[0].HiddenCount),
		make(gocards.Cards, tab[1].HiddenCount),
		make(gocards.Cards, tab[2].HiddenCount),
		make(gocards.Cards, tab[3].HiddenCount),
		make(gocards.Cards, tab[4].HiddenCount),
		make(gocards.Cards, tab[5].HiddenCount),
		make(gocards.Cards, tab[6].HiddenCount),
		make(gocards.Cards, tab[7].HiddenCount),
		make(gocards.Cards, tab[8].HiddenCount),
		make(gocards.Cards, tab[9].HiddenCount),
	}

	var row int
	var ok bool
	for col := 0; col < types.TableauWidth; col++ {
		if row < tab[col].HiddenCount {
			hid[col][row], ok = deck.Next()
			if !ok {
				panic("initialDeal")
			}
		} else if row == tab[col].HiddenCount {
			tab[row].Cards[0], ok = deck.Next()
			if !ok {
				panic("initialDeal")
			}
		}
		row++
	}

	return tab, hid
}
