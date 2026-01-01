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

/*
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

	// print all cards and values without comama anything raw data
	// cards.print()
	// it should call function and put all values there its syntax correct ideally
	// how to assign multiple values in multi value return function

	// deal(cards , 5)

	// assigning values.   1st value will be assigned to first here like that

	hand, remianingdeck := deal(cards, 7)
	fmt.Println("printing 7 cards ")
	hand.print()
	fmt.Println("remaining in deck")
	remianingdeck.print()

	// EXMAPLE of TYPE CONVERSION.  we write the type to convert to and the string like []byte (greetings)
	greetings := "Hi There"
	fmt.Println(greetings)
	fmt.Println("converted")

	fmt.Println([]byte(greetings))

}
*/

//sample function declaration

// func newCard() string {
// 	return "Five of Diamonds"
// }

// FUNCTION MAIN AGAIN  for FILE HANDLING

func main() {

	cards := newdeck()
	// printing the whole string connected by ,

	fmt.Println(cards.toString())

}
