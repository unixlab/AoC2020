package day15

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Run() {
	file, _ := os.Open("day15/input.txt")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var last int
		lastSpoken := make(map[int]int, 1000)

		for k, v := range strings.Split(scanner.Text(), ",") {
			last, _ = strconv.Atoi(v)
			lastSpoken[last] = k
		}

		last = 0

		for i := len(lastSpoken); i < 30e6; i++ {
			pos, ok := lastSpoken[last]
			lastSpoken[last] = i
			if ok {
				last = i - pos
			} else {
				last = 0
			}
			if i == 2018 {
				fmt.Printf("part 1 => %d\n", last)
			}
			if i == 29999998 {
				fmt.Printf("part 1 => %d\n", last)
			}
		}
	}
}
