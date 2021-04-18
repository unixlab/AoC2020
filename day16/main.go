package day16

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Rule struct {
	Name string
	Min1 int
	Max1 int
	Min2 int
	Max2 int
}

func (r Rule) check(value int) bool {
	if (value >= r.Min1 && value <= r.Max1) || (value >= r.Min2 && value <= r.Max2) {
		return true
	}
	return false
}

func Run() {
	var rules []Rule
	var validTickets [][]int

	file, _ := os.Open("day16/input.txt")
	scanner := bufio.NewScanner(file)

	rulesRegex, _ := regexp.Compile("^([a-z ]+): ([0-9]+)-([0-9]+) or ([0-9]+)-([0-9]+)$")

	sum := 0
	nearby := false
	yourTicket := false
	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "nearby tickets") {
			nearby = true
			continue
		}

		if yourTicket {
			var ticket []int
			for _, v := range strings.Split(line, ",") {
				value, _ := strconv.Atoi(v)
				ticket = append(ticket, value)
			}
			validTickets = append(validTickets, ticket)
			yourTicket = false
		}

		if strings.HasPrefix(line, "your ticket") {
			yourTicket = true
			continue
		}

		if nearby {
			var ticket []int
			validTicket := true
			for _, v := range strings.Split(line, ",") {
				value, _ := strconv.Atoi(v)
				ticket = append(ticket, value)
				match := false
				for _, rule := range rules {
					if rule.check(value) {
						match = true
						break
					}
				}
				if !match {
					sum += value
					validTicket = false
				}
			}
			if validTicket {
				validTickets = append(validTickets, ticket)
			}
		} else {
			regexResult := rulesRegex.FindAllStringSubmatch(line, -1)

			if len(regexResult) > 0 {
				ruleName := regexResult[0][1]
				min1, _ := strconv.Atoi(regexResult[0][2])
				max1, _ := strconv.Atoi(regexResult[0][3])
				min2, _ := strconv.Atoi(regexResult[0][4])
				max2, _ := strconv.Atoi(regexResult[0][5])

				rules = append(rules, Rule{ruleName, min1, max1, min2, max2})
			}
		}
	}
	fmt.Printf("part 1 => %d\n", sum)

	possibilities := make(map[string][]int, len(rules))
	for i := 0; i < len(validTickets[0]); i++ {
		for _, r := range rules {
			matchAll := true
			for j := 0; j < len(validTickets); j++ {
				if !r.check(validTickets[j][i]) {
					matchAll = false
				}
			}
			if matchAll {
				possibilities[r.Name] = append(possibilities[r.Name], i)
			}
		}
	}

	var sumPart2 int

	for len(possibilities) > 0 {
		for field, col := range possibilities {
			if len(col) == 1 {
				if strings.HasPrefix(field, "departure") {
					if sumPart2 == 0 {
						sumPart2 = validTickets[0][col[0]]
					} else {
						sumPart2 = sumPart2 * validTickets[0][col[0]]
					}
				}
				for f, c := range possibilities {
					var newCols []int
					for _, v := range c {
						if v != col[0] {
							newCols = append(newCols, v)
						}
					}
					possibilities[f] = newCols
				}
				delete(possibilities, field)
			}
		}
	}

	fmt.Printf("part 2 => %d\n", sumPart2)
}
