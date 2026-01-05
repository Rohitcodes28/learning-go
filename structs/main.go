package main

import "fmt"

// embedded struct examples

type contactInfo struct {
	email   string
	zipcode int
}

type Person struct {
	FirstName string
	LastName  string
}

type jim struct {
	FirstName string
	LastName  string
	contact   contactInfo
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

	//Make sure that a struct always ends a line with a " , "  never forget that even if the last line is there the comma should be there.
	RealName := jim{
		FirstName: "jim",
		LastName:  "anderson",
		contact: contactInfo{email: "jim@gmail.com",
			zipcode: 140021,
		},
	}
	fmt.Printf("%+v", RealName)

}
