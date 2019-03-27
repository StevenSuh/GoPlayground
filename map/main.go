package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	for color, hex := range colors {
		fmt.Println(color, hex)
	}
}
