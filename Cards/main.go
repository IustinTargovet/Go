package main

import "fmt"

func main() {
	deck := newDeck()
	hand, remainingDeck := deal(deck, 7)
	fmt.Println(hand)
	fmt.Println(remainingDeck)
}
