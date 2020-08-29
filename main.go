package main

import "fmt"

func main() {
	// Initialize new card variable
	// card := newCard()
	// fmt.Println(cards)

	// Cards "Slice"
	// The append function does not modify existing slice
	cards := []string{
		newCard(),
		"Ace of Hearts",
	}

	cards = append(cards, "Three of Spades")
	// fmt.Println(cards)

	// Iterate over slice
	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Five of Diamonds"
}
