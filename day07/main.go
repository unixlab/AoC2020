package day07

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func getShinyGold(bags map[string][]string, bag string) bool {
	if len(bags[bag]) == 0 {
		return false
	}
	if bag == "shiny_gold" {
		return true
	}
	for _, nextBag := range bags[bag] {
		if getShinyGold(bags, nextBag) {
			return true
		}
	}
	return false
}

func Run() {
	file, _ := os.Open("day07/input.txt")
	scanner := bufio.NewScanner(file)

	bagRegex, _ := regexp.Compile("([0-9]+ [A-Za-z]+ [A-Za-z]+ bag[s]?)")

	bags := make(map[string][]string, 1000)

	for scanner.Scan() {
		line := scanner.Text()

		bag := strings.Join(strings.Split(line, " ")[0:2], "_")
		bagRegexRes := bagRegex.FindAllStringSubmatch(line, -1)

		var inside []string

		for _, bagInside := range bagRegexRes {
			inside = append(inside, strings.Join(strings.Split(bagInside[0], " ")[1:3], "_"))
		}

		bags[bag] = inside
	}

	bagCounter := 0
	for bag := range bags {
		if bag == "shiny_gold" {
			continue
		}
		if getShinyGold(bags, bag) {
			bagCounter++
		}
	}

	fmt.Printf("part 1 => %d\n", bagCounter)
}
