package day01

import (
	"bufio"
	"fmt"
	"gonum.org/v1/gonum/stat/combin"
	"os"
	"strconv"
)

func Run() {
	file, _ := os.Open("day01/input.txt")
	scanner := bufio.NewScanner(file)

	var numbers []int

	for scanner.Scan() {
		line := scanner.Text()

		number, _ := strconv.Atoi(line)
		numbers = append(numbers, number)
	}

	// Part 1
	combinationsIndex := combin.Combinations(len(numbers), 2)
	for _, combinationIndex := range combinationsIndex {
		if numbers[combinationIndex[0]]+numbers[combinationIndex[1]] == 2020 {
			fmt.Printf("part 1 => %d\n", numbers[combinationIndex[0]]*numbers[combinationIndex[1]])
			break
		}
	}

	// Part 2
	combinationsIndex = combin.Combinations(len(numbers), 3)
	for _, combinationIndex := range combinationsIndex {
		if numbers[combinationIndex[0]]+numbers[combinationIndex[1]]+numbers[combinationIndex[2]] == 2020 {
			fmt.Printf("part 2 => %d\n", numbers[combinationIndex[0]]*numbers[combinationIndex[1]]*numbers[combinationIndex[2]])
			break
		}
	}
}
