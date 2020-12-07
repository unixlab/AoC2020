package day07

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Bag struct {
	Name   string
	Number int
}

func getShinyGold(bags map[string][]Bag, bag string) bool {
	if len(bags[bag]) == 0 {
		return false
	}
	if bag == "shiny_gold" {
		return true
	}
	for _, nextBag := range bags[bag] {
		if getShinyGold(bags, nextBag.Name) {
			return true
		}
	}
	return false
}

func getBagChain(bags map[string][]Bag, bag string) int {
	var sum int
	for _, curBag := range bags[bag] {
		sum += curBag.Number
		sum += curBag.Number * getBagChain(bags, curBag.Name)
	}
	return sum
}

func Run() {
	file, _ := os.Open("day07/input.txt")
	scanner := bufio.NewScanner(file)

	bagRegex, _ := regexp.Compile("([0-9]+ [A-Za-z]+ [A-Za-z]+ bag[s]?)")

	bags := make(map[string][]Bag, 1000)

	for scanner.Scan() {
		line := scanner.Text()

		bag := strings.Join(strings.Split(line, " ")[0:2], "_")
		bagRegexRes := bagRegex.FindAllStringSubmatch(line, -1)

		var newBags []Bag

		for _, bagInside := range bagRegexRes {
			var newBag Bag
			newBag.Name = strings.Join(strings.Split(bagInside[0], " ")[1:3], "_")
			newBag.Number, _ = strconv.Atoi(strings.Split(bagInside[0], " ")[0])
			newBags = append(newBags, newBag)
		}

		bags[bag] = newBags
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
	fmt.Printf("part 2 => %d\n", getBagChain(bags, "shiny_gold"))
}
