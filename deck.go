package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	suits := []string{"♦", "♥"}
	values := []string{"A", "2", "3", "4"}
	newDeck := deck{}

	for _, suit := range suits {
		for _, value := range values {
			newDeck = append(newDeck, value+suit)
		}
	}

	return newDeck
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	dealt := d[:handSize]
	remaining := d[handSize:]

	return dealt, remaining
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	byteslice := []byte(d.toString())
	return ioutil.WriteFile(filename, byteslice, 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("[!] ERROR:", err)
		os.Exit(1)
	}

	ss := strings.Split(string(bs), ",")
	return deck(ss)
}

func (d deck) shuffle() deck {
	seed := time.Now().UnixNano()
	source := rand.NewSource(seed)
	rng := rand.New(source)

	for i := range d {
		newPosition := rng.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
	return d
}
