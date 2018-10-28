package main

import (
	"fmt"
)

func main() {
	// var card string = "Ace of Spades"
	cards := newDeck()

	hand, remainingCards := deal(cards, 5)

	hand.print()
	fmt.Println("----------")
	remainingCards.print()

}
