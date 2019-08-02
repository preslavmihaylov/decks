# decks

Package decks provides a simple API for creating and customizing decks of cards in Go.

## Installation
```
go get github.com/preslavmihaylov/decks
```

## Quickstart
```go
// get a default, ordered deck of 52 cards
d, err := decks.New()
// handle error

// Shuffle deck of cards
decks.Shuffle()

// draw cards
myHand := []decks.Card{}
myHand = append(myHand, d.Draw())
myHand = append(myHand, d.Draw())

// print hand
for _, c := range myHand {
  fmt.Println(c)
}

// discard hand
d.InsertBottom(myHand)
```

[Try it on Go Playground](https://play.golang.org/p/0NS_9C5DlYU)
