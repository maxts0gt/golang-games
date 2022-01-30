package main

import (
	"fmt"
	"math/rand"
	"time"
	"unicode"
)

var dictionary = []string{
	"Zombie",
	"Gopher",
	"USA",
	"South Korea",
	"United States of America",
	"Nazism",
	"Apple",
	"Book",
	"Computer",
}

func main() {
	rand.Seed(time.Now().UnixNano())
	randomWord := getRandomWord()
	randomWord = "United States of America"
	guessedLetters := initializeGuessedLetters(randomWord)
	printGameState(randomWord, guessedLetters)
	guessedLetters['s'] = true
	printGameState(randomWord, guessedLetters)

}

func initializeGuessedLetters(randomWord string) map[rune]bool {
	guessedLetters := map[rune]bool{}
	guessedLetters[unicode.ToLower(rune(randomWord[0]))] = true
	guessedLetters[unicode.ToLower(rune(randomWord[len(randomWord)-1]))] = true
	return guessedLetters
}

func getRandomWord() string {
	wordToGuess := dictionary[rand.Intn(len(dictionary))]
	return wordToGuess
}

func printGameState(randomWord string, guessedLetters map[rune]bool) {
	for _, ch := range randomWord {
		if ch == ' ' {
			fmt.Print(" ")
		} else if guessedLetters[unicode.ToLower(ch)] {
			fmt.Printf("%c", ch)
		} else {
			fmt.Printf("_")
		}

		fmt.Print(" ")
	}
}
