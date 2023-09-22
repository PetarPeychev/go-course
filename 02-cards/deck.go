package main

import (
	"fmt"
	"os"
	"strings"
)

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"♠", "♤", "♥", "♡", "♦", "♢", "♣", "♧"}
	cardValues := []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, "|"+value+" "+suit+"|")
		}
	}

	return cards
}

func (d deck) print() {
	fmt.Println(d.toString())
}

func (d deck) deal(handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), "   ")
}

func deckFromString(s string) deck {
	return deck(strings.Split(s, "   "))
}

func (d deck) toFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func deckFromFile(filename string) (deck, error) {
	bytes, err := os.ReadFile(filename)
	return deckFromString(string(bytes)), err
}
