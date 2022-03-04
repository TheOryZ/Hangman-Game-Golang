package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
)

var words = []string{
	"TURKEY",
	"HOLLAND",
	"RUSSIA",
	"GERMANY",
	"GEORGIA",
	"SPAIN",
	"AZERBAIJAN",
	"JAPAN",
}

func main() {
	color.Red("Welcome to the Hangman game.")
	//Random word generation
	rand.Seed(time.Now().UnixNano())
	word := words[rand.Intn((len(words)-1)-0+1)]
	var guessWords string
	isTurn := true
	s := 2
	//Create a string as long as word
	resultWord := make([]string, len(word), len(word))
	for k := range resultWord {
		resultWord[k] = "-"
	}
	for isTurn {
		fmt.Print("Word : ", resultWord)
		fmt.Println(" Word Length :", len(word))
		//Get currency state
		data, err := ioutil.ReadFile("states/hangman" + strconv.Itoa(s))
		if err != nil {
			fmt.Println("File reading error", err)
		}
		fmt.Printf("Status of game (state %v):\n", s)
		color.Green(string(data))
		//User guessing
		inputReader := bufio.NewReader(os.Stdin)
		text, _ := inputReader.ReadString('\n')
		text = strings.TrimSpace(text)
		guessWords += strings.ToUpper(text) + "/"
		guessStatus := false
		if len(text) > 1 {
			if strings.ToUpper(text) != word {
				color.Yellow("Wrong !!")
				WrongAnswer(&s)
			} else {
				color.Green("You Win!! Congratulations!!")
				isTurn = false
				guessStatus = true
			}
		} else {
			upperText := strings.ToUpper(text)
			if strings.Contains(word, upperText) {
				for k, v := range word {
					if string(v) == upperText {
						resultWord[k] = upperText
						guessStatus = true
					}
				}
			}
			if !guessStatus {
				WrongAnswer(&s)
			}
		}
		//Set turn value
		if s >= 9 {
			isTurn = false
			color.Red("You lose!!")
		}
		var _resultWord string
		for k := range resultWord {
			_resultWord += resultWord[k]
		}
		if word == _resultWord {
			isTurn = false
			color.Green("You Win!! Congratulations!!")
		}
		fmt.Println("Your Guesses : ", guessWords)
	}
	color.Yellow(strings.Repeat("*", 25))
	//Get currency state again
	data, err := ioutil.ReadFile("states/hangman" + strconv.Itoa(s))
	if err != nil {
		fmt.Println("File reading error", err)
	}
	fmt.Printf("Status of game (state %v):\n", s)
	if s < 9 {
		color.Green(string(data))
	} else {
		color.Red(string(data))
	}

}

func WrongAnswer(s *int) {
	*s++
}
