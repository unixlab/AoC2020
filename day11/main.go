package day11

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	LENGTH = 98
)

func getNumberOfAdjacent(grid [LENGTH][LENGTH]int, y int, x int) int {
	var sum int
	for i := y - 1; i <= y+1; i++ {
		for j := x - 1; j <= x+1; j++ {
			if i < 0 || i >= LENGTH {
				continue
			}
			if j < 0 || j >= LENGTH {
				continue
			}
			if i == y && j == x {
				continue
			}
			if grid[i][j] == 2 {
				sum++
			}
		}
	}
	return sum
}

func Run() {
	file, _ := os.Open("day11/input.txt")
	scanner := bufio.NewScanner(file)

	var grid [LENGTH][LENGTH]int

	y := 0
	for scanner.Scan() {
		line := scanner.Text()

		seats := strings.Split(line, "")

		for x, seat := range seats {
			if seat == "." {
				grid[y][x] = 0
			} else if seat == "L" {
				grid[y][x] = 1
			} else {
				grid[y][x] = 2
			}
		}
		y++
	}

	changes := 1
	for changes > 0 {
		changes = 0

		var oldGrid [LENGTH][LENGTH]int
		for i := 0; i < LENGTH; i++ {
			for j := 0; j < LENGTH; j++ {
				oldGrid[i][j] = grid[i][j]
			}
		}

		for y, row := range oldGrid {
			for x, _ := range row {
				if grid[y][x] == 0 {
					continue
				}
				numberOfSeats := getNumberOfAdjacent(oldGrid, y, x)
				if numberOfSeats == 0 {
					if grid[y][x] != 2 {
						grid[y][x] = 2
						changes++
					}
				}
				if numberOfSeats >= 4 {
					if grid[y][x] != 1 {
						grid[y][x] = 1
						changes++
					}
				}
			}
		}
	}

	var sum int
	for i := 0; i < LENGTH; i++ {
		for j := 0; j < LENGTH; j++ {
			if grid[i][j] == 2 {
				sum++
			}
		}
	}

	fmt.Printf("part 1 => %d\n", sum)
}
