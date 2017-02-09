package main

import (
	"fmt"

	"github.com/dougfort/arachne/types"
	"github.com/dougfort/gocards"
	"github.com/dougfort/gocards/standard"
)

const displayCardBack = '🂠'

// displayCardsMap maps cards to unicode images
type displayCardsMap map[gocards.Card]rune

var (
	displayCards = displayCardsMap{
		gocards.Card{Suit: standard.Clubs, Rank: standard.Ace}:      '🃑',
		gocards.Card{Suit: standard.Clubs, Rank: standard.Two}:      '🃒',
		gocards.Card{Suit: standard.Clubs, Rank: standard.Three}:    '🃓',
		gocards.Card{Suit: standard.Clubs, Rank: standard.Four}:     '🃔',
		gocards.Card{Suit: standard.Clubs, Rank: standard.Five}:     '🃕',
		gocards.Card{Suit: standard.Clubs, Rank: standard.Six}:      '🃖',
		gocards.Card{Suit: standard.Clubs, Rank: standard.Seven}:    '🃗',
		gocards.Card{Suit: standard.Clubs, Rank: standard.Eight}:    '🃘',
		gocards.Card{Suit: standard.Clubs, Rank: standard.Nine}:     '🃙',
		gocards.Card{Suit: standard.Clubs, Rank: standard.Ten}:      '🃚',
		gocards.Card{Suit: standard.Clubs, Rank: standard.Jack}:     '🃛',
		gocards.Card{Suit: standard.Clubs, Rank: standard.Queen}:    '🃝',
		gocards.Card{Suit: standard.Clubs, Rank: standard.King}:     '🃞',
		gocards.Card{Suit: standard.Diamonds, Rank: standard.Ace}:   '🃁',
		gocards.Card{Suit: standard.Diamonds, Rank: standard.Two}:   '🃂',
		gocards.Card{Suit: standard.Diamonds, Rank: standard.Three}: '🃃',
		gocards.Card{Suit: standard.Diamonds, Rank: standard.Four}:  '🃄',
		gocards.Card{Suit: standard.Diamonds, Rank: standard.Five}:  '🃅',
		gocards.Card{Suit: standard.Diamonds, Rank: standard.Six}:   '🃆',
		gocards.Card{Suit: standard.Diamonds, Rank: standard.Seven}: '🃇',
		gocards.Card{Suit: standard.Diamonds, Rank: standard.Eight}: '🃈',
		gocards.Card{Suit: standard.Diamonds, Rank: standard.Nine}:  '🃉',
		gocards.Card{Suit: standard.Diamonds, Rank: standard.Ten}:   '🃊',
		gocards.Card{Suit: standard.Diamonds, Rank: standard.Jack}:  '🃋',
		gocards.Card{Suit: standard.Diamonds, Rank: standard.Queen}: '🃍',
		gocards.Card{Suit: standard.Diamonds, Rank: standard.King}:  '🃎',
		gocards.Card{Suit: standard.Hearts, Rank: standard.Ace}:     '🂱',
		gocards.Card{Suit: standard.Hearts, Rank: standard.Two}:     '🂲',
		gocards.Card{Suit: standard.Hearts, Rank: standard.Three}:   '🂳',
		gocards.Card{Suit: standard.Hearts, Rank: standard.Four}:    '🂴',
		gocards.Card{Suit: standard.Hearts, Rank: standard.Five}:    '🂵',
		gocards.Card{Suit: standard.Hearts, Rank: standard.Six}:     '🂶',
		gocards.Card{Suit: standard.Hearts, Rank: standard.Seven}:   '🂷',
		gocards.Card{Suit: standard.Hearts, Rank: standard.Eight}:   '🂸',
		gocards.Card{Suit: standard.Hearts, Rank: standard.Nine}:    '🂹',
		gocards.Card{Suit: standard.Hearts, Rank: standard.Ten}:     '🂺',
		gocards.Card{Suit: standard.Hearts, Rank: standard.Jack}:    '🂻',
		gocards.Card{Suit: standard.Hearts, Rank: standard.Queen}:   '🂽',
		gocards.Card{Suit: standard.Hearts, Rank: standard.King}:    '🂾',
		gocards.Card{Suit: standard.Spades, Rank: standard.Ace}:     '🂡',
		gocards.Card{Suit: standard.Spades, Rank: standard.Two}:     '🂢',
		gocards.Card{Suit: standard.Spades, Rank: standard.Three}:   '🂣',
		gocards.Card{Suit: standard.Spades, Rank: standard.Four}:    '🂤',
		gocards.Card{Suit: standard.Spades, Rank: standard.Five}:    '🂥',
		gocards.Card{Suit: standard.Spades, Rank: standard.Six}:     '🂦',
		gocards.Card{Suit: standard.Spades, Rank: standard.Seven}:   '🂧',
		gocards.Card{Suit: standard.Spades, Rank: standard.Eight}:   '🂨',
		gocards.Card{Suit: standard.Spades, Rank: standard.Nine}:    '🂩',
		gocards.Card{Suit: standard.Spades, Rank: standard.Ten}:     '🂪',
		gocards.Card{Suit: standard.Spades, Rank: standard.Jack}:    '🂫',
		gocards.Card{Suit: standard.Spades, Rank: standard.Queen}:   '🂭',
		gocards.Card{Suit: standard.Spades, Rank: standard.King}:    '🂮',
	}
)

func displayTableau(game gameData) {
	var row int
	var image [types.TableauWidth]rune

ROW_LOOP:
	for {
		var found bool
		for col := 0; col < types.TableauWidth; col++ {
			if row < game.remote.Tableau[col].HiddenCount {
				image[col] = displayCardBack
				found = true
			} else {
				visibleRow := row - game.remote.Tableau[col].HiddenCount
				if visibleRow < len(game.remote.Tableau[col].Cards) {
					image[col] = displayCards[game.remote.Tableau[col].Cards[visibleRow]]
					found = true
				} else {
					image[col] = ' '
				}
			}
		}
		if !found {
			break ROW_LOOP
		}
		fmt.Printf("%c %c %c %c %c %c %c %c %c %c\n",
			image[0], image[1], image[2], image[3], image[4],
			image[5], image[6], image[7], image[8], image[9],
		)

		row++
	}
}
