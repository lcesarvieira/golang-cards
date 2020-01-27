//defines an executable package
package main

import "fmt"

func main() {
	cards := newDeck()
	fmt.Println(cards.toString())

	//deal hand of 3 cards
	dealHand, otherCards := deal(cards, 3)
	dealHand.print()
	otherCards.print()

	cards.shuffle()
	cards.saveToFile("__deck_main")

	cards.print()
}
