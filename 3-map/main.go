package main

import "fmt"

func main() {
	colors := map[string]string {
		"red": "#ff0000",
		"white": "#000000",
	}

	// Empty map
	// var colors map[string]string
	// or
	// colors := make(map[string]string)

	colors["test"] = "green"

	fmt.Println(colors)

	delete(colors, "test")

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
