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

func TestFilter(t *testing.T) {
	filter := func(card Card) bool {
		return card.Rank == Two || card.Rank == Three
	}
	filteredCards := New(Filter(filter))

	for _, filteredCard := range filteredCards {
		if filteredCard.Rank == Two || filteredCard.Rank == Three {
			t.Error("Cards of Rank Two or Three are not filered correctly")
		}
	}
}

func TestDeck(t *testing.T) {
	deckOfCards := New(Deck(4))
	epectedLength := 13 * 4 * 4
	if epectedLength != len(deckOfCards) {
		t.Errorf("Deck cards count required: %d, found: %d", 13*4*4, len(deckOfCards))
	}
}
