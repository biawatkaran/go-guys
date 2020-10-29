package main

import "fmt"

func main() {

	//var card string = "Ace of Spades" //long format
	//card := "Ace of Spades" //short format
	//cards := deck{newCard(), newCard()}
	//cards = append(cards, "A King Of Heart")

	cards := newDeck()
	fmt.Println("<<<<<<<<<< Complete Deck >>>>>>>>>>")
	cards.print()

	hand, remainingCards := deal(cards, 5)
	fmt.Println("<<<<<<<<<< Hand >>>>>>>>>>")
	hand.print()
	fmt.Println("<<<<<<<<<< Remaining >>>>>>>>>>")
	remainingCards.print()

	fmt.Println("<<<<<<<<<< To String >>>>>>>>>>")
	fmt.Println(cards.toString())

	filename := "my_cards"
	fmt.Println("<<<<<<<<<< Save To File >>>>>>>>>>")
	cards.saveToFile(filename)

	fmt.Println("<<<<<<<<<< Read From File >>>>>>>>>>")
	cardsFromFile := newDeckFromFile(filename)
	cardsFromFile.print()

	fmt.Println("<<<<<<<<<< Shuffle >>>>>>>>>>")
	cards.shuffle()
	cards.shuffleWithRandomness()
	cards.print()
}
