package deck

import (
	"fmt"
	"testing"
)

func TestCard(t *testing.T) {
	fmt.Println(Card{Suit: Heart, Rank: King})
	fmt.Println(Card{Suit: Joker, Rank: Ace})
	//Output:
	// King of Hearts
	// Joker
}

func TestNew(t *testing.T) {
	newDeck := New()
	if len(newDeck) != 52 {
		t.Error("Not required number of cards in the deck")
	}
}
