package day13

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"
)

//Generated by AoC helper tool 🎄
//In order to get started, extend main.go so it discovers and runs this file as well

type Day13Solver struct{}

type coordinate struct {
	x int
	y int
}

type machine struct {
	button_a coordinate
	button_b coordinate
	prize    coordinate
}

func calcMachine(machine machine) (int, int) {
	det := machine.button_a.x*machine.button_b.y - machine.button_b.x*machine.button_a.y
	detX := machine.prize.x*machine.button_b.y - machine.prize.y*machine.button_b.x
	detY := machine.prize.y*machine.button_a.x - machine.prize.x*machine.button_a.y

	if detX%det != 0 || detY%det != 0 {
		return 0, 0
	}

	return detX / det, detY / det
}

func (d Day13Solver) Solve(filename string, answerChan chan int, doneChan chan bool, errorChan chan error) {
	defer close(answerChan)
	defer close(errorChan)
	file, err := os.Open(filename)

	if err != nil {
		errorChan <- err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	reButton := regexp.MustCompile(`Button (A|B): X[+=](\d+), Y[+=](\d+)`)
	rePrize := regexp.MustCompile(`Prize: X=(\d+), Y=(\d+)`)

	var machines []machine
	var currentMachine machine

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if matches := reButton.FindStringSubmatch(line); matches != nil {
			x, _ := strconv.Atoi(matches[2])
			y, _ := strconv.Atoi(matches[3])
			coord := coordinate{x: x, y: y}
			if matches[1] == "A" {
				currentMachine.button_a = coord
			} else if matches[1] == "B" {
				currentMachine.button_b = coord
			}
			continue
		}

		if matches := rePrize.FindStringSubmatch(line); matches != nil {
			x, _ := strconv.Atoi(matches[1])
			y, _ := strconv.Atoi(matches[2])
			currentMachine.prize = coordinate{x: x, y: y}

			machines = append(machines, currentMachine)
			currentMachine = machine{}
		}
	}

	res1 := 0
	for _, m := range machines {
		a, b := calcMachine(m)
		res1 += a * 3
		res1 += b
	}

	res2 := 0
	for _, m := range machines {
		newMachine := machine{
			button_a: m.button_a,
			button_b: m.button_b,
			prize: coordinate{
				x: m.prize.x + 10000000000000,
				y: m.prize.y + 10000000000000,
			},
		}
		a, b := calcMachine(newMachine)
		res2 += a * 3
		res2 += b
	}

	answerChan <- res1
	answerChan <- res2

	doneChan <- true
}
