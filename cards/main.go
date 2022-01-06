package main

import "fmt"

func main() {
	// var card string = "Ace of Spades"
	// card := newDeck()

	// hand, deck_stk := deal(card, 5)

	// hand.print()
	// deck_stk.print()
	// hand.saveToFile("deck_hand")

	card := deckFromFile("deck_hand")
	card.print()
	card.shuffle()
	fmt.Println("")
	card.print()
	// card.print()

}

func newCard() string {
	return "Five of hearts"
}
