package main

import (
	_ "embed"
	"flag"
	"fmt"
	"log"
	"strings"
)

const (
	Win  = 6
	Draw = 3
	Loss = 0

	Rock     = 1
	Paper    = 2
	Scissors = 3
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

	choices := map[string]int{
		"X": Rock,
		"Y": Paper,
		"Z": Scissors,
	}

	score := 0

	for _, l := range lines {
		if _, ok := choices[l[1]]; !ok {
			log.Fatalln("error - choice is not valid")
		}

		score += choices[l[1]]

		switch l[1] {
		case "X": // i played rock
			switch l[0] {
			case "A":
				score += Draw
			case "B":
				score += Loss
			case "C":
				score += Win
			default:
				log.Fatalln("error - choice is not valid - ", l[0])
			}
		case "Y": // i played paper
			switch l[0] {
			case "A": // rock
				score += Win
			case "B": // paper
				score += Draw
			case "C": // scissors
				score += Loss
			default:
				log.Fatalln("error - choice is not valid - ", l[0])
			}
		case "Z": // i played scissors
			switch l[0] {
			case "A": // rock
				score += Loss
			case "B": // paper
				score += Win
			case "C": // scissors
				score += Draw
			default:
				log.Fatalln("error - choice is not valid - ", l[0])
			}
		}
	}

	return score
}

func part2(input string) int {
	lines := parseInput(input)

	results := map[string]int{
		"X": Loss,
		"Y": Draw,
		"Z": Win,
	}

	score := 0

	for _, l := range lines {
		if _, ok := results[l[1]]; !ok {
			log.Fatalln("error - result is not valid")
		}

		score += results[l[1]]

		switch l[0] {
		case "A": // opp: rock
			switch l[1] {
			case "X": // lose
				score += Scissors
			case "Y": // draw
				score += Rock
			case "Z": // win
				score += Paper
			default:
				log.Fatalln("error - choice is not valid - ", l[0])
			}
		case "B": // opp: paper
			switch l[1] {
			case "X": // lose
				score += Rock
			case "Y": // draw
				score += Paper
			case "Z": // win
				score += Scissors
			default:
				log.Fatalln("error - choice is not valid - ", l[0])
			}
		case "C": // opp: scissors
			switch l[1] {
			case "X": // lose
				score += Paper
			case "Y": // draw
				score += Scissors
			case "Z": // win
				score += Rock
			default:
				log.Fatalln("error - choice is not valid - ", l[0])
			}
		}
	}

	return score
}

func parseInput(input string) (parsed [][]string) {
	for _, line := range strings.Split(input, "\n") {
		parsed = append(parsed, strings.Split(line, " "))
	}
	return parsed
}
