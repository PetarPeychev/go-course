package main

import "fmt"

func main() {
	var hand deck
	deck := newDeck()

	for i := 0; i < 3; i++ {
		fmt.Println("Hand", i+1, ":")
		hand, deck = deck.deal(5)
		hand.print()
		fmt.Println()
	}
}
