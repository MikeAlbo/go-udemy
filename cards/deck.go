package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// create a new typw of deck that is a slice of strings

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Hearts", "Spades", "Diamonds", "Clubs"}
	cardValue := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValue {
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

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(fileName string) deck {
	bs, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Println("error: ", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}
