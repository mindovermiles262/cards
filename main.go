package main

// func (r receiver) functionName(arg argtype) returntype {
// 	 Do something
// }

func main() {

	// cards := deck{ // Create a new slice of strings, with two cards
	// 	newCard(),
	// 	createCard("King of Hearts"),
	// }

	// cards = append(cards, "Six of Spades") // append() returns a new slice
	// cards.print()                          // Call the cards 'print' method defined in ./deck.go

	deck1 := newDeck()
	// deck1.print()

	// dealt, remaining := deal(deck1, 3)
	// dealt.print()
	// remaining.print()

	// fmt.Println(deck1.toString())
	// deck1.saveToFile("deck1.deck")

	// deck2 := newDeckFromFile("deck1.deck")
	// fmt.Println("[*] D2: ", deck2)
	// deck3 := newDeckFromFile("nope.deck")
	// fmt.Println("[*] D3: ", deck3)

	deck1.shuffle()
	deck1.print()
}

// func newCard() string {
// 	return "Five of Diamonds"
// }

// func createCard(cardName string) string {
// 	return cardName
// }
