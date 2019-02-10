package main

func main() {
	myDeck := newDeck()
	myDeck.saveToFile("test")

	newDeck := newDeckFromFile("test")
	newDeck.shuffle()
	newDeck.print()
}
