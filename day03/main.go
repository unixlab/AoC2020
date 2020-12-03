package day03

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func RunP1() {
	fmt.Println(Slope(3, 1))
}

func Slope(right int, down int) int {
	file, _ := os.Open("day03/input.txt")
	scanner := bufio.NewScanner(file)

	var grid [323][31]bool
	var y, x, treeCount int

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, "")
		for x, field := range fields {
			if field == "#" {
				grid[y][x] = true
			} else {
				grid[y][x] = false
			}
		}
		y++
	}

	for y := 0; y < len(grid); y += down {
		if grid[y][x] {
			treeCount++
		}
		x += right
		if x >= len(grid[y]) {
			x = x - len(grid[y])
		}
	}
	return treeCount
}
