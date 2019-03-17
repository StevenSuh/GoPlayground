package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
)

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{
		"Spades",
		"Diamonds",
		"Hearts",
	}
	cardValues := []string{
		"Ace",
		"Two",
		"Three",
	}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) toString(separator string) string {
	sep := ","
	if separator != "" {
		sep = separator
	}

	return strings.Join([]string(d), sep)
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[handSize:], d[:handSize]
}

func (d deck) saveToFile(params ...string) error {
	if len(params) == 0 {
		return errors.New("Params is too short")
	} else if len(params) > 1 {
		return ioutil.WriteFile(params[0], []byte(d.toString(params[1])), 0666)
	}
	return ioutil.WriteFile(params[0], []byte(d.toString("")), 0666)
}

func newDeckFromFile(params ...string) (deck, error) {
	byteSlice, error := ioutil.ReadFile(params[0])
	if error != nil {
		return nil, errors.New("Poopie")
	}

	separator := ","
	if len(params) > 1 {
		separator = params[1]
	}
	strArr := strings.Split(string(byteSlice), separator)
	return deck(strArr), nil
}
