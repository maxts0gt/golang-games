package main

import (
	"fmt"
	"io/ioutil"
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
	guessedLetters := initializeGuessedLetters(randomWord)
	hangmanState := 0
	printGameState(randomWord, guessedLetters, hangmanState)

}

func printGameState(randomWord string, guessedLetters map[rune]bool, hangmanState int) {
	fmt.Println(getGuessingProgress(randomWord, guessedLetters))
	fmt.Print(" ")
	fmt.Println(getHangmanParts(hangmanState))
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

func getGuessingProgress(randomWord string, guessedLetters map[rune]bool) string {

	result := ""

	for _, ch := range randomWord {
		if ch == ' ' {
			result += " "
		} else if guessedLetters[unicode.ToLower(ch)] {
			result += fmt.Sprintf("%c", ch)
		} else {
			result += "_"
		}

		result += " "

	}
	return result
}

func getHangmanParts(hangmanState int) string {
	data, err := ioutil.ReadFile(fmt.Sprintf("states/hangman%d", hangmanState))
	if err != nil {
		panic(err)
	}
	return string(data)
}
