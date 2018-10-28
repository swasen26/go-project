package main

import "fmt"

func main() {
	// var card string = "Ace of Spades"
	card := newCard()

	cards := []string{"Ace of Diamonds", newCard()}

	cardsDeck := deck{"Ace of Diamonds", newCard()}

	cards = append(cards, card)

	fmt.Println(card)

	for i, card := range cards {
		fmt.Println(i, card)
	}

	cardsDeck.print()

}

func newCard() string {

	return "Five of Diamonds"
}
