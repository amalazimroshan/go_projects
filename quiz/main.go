package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type Question struct {
	problem string
	answer  int
}

func main() {

	numberOfQuestions := flag.Int("questions", -1, "number of questions")
	timeLimit := flag.Int("limit", 30, "time limit for one question")
	flag.Parse()

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

	if *numberOfQuestions < 0 {
		*numberOfQuestions = len(questions)
	}

	fmt.Println("Press enter to start the quizz")
	fmt.Scanln()

	score := 0
	timeDuration := time.Duration(*timeLimit**numberOfQuestions) * time.Second
	timer := time.NewTimer(timeDuration)
	done := make(chan bool)

	fmt.Printf("you have a total of %v to complete the quiz\n", timeDuration)
	go func() {
		for i := 0; i < *numberOfQuestions; i++ {
			select {
			case <-done:
				return
			default:
				fmt.Printf("Solve: %s =", questions[i].problem)
				var userAnswer int
				fmt.Scanln(&userAnswer)
				if userAnswer == questions[i].answer {
					score++
				} else {
					fmt.Println("incorrect! the correct answer was:", questions[i].answer)
				}
			}
		}
		done <- true
	}()

	select {
	case <-timer.C:
		fmt.Println("\n times up!")
	case <-done:
		fmt.Println("\nquiz completed")
	}
	fmt.Printf("you scored %d out of %d", score, *numberOfQuestions)
}
