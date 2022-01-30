package main

import (
	"fmt"
	"math/rand"
	"time"
)

var dictionary = []string{
	"Zombie",
	"Gopher",
	"USA",
	"South Korea",
	"Nazism",
	"Apple",
	"Book",
	"Computer",
}

func main() {
	rand.Seed(time.Now().UnixNano())
	randomWord := getRandomWord()
	fmt.Println(randomWord)
	// Printing game state
	// print word  u r guessing
	// print hangman state

}

func getRandomWord() string {
	wordToGuess := dictionary[rand.Intn(len(dictionary))]
	return wordToGuess
}
