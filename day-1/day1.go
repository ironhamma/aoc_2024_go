package day1

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Day1Solver struct{}

func (d Day1Solver) Solve(filename string, answerChan chan int, doneChan chan bool, errorChan chan error) {
	defer close(answerChan)
	defer close(errorChan)
	defer close(doneChan)

	file, err := os.Open(filename)

	if err != nil {
		errorChan <- err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	leftData := []int{}
	rightData := []int{}

	for scanner.Scan() {
		newPair := strings.Split(scanner.Text(), "   ")
		num, err := strconv.ParseInt(newPair[0], 10, 64)
		if err != nil {
			errorChan <- err
		}
		leftData = append(leftData, int(num))

		numRight, err := strconv.ParseInt(newPair[1], 10, 64)
		if err != nil {
			errorChan <- err
		}
		rightData = append(rightData, int(numRight))
	}

	if err := scanner.Err(); err != nil {
		errorChan <- err
	}

	sort.Ints(leftData)
	sort.Ints(rightData)

	sum := 0

	occurance := make(map[int]int)

	for index, v := range leftData {
		_, exists := occurance[rightData[index]]
		if exists {
			occurance[rightData[index]] += 1
		} else {
			occurance[rightData[index]] = 1
		}

		product := v - rightData[index]
		if product < 0 {
			product = product * -1
		}
		sum += product
	}
	answerChan <- sum

	similarity := 0

	for _, v := range leftData {
		_, exists := occurance[v]
		if exists {
			similarity += occurance[v] * v
		} else {
			similarity += 0
		}
	}

	answerChan <- similarity

	time.Sleep(time.Second * 4)
	doneChan <- true
}
