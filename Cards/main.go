package main

import "fmt"

func main() {
	deck := newDeck()
	hand, remainingDeck := deal(deck, 5)
	fmt.Println(hand)
	fmt.Println(remainingDeck)
}
