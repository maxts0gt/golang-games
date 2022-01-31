package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
	"unicode"
)

var inputReader = bufio.NewReader(os.Stdin)
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
	for {
		printGameState(randomWord, guessedLetters, hangmanState)
		input := readInput()
		if len(input) != 1 {
			fmt.Println("Invalid input. Please use letters only.")
			continue
		}
		letter := rune(input[0])
		if isCorrectWord(randomWord, letter) {
			guessedLetters[letter] = true
		} else {
			hangmanState++
		}
	}

}

func getRandomWord() string {
	wordToGuess := dictionary[rand.Intn(len(dictionary))]
	return wordToGuess
}

func initializeGuessedLetters(randomWord string) map[rune]bool {
	guessedLetters := map[rune]bool{}
	guessedLetters[unicode.ToLower(rune(randomWord[0]))] = true
	guessedLetters[unicode.ToLower(rune(randomWord[len(randomWord)-1]))] = true
	return guessedLetters
}

func printGameState(randomWord string, guessedLetters map[rune]bool, hangmanState int) {
	fmt.Println(getGuessingProgress(randomWord, guessedLetters))
	fmt.Print(" ")
	fmt.Println(getHangmanParts(hangmanState))
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

func readInput() string {
	fmt.Print("> ")
	input, err := inputReader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	return strings.TrimSpace(input)
}

func isCorrectWord(randomWord string, letter rune) bool {
	return strings.ContainsRune(randomWord, letter)
}
