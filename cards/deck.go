package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
)

//	Create a new Deck Type Variable
//	slice of strings

type deck []string

func newDeck() deck {

	cards := deck{}

	cardSuits := []string{"Hearts", "Clubs", "Diamonds", "Spades"}
	cardValues := []string{"Ace", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func deal(d deck, handSize int) (deck, deck) {

	return d[:handSize], d[handSize:]

}

// func (receiver) func_name return_type{}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) toByteSlice() []byte {
	stringElement := d.toString()
	return []byte(stringElement)
}

func (d deck) saveToFile(fileName string) error {
	// 0666 -> permission: any one can access the file
	return ioutil.WriteFile(fileName, d.toByteSlice(), 0666)
}

func deckFromFile(fileName string) deck {
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error: ", err)
		//return newDeck()
		os.Exit(1)
	}
	return deck(strings.Split(string(bs), ","))
}

func (d deck) shuffleDeck() deck {
	for i := range d {
		newPosition := rand.Intn(len(d) - 1)
		// temp := d[i]
		// d[i] = d[newPosition]
		// d[newPosition] = temp

		d[i], d[newPosition] = d[newPosition], d[i]
	}
	return d
}
