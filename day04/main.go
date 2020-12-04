package day04

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Run() {
	file, _ := os.Open("day04/input.txt")
	scanner := bufio.NewScanner(file)

	passport := make(map[string]string, 8)
	validPassportsPart1 := 0
	validPassportsPart2 := 0

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

				byr, _ := strconv.Atoi(passport["byr"])
				if byr < 1920 || byr > 2002 {
					passportOk = false
				}

				iyr, _ := strconv.Atoi(passport["iyr"])
				if iyr < 2010 || iyr > 2020 {
					passportOk = false
				}

				eyr, _ := strconv.Atoi(passport["eyr"])
				if eyr < 2020 || eyr > 2030 {
					passportOk = false
				}

				hgt := passport["hgt"]
				if strings.Index(hgt, "cm") >= 0 {
					hgt = strings.Replace(hgt, "cm", "", 1)
					hgtInt, _ := strconv.Atoi(hgt)
					if hgtInt > 193 || hgtInt < 150 {
						passportOk = false
					}
				} else {
					hgt = strings.Replace(hgt, "in", "", 1)
					hgtInt, _ := strconv.Atoi(hgt)
					if hgtInt < 59 || hgtInt > 76 {
						passportOk = false
					}
				}

				hcl := passport["hcl"]
				colorRegex, _ := regexp.Compile("^#[0-9a-f]{6}$")
				if !colorRegex.MatchString(hcl) {
					passportOk = false
				}

				ecl := passport["ecl"]
				if ecl != "amb" && ecl != "blu" && ecl != "brn" && ecl != "gry" && ecl != "grn" && ecl != "hzl" && ecl != "oth" {
					passportOk = false
				}

				pid := passport["pid"]
				pidRegex, _ := regexp.Compile("^[0-9]{9}$")
				if !pidRegex.MatchString(pid) {
					passportOk = false
				}

				if passportOk {
					validPassportsPart2++
				}
			}
			passport = make(map[string]string, 8)
			continue
		}

		for _, e := range strings.Split(line, " ") {
			passport[strings.Split(e, ":")[0]] = strings.Split(e, ":")[1]
		}
	}

	fmt.Printf("part 1 => %d\n", validPassportsPart1)
	fmt.Printf("part 2 => %d\n", validPassportsPart2)
}
