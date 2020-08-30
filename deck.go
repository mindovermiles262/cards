package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// We will create a "deck" type to hold the information for a deck of cards
// Create a new type ('deck'), which is a slice of string

type deck []string

func newDeck() deck {
	// Creates a new "deck" of cards
	cards := deck{}
	cardSuits := []string{
		"Hearts",
		"Diamonds",
		"Spades",
		"Clubs",
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
	// Prints contents of deck, each card on its own line
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	// Converts deck object into CSV string
	return strings.Join([]string(d), ",")
}

func (d deck) save(fileName string) error {
	// Writes deck to CSV file
	msg := []byte(d.toString())
	err := ioutil.WriteFile(fileName, msg, 0644)
	return err
}

func openFromFile(filename string) deck {
	byteslice, err := ioutil.ReadFile(filename)
	if err != nil {
		// Error handling, return new deck if there is an error
		fmt.Println("[!]", err)
		return newDeck()
	}
	strDeck := string([]byte(byteslice))
	strSlice := strings.Split(strDeck, ",")
	return deck(strSlice)
}
