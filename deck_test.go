package main

import ( 
  "testing"
  "os"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
  firstCard := d[0]
  lastCard := d[len(d)-1] 

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %d", len(d))
	}
  
  if firstCard != "Ace of Hearts" {
    t.Errorf("Expected first card to be 'Ace of Hearts' but got %v", firstCard)
  }
 
  if lastCard != "Four of Clubs" {
    t.Errorf("Expected last card to be 'Four of Clubs' but got %v", lastCard)
  }
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
  saveFile := "_decktesting"
  os.Remove(saveFile) 
  deck := newDeck()
  deck.save(saveFile)
  deckFromFile := openFromFile(saveFile)

  if len(deckFromFile) != 16 {
    t.Errorf("Expected 16 cards in deck but got %d", len(deckFromFile))
  }

  os.Remove(saveFile)
}
