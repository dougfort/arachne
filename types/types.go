package types

// Suit represents the suit of a card
type Suit uint8

// Suit values
const (
	NullSuit Suit = iota
	Clubs
	Diamonds
	Hearts
	Spades
)

// Rank represents the order of a Suit
type Rank uint8

// Rank Values
const (
	NullRank Rank = iota
	Ace
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

// CardType represents the value of a single card
type CardType struct {
	Suit Suit
	Rank Rank
}

// CardsType is a group of Card(s)
type CardsType []CardType

// StackType reresesnts on stack of cards in the Tableau
type StackType struct {
	HiddenCount int
	Cards       CardsType
}

// TableauWidth is the number of stacks in the Tableau
const TableauWidth = 10

// Tableau is the game layout
type Tableau []StackType
