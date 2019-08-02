package decks

import (
	"math/rand"
	"sort"
)

// Comparator is used to sort a deck of cards based on a user-defined comparator.
type Comparator func(deck *Deck) func(i, j int) bool

// Sort is a functional option for sorting the deck based on a user-defined comparator.
func Sort(comp Comparator) Option {
	return func(deck *Deck) error {
		sort.Slice(deck.Cards, comp(deck))
		return nil
	}
}

// DefaultComparator is the default comparator used for sorting a deck of cards.
func DefaultComparator(deck *Deck) func(i, j int) bool {
	return func(i, j int) bool {
		cards := deck.Cards
		if cards[i].IsJoker {
			return true
		} else if cards[j].IsJoker {
			return false
		}

		if cards[i].Suit != cards[j].Suit {
			return cards[i].Suit < cards[j].Suit
		}

		return cards[i].Rank < cards[j].Rank
	}
}

// WithDecks is a functional option for specifying the amount of decks to include.
func WithDecks(cnt int) Option {
	return func(deck *Deck) error {
		var result []Card
		for i := 0; i < cnt; i++ {
			result = append(result, deck.Cards...)
		}

		deck.Cards = result
		return nil
	}
}

// WithJokers is a functional option for specifying how many jokers to include.
func WithJokers(cnt int) Option {
	return func(deck *Deck) error {
		for i := 0; i < cnt; i++ {
			deck.Cards = append(deck.Cards, Card{0, Clovers, true})
		}

		return nil
	}
}

// Shuffle is a functional option for shuffling the deck in a random way.
func Shuffle() Option {
	return func(deck *Deck) error {
		var shuffled []Card

		deckSize := len(deck.Cards)
		for i := 0; i < deckSize; i++ {
			pick := rand.Intn(len(deck.Cards))

			shuffled = append(shuffled, deck.Cards[pick])
			deck.Cards = append(deck.Cards[:pick], deck.Cards[pick+1:]...)
		}

		deck.Cards = shuffled
		return nil
	}
}

// ShuffleWithSeed is a functional option for shuffling the deck based on a random seed.
func ShuffleWithSeed(s int64) Option {
	rand.Seed(s)
	return Shuffle()
}

// Filter is a functional option for filtering given cards from the deck.
func Filter(filterFunc func(fc Card) bool) Option {
	return func(deck *Deck) error {
		var filtered []Card
		for _, c := range deck.Cards {
			if !filterFunc(c) {
				filtered = append(filtered, c)
			}
		}

		deck.Cards = filtered
		return nil
	}
}

func defaultDeck() *Deck {
	deck := Deck{}
	deck.Cards = make([]Card, DeckSize)
	for s := Clovers; s < SuitsCnt; s++ {
		for r := Ace; r < RanksCnt; r++ {
			i := int(s)*int(RanksCnt) + int(r)
			deck.Cards[i] = Card{r, s, false}
		}
	}

	return &deck
}
