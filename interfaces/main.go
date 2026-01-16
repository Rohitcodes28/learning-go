package main

import "fmt"

// interface declaration
type bot interface {
	getGreeting() string
}

// we will be working on the interfaces  example using a bot english and spanish bot

// empty struct. with no params
type englishBot struct{}
type spanishBot struct{}

func main() {

	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb)

}

// we are not reusing the code of get Greeting its a seprate custom code function for both bots.
// we will make a common function as well let's assusme this is a very custom logic function

// we can write eb englishBIot or we can also write the englishBOt only since we are not using the eb anywhere its totally fine.
func (eb englishBot) getGreeting() string {
	return "Hi there!"
}

func (sb spanishBot) getGreeting() string {
	return "Hola!"
}

// we can not have same names if we are passing the differnt values and function is non receiver function
// func printGreeting(eb englishBot) {
// 	fmt.Println(eb.getGreeting())
// }

// func printGreeting(sb spanishBot) {
// 	fmt.Println(sb.getGreeting())
// }

// instead we will have a better function now that we have defined interface on top
// bot type has both englishBot and spanishBot as they are both having getGreeting() string. function
// any other type that also has a getGreeting() string.  will be included automatically in bot type.

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
