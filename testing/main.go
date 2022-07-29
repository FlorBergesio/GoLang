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

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts"}
	cardValues := []string{"Ace", "Two", "Three"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

func newDeckFromFile(fileName string) deck {
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	randomInt := time.Now().UnixNano()
	source := rand.NewSource(randomInt)
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

func main() {
	cards := newDeck()

	fmt.Println(cards.toString())

	cards.saveToFile("my_file")

	fmt.Println("-------")
	fmt.Println("File Saved Containing deck")
	fmt.Println("-------")
	fmt.Println("Reading data from file")
	fmt.Println("-------")

	cardsRead := newDeckFromFile("my_file")
	fmt.Println(cardsRead.toString())

	fmt.Println("-------")
	fmt.Println("Shuffling Cards")
	fmt.Println("-------")

	cardsRead.shuffle()
	fmt.Println(cardsRead.toString())
}
