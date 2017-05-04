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

// Replay a game for a specific seed
func Replay(seed int64) *Game {
	var game Game

	d := standard.NewDecks(deckCount)
	game.Deck = d.SeededShuffle(seed)
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
// returns 'true' if a run of cards is captured
func (g *Game) Move(m MoveType) (bool, error) {
	var s gocards.Cards
	var suitMatchCount int
	var capture bool
	var ok bool
	var err error

	if s, err = g.Tableau.getSliceToMove(m); err != nil {
		return false, errors.Wrap(err, "getSliceToMove failed")
	}

	if !toColValid(m) {
		return false, errors.Errorf("invalid ToCol: %d", m.FromCol)
	}

	if suitMatchCount, ok = g.Tableau.evaluateMoveDest(m, s); !ok {
		return false, errors.Errorf("move slice does not fit ToCol")
	}

	// if the move results in a run of all cards from Ace...King
	// we can capture those cards and remove them from the tableau.
	if len(s)+suitMatchCount == standard.CardsOfSuit {
		// remove the cards from the 'To' stack
		// if the From stack now has no visible cards
		// bring in the outtermost hidden card
		captureRow := len(g.Tableau[m.ToCol].Cards) - suitMatchCount
		if err = g.removeCardsAtRow(m.ToCol, captureRow); err != nil {
			return false,
				errors.Wrapf(err, "removeCardsAtRow(%d, %d) %s",
					m.ToCol, captureRow, m)
		}
		capture = true
	} else {
		g.Tableau[m.ToCol].Cards =
			append(g.Tableau[m.ToCol].Cards, s...)
	}

	// remove the cards from the 'From' stack
	// if the From stack now has no visible cards
	// bring in the outtermost hidden card
	row := g.Tableau.computeCardsRow(m)
	if err = g.removeCardsAtRow(m.FromCol, row); err != nil {
		return false,
			errors.Wrapf(err, "removeCardsAtRow(%d, %d) %s",
				m.FromCol, row, m)
	}

	return capture, nil
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

func (g *Game) removeCardsAtRow(col, row int) error {
	if !(row >= 0 && row < len(g.Tableau[col].Cards)) {
		return errors.Errorf("row %d invalid: col=%d", row, col)
	}
	g.Tableau[col].Cards = g.Tableau[col].Cards[:row]
	if len(g.Tableau[col].Cards) == 0 {
		if g.Tableau[col].HiddenCount > 0 {
			if g.Tableau[col].HiddenCount != len(g.HiddenCards[col]) {
				return errors.Errorf("hidden card mismatch %d != %d",
					g.Tableau[col].HiddenCount,
					len(g.HiddenCards[col]),
				)
			}
			r := g.Tableau[col].HiddenCount - 1
			g.Tableau[col].Cards = g.HiddenCards[col][r:]
			g.HiddenCards[col] = g.HiddenCards[col][:r]
			g.Tableau[col].HiddenCount--
		}
	}

	return nil
}
