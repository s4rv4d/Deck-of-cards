package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

//create a new type of 'deck'
//which is a slice of strings

type deck []string

//create a new slice with card names and return it(its of type deck)
func newDeck() deck {
	//initialize a slice of type deck
	cards := deck{}

	//defining the slice of suits which is a string
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	//defining the slice of values which is a string. later on we use both value and suit to make a card
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	//now joining them and appending them into cards
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	//returning the deck
	return cards
}

//create a function to loop through the cards

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

//create a function to deal n number of cards with arguments consisiting of deck slick and n
//we need to be returning two slices as final output of type deck

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

//helper function to take a deck and return the string representation
func (d deck) toString() string {
	singleStringRepresentation := strings.Join([]string(d), ",")
	return singleStringRepresentation
}

//create a function to convert to save to file
func (d deck) saveToFile(filename string) error {
	//0666 its a permission basically telling anyone can read or write
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

//retrieving a new deck from the file
func newDeckFromFile(filename string) deck {
	nd, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		//to exit the project we use the package os to exit with a error code which is any number otehr than 0
		os.Exit(1)
	}
	//order for back conversion is []byte -> string -> []string == []deck
	s := strings.Split(string(nd), ",")
	d := deck(s)
	return d
}

//shuffling procedure
func (d deck) shuffle() {
	for i := range d {
		//generating a random int64 value through time packages unixnano
		it64 := time.Now().UnixNano()
		//generating random new seed
		s := rand.NewSource(it64)
		//creating new rand
		r := rand.New(s)

		//generating random number using new seed
		np := r.Intn(len(d) - 1)

		//swapping values
		d[i], d[np] = d[np], d[i]
	}
	//this method or randomize the seed
	// d.saveToFile("my_cards")

}
