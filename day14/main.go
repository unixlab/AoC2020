package day14

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func calculateNewValue(currentMask string, value string) string {
	var newValue strings.Builder
	for i := 0; i < 36; i++ {
		switch currentMask[i] {
		// X
		case 88:
			newValue.WriteString(value[i : i+1])
		// 0
		case 48:
			newValue.WriteString("0")
		// 1
		case 49:
			newValue.WriteString("1")
		default:
			panic("meeep")
		}
	}
	return newValue.String()
}

func intToBit(i int) string {
	var b strings.Builder
	for j := 35; j >= 0; j-- {
		b.WriteString(strconv.Itoa(i / powInt(2, j)))
		i = i % powInt(2, j)
	}
	return b.String()
}

func bitToInt(b string) int {
	v := 0
	for i := 0; i < 36; i++ {
		if b[i] == 48 {
			continue
		}
		v += powInt(2, 35-i)
	}
	return v
}

func powInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func Run() {
	file, _ := os.Open("day14/input.txt")
	scanner := bufio.NewScanner(file)

	var currentMask string

	memory := make(map[string]string, 100)

	maskRegex, _ := regexp.Compile("^mask = ([A-Za-z0-9]{36})$")
	memoryRegex, _ := regexp.Compile("^mem\\[([0-9]+)\\] = ([0-9]+)$")

	for scanner.Scan() {
		line := scanner.Text()

		maskResult := maskRegex.FindAllStringSubmatch(line, -1)
		if len(maskResult) > 0 {
			currentMask = maskResult[0][1]
			continue
		}

		memoryResult := memoryRegex.FindAllStringSubmatch(line, -1)
		if len(memoryResult) > 0 {
			memoryAddress := memoryResult[0][1]
			valueInt, _ := strconv.Atoi(memoryResult[0][2])
			value := intToBit(valueInt)

			memory[memoryAddress] = calculateNewValue(currentMask, value)
		}
	}

	sum := 0

	for _, v := range memory {
		sum += bitToInt(v)
	}

	fmt.Printf("part 1 => %d\n", sum)
}
