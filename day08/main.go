package day08

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Instruction struct {
	Command string
	Value   int
}

func countAllZero(counter []int) bool {
	for _, c := range counter {
		if c > 1 {
			return false
		}
	}
	return true
}

func Run() {
	file, _ := os.Open("day08/input.txt")
	scanner := bufio.NewScanner(file)

	var instructions []Instruction

	for scanner.Scan() {
		line := scanner.Text()
		splitRes := strings.Split(line, " ")

		var newInstruction Instruction
		newInstruction.Command = splitRes[0]
		newInstruction.Value, _ = strconv.Atoi(splitRes[1])

		instructions = append(instructions, newInstruction)
	}

	instructionCounter := make([]int, len(instructions))

	offset := 0
	accumulator := 0
	for offset < len(instructions) {
		instructionCounter[offset]++
		if !countAllZero(instructionCounter) {
			break
		}
		switch instructions[offset].Command {
		case "nop":
			offset++
		case "acc":
			accumulator += instructions[offset].Value
			offset++
		case "jmp":
			offset += instructions[offset].Value
		}
	}

	fmt.Printf("part 1 => %d\n", accumulator)

	for key, inst := range instructions {
		offset = 0
		accumulator = 0
		instructionCounter = make([]int, len(instructions))
		orgInst := inst
		switch inst.Command {
		case "nop":
			instructions[key].Command = "jmp"
		case "acc":
			continue
		case "jmp":
			instructions[key].Command = "nop"
		}
		for offset < len(instructions) {
			instructionCounter[offset]++
			if !countAllZero(instructionCounter) {
				break
			}
			switch instructions[offset].Command {
			case "nop":
				offset++
			case "acc":
				accumulator += instructions[offset].Value
				offset++
			case "jmp":
				offset += instructions[offset].Value
			}
		}
		if offset >= len(instructions) {
			fmt.Printf("part 2 => %d\n", accumulator)
		}
		instructions[key].Command = orgInst.Command
	}
}
