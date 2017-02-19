package main

import (
	"fmt"

	"github.com/dougfort/gocards/standard"

	"github.com/dougfort/arachne/game"
)

func displayTableauRunes(g gameData) {
	var row int
	image := make([]rune, game.TableauWidth)

ROW_LOOP:
	for {
		var found bool
		for col := 0; col < game.TableauWidth; col++ {
			if row < g.remote.Tableau[col].HiddenCount {
				image[col] = standard.RuneBack
				found = true
			} else {
				visibleRow := row - g.remote.Tableau[col].HiddenCount
				if visibleRow < len(g.remote.Tableau[col].Cards) {
					image[col] = standard.Runes[g.remote.Tableau[col].Cards[visibleRow]]
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
