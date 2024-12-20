package main

import (
	"fmt"
	"os"
)

func CreateDayFiles(day string) {
	fmt.Println("🔍 Could not find files for day " + day + " task.")
	dirName := "./day-" + day
	filename := dirName + "/day" + day + ".go"

	testFileName := dirName + "/test.txt"
	inputFileName := dirName + "/input.txt"

	// Make directory
	err := os.MkdirAll(dirName, 0755)
	if err != nil {
		fmt.Println("❌ Error creating directory:", err)
	}

	// Make test.txt
	testFile, err := os.Create(testFileName)
	if err != nil {
		fmt.Println("❌ Error creating test file:", err)
		return
	}
	defer testFile.Close()

	// Make input.txt
	inputFile, err := os.Create(inputFileName)
	if err != nil {
		fmt.Println("❌ Error creating input file:", err)
		return
	}
	defer inputFile.Close()

	// Make go file
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("❌ Error creating file:", err)
		return
	}

	defer file.Close()

	// Write in go code
	_, err = file.WriteString("package day" + day + "\n\n//Generated by AoC helper tool 🎄\n//In order to get started, extend main.go so it discovers and runs this file as well\n\ntype Day" + day + "Solver struct{}\n\n func (d Day" + day + "Solver) Solve(filename string, answerChan chan int, doneChan chan bool, errorChan chan error){\ndefer close(answerChan)\n defer close(errorChan)\nfile, err := os.Open(filename)\n\n if err != nil {\nerrorChan <- err\n}\n\n defer file.Close()\n\n//scanner := bufio.NewScanner(file)\n\nanswerChan <- 0\nanswerChan <- 0\n\ndoneChan <- true\n}\n")
	if err != nil {
		fmt.Println("❌ Error writing file:", err)
		return
	}

	fmt.Printf("💫💫 Task directory created: %v! 💫💫\n\n", filename)
}
