package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}
	colors["white"] = "#ffffff"
	delete(colors, "blue")

	phuc := make(map[string]string)
	phuc["first"] = "Phuc"
	phuc["last"] = "Nguyen"

	fmt.Println(colors)
	fmt.Println(phuc)

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}