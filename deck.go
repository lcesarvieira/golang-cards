package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck'
// which is a slice of strings (extends)
type deck []string

// Deck constructor
func newDeck() deck {

	//starts slice as empty
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {

		//printing current suit
		//fmt.Println("suit:" + suit)

		for _, value := range cardValues {

			//printing current value
			//fmt.Println("value:" + value)

			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// Deck type receiver function
// for printing decks onto terminal
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// Deck type receiver function
//for dealing a hand of cards
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// deck type receiver function
// for joining deck of cards as a string
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// Deck type receiver function
// for saving deck of cards into file
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

// Deck type receiver function
// for loading deck of cards from file
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		// Option #1 - log the error and return a call to newDeck()
		// Option #2 - Log the error and entirely quit the program
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}

//deck type receiver function
//for shuffling deck hand of cards
func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
