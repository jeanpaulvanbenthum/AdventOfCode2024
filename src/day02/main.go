package main

import (
	"bufio"
	"fmt"
	"github.com/gookit/goutil/mathutil"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	list1, list2 := getLocationIds()
	part1(list1, list2)
	part2(list1, list2)
}

func getLocationIds() ([]int, []int) {
	file, err := os.Open("input1.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)

	list1 := make([]int, 0)
	list2 := make([]int, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		ids := strings.Split(line, "   ")
		id1, _ := strconv.Atoi(ids[0])
		id2, _ := strconv.Atoi(ids[1])
		list1 = append(list1, int(id1))
		list2 = append(list2, int(id2))
	}

	sort.Ints(list1)
	sort.Ints(list2)

	return list1, list2
}

func part1(list1 []int, list2 []int) {
	var delta int
	for i, _ := range list1 {
		delta += mathutil.Abs(list1[i] - list2[i])
	}

	fmt.Printf("The total distane between the lists is: %d\n\n", delta)
}

func part2(list1 []int, list2 []int) {
	var similarityScore int
	for i, _ := range list1 {
		number := list1[i]
		similarityScore += number * occurrences(list2, number)
	}

	fmt.Printf("The total similarity score is: %d\n\n", similarityScore)
}

func occurrences(list []int, value int) int {
	count := 0
	for _, v := range list {
		if v == value {
			count++
		}
	}
	return count
}
