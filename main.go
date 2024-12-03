package main

import (
	day1 "aoc/day-1"
	day2 "aoc/day-2"
	day3 "aoc/day-3"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Solver interface {
	Solve(string, chan int, chan bool, chan error)
}

func main() {
	fmt.Printf("\n--------------------------------------------\n")
	fmt.Printf("\n❄️  Hi there! Welcome to Advent of Code solving in Go! 🎅\n\n")
	fmt.Printf("--------------------------------------------\n\n")
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("🚀 Please provide the day you want to run: ")

	scanner.Scan()
	day := scanner.Text()

	dayNum, err := strconv.ParseInt(day, 10, 64)
	if err != nil {
		log.Fatalf("❌ Some error happened during parsing the day number: %v", err)
	}

	fmt.Printf("📆 You chose day: %v\n", dayNum)

	var taskSolver Solver

	switch dayNum {
	case 1:
		dayOne := day1.Day1Solver{}
		taskSolver = dayOne
	case 2:
		dayTwo := day2.Day2Solver{}
		taskSolver = dayTwo
	case 3:
		dayThree := day3.Day3Solver{}
		taskSolver = dayThree
	default:
		if dayNum < 25 && dayNum > 0 {
			CreateDayFiles(day)
			return
		}
		fmt.Println("🟨 This day was not yet implemented! Bummer...")
		return
	}

	fmt.Print("🚧 Do you want to run in test mode? Type 1 to run in test mode. ")

	scanner.Scan()
	isTest := scanner.Text()

	testBool, err := strconv.ParseInt(isTest, 10, 64)
	if err != nil {
		testBool = 0
	}

	fileName := "./day-" + day

	if testBool == 1 {
		fileName += "/test.txt"
	} else {
		fileName += "/input.txt"
	}

	fmt.Printf("\n------------------------------------\n")
	fmt.Printf("🎄 Advent of Code 2024: Day %v 🎄\n", dayNum)
	fmt.Printf("------------------------------------\n\n")

	answerChan := make(chan int, 2)
	errorChan := make(chan error)
	doneChan := make(chan bool)

	switch dayNum {
	case 1, 2, 3:
		go taskSolver.Solve(fileName, answerChan, doneChan, errorChan)
	default:
		fmt.Println("🟨 This day was not yet implemented! Bummer...")
		return
	}

	taskNum := 1

	isDone := false

	for {
		if isDone {
			break
		}
		select {
		case answer, ok := <-answerChan:
			if !ok {
				log.Fatalf("❌ Channel closed, f...")
				return
			}
			fmt.Printf("⛄ Task %v answer: %v ✅\n", taskNum, answer)
			taskNum += 1
		case err, ok := <-errorChan:
			if !ok {
				log.Fatalf("❌ Channel closed, f...")
				return
			}
			log.Fatalf("❌ An error appeared during task running: %v", err)
		case status, ok := <-doneChan:
			if !ok {
				log.Fatalf("❌ Channel closed, f...")
				return
			}
			isDone = status
		}
	}

	fmt.Printf("\n--------------------------------------------\n")
	fmt.Println("🎁 Don't forget to submit your results! 🎁")
	fmt.Printf("--------------------------------------------\n\n")
}
