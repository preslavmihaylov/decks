package deck

import (
	"math/rand"
	"sort"
	"time"
)

// Constants used to define boundaries for normal decks of cards
const (
	DefaultDeckSize = int(SuitsCnt) * int(RanksCnt)
)

// Option is a function alias used for constructing a functional-options constructor for a deck.
type Option func([]Card) ([]Card, error)

// New deck of cards, based on the options provided.
// In case none are passed, a sorted default deck of cards (without jokers) will be returned.
func New(opts ...Option) ([]Card, error) {
	cards := defaultDeck()
	for _, opt := range opts {
		var err error
		cards, err = opt(cards)
		if err != nil {
			return nil, err
		}
	}

	return cards, nil
}

// Comparator is used to sort a deck of cards based on a user-defined comparator.
type Comparator func(cards []Card) func(i, j int) bool

// Sort is a functional option for sorting the deck based on a user-defined comparator.
func Sort(comp Comparator) Option {
	return func(cards []Card) ([]Card, error) {
		sort.Slice(cards, comp(cards))
		return cards, nil
	}
}

// DefaultComparator is the default comparator used for sorting a deck of cards.
func DefaultComparator(cards []Card) func(i, j int) bool {
	return func(i, j int) bool {
		if cards[i].IsJoker {
			return true
		} else if cards[j].IsJoker {
			return false
		}

		if cards[i].Suit == cards[j].Suit {
			return cards[i].Rank < cards[j].Rank
		}

		return cards[i].Suit < cards[j].Suit
	}
}

// WithDecks is a functional option for specifying the amount of decks to include.
func WithDecks(cnt int) Option {
	return func(cards []Card) ([]Card, error) {
		cards = make([]Card, 0)
		for i := 0; i < cnt; i++ {
			cards = append(cards, defaultDeck()...)
		}

		return cards, nil
	}
}

// WithJokers is a functional option for specifying how many jokers to include.
func WithJokers(cnt int) Option {
	return func(cards []Card) ([]Card, error) {
		for i := 0; i < cnt; i++ {
			cards = append(cards, Card{0, Clovers, true})
		}

		return cards, nil
	}
}

// Shuffle is a functional option for shuffling the deck in a random way.
func Shuffle() Option {
	return ShuffleWithSeed(time.Now().UnixNano())
}

// ShuffleWithSeed is a functional option for shuffling the deck based on a random seed.
func ShuffleWithSeed(s int64) Option {
	return func(cards []Card) ([]Card, error) {
		shuffled := []Card{}

		rand.Seed(s)
		deckSize := len(cards)
		for i := 0; i < deckSize; i++ {
			pick := rand.Intn(len(cards))
			shuffled = append(shuffled, cards[pick])
			cards = append(cards[:pick], cards[pick+1:]...)
		}

		return shuffled, nil
	}
}

// Filter is a functional option for filtering given cards from the deck.
func Filter(filteredDecks ...[]Card) Option {
	return func(cards []Card) ([]Card, error) {
		for _, fd := range filteredDecks {
			for _, fc := range fd {
				for i, c := range cards {
					if c.Rank == fc.Rank && c.Suit == fc.Suit {
						cards = append(cards[:i], cards[i+1:]...)
						break
					}
				}
			}
		}

		return cards, nil
	}
}

// Aces in a standard deck
func Aces() []Card { return cardsWithValue(Ace) }

// Twos in a standard deck
func Twos() []Card { return cardsWithValue(Two) }

// Threes in a standard deck
func Threes() []Card { return cardsWithValue(Three) }

// Fours in a standard deck
func Fours() []Card { return cardsWithValue(Four) }

// Fives in a standard deck
func Fives() []Card { return cardsWithValue(Five) }

// Sixes in a standard deck
func Sixes() []Card { return cardsWithValue(Six) }

// Sevens in a standard deck
func Sevens() []Card { return cardsWithValue(Seven) }

// Eights in a standard deck
func Eights() []Card { return cardsWithValue(Eight) }

// Nines in a standard deck
func Nines() []Card { return cardsWithValue(Nine) }

// Tens in a standard deck
func Tens() []Card { return cardsWithValue(Ten) }

// Jacks in a standard deck
func Jacks() []Card { return cardsWithValue(Jack) }

// Queens in a standard deck
func Queens() []Card { return cardsWithValue(Queen) }

// Kings in a standard deck
func Kings() []Card { return cardsWithValue(King) }

func cardsWithValue(r Rank) []Card {
	cs := make([]Card, SuitsCnt)
	for s := Clovers; s < SuitsCnt; s++ {
		cs[int(s)] = Card{r, s, false}
	}

	return cs
}

func defaultDeck() []Card {
	cards := make([]Card, DefaultDeckSize)
	for s := Clovers; s < SuitsCnt; s++ {
		for r := Ace; r < RanksCnt; r++ {
			i := int(s)*int(RanksCnt) + int(r)
			cards[i] = Card{r, s, false}
		}
	}

	return cards
}
