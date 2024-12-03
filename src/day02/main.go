package main

import (
	"bufio"
	"fmt"
	"github.com/gookit/goutil/mathutil"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	reports(false)
	reports(true)
}

func reports(tolerate bool) {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	scanner := bufio.NewScanner(file)
	safeCounter := 0

scannerLoop:
	for scanner.Scan() {
		report := scanner.Text()
		originalLevels := stringSliceToIntSlice(strings.Split(report, " "))

		var newLevels [][]int
		if tolerate {
			newLevels = generateSlices(originalLevels)
		} else {
			newLevels = [][]int{originalLevels}
		}

	tolerateLoop:
		for _, levels := range newLevels {
			if !checkIncreasingOrDecreasing(levels) {
				continue tolerateLoop
			}

			for i := range levels {
				if i == 0 {
					diff := mathutil.Abs(levels[1] - levels[0])

					// Safe when any two adjacent levels differ by at least one and at most three.
					if diff < 1 || diff > 3 {
						continue tolerateLoop
					}
				}

				if i > 0 && i < len(levels)-1 {
					diff := mathutil.Abs(levels[i] - levels[i-1])
					if diff < 1 || diff > 3 {
						continue tolerateLoop
					}
					diff = mathutil.Abs(levels[i] - levels[i+1])
					if diff < 1 || diff > 3 {
						continue tolerateLoop
					}
				}

				if i == len(levels)-1 {
					diff := mathutil.Abs(levels[i] - levels[i-1])

					// Safe when any two adjacent levels differ by at least one and at most three.
					if diff < 1 || diff > 3 {
						continue tolerateLoop
					}
				}
			}

			safeCounter++
			continue scannerLoop
		}
	}

	fmt.Printf("Safe couter is: %d\n\n", safeCounter)
}

func stringSliceToIntSlice(strings []string) []int {
	var intSlice []int
	for _, stringVal := range strings {
		// Convert each string to an integer
		integer, err := strconv.Atoi(stringVal)
		if err != nil {
			log.Fatalf("Error converting %s to int: %v\n\n", stringVal, err)
		}
		intSlice = append(intSlice, integer)
	}

	return intSlice
}

func checkIncreasingOrDecreasing(values []int) bool {
	ascendingValues := make([]int, len(values))
	descendingValues := make([]int, len(values))
	copy(ascendingValues, values)
	copy(descendingValues, values)

	slices.Sort(ascendingValues)
	if slices.Compare(ascendingValues, values) == 0 {
		return true
	}

	slices.Sort(descendingValues)
	slices.Reverse(descendingValues)
	if slices.Compare(descendingValues, values) == 0 {
		return true
	}

	return false
}

func generateSlices(original []int) [][]int {
	result := make([][]int, 0, len(original))
	result = append(result, original)

	for i := range original {
		newSlice := append([]int{}, original[:i]...)
		newSlice = append(newSlice, original[i+1:]...)
		result = append(result, newSlice)
	}

	return result
}
