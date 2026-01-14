package main

import "fmt"

// basic defination of the maps
func main() {
	// 1st way to declare map

	// var colors map[string]string

	// 2nd way to make a map

	// colors := make(map[string]string)

	// colors["white"] = "#ffffff"

	// 3rd way to declare
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}
	fmt.Println(colors)

	//deleting the  value in map

	delete(colors, "white")

	fmt.Println(colors)

	printMap(colors)

}

// write a function to iterate over the map values and make sure that the map can be read

// format is like func name (variable , type_of_map)

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println(" Hex code for ", color, "is ", hex)

	}
}
