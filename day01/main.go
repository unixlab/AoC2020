package day01

import (
	"bufio"
	"fmt"
	"os"
)

func RunP1() {
	file, _ := os.Open("day01/input.txt")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
}