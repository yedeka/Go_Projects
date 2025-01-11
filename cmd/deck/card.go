//go:generate stringer -type=Suit,Rank
package deck

import (
	"fmt"
	"math/rand/v2"
	"sort"
	"time"
)

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

type CardOptions func([]Card) []Card

type comparator func(i, j int) bool

type sorter func(cards []Card) comparator

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

func New(opts ...CardOptions) []Card {
	var cards []Card

	for _, suit := range suits {
		for rank := minRank; rank <= maxRank; rank += 1 {
			cards = append(cards, Card{Suit: suit, Rank: rank})
		}
	}

	for _, opt := range opts {
		cards = opt(cards)
	}

	return cards
}

func DefaultSort(cards []Card) []Card {
	sort.Slice(cards, Less(cards))
	return cards
}

var Less sorter = func(cards []Card) comparator {
	return func(i, j int) bool {
		return absRank(cards[i]) < absRank(cards[j])
	}
}

var ReverseLess sorter = func(cards []Card) comparator {
	return func(i, j int) bool {
		return absRank(cards[i]) > absRank(cards[j])
	}
}

func absRank(card Card) int {
	return int(card.Suit)*int(maxRank) + int(card.Rank)
}

func CustomSort(less sorter) CardOptions {
	return func(cards []Card) []Card {
		sort.Slice(cards, less(cards))
		return cards
	}
}

var Shuffle CardOptions = func(cards []Card) []Card {
	shuffledCards := make([]Card, len(cards))
	rand.New(rand.NewPCG(uint64(time.Now().Unix()), 15))
	perm := rand.Perm(len(cards))
	for i, j := range perm {
		shuffledCards[i] = cards[j]
	}
	return shuffledCards
}

func Jokers(n int) CardOptions {
	return func(cards []Card) []Card {
		for i := 0; i < n; i++ {
			cards = append(cards, Card{
				Suit: Joker,
				Rank: Rank(i),
			})
		}
		return cards
	}
}
