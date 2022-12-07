package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	timeLimit := flag.Int("timer", 30, "seconds on the timer (Default: 30)")
	flag.Parse()

	file, err := os.Open("./problems.csv")
	if err != nil {
		log.Fatal("Unable to read input file: ")
	}
	defer file.Close()
	csv_reader := csv.NewReader(file)
	var q_num int
	var totalQs int
	var correctAs int

	fmt.Println("---- Cris' Quiz Game ----")
	fmt.Println(*timeLimit)
	fmt.Print("Begin?")
	fmt.Scanln()

	*timeLimit--
	fmt.Println(*timeLimit)

	for range time.Tick(1 * time.Second) {
		*timeLimit--

		if *timeLimit <= 0 {
			fmt.Println("Time limit reached!")
			os.Exit(1)
		}
	}

	for {
		question_answer, err := csv_reader.Read()

		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal("Unable to read input file")
		}

		question := question_answer[0]
		answer, err := strconv.Atoi(question_answer[1])

		if err != nil {
			log.Fatal("Could not cast str to int")
		}
		q_num++

		fmt.Printf("%d) %v\n", q_num, question)
		fmt.Print("Answer:")
		var user_input string
		fmt.Scanln(&user_input)
		user_answer, err := strconv.Atoi(user_input)

		if err != nil {
			log.Fatal("Error with answer")
		}

		if user_answer == answer {
			correctAs++
		}
		totalQs++
	}
	fmt.Println("Score:")
	fmt.Printf("%d/%d", correctAs, totalQs)
}
