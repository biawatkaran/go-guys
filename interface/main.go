package main

import "fmt"

type bot interface {
	getGreeting() string
}

// interfaces are implicit i.e. like other languages `type englishBot struct{} extends/implements bot` - is not needed
type englishBot struct{}
type spanishBot struct{}

func main() {

	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

	readSiteData()
}

func printGreeting(b bot) {

	fmt.Println(b.getGreeting())
}

/*
//BAD Way 1. go doesn't allow 2. functioning wise same - that's where interface comes :)
func printGreeting(sb spanishBot) {

	fmt.Println(sb.getGreeting())
}

func printGreeting(eb englishBot) {

	fmt.Println(eb.getGreeting())
}

*/

func (englishBot) getGreeting() string {

	return "Hi There!"
}

func (spanishBot) getGreeting() string {

	return "Hola!"
}
