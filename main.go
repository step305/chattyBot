package main

import (
	"errors"
	"fmt"
)

const GreetingTemplate string = "Hello! My name is %s.\nI was created in %d.\n"
const BotName string = "Chatty"
const BotBirthYear int = 2022
const askUserNameTemplate string = "Please, remind me your name.\n"
const confirmUserNameTemplate string = "What a great name you have, %s!\n"
const guessAgeStartTemplate string = "Let me guess your age.\nEnter remainders of dividing your age by 3, 5 and 7.\n"
const guessAgeEndTemplate string = "Your age is %d; that's a good time to start programming!\n"
const countGameStartTemplate string = "Now I will prove to you that I can count to any number you want.\n"
const countGameEndTemplate string = "Completed, have a nice day!\n"
const quizGameStartTemplate string = "Let's test your programming knowledge.\n"
const quizGameTryAgainTemplate string = "Please, try again.\n"
const quizGameEndTemplate string = "Congratulations, have a nice day!\n"
const errorInvalidInput string = "[Error] invalid input"

type Quiz struct {
	question string
	choices  []string
	answer   int
}

var quiz Quiz = Quiz{
	question: "Why do we use methods?",
	choices: []string{
		"To repeat a statement multiple times.",
		"To decompose a program into several small subroutines.",
		"To determine the execution time of a program.",
		"To interrupt the execution of a program.",
	},
	answer: 2,
}

func greeting() {
	fmt.Printf(GreetingTemplate, BotName, BotBirthYear)
}

func askUserName() (string, error) {
	var name string
	fmt.Print(askUserNameTemplate)
	_, err := fmt.Scanln(&name)
	if err != nil {
		return "", errors.New(errorInvalidInput)
	}
	fmt.Printf(confirmUserNameTemplate, name)
	return name, nil
}

func guessAgeGame() error {
	var remainders [3]int
	fmt.Print(guessAgeStartTemplate)
	for i := range remainders {
		_, err := fmt.Scanln(&remainders[i])
		if err != nil {
			return errors.New("[Error] cannot input remainders")
		}
	}
	guessAge := (remainders[0]*70 + remainders[1]*21 + remainders[2]*15) % 105
	fmt.Printf(guessAgeEndTemplate, guessAge)
	return nil
}

func countGame() error {
	var limit int
	fmt.Print(countGameStartTemplate)
	_, err := fmt.Scanln(&limit)
	if err != nil {
		return errors.New("[Error] counter limit input error")
	}
	for i := 0; i <= limit; i++ {
		fmt.Printf("%d !\n", i)
	}
	fmt.Print(countGameEndTemplate)
	return nil
}

func quizGame(newQuiz Quiz) error {
	var guess int = 0
	fmt.Print(quizGameStartTemplate)
	fmt.Println(newQuiz.question)
	for i, choice := range newQuiz.choices {
		fmt.Printf("%d. %s\n", i+1, choice)
	}
	for {
		_, err := fmt.Scanln(&guess)
		if err != nil {
			return errors.New(errorInvalidInput)
		}
		if guess == quiz.answer {
			break
		} else {
			fmt.Print(quizGameTryAgainTemplate)
		}
	}
	fmt.Print(quizGameEndTemplate)
	return nil
}

func main() {
	greeting()
	_, err := askUserName()
	if err != nil {
		fmt.Println(err)
	}
	err = guessAgeGame()
	if err != nil {
		fmt.Println(err)
	}
	err = countGame()
	if err != nil {
		fmt.Println(err)
	}
	err = quizGame(quiz)
	if err != nil {
		fmt.Println(err)
	}
}
