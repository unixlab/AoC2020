package day13

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func RunP1() {
	file, _ := os.Open("day13/input.txt")
	scanner := bufio.NewScanner(file)

	var startTime int
	var busLines []int

	for scanner.Scan() {
		line := scanner.Text()

		fields := strings.Split(line, ",")

		if len(fields) == 1 {
			startTime, _ = strconv.Atoi(fields[0])
		} else {
			for _, field := range fields {
				if field == "x" {
					continue
				}
				bus, _ := strconv.Atoi(field)
				busLines = append(busLines, bus)
			}
		}
	}

	minTime := -1
	minBus := -1
	for _, bus := range busLines {
		diff := bus - (startTime % bus)

		if diff < minTime || minTime == -1 {
			minTime = diff
			minBus = bus
		}
	}

	part1 := minTime * minBus

	fmt.Printf("part 1 => %d\n", part1)
}
