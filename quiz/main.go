package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
)

type Question struct {
	problem string
	answer  int
}

func main() {

	args := os.Args[1:]
	var numberOfQuestions int

	file, err := os.Open("problems.csv")
	if err != nil {
		log.Fatal("Error opening file: ", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal("Error reading CSV: ", err)
	}

	var questions []Question
	for _, record := range records {
		answer, err := strconv.Atoi(record[1])
		if err != nil {
			log.Fatal("Error parsing answer: ", err)
		}
		questions = append(questions, Question{problem: record[0], answer: answer})
	}

	rand.Shuffle(len(questions), func(i, j int) {
		questions[i], questions[j] = questions[j], questions[i]
	})

	if len(args) > 0 {
		num, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("invalid argument Switching to default state")
			numberOfQuestions = len(questions)
		} else {
			numberOfQuestions = num
		}
	} else {
		numberOfQuestions = len(questions)
	}

	score := 0
	for i := 0; i < numberOfQuestions; i++ {
		fmt.Printf("Solve: %s =", questions[i].problem)
		var userAnswer int
		fmt.Scanln(&userAnswer)
		if userAnswer == questions[i].answer {
			score++
		} else {
			fmt.Println("incorrect! the correct answer was:", questions[i].answer)
		}
	}

	fmt.Printf("you scored %d out of %d", score, numberOfQuestions)
}
