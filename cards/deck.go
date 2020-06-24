package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

//Create a new type deck

type deck []string

func newDeck() deck { //deck is the type for the returning function value
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clutch"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits { //underscore is used to replace the unused index to avoid golang error
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() { //a receiver function //any variable of type deck now has access to the print method.
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) { //function signature
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",") //convers a deck of string to string

}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)

	//error handling
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1) //exit(0) means exit without an error. exit(1) means exit with error.

	}

	s := strings.Split(string(bs), ",") //turn byte slice to []string
	//turn string s to deck
	return deck(s)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}

}
