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

	for i := 0; i < len(numbers); i++ {
		var part2Numbers []int
		var sum int
		offset := i
		for sum < resultPart1 && offset < len(numbers) {
			part2Numbers = append(part2Numbers, numbers[offset])
			sum += numbers[offset]
			offset++
		}
		if sum == resultPart1 {
			min := -1
			max := 0
			for _, number := range part2Numbers {
				if number > max {
					max = number
				}
				if number < min || min == -1 {
					min = number
				}
			}
			fmt.Printf("part 2 => %d\n", min+max)
			break
		}
	}
}
