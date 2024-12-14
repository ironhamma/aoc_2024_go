package day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	LEAST = 1
	MOST  = 3
)

type Day2Solver struct{}

func checkRow(row []int) bool {
	isValid := false
	direction := 0

	for j := 0; j < len(row)-1; j++ {
		current := row[j]
		next := row[j+1]
		sub := current - next
		isGood := abs(sub) <= MOST && abs(sub) >= LEAST

		if direction == 0 && sub > 0 {
			direction = 1
		}

		if direction == 0 && sub < 0 {
			direction = -1
		}

		if (direction == -1 && isGood && sub < 0) || (direction == 1 && isGood && sub > 0) {
			isValid = true
			continue
		}

		isValid = false
		break
	}

	return isValid
}

func abs(value int) int {
	if value < 0 {
		return -value
	}
	return value
}

func (d Day2Solver) Solve(filename string, answerChan chan int, doneChan chan bool, errorChan chan error) {
	start := time.Now()

	defer close(answerChan)
	defer close(errorChan)
	defer close(doneChan)

	file, err := os.Open(filename)
	if err != nil {
		errorChan <- err
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var assets [][]int

	for scanner.Scan() {
		line := scanner.Text()
		row := strings.Fields(line)
		var nums []int
		for _, num := range row {
			val, err := strconv.Atoi(num)
			if err != nil {
				errorChan <- err
				return
			}
			nums = append(nums, val)
		}
		assets = append(assets, nums)
	}

	if err := scanner.Err(); err != nil {
		errorChan <- err
		return
	}

	countValidAssets := 0
	countFixable := 0

	for _, row := range assets {
		isValid := checkRow(row)

		if isValid {
			countValidAssets++
			countFixable++
			continue
		}

		foundValid := false
		for j := 0; j < len(row); j++ {
			rcp := append([]int{}, row[:j]...)
			rcp = append(rcp, row[j+1:]...)

			if checkRow(rcp) {
				foundValid = true
				break
			}
		}

		if foundValid {
			countFixable++
		}
	}

	answerChan <- countValidAssets
	answerChan <- countFixable

	end := time.Now()
	elapsed := end.Sub(start)
	fmt.Printf("\n\n⏱️ Execution took %v time! ⏱️\n\n", elapsed)

	time.Sleep(time.Second)
	doneChan <- true
}
