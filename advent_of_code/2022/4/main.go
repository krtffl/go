package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func init() {
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		log.Fatalln("error - input is empty")
	}
}

func main() {
	var part int

	flag.IntVar(&part, "part", 1, "solve part 1 or part 2")
	flag.Parse()

	fmt.Println("running part", part)

	if part == 1 {
		fmt.Println("answer: ", part1(input))
	}

	if part == 2 {
		fmt.Println("answer: ", part2(input))
	}
}

func part1(input string) int {
	lines := parseInput(input)
	count := 0

	for _, line := range lines {
		if doesLeftContain(line[:2], line[2:]) || doesLeftContain(line[2:], line[:2]) {
			count++
		}
	}

	return count
}

func part2(input string) int {
	lines := parseInput(input)
	count := 0

	for _, line := range lines {
		if doesOverlap(line[:2], line[2:]) {
			count++
		}
	}

	return count
}

func doesOverlap(left, right []int) bool {
	if left[0] > right[0] {
		left, right = right, left
	}
	return left[1] >= right[0]
}

func doesLeftContain(left, right []int) bool {
	return left[0] <= right[0] && left[1] >= right[1]
}

func parseInput(input string) (parsed [][]int) {
	for _, line := range strings.Split(input, "\n") {
		sections := strings.Split(line, ",")

		leftSection := strings.Split(sections[0], "-")
		rightSection := strings.Split(sections[1], "-")

		parsed = append(parsed, []int{
			parseInt(leftSection[0]),
			parseInt(leftSection[1]),
			parseInt(rightSection[0]),
			parseInt(rightSection[1]),
		})
	}

	return parsed
}

func parseInt(input string) int {
	value, err := strconv.Atoi(input)

	if err != nil {
		log.Fatalln("error - cannot parse to int")
	}

	return value
}
