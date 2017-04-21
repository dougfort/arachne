package game

import (
	"github.com/pkg/errors"

	"github.com/dougfort/gocards"
	"github.com/dougfort/gocards/standard"
)

// HiddenCards represents the Cards that are not visible in the Tableau
type HiddenCards [TableauWidth]gocards.Cards

// Game represents a complete, playable, game
type Game struct {
	Deck        gocards.PlayableDeck
	Tableau     Tableau
	HiddenCards HiddenCards
}

const deckCount = 2

// New starts a new game
func New() *Game {
	var game Game

	d := standard.NewDecks(deckCount)
	game.Deck = d.Shuffle()
	game.Tableau, game.HiddenCards = initialDeal(game.Deck)

	return &game
}

// initialDeal deals the cards the way a human would
// returning the Tableau and HiddenCards
// . . . . . . . . . .
// . . . . . . . . . .
// X     X     X     X
//
// where '.' is a hidden card and 'X' is visible
func initialDeal(deck gocards.PlayableDeck) (Tableau, HiddenCards) {
	var tab = Tableau{
		StackType{HiddenCount: 5, Cards: make(gocards.Cards, 1)},
		StackType{HiddenCount: 4, Cards: make(gocards.Cards, 1)},
		StackType{HiddenCount: 4, Cards: make(gocards.Cards, 1)},
		StackType{HiddenCount: 5, Cards: make(gocards.Cards, 1)},
		StackType{HiddenCount: 4, Cards: make(gocards.Cards, 1)},
		StackType{HiddenCount: 4, Cards: make(gocards.Cards, 1)},
		StackType{HiddenCount: 5, Cards: make(gocards.Cards, 1)},
		StackType{HiddenCount: 4, Cards: make(gocards.Cards, 1)},
		StackType{HiddenCount: 4, Cards: make(gocards.Cards, 1)},
		StackType{HiddenCount: 5, Cards: make(gocards.Cards, 1)},
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
ROW_LOOP:
	for {
		var found bool
		for col := 0; col < TableauWidth; col++ {
			var ok bool
			if row < tab[col].HiddenCount {
				hid[col][row], ok = deck.Next()
				if !ok {
					panic("initialDeal")
				}
				found = true
			} else if row == tab[col].HiddenCount {
				tab[col].Cards[0], ok = deck.Next()
				if !ok {
					panic("initialDeal")
				}
				found = true
			}
		}
		if !found {
			break ROW_LOOP
		}
		row++
	}

	return tab, hid
}

// Move takes a slice of cards from one Stack and appends it to another
// If neccessary a hidden card will be exposed in the Stack from which the
// move originates
func (g *Game) Move(m MoveType) error {
	var s gocards.Cards
	var err error

	if s, err = g.Tableau.getSliceToMove(m); err != nil {
		return errors.Wrap(err, "getSliceToMove failed")
	}

	if !toColValid(m) {
		return errors.Errorf("invalid ToCol: %d", m.FromCol)
	}

	if _, ok := g.Tableau.evaluateMoveDest(m, s); !ok {
		return errors.Errorf("move slice does not fit ToCol")
	}

	g.Tableau[m.ToCol].Cards = append(g.Tableau[m.ToCol].Cards, s...)

	// remove the cards from the 'From' stack
	// if the From stack now has no visible cards
	// bring in the outtermost hidden card
	row := g.Tableau.computeCardsRow(m)
	if !(row >= 0 && row < len(g.Tableau[m.FromCol].Cards)) {
		return errors.Errorf("computeCardsRow %d invalid: %s", row, m)
	}
	g.Tableau[m.FromCol].Cards = g.Tableau[m.FromCol].Cards[:row]
	if len(g.Tableau[m.FromCol].Cards) == 0 {
		if g.Tableau[m.FromCol].HiddenCount > 0 {
			if g.Tableau[m.FromCol].HiddenCount != len(g.HiddenCards[m.FromCol]) {
				return errors.Errorf("hidden card mismatch %d != %d",
					g.Tableau[m.FromCol].HiddenCount,
					len(g.HiddenCards[m.FromCol]),
				)
			}
			r := g.Tableau[m.FromCol].HiddenCount - 1
			g.Tableau[m.FromCol].Cards = g.HiddenCards[m.FromCol][r:]
			g.HiddenCards[m.FromCol] = g.HiddenCards[m.FromCol][:r]
			g.Tableau[m.FromCol].HiddenCount--
		}
	}

	return nil
}

// Deal appends a card to each stack
func (g *Game) Deal() error {
	for col := 0; col < TableauWidth; col++ {
		card, ok := g.Deck.Next()
		if !ok {
			return errors.Errorf("deck exhausted")
		}
		g.Tableau[col].Cards = append(g.Tableau[col].Cards, card)
	}

	return nil
}
