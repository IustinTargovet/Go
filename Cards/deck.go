package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

//Create a new type "deck"
//Slice of strings

type deck []string

func newDeck() deck {
	newDeck := deck{}
	numbers := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	suits := []string{"Spades", "Clubs", "Hearts", "Diamonds"}

	for _, number := range numbers {
		for _, suit := range suits {
			card := number + " of " + suit
			newDeck = append(newDeck, card)
		}
	}

	return newDeck
}

func deal(d deck, handSize int) (deck, deck) {
	hand := d[:handSize]
	newDeck := d[handSize:]
	return hand, newDeck
}

func (d deck) toString() string {
	return strings.Join([]string(d), ", ")
}

func (d deck) toByteSlice() []byte {
	return []byte(d.toString())
}

func (d deck) saveToFile() {
	err := ioutil.WriteFile("deck", d.toByteSlice(), 0644)
	if err != nil {
		fmt.Println(err)
	}
}

func deckFromByteSlice(bs []byte) deck {
	return deck(strings.Split(string(bs), ", "))
}

func newDeckFromFile(fileName string) deck {
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return deckFromByteSlice(bs)
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
