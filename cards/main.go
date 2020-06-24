package main

func main() {
	// cards := newDeck()

	// // deck{"Ace of Diamonds", newCard()} //cards of type slice that contains thee new card and Ace elements

	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// remainingCards.print()

	//////////////////////////////////////
	// cards := newDeck()

	// fmt.Println(cards.toString())

	// cards.saveToFile("my_cards")

	///////////////////////////////////////
	// cards := newDeckFromFile("my_cards")
	// cards.print()
	//////////////////////////////////////

	cards := newDeck()
	cards.shuffle()
	cards.print()
}
