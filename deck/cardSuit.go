package deck

// CardSuit represents a suit from a normal deck of cards.
type CardSuit int

// Card suits from a normal deck of cards
const (
	Clovers CardSuit = iota
	Diamonds
	Hearts
	Spades
	SuitsCnt int = iota
)

var suitNames = []string{"Clovers", "Diamonds", "Hearts", "Spades"}

func (cs CardSuit) String() string {
	return suitNames[cs]
}
