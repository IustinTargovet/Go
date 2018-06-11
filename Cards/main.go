package main

import "fmt"

func main() {
	deck := newDeck()
	hand, remainingDeck := deal(deck, 1)
	fmt.Println(hand)
	fmt.Println(remainingDeck)
}
