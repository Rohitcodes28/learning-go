package main

import "fmt"

// create a new type deck  which is just a slice of strings

type deck []string

// creating a finction newDeck that makes a completly new array with all possoble card combinations

func newdeck() deck {
	cards := deck{}

	cardType := []string{"spades", "diamonds", "clubs", "hearts"}
	cardValue := []string{"Ace", "two", "three", "four"}

	for _, suit := range cardType {
		for _, value := range cardValue {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() {
	fmt.Println(d)
	fmt.Println("printing from the loop")
	for i, card := range d {
		fmt.Println(i, card)
	}
}
