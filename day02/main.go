package day02

import (
	"bufio"
	"fmt"
	"os"
)

func Run() {
	file, _ := os.Open("day02/input.txt")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
}
