package main

import "fmt"

// GOLANG DOES NOT HAVE CLASSES ONLY FUNCTIONS! //

func main() {
	//card := "One of Diamonds"
	//card := NewCard()
	//fmt.Println(card)
	//cards := deck{NewCard(), NewCard(), NewCard()}
	//cards = append(cards, "Six of Kings,")
	//fmt.Println(cards)
	//for i, card := range cards {
	//	fmt.Println(i, card)
	//}
	//cards.print()
	//cards.print()

	//cards := newDeck()
	//hand, remainingDeck := deal(cards, 5)
	//hand.print()
	//remainingDeck.print()

	//cards := newDeck()
	//fmt.Println(cards.ToString())

	//cards := newDeck()
	//cards.saveToFile("MyCards")

	cards := newDeckFromFile("MyCards")
	fmt.Println("Shuffled deck:", cards.shuffleDeck())

}
