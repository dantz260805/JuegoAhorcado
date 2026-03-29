// Go program for the hangman game
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// List of words to choose from
var words = []string{
	"hangman",
	"programming",
	"golang",
	"development",
	"algorithm",
}

// Function to choose a random word
func chooseWord() string {
	rand.Seed(time.Now().UnixNano())
	return words[rand.Intn(len(words))]
}

// Main function
func main() {
	word := chooseWord()
	fmt.Printf("The word has %d letters.\n", len(word))
	// Implement the rest of the hangman game logic here
}