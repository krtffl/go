package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
    played := map[string]string{
        "A": "rock",
        "B": "paper",
        "C": "scissors",
        "X": "rock",
        "Y": "paper",
        "Z": "scissors",
    }

    points := map[string]int{
        "rock": 1,
        "paper": 2,
        "scissors": 3,
    }

    outcome := map[string]int{
        "X": -1,
        "Y": 0,
        "Z": 1,
    }

    input, err := os.ReadFile("./input.txt")

    if err != nil {
        log.Fatal(err)
    }

    stringifiedInput :=strings.TrimSpace(string(input))
    data := strings.Split(stringifiedInput, "\n")

    score := 0

    for i:= 0; i < len(data); i++ {
        plays := make([]string, 2)
        play := strings.Split(data[i], " ")

        plays[0] = played[play[0]]
        plays[1] = choose(outcome[play[1]], plays[0])

        score += duel(plays[0], plays[1])
        score += points[plays[1]]
    }

    fmt.Println(score)
}

func duel(left, right string) int {
    if left == right {
        return 3
    }

    if (left == "rock" && right == "paper") ||  (left == "paper" && right == "scissors") || (left == "scissors" && right == "rock") {
        return 6
    }
    return 0
}

func choose(outcome int, against string) string {
    if (outcome == 1 && against == "paper") || (outcome == -1 && against == "rock") {
        return "scissors"
    }

    if (outcome == 1 && against == "rock") || (outcome == -1 && against == "scissors") {
        return "paper"
    }

    if (outcome == 1 && against == "scissors") || (outcome == -1 && against == "paper") {
        return "rock"
    }
    return against
}

