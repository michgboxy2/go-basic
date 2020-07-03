package main

import "fmt"

func main() {
	// var colors map[string]string

	// colors := make(map[int]string)

	colors := map[string]string{ //declaring a map where all the keys are of type string and all the values are of type string
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "ffffff",
	}

	// colors[10] = "#fff"

	// delete(colors, 10) //delete from a map

	// fmt.Println(colors)
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
