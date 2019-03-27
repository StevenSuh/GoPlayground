package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct {
	name string
}
type spanishBot struct {
	name string
}

func main() {
	eb := englishBot{name: "Steven"}
	sb := spanishBot{name: "Esteban"}

	printGreeting(eb)
	printGreeting(sb)
}

func (eb englishBot) getGreeting() string {
	return eb.name
}

func (sb spanishBot) getGreeting() string {
	return sb.name
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
