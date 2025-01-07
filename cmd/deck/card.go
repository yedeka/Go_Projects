package deck

type Suit uint8

type Rank uint8

const (
	Spade Suit = iota
	Diamond
	Club
	Hart
	Joker
)

const (
	// _ is used to have 0 value to be unused so that we can start with value of 1 for Ace
	_ Rank = iota
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

type Card struct {
	Suit
	Rank
}
