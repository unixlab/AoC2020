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

func RunP2() {
	file, _ := os.Open("day13/input.txt")
	scanner := bufio.NewScanner(file)

	var busLines []int

	for scanner.Scan() {
		line := scanner.Text()

		fields := strings.Split(line, ",")

		if len(fields) == 1 {
			continue

		}
		for _, field := range fields {
			if field == "x" {
				busLines = append(busLines, 0)
			} else {
				bus, _ := strconv.Atoi(field)
				busLines = append(busLines, bus)
			}
		}
	}

	var maxBus int
	var maxBusOffset int
	for offset, bus := range busLines {
		if bus > maxBus {
			maxBus = bus
			maxBusOffset = offset
		}
	}

	var i uint64
	i = uint64(maxBus - maxBusOffset)

	maxBusUint64 := uint64(maxBus)

	for {
		brokeLoop := false
		for offset, busId := range busLines {
			if busId == 0 {
				continue
			}
			if (i+uint64(offset))%uint64(busId) != 0 {
				brokeLoop = true
				break
			}
		}
		if !brokeLoop {
			break
		}
		i += maxBusUint64
	}

	fmt.Printf("part 2 => %d\n", i)
}
