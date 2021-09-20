package main

import (
	"flag"

	utils "github.com/Zeddling/gophercises/quiz_game/utils"
)

func main() {
	var path string
	flag.StringVar(&path, "f", "problems.csv", "Defaulted to problems.csv")
	timer := flag.Int("t", 30, "the time limit for the quiz in seconds")
	println(path)

	questions, _ := utils.NewQuestions(path)
	questions.Test(*timer)
}
