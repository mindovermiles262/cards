package main

import "fmt"

// We will create a "deck" type to hold the information for a deck of cards
// Create a new type ('deck'), which is a slice of string

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{
		"Hearts",
		"Diamonds",
		// "Spades",
		// "Clubs",
	}
	cardValues := []string{
		"Ace",
		"Two",
		"Three",
		"Four",
		// "Five",
		// "Six",
		// "Seven",
		// "Eight",
		// "Nine",
		// "Ten",
		// "Jack",
		// "Queen",
		// "King",
	}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
