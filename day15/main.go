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
		var numbersSpoken [2021]int

		var maxK int
		for k, v := range strings.Split(scanner.Text(), ",") {
			numbersSpoken[k], _ = strconv.Atoi(v)
			maxK = k
		}

		for i := maxK; i < 2020; i++ {
			numberToSpeak := 0
			for j := i - 1; j >= 0; j-- {
				if numbersSpoken[j] == numbersSpoken[i] {
					numberToSpeak = i - j
					break
				}
			}
			numbersSpoken[i+1] = numberToSpeak
		}

		fmt.Println(numbersSpoken[2019])
	}
}
