package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	program := getProgram()
	fmt.Print("Part 1: ")
	runProgram(program, true)

	fmt.Print("Part 2: ")
	runProgram(program, false)
}

func runProgram(program string, alwaysEnabled bool) {
	instructions := getInstructions(program)

	var result int
	var enabled bool
	enabled = true
	for _, instruction := range instructions {
		if strings.Contains(instruction, "do()") {
			enabled = true
			continue
		}
		if strings.Contains(instruction, "don't()") {
			enabled = false
			continue
		}

		if enabled || alwaysEnabled {
			_, arg1, arg2 := parseInstruction(instruction)

			result += mul(arg1, arg2)
		}
	}

	fmt.Printf("The total of all the multiplications is: %d\n\n", result)
}

func getProgram() string {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	return string(content)
}

func getInstructions(program string) []string {
	re := regexp.MustCompile(`(mul\(\d{1,3},\d{1,3}\)|do\(\)|don't\(\))`)
	matches := re.FindAllString(program, -1)

	return matches
}

func parseInstruction(instruction string) (string, int, int) {
	// Extract function name and arguments from the instruction
	re := regexp.MustCompile(`(\w+)\((\d+),(\d+)\)`)
	matches := re.FindStringSubmatch(instruction)
	if len(matches) != 4 {
		log.Fatalf("Invalid input: %s", instruction)
	}

	funcName := matches[1]
	arg1, _ := strconv.Atoi(matches[2])
	arg2, _ := strconv.Atoi(matches[3])

	return funcName, arg1, arg2
}

func mul(a, b int) int {
	return a * b
}
