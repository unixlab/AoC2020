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
	Min int
	Max int
}

func Run() {
	var rules []Rule

	file, _ := os.Open("day16/input.txt")
	scanner := bufio.NewScanner(file)

	rulesRegex, _ := regexp.Compile("^[a-z ]+: ([0-9]+)-([0-9]+) or ([0-9]+)-([0-9]+)$")

	sum := 0
	nearby := false
	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "nearby tickets") {
			nearby = true
			continue
		}

		if nearby {
			for _, v := range strings.Split(line, ",") {
				value, _ := strconv.Atoi(v)
				match := false
				for _, rule := range rules {
					if value >= rule.Min && value <= rule.Max {
						match = true
						break
					}
				}
				if !match {
					sum += value
				}
			}
		} else {
			regexResult := rulesRegex.FindAllStringSubmatch(line, -1)

			if len(regexResult) > 0 {
				min1, _ := strconv.Atoi(regexResult[0][1])
				max1, _ := strconv.Atoi(regexResult[0][2])
				min2, _ := strconv.Atoi(regexResult[0][3])
				max2, _ := strconv.Atoi(regexResult[0][4])

				rules = append(rules, Rule{min1, max1,})
				rules = append(rules, Rule{min2, max2})
			}
		}
	}
	fmt.Printf("part 1 => %d\n", sum)
}
