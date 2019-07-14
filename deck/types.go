//go:generate stringer -type=Suit
//go:generate stringer -type=Rank

package deck

// Suit from a normal deck of cards.
// i.e. Clovers, Diamonds, etc...
type Suit uint8

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

func (c *Card) String() string {
	if c.IsJoker {
		return "Joker"
	}

	return c.Rank.String() + " of " + c.Suit.String()
}
