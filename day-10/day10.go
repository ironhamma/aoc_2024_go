package day10

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Day10Solver struct{}

type Node struct {
	id         string
	value      int
	neighbours []*Node
}

type NodeMap struct {
	nodes map[string]*Node
}

func (nmap *NodeMap) parseInput(input [][]string) {
	for i, row := range input {
		for j := range row {
			id := fmt.Sprintf("%d,%d", i, j)
			if _, exists := nmap.nodes[id]; !exists {
				value, _ := strconv.Atoi(row[j])
				nmap.nodes[id] = &Node{
					id:    id,
					value: value,
				}
			}

			current := nmap.nodes[id]
			neighbors := [][]int{
				{i - 1, j},
				{i + 1, j},
				{i, j - 1},
				{i, j + 1},
			}
			for _, n := range neighbors {
				ni, nj := n[0], n[1]
				if ni >= 0 && ni < len(input) && nj >= 0 && nj < len(row) {
					neighborID := fmt.Sprintf("%d,%d", ni, nj)
					if neighbor, exists := nmap.nodes[neighborID]; exists {
						current.neighbours = append(current.neighbours, neighbor)
					} else {
						neighborValue, _ := strconv.Atoi(input[ni][nj])
						neighbor := &Node{
							id:    neighborID,
							value: neighborValue,
						}
						nmap.nodes[neighborID] = neighbor
						current.neighbours = append(current.neighbours, neighbor)
					}
				}
			}
		}
	}
}

func (nmap NodeMap) bfs(startNode *Node, useVisited bool) int {
	visited := make(map[string]bool)
	queue := []*Node{startNode}
	count := 0

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if visited[current.id] && useVisited {
			continue
		}

		if useVisited {
			visited[current.id] = true
		}

		if current.value == 9 {
			count++
			continue
		}

		if useVisited {
			for _, neighbor := range current.neighbours {
				if !visited[neighbor.id] && neighbor.value == current.value+1 {
					queue = append(queue, neighbor)
				}
			}
		} else {
			for _, neighbor := range current.neighbours {
				if neighbor.value == current.value+1 {
					queue = append(queue, neighbor)
				}
			}
		}
	}

	return count
}

func (nmap NodeMap) CountGood(useVisited bool) int {
	totalScore := 0
	for _, node := range nmap.nodes {
		if node.value == 0 {
			totalScore += nmap.bfs(node, useVisited)
		}
	}
	return totalScore
}

func (d Day10Solver) Solve(filename string, answerChan chan int, doneChan chan bool, errorChan chan error) {
	defer close(answerChan)
	defer close(errorChan)

	file, err := os.Open(filename)
	if err != nil {
		errorChan <- err
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	nmap := NodeMap{nodes: make(map[string]*Node)}
	input := [][]string{}

	for scanner.Scan() {
		line := scanner.Text()
		row := []string{}
		for _, c := range line {
			row = append(row, string(c))
		}
		input = append(input, row)
	}

	nmap.parseInput(input)

	res1 := nmap.CountGood(true)
	res2 := nmap.CountGood(false)
	answerChan <- res1
	answerChan <- res2

	doneChan <- true
}
