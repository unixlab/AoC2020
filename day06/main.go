package day06

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Run() {
	file, _ := os.Open("day06/input.txt")
	scanner := bufio.NewScanner(file)

	var sumPart1 int
	var sumPart2 int
	var people int
	questions := make(map[string]int, 26)

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			sumPart1 += len(questions)

			for _, v := range questions {
				if v == people {
					sumPart2++
				}
			}

			questions = make(map[string]int, 26)
			people = 0
			continue
		}

		for _, c := range strings.Split(line, "") {
			questions[c]++
		}
		people++
	}

	sumPart1 += len(questions)
	for _, v := range questions {
		if v == people {
			sumPart2++
		}
	}

	fmt.Printf("part 1 => %d\n", sumPart1)
	fmt.Printf("part 2 => %d\n", sumPart2)
}
