package main

import (
	"fmt"
)

func main() {
	evenodd := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i := range evenodd {
		if evenodd[i]%2 == 0 {
			fmt.Printf("The numeber is even :- %v. ", evenodd[i])

		} else {
			fmt.Printf("The numeber is odd :- %v. ", evenodd[i])
		}
	}
}
