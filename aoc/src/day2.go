package src

import (
	"fmt"
	"krtffl/aoc/helpers"
	"log"
	"strings"
)

func Solve_Day_2() {
    map_play := map[string]string{
        "A": "rock",
        "B": "paper",
        "C": "scissors",
        "X": "rock",
        "Y": "paper",
        "Z": "scissors",
    }

    map_points := map[string]int{
        "rock": 1,
        "paper": 2,
        "scissors": 3,
    }

    map_outcome := map[string]int{
        "X": -1,
        "Y": 0,
        "Z": 1,
    }


    data, error := helpers.ReadInput("day2.txt")
    if error != nil {
        log.Fatal(error)
    }

    score := 0

    for i:= 0; i < len(data); i++ {
        play := strings.Split(data[i], " ")
        plays := make([]string, 2)

        plays[0] = map_play[play[0]]
        plays[1] = choose_draw(map_outcome[play[1]], plays[0])

        score += get_points(plays[0], plays[1])
        score += map_points[plays[1]]
    }

    fmt.Println(score)
}

func get_points(left, right string) int {
    if left == right {
        return 3
    }

    if (left == "rock" && right == "paper") ||  (left == "paper" && right == "scissors") || (left == "scissors" && right == "rock") {
        return 6
    }

    return 0
}

func choose_draw(outcome int, against string) string {
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

