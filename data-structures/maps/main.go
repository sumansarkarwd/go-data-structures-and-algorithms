package main

import "fmt"

func main() {

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#554488",
	}

	printColors(colors)

	// var colors map[string]string

	// colors := make(map[int]string)

	// colors[10] = "red"

	// delete(colors, 10)

	// fmt.Println(colors)

}

func printColors(c map[string]string) {
	for color, hex := range c {
		fmt.Printf("%v: %v", color, hex)
		fmt.Println("")
	}
}
