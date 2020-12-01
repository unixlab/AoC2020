package dayXX

import (
	"bufio"
	"fmt"
	"os"
)

func Run() {
	file, _ := os.Open("dayXX/input.txt")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
}
