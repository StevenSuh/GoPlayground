package main

import "fmt"

func main() {
	cards := newDeck()
	cards.saveToFile("myCards.txt", "\n")

	loadedCards, error := newDeckFromFile("myCards.txt", "\n")
	if error != nil {
		fmt.Println("Error:", error)
	} else {
		loadedCards.print()
	}
}
