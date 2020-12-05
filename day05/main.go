package day05

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Partition struct {
	LowerBound int
	UpperBound int
}

func getNewPartition(oldPartition Partition, key string) Partition {
	var newPartition Partition
	if key == "F" || key == "L" {
		newPartition.LowerBound = oldPartition.LowerBound
		newPartition.UpperBound = oldPartition.UpperBound - ((oldPartition.UpperBound - oldPartition.LowerBound) / 2) - 1
	} else if key == "B" || key == "R" {
		newPartition.LowerBound = oldPartition.LowerBound + ((oldPartition.UpperBound - oldPartition.LowerBound) / 2) + 1
		newPartition.UpperBound = oldPartition.UpperBound
	}
	return newPartition
}

func Run() {
	file, _ := os.Open("day05/input.txt")
	scanner := bufio.NewScanner(file)

	var maxSeatId int
	var allPart2Seats []int

	for scanner.Scan() {
		line := scanner.Text()

		rowString := line[:7]
		seatString := line[7:]

		rowPartition := Partition{0, 127}
		seatPartition := Partition{0, 7}

		for i := 0; i < 7; i++ {
			rowPartition = getNewPartition(rowPartition, rowString[i:i+1])
		}

		for i := 0; i < 3; i++ {
			seatPartition = getNewPartition(seatPartition, seatString[i:i+1])
		}

		seatId := rowPartition.UpperBound*8 + seatPartition.UpperBound

		if seatId > maxSeatId {
			maxSeatId = seatId
		}

		allPart2Seats = append(allPart2Seats, seatId)
	}
	fmt.Printf("part 1 => %d\n", maxSeatId)

	sort.Ints(allPart2Seats)
	for i := 0; i < len(allPart2Seats)-1; i++ {
		if allPart2Seats[i+1]-allPart2Seats[i] > 1 {
			fmt.Printf("part 2 => %d\n", allPart2Seats[i]+1)
		}
	}
}
