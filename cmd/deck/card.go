//go:generate stringer -type=Suit,Rank
package deck

import "fmt"

type Suit uint8

type Rank uint8

const (
	Spade Suit = iota
	Diamond
	Club
	Heart
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

var suits = [...]Suit{Spade, Diamond, Club, Heart}

const (
	minRank = Ace
	maxRank = King
)

func (c Card) String() string {
	if c.Suit == Joker {
		return c.Suit.String()
	}
	return fmt.Sprintf("%s of %ss", c.Rank.String(), c.Suit.String())
}

func New() []Card {
	var cards []Card

	for _, suit := range suits {
		for rank := minRank; rank <= maxRank; rank += 1 {
			cards = append(cards, Card{Suit: suit, Rank: rank})
		}
	}
	return cards
}
