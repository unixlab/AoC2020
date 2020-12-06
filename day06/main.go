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

	var sum int
	questions := make(map[string]int, 26)

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			sum += len(questions)
			questions = make(map[string]int, 26)
		}

		for _, c := range strings.Split(line, "") {
			questions[c]++
		}
	}

	sum += len(questions)

	fmt.Printf("part 1 => %d\n", sum)
}
