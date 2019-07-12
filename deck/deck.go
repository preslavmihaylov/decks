package deck

import (
	"math/rand"
	"sort"
	"time"
)

// Constants used to define boundaries for normal decks of cards
const (
	DefaultCardsInSuit = 13
	DefaultDeckSize    = int(SuitsCnt) * DefaultCardsInSuit
)

// Card represents a playing card from a deck of cards.
// It contains a value and a suit. The value is represented as an integer,
// where 1 stands for Ace, 11 for Jack, 12 for Queen, 13 for King.
// It also contains a boolean flag denoting this card as a joker.
type Card struct {
	Value   int
	Suit    CardSuit
	IsJoker bool
}

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
func Sort(comparator Comparator) Option {
	return func(cards []Card) ([]Card, error) {
		sort.Slice(cards, comparator(cards))
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
			return cards[i].Value < cards[j].Value
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
					if c.Value == fc.Value && c.Suit == fc.Suit {
						cards = append(cards[:i], cards[i+1:]...)
						break
					}
				}
			}
		}

		return cards, nil
	}
}

// Aces - all aces in a standard deck
func Aces() []Card { return cardsWithValue(1) }

// Twos - all twos in a standard deck
func Twos() []Card { return cardsWithValue(2) }

// Threes - all threes in a standard deck
func Threes() []Card { return cardsWithValue(3) }

// Fours - all fours in a standard deck
func Fours() []Card { return cardsWithValue(4) }

// Fives - all fives in a standard deck
func Fives() []Card { return cardsWithValue(5) }

// Sixes - all sixes in a standard deck
func Sixes() []Card { return cardsWithValue(6) }

// Sevens - all sevens in a standard deck
func Sevens() []Card { return cardsWithValue(7) }

// Eights - all eights in a standard deck
func Eights() []Card { return cardsWithValue(8) }

// Nines - all nines in a standard deck
func Nines() []Card { return cardsWithValue(9) }

// Tens - all tens in a standard deck
func Tens() []Card { return cardsWithValue(10) }

// Jacks - all jacks in a standard deck
func Jacks() []Card { return cardsWithValue(11) }

// Queens - all queens in a standard deck
func Queens() []Card { return cardsWithValue(12) }

// Kings - all kings in a standard deck
func Kings() []Card { return cardsWithValue(13) }

func cardsWithValue(v int) []Card {
	cs := make([]Card, SuitsCnt)
	for s := Clovers; s <= Spades; s++ {
		cs[int(s)] = Card{v, s, false}
	}

	return cs
}

func defaultDeck() []Card {
	cards := make([]Card, DefaultDeckSize)
	for s := Clovers; s <= Spades; s++ {
		for v := 0; v < DefaultCardsInSuit; v++ {
			cards[int(s)*DefaultCardsInSuit+v] = Card{v + 1, s, false}
		}
	}

	return cards
}
