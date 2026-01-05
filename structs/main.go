package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
}

func main() {
	// one way of defining without mentioning anythng it  will assume as in the list order provided above.

	alex := Person{"Alex", "Anderson"}
	fmt.Println(alex)

	// if the order changes above the values will be mapped wrong so another better way is

	Alex := Person{FirstName: "Alex", LastName: "Anderson"}
	fmt.Println(Alex)

	//defining the variable without values

	var Alex1 Person
	fmt.Println(Alex1)

	// add %+v to add both the names and values all in a string

	fmt.Printf("%+v", Alex1)

	// we can also define the struct values like

	Alex1.FirstName = "Alex"

	Alex1.LastName = "Anderson"

	fmt.Printf("%+v", Alex1)

}
