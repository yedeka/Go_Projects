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

func TestDefaultSort(t *testing.T) {
	newDeck := New(DefaultSort)
	expectedCard := Card{Suit: Spade, Rank: Ace}
	if expectedCard != newDeck[0] {
		t.Error("Expected Ace of Spade as first card but received", newDeck[0])
	}
}

func TestCustomSort(t *testing.T) {
	newDeck := New(CustomSort(ReverseLess))
	expectedCard := Card{Suit: Heart, Rank: King}
	if expectedCard != newDeck[0] {
		t.Error("Expected Ace of Spade as first card but received", newDeck[0])
	}
}

func TestJokers(t *testing.T) {
	newDeckWithJokers := New(Jokers(3))
	var cnt uint
	for _, j := range newDeckWithJokers {
		if j.Suit == Joker {
			cnt += 1
		}
	}
	if cnt != 3 {
		t.Error("Expected Jokers: 4, Received : ", cnt)
	}
}
