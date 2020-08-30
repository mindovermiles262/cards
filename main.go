package main

func main() {
	cards := newDeck()
	// cards.print()
	// hand, remainingDeck := deal(cards, 5)

	// hand.print()
	// remainingDeck.print()

	// fmt.Println(cards.toString())
	cards.save("deck.csv")

	restoredDeck := openFromFile("deck.csv")
	restoredDeck.print()
}
