package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct {}
type spanishBot struct {}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func (englishBot) getGreeting() string {
	// Very custom logic for returning english greeting
	return "Hi there"
}

func (spanishBot) getGreeting() string {
	// VERY custom logic for returning spanish  greeting
	return "Hola"
}

//func printGreeting(eb englishBot) {
//	fmt.Println(eb.getGreeting())
//}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}