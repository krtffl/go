package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log"
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
	priorities := 0
	rucksacks := parseInput(input)

	for _, line := range rucksacks {
		rightHalf := line[len(line)/2:]

		if len(rightHalf) != len(line)/2 {
			log.Fatalln("error - length cannot be split in half", line)
		}

		for i := 0; i < len(line)/2; i++ {
			leftChar := line[i]

			if strings.Contains(rightHalf, string(leftChar)) {
				if priority, ok := charToPriority(leftChar); ok {
					priorities += priority
					break
				}

				log.Fatalln("error - invalid char", leftChar)
				break
			}
		}
	}

	return priorities
}

func part2(input string) int {
	priorities := 0
	rucksacks := parseInput(input)

	for i := 0; i < len(rucksacks); i += 3 {
		set2, set3 := stringToSet(rucksacks[i+1]), stringToSet(rucksacks[i+2])
		for j := 0; j < len(rucksacks[i]); j++ {
			char := string(rucksacks[i][j])
			if set2[char] && set3[char] {
				if priority, ok := charToPriority(rucksacks[i][j]); ok {
					priorities += priority
					break
				}

				log.Fatalln("error - invalid char", char)
				break
			}
		}
	}

	return priorities
}

func parseInput(input string) (parsed []string) {
	return strings.Split(input, "\n")
}

func stringToSet(s string) map[string]bool {
	set := map[string]bool{}
	for _, char := range strings.Split(s, "") {
		set[char] = true
	}
	return set
}

func charToPriority(char byte) (int, bool) {
	if char >= 97 && char <= 122 {
		return int(char % 96), true
	}

	if char >= 65 && char <= 90 {
		return int(char%64 + 26), true
	}

	return 0, false

}
