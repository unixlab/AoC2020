package day12

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func getDistance(x1 int, y1 int, x2 int, y2 int) int {
	return int(math.Abs(float64(x1)-float64(x2)) + math.Abs(float64(y1)-float64(y2)))
}

func RunP1() {
	file, _ := os.Open("day12/input.txt")
	scanner := bufio.NewScanner(file)

	direction := 90
	x := 0
	y := 0

	for scanner.Scan() {
		line := scanner.Text()

		action := line[0:1]
		value, _ := strconv.Atoi(line[1:])

		switch action {
		case "N":
			y += value
		case "S":
			y -= value
		case "E":
			x += value
		case "W":
			x -= value
		case "L":
			direction -= value
		case "R":
			direction += value
		case "F":
			switch direction % 360 / 90 {
			case 0:
				y += value
			case 1:
				x += value
			case 2:
				y -= value
			case 3:
				y += value
			}
		default:
			panic("unknown action")
		}
	}
	fmt.Printf("part 1 => %d\n", getDistance(0, y, 0, x))
}

func RunP2() {
	file, _ := os.Open("day12/input.txt")
	scanner := bufio.NewScanner(file)

	y := 0
	x := 0

	wayY := 1
	wayX := 10

	for scanner.Scan() {
		line := scanner.Text()

		action := line[0:1]
		value, _ := strconv.Atoi(line[1:])

		switch action {
		case "N":
			wayY += value
		case "S":
			wayY -= value
		case "E":
			wayX += value
		case "W":
			wayX -= value
		case "L":
			for i := value; i > 0; i -= 90 {
				newWayY := wayX
				newWayX := -wayY
				wayX = newWayX
				wayY = newWayY
			}
		case "R":
			for i := value; i > 0; i -= 90 {
				newWayY := -wayX
				newWayX := wayY
				wayX = newWayX
				wayY = newWayY
			}
		case "F":
			y += wayY * value
			x += wayX * value
		default:
			panic("unknown action")
		}
	}
	fmt.Printf("part 2 => %d\n", getDistance(0, y, 0, x))
}
