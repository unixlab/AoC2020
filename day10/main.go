package day10

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func Run() {
	file, _ := os.Open("day10/input.txt")
	scanner := bufio.NewScanner(file)

	var adapters []int

	adapters = append(adapters, 0)

	for scanner.Scan() {
		line := scanner.Text()
		adapter, _ := strconv.Atoi(line)
		adapters = append(adapters, adapter)
	}

	sort.Ints(adapters)

	adapters = append(adapters, adapters[len(adapters)-1]+3)

	var oneJoltDiff int
	var threeJoltDiff int

	for i := 1; i < len(adapters); i++ {
		diff := adapters[i] - adapters[i-1]
		if diff == 1 {
			oneJoltDiff++
		}
		if diff == 3 {
			threeJoltDiff++
		}
	}

	part1 := oneJoltDiff * threeJoltDiff
	fmt.Printf("part 1 => %d\n", part1)
}
