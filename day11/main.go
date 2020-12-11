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

func getNumberOfAdjacentP2(grid [LENGTH][LENGTH]int, y int, x int) int {
	var sum int

	for i := y - 1; i >= 0; i-- {
		if grid[i][x] == 2 {
			sum++
			break
		} else if grid[i][x] == 1 {
			break
		}
	}

	for i := y + 1; i < LENGTH; i++ {
		if grid[i][x] == 2 {
			sum++
			break
		} else if grid[i][x] == 1 {
			break
		}
	}

	for i := x - 1; i >= 0; i-- {
		if grid[y][i] == 2 {
			sum++
			break
		} else if grid[y][i] == 1 {
			break
		}
	}

	for i := x + 1; i < LENGTH; i++ {
		if grid[y][i] == 2 {
			sum++
			break
		} else if grid[y][i] == 1 {
			break
		}
	}

	{
		i := y
		j := x
		for i > 0 && j > 0 {
			i--
			j--
			if grid[i][j] == 2 {
				sum++
				break
			} else if grid[i][j] == 1 {
				break
			}
		}
	}

	{
		i := y
		j := x
		for i < LENGTH-1 && j < LENGTH-1 {
			i++
			j++
			if grid[i][j] == 2 {
				sum++
				break
			} else if grid[i][j] == 1 {
				break
			}
		}
	}

	{
		i := y
		j := x
		for i < LENGTH-1 && j > 0 {
			i++
			j--
			if grid[i][j] == 2 {
				sum++
				break
			} else if grid[i][j] == 1 {
				break
			}
		}
	}

	{
		i := y
		j := x
		for i > 0 && j < LENGTH-1 {
			i--
			j++
			if grid[i][j] == 2 {
				sum++
				break
			} else if grid[i][j] == 1 {
				break
			}
		}
	}

	return sum
}

func RunP1() {
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

func RunP2() {
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
				numberOfSeats := getNumberOfAdjacentP2(oldGrid, y, x)
				if numberOfSeats == 0 {
					if grid[y][x] != 2 {
						grid[y][x] = 2
						changes++
					}
				}
				if numberOfSeats >= 5 {
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

	fmt.Printf("part 2 => %d\n", sum)
}
