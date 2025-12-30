package main

import "fmt"

// function that prints the value of the  variables and variable definiations
/*
func main() {

	card := newCard()

	cardint := newCardint()

	fmt.Println("Hi there!")
	fmt.Println("the card is", card)
	fmt.Println("the card is", cardint)
}
*/

// defining the function with return type string

// func newCard() string {
// 	return "Five of Diamonds"
// }

// func newCardint() int {
// 	return 12
// }

// Array and Slices

// array is of fixed length and slice can have variable lengths.

func main() {

	// defining a slice and adding a new value using append, using function to add value to slice
	// make sure correct data type is provided its string not strings.

	//deck is defined as slice of strings in the deck.go or wrtite []string if its not working
	// cards := deck{"ace of spades", newCard(), "four of diamonds"}
	// cards := []string{"ace of spades", newCard(), "four of diamonds"}

	// fmt.Println("values in original slice is", cards)

	// cards = append(cards, "king of diamonds")

	// fmt.Println("values after appending original slice is", cards)

	// as we are using the type deck we can add it as a receiver function

	// for i, card := range cards {
	// 	fmt.Println("printing from the loop")
	// 	fmt.Println(i, card)
	// }

	// cards.print()

	//using the dunction defined in deck.go to get a list of cards

	cards := newdeck()
	total_deck_cards := len(cards)
	fmt.Println("total cards in deck are", total_deck_cards)

	cards.print()

}

//sample function declaration

// func newCard() string {
// 	return "Five of Diamonds"
// }
