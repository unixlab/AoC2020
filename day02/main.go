package day02

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Run() {
	file, _ := os.Open("day02/input.txt")
	scanner := bufio.NewScanner(file)

	validCounterPart1 := 0

	for scanner.Scan() {
		line := scanner.Text()
		sections := strings.Split(line, " ")
		min, _ := strconv.Atoi(strings.Split(sections[0], "-")[0])
		max, _ := strconv.Atoi(strings.Split(sections[0], "-")[1])
		char := strings.Replace(sections[1], ":", "", -1)
		password := sections[2]

		charCount := strings.Count(password, char)

		if charCount >= min && charCount <= max {
			validCounterPart1++
		}

	}

	fmt.Printf("part 1 => %d\n", validCounterPart1)
}
