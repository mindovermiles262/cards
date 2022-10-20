package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"white": "#ffffff",
		"black": "#000000",
	}

	// var colors map[string]string

	// colors := make(map[string]string)

	// Add data to map
	colors["green"] = "#00ff00"

	// Remove data from map
	delete(colors, "green")

	// fmt.Println(colors)

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("The color", color, "has the hex code of:", hex)
	}
}
