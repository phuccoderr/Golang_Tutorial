package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace","Two","Three","Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards,value + " of " + suit)
		}
	}

	return cards
}

func deal(this deck, handSize int) (deck, deck) {
	return this[:handSize],this[handSize:]
}

func (this deck) print() {
	for i, card := range this {
		fmt.Println(i, card)
	}
}

func (this deck) toString() string {
	return strings.Join([]string(this), ",")
}

func (this deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename,[]byte(this.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	resuslt,err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(0);
	}

	s := strings.Split(string(resuslt), ",")// Ace of Spades, Two of Spades...
	return deck(s)
}

func (this deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range this {
		newPosition := r.Intn(len(this) - 1)

		this[i], this[newPosition] = this[newPosition], this[i]
	}
}