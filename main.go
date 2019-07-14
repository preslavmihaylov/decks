package main

import (
	"fmt"
	"log"

	"github.com/preslavmihaylov/learn-golang/gophercises/ex09-deck/decks"
)

func main() {
	d, err := decks.New(decks.WithJokers(3), decks.Shuffle(), decks.Sort(decks.DefaultComparator))
	if err != nil {
		log.Fatalf("Something went wrong: %s", err)
	}

	d.InsertBottom([]decks.Card{decks.Card{
		Rank: decks.Ace,
		Suit: decks.Clovers,
	}})

	for _, c := range d.Cards {
		fmt.Println(c.String())
	}
}
