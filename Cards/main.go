package main

func main() {
	cards := newDeck()

	hand, remainingCards := deal(cards, 5)
	cards = remainingCards
	hand.print()
	cards.print()
	hand.saveToFile()
}
