package deck

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()
}