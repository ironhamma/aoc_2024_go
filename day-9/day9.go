package day9

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
)

//Generated by AoC helper tool 🎄
//In order to get started, extend main.go so it discovers and runs this file as well

type Day9Solver struct{}

type File struct {
	id    int
	count int
	start int
	len   int
}

type Memory struct {
	files []File
	empty []File
}

func (m Memory) CountCheckSum() int {
	sum := 0
	for _, file := range m.files {
		for i := 0; i < file.count; i++ {
			sum += file.id * (file.start + (i * file.len))
		}
	}
	return sum
}

func (m Memory) Print() {
	sortedCopy := make([]File, len(m.files))
	copy(sortedCopy, m.files)

	slices.SortFunc(sortedCopy, func(a, b File) int {
		return a.start - b.start
	})

	for _, file := range sortedCopy {
		for i := 0; i < file.count; i++ {
			fmt.Print(file.id)
		}
	}
	fmt.Println()
}

func (d Day9Solver) Solve(filename string, answerChan chan int, doneChan chan bool, errorChan chan error) {
	defer close(answerChan)
	defer close(errorChan)
	file, err := os.Open(filename)

	if err != nil {
		errorChan <- err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()

	idCounter := 0
	memory := Memory{}
	indexCounter := 0
	for index, char := range line {
		num, _ := strconv.Atoi(string(char))
		if num == 0 {
			continue
		}
		if index%2 == 0 {
			idAsString := strconv.Itoa(idCounter)
			file := File{id: idCounter, count: num, start: indexCounter, len: len(idAsString)}
			memory.files = append(memory.files, file)
			idCounter++
			indexCounter += num
		} else {
			empty := File{id: -1, count: num, start: indexCounter, len: 1}
			memory.empty = append(memory.empty, empty)
			indexCounter += num
		}
	}

	sortedCopy := make([]File, len(memory.files))
	copy(sortedCopy, memory.files)

	// Solve part 2
	slices.SortFunc(sortedCopy, func(a, b File) int {
		return b.id - a.id
	})

	for _, file := range sortedCopy {
		for i, _ := range memory.empty {
			empty := &memory.empty[i]
			if empty.count >= file.count*file.len {
				for i, f := range memory.files {
					originalFile := &memory.files[i]
					if f.id == file.id {
						originalFile.start = empty.start
						break
					}
				}
				empty.count -= file.count * file.len
				empty.start += file.count * file.len
				break
			}
		}
	}

	res2 := memory.CountCheckSum()
	fmt.Println(memory.files)

	answerChan <- 0
	answerChan <- res2

	doneChan <- true
}
