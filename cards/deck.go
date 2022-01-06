package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// create a new type of var "deck"
// which is a slice of string

type deck []string

func (d deck) print() {
	for i, one_card := range d {
		fmt.Println(i, one_card)
	}
}

func newDeck() deck {
	card := deck{}
	cardSuits := []string{"Hearts", "Spades", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suits := range cardSuits {
		for _, values := range cardValues {
			card = append(card, values+" of "+suits)
		}
	}

	return card
}

func deal(d deck, handsize int) (deck, deck) {

	return d[:handsize], d[handsize:]
}

func (d deck) deckToString() string {

	return strings.Join([]string(d), ", ")
}

func (d deck) saveToFile(filename string) error {

	return ioutil.WriteFile(filename, []byte(d.deckToString()), 0666)
}

func deckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	str_deck := string(bs)

	str_slice := strings.Split(str_deck, ",")
	return deck(str_slice)
}

func (d deck) shuffle() {

	card_len := len(d)
	rand.Seed(time.Now().UnixNano())
	var rand_int int
	for i := range d {
		rand_int = rand.Intn(card_len - 1)
		d[rand_int], d[i] = d[i], d[rand_int]
	}
}
