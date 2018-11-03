package main

func main() {
	// var card string = "Ace of Spades"
	cards := newDeck()

	hand, remainingCards := deal(cards, 5)

	hand.print()
	remainingCards.print()

	// greet := "Hi there!"
	// string -> byte slice
	// fmt.Println([]byte(greet))

	// fmt.Println(cards.toString())
	// fmt.Println(cards.toByteSlice())

	// cards.saveToFile("my_deck")

	cardsFromFile := deckFromFile("my_decs")
	cardsFromFile.print()
}
