package main

import (
	"fmt"
	"strings"
)

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

// will only accept the variables of type deck  and will be used to print them=
// it is designed to print both the index and the value

func (d deck) print() {
	// fmt.Println(d)
	fmt.Println("printing from the print function below is your result of data type deck ")
	for i, card := range d {
		fmt.Println(i, card)
	}
}

//it will return multiple values has 2 deck deck  both values will be of type deck only as
//we are splitting that only but will retun 2 values.
//basically this function when used in main.go will split the deck in 2 parts and share that

func deal(d deck, hand int) (deck, deck) {
	return d[:hand], d[hand:]

}

// sample program below showing multiple retun values of the function
/*
	package main

	import "fmt"

	func main() {
		color1, color2, color3 := colors()

		fmt.Println(color1, color2, color3)
	}

	func colors() (string, string, string) {
		return "red", "yellow", "blue"
	}


*/

// FILE HANDLING

// we need a function that can combine a slice of string into string and make sure that we get a final string before writing to file

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
	// we are converting the deck back to slice of strings as deck is our datatype not generic to go
}
