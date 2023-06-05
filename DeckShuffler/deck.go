package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// is like a class for OOP --> called 'type declaration'
type deck []string

// is like a method for the class in OOP --> called 'receivers'
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Swords", "Axes", "Shields", "Spears"}
	cardNumbers := []string{"Ace", "Two", "Three", "Four", "Five", "Six"}

	for _, suit := range cardSuits {
		for _, number := range cardNumbers {
			cards = append(cards, number+" of "+suit)
		}
	}
	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) ToString() string {
	return strings.Join([]string(d), ",")

}

// 0666: everyone has access, read, write, delete permissions
func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.ToString()), 0666)
}

func newDeckFromFile(filename string) deck {
	byte_string, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	split_string := strings.Split(string(byte_string), ",")
	return deck(split_string)
}

func (d deck) shuffleOnce(random *rand.Rand) {
	for i := range d {
		j := random.Intn(len(d))
		d[i], d[j] = d[j], d[i]
	}
}

func (d deck) shuffleDeck() deck {
	// Create a private random number generator with a specific seed
	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Perform multiple shuffles
	for i := 0; i < 5; i++ {
		d.shuffleOnce(random)
	}

	return d
}
