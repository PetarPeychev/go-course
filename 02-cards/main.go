package main

import (
	"fmt"
	"os"
)

func main() {
	var hand deck
	deck := newDeck()
	deck.shuffle()

	for i := 0; i < 3; i++ {
		fmt.Println("Hand", i+1, ":")
		hand, deck = deck.deal(5)
		hand.print()
		fmt.Println()
	}

	deck.toFile("deck.txt")
	deck, err := deckFromFile("cards.txt")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	} else {
		deck.print()
	}
}
