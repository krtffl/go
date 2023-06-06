package src 

import (
	"fmt"
	"krtffl/aoc/helpers"
	"log"
)

func Solve_Day_3() {
    data, error := helpers.ReadInput("day3.txt")
    if error != nil {
        log.Fatal(error) 
    }

    priorities := 0

    for i:= 0; i < len(data); i += 3 {
        rucksackA := data[i]
        rucksackB := data[i + 1]
        rucksackC := data[i + 2]


        var match byte

        for i:= 0; i < len(rucksackA); i++ {
            for j:= 0; j < len(rucksackB); j++ {
                for z:=0; z < len(rucksackC); z ++ {
                    if rucksackA[i] == rucksackB[j] && rucksackA[i] == rucksackC[z] {
                        match = rucksackA[i]
                    }
                }
            }
        }

        priorities += get_priority(match)
    }

    fmt.Println(priorities)
}


func get_priority(object byte) int {
    if object >= 97 { 
        return int(object % 96)
    }

    return int(object % 64 + 26)
}
