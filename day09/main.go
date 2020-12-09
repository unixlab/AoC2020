package day09

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const (
	SIZE = 25
)

func Run() {
	file, _ := os.Open("day09/input.txt")
	scanner := bufio.NewScanner(file)

	var numbers []int
	var resultPart1 int

	for scanner.Scan() {
		line := scanner.Text()
		number, _ := strconv.Atoi(line)

		numbers = append(numbers, number)

		start := len(numbers) - SIZE - 1

		if start < 0 {
			continue
		}

		valid := false

		for i := start; i < len(numbers); i++ {
			for j := start; j < len(numbers); j++ {
				if i == j {
					continue
				}
				if numbers[i]+numbers[j] == number {
					valid = true
				}
			}
		}

		if !valid {
			resultPart1 = number
		}
	}

	fmt.Printf("part 1 => %d\n", resultPart1)
}
