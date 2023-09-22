package main

import "fmt"

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"♠", "♤", "♥", "♡", "♦", "♢", "♣", "♧"}
	cardValues := []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}
