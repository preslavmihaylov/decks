//go:generate stringer -type=Suit
//go:generate stringer -type=Rank

package decks

// Suit from a normal deck of cards.
// i.e. Clovers, Diamonds, etc...
type Suit uint8

const (
	// DeckSize of a normal deck of cards.
	DeckSize = int(SuitsCnt) * int(RanksCnt)
)

// Suits from a normal deck of cards
const (
	Clovers Suit = iota
	Diamonds
	Hearts
	Spades
	SuitsCnt
)

// Rank from a normal deck of cards.
// i.e. Ace, Two, Jack, etc...
type Rank uint8

// The ranks from a normal deck of cards.
const (
	Ace Rank = iota
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
	RanksCnt
)

// Card represents a playing card from a deck of cards.
// It contains a rank (e.g. Ace) and a suit (e.g. Clovers).
// It also contains a boolean flag denoting this card as a joker.
type Card struct {
	Rank
	Suit
	IsJoker bool
}

func (c Card) String() string {
	if c.IsJoker {
		return "Joker"
	}

	return c.Rank.String() + " of " + c.Suit.String()
}

// Deck encapsulates a normal deck of cards and defines some operations on them.
type Deck struct {
	Cards []Card
}

// Option is a function alias used for constructing a functional-options constructor for a deck.
type Option func(*Deck) error

// New deck of cards, based on the options provided.
// In case none are passed, a sorted default deck of cards (without jokers) will be returned.
func New(opts ...Option) (*Deck, error) {
	deck := defaultDeck()
	for _, opt := range opts {
		var err error
		err = opt(deck)
		if err != nil {
			return nil, err
		}
	}

	return deck, nil
}

// Draw a card from the deck.
// The drawn card is returned and removed from the deck.
func (d *Deck) Draw() Card {
	c := d.Cards[0]
	d.Cards = d.Cards[1:]

	return c
}

// Shuffle the deck randomly.
func (d *Deck) Shuffle() error {
	var err error
	err = Shuffle()(d)
	if err != nil {
		return err
	}

	return nil
}

// InsertBottom puts the given cards at the bottom of the deck.
func (d *Deck) InsertBottom(cs []Card) {
	d.Cards = append(d.Cards, cs...)
}
