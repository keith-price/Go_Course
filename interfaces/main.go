package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type indonesianBot struct{}

func main() {
	eb := englishBot{}
	ib := indonesianBot{}

	printGreeting(eb)

	printGreeting(ib)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	return "Good morning!"
}

func (indonesianBot) getGreeting() string {
	return "Selamat pagi!"
}
