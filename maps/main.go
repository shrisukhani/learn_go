package main

import "fmt"

func main() {
	colors := make(map[string]string)
	printMap(colors)
	colors["Red"] = "#ff0000"
	colors["Green"] = "#00ff00"
	printMap(colors)
	delete(colors, "Green")
	printMap(colors)
}

func printMap(c map[string]string) {
	for clr, hex := range c {
		fmt.Println("The hex value for", clr, "is", hex)
	}
	fmt.Println("---- x ----")
}
