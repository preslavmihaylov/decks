package main

import (
	"fmt"
	"log"

	"github.com/preslavmihaylov/learn-golang/gophercises/ex09-deck/deck"
)

func main() {
	cards, err := deck.New(deck.WithJokers(3), deck.Shuffle(), deck.Sort(deck.DefaultComparator))
	if err != nil {
		log.Fatalf("Something went wrong: %s", err)
	}

	for _, c := range cards {
		fmt.Println(c.String())
	}
}
