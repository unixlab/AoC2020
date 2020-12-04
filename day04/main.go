package day04

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Run() {
	file, _ := os.Open("day04/input.txt")
	scanner := bufio.NewScanner(file)

	passport := make(map[string]string, 8)
	validPassportsPart1 := 0

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {
			passportOk := true
			for _, f := range []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"} {
				_, exists := passport[f]
				if !exists {
					passportOk = false
				}
			}

			if passportOk {
				validPassportsPart1++
			}
			passport = make(map[string]string, 8)
			continue
		}

		for _, e := range strings.Split(line, " ") {
			passport[strings.Split(e, ":")[0]] = strings.Split(e, ":")[1]
		}
	}

	fmt.Printf("part 1 => %d\n", validPassportsPart1)
}
