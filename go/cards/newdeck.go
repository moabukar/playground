package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	// loop through cardSuits and cardValues and create a card for each
	// replace unused variables with underscore
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	//cards.print()
	return cards
}

// return two values of type deck (hand and remaining cards)
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// function to convert deck	to string
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// function to save deck to file
func (d deck) saveToFile(name string) error {
	return os.WriteFile(name, []byte(d.toString()), 0666)
}

func newDeckFromFile(name string) deck {
	bs, err := os.ReadFile(name)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	s := deck(strings.Split(string(bs), ","))
	return s
}

func (d deck) shuffle() {
	// simple shuffle logic
	// loop through deck
	// generate random number between 0 and length of deck
	// swap current card with card at random number index

	// create new source for random number generator (in this case time otherwise it will generate same random number)
	rand.Seed(time.Now().UnixNano())

	for i := range d {
		newPosition := rand.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

func main() {
	// cards := newDeck()
	// cards.saveToFile("my_cards.txt")
	//////////////////////////////////
	// createCards := newDeckFromFile("my_cards.txt")
	// fmt.Println(createCards)
	// //////////////////////////////////
	// readCards := newDeck()
	// readCards.saveToFile("newcards.txt")
	///
	// shuffle
	cards := newDeck()
	cards.shuffle()
	cards.print()
}
