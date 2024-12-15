package day7

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

//Generated by AoC helper tool 🎄
//In order to get started, extend main.go so it discovers and runs this file as well

type Day7Solver struct{}

type Equation struct {
	result    int
	parts     []int
	operators []rune
}

func generateCombinations(operators []rune, n int, current string, result *[]string) {
	if len(current) == n-1 {
		*result = append(*result, current)
		return
	}

	for _, op := range operators {
		generateCombinations(operators, n, current+string(op), result)
	}
}

func evaluateExpression(numbers []int, operators string) int {
	result := numbers[0]
	for i, op := range operators {
		if op == '+' {
			result += numbers[i+1]
		} else if op == '*' {
			result *= numbers[i+1]
		} else if op == '|' {
			conc := strconv.Itoa(result) + strconv.Itoa(numbers[i+1])
			parsed, _ := strconv.ParseInt(conc, 10, 64)
			result = int(parsed)
		}
	}
	return result
}

func (e *Equation) compute() bool {
	operators := e.operators
	var combinations []string

	generateCombinations(operators, len(e.parts), "", &combinations)

	wasGood := false
	for _, ops := range combinations {
		result := evaluateExpression(e.parts, ops)
		if result == e.result {
			wasGood = true
		}
	}
	return wasGood
}

func (d Day7Solver) Solve(filename string, answerChan chan int, doneChan chan bool, errorChan chan error) {
	defer close(answerChan)
	defer close(errorChan)
	file, err := os.Open(filename)

	if err != nil {
		errorChan <- err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	equations := []Equation{}

	for scanner.Scan() {
		line := scanner.Text()
		splits := strings.Split(line, ":")
		this := Equation{}

		result, _ := strconv.ParseInt(splits[0], 10, 64)
		this.result = int(result)

		parts := strings.Split(splits[1][1:], " ")

		for _, part := range parts {
			parsed, _ := strconv.ParseInt(part, 10, 64)
			this.parts = append(this.parts, int(parsed))
		}

		this.operators = append(this.operators, '+')
		this.operators = append(this.operators, '*')
		equations = append(equations, this)
	}

	countDoable := 0
	for _, equation := range equations {
		if equation.compute() {
			countDoable += equation.result
		}
	}

	// round 2

	countR2Doable := 0
	for _, equation := range equations {
		equation.operators = append(equation.operators, '|')
		if equation.compute() {
			countR2Doable += equation.result
		}
	}

	answerChan <- countDoable
	answerChan <- countR2Doable

	doneChan <- true
}
